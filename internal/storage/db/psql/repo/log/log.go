package log

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/Egor123qwe/logs-storage/internal/storage/model"
	repo2 "github.com/Egor123qwe/logs-storage/internal/storage/repo"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt/transaction"
)

const (
	requestTimeout = 15 * time.Second
)

type repo struct {
	db *sqlt.DB
}

func New(db *sqlt.DB) repo2.Log {
	return &repo{
		db: db,
	}
}

func (r repo) WithTransaction(ctx context.Context) (repo2.Log, transaction.Service, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	db, transactor := sqlt.NewTX(tx)

	return New(db), transactor, nil
}

func (r repo) AddLogs(ctx context.Context, logs ...model.LogReq) error {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil && !errors.Is(err, sqlt.ErrTxAlreadyStarted) {
		return err
	}

	// if error is not ErrTxAlreadyStarted
	if err == nil {
		defer tx.Commit()
	}

	query := `INSERT INTO logs 
        				(trace_id, module_id, time, level, message)
      		  	  VALUES ($1, $2, $3, $4, $5)`

	for _, log := range logs {
		err = r.db.QueryRowContext(ctx, query,
			log.TraceID,
			log.ModuleID,
			log.Time,
			log.Level.String(),
			log.Message,
		).Err()

		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return nil
}

func (r repo) GetLogs(ctx context.Context, req model.LogFilter) (model.LogResp, error) {
	result := model.LogResp{}

	sb := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	withFilter := func(builder sq.SelectBuilder) sq.SelectBuilder {
		result := builder.
			From("logs").
			InnerJoin("modules ON modules.id = logs.module_id").
			Where("message LIKE ?", "%"+req.Message+"%")

		if req.TraceID != nil {
			result = result.Where("trace_id = ?", *req.TraceID)
		}

		if req.ModuleID != nil {
			result = result.Where("module_id = ?", *req.ModuleID)
		}

		if req.Level != nil {
			result = result.Where("level = ?", (*req.Level).String())
		}

		startTime := time.Time{}
		endTime := time.Now()

		if req.StartTime != nil {
			startTime = *req.StartTime
		}

		if req.EndTime != nil {
			endTime = *req.EndTime
		}

		result = result.Where("time BETWEEN ? AND ?", startTime, endTime)

		return result
	}

	err := withFilter(sb.Select("COUNT(*)")).RunWith(r.db).QueryRowContext(ctx).Scan(&result.Total)
	if err != nil {
		return model.LogResp{}, err
	}

	request := sb.Select("logs.id", "logs.trace_id", "modules.name", "logs.time", "logs.level", "logs.message")
	request = withFilter(request).Limit(uint64(req.CountOnPage)).Offset(uint64((req.Page - 1) * req.CountOnPage))

	rows, err := request.RunWith(r.db).QueryContext(ctx)
	if err != nil {
		return model.LogResp{}, err
	}

	defer rows.Close()

	for rows.Next() {
		log := model.Log{}

		if err := rows.Scan(
			&log.ID,
			&log.TraceID,
			&log.Module,
			&log.Time,
			&log.Level,
			&log.Message,
		); err != nil {
			return model.LogResp{}, err
		}

		result.Logs = append(result.Logs, log)
	}

	return result, nil
}
