package transactor

import (
	"github.com/jmoiron/sqlx"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt/transaction"
)

type service struct {
	tx *sqlx.Tx
}

func New(tx *sqlx.Tx) transaction.Service {
	return &service{
		tx: tx,
	}
}

func (s *service) Rollback() error {
	return s.tx.Rollback()
}

func (s *service) Commit() error {
	return s.tx.Commit()
}
