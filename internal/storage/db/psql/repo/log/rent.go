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
	return nil
}
