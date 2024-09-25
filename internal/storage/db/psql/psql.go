package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Egor123qwe/logs-storage/internal/storage/repo/log"
)

type Store interface {
	Log() log.Log

	Close() error
}

type store struct {
	db  *sqlx.DB
	log log.Log
}

func New(config Config) (Store, error) {
	db, err := sqlx.Connect(config.logStorage.Driver, config.logStorage.URL)
	if err != nil {
		return nil, err
	}

	storage := store{
		db: db,
	}

	return storage, nil
}

func (s store) Close() error {
	return s.db.Close()
}

func (s store) Log() log.Log {
	return s.log
}
