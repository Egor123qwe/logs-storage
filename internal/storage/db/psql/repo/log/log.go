package log

import (
	"context"
	"errors"
	"time"

	logmodel "github.com/Egor123qwe/logs-storage/internal/model/log"
	"github.com/Egor123qwe/logs-storage/internal/storage/repo/log"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt/transaction"
)

const (
	requestTimeout = 15 * time.Second
)

var (
	ErrRentNotFound = errors.New("log not found")
)

type repo struct {
	db *sqlt.DB
}

func New(db *sqlt.DB) log.Log {
	return &repo{
		db: db,
	}
}

func (r repo) WithTransaction(ctx context.Context) (log.Log, transaction.Service, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	db, transactor := sqlt.NewTX(tx)

	return New(db), transactor, nil
}

func (r repo) Add(ctx context.Context, logs ...logmodel.Log) error {
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

	logsQuery := `INSERT INTO logs 
        				(trace_id, time, module, level, message)
      		  	  VALUES ($1, $2, $3, $4, $5)`

	for _, log := range logs {
		err = r.db.QueryRowContext(ctx, logsQuery,
			log.TraceID,
			log.Time,
			log.Module,
			log.Level,
			log.Message,
		).Scan()

		if err != nil {
			tx.Rollback()

			return err
		}
	}

	return nil
}
