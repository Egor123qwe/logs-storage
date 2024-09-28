package log

import (
	"context"

	"github.com/lib/pq"

	"github.com/Egor123qwe/logs-storage/internal/storage/model"
)

const (
	uniqueViolationErr = "23505"
)

func (r repo) InitModule(ctx context.Context, module string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	var id int64

	query := `INSERT INTO modules 
        				(name)
      		  VALUES ($1) RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		module,
	).Scan(
		&id,
	)

	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == uniqueViolationErr {
			return r.getModule(ctx, module)
		}

		return 0, err
	}

	return id, nil
}

func (r repo) GetModules(ctx context.Context, req model.ModuleReq) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	var result []string

	query := `SELECT name FROM modules 
              WHERE name LIKE $1`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		"%"+req.NameFilter+"%",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var moduleName string

		if err := rows.Scan(&moduleName); err != nil {
			return nil, err
		}

		result = append(result, moduleName)
	}

	return result, nil
}

func (r repo) getModule(ctx context.Context, module string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	var id int64

	query := `SELECT id FROM modules WHERE name = $1`

	err := r.db.QueryRowContext(ctx, query,
		module,
	).Scan(
		&id,
	)

	if err != nil {
		return 0, err
	}

	return id, nil
}
