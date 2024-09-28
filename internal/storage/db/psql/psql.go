package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	logrepo "github.com/Egor123qwe/logs-storage/internal/storage/db/psql/repo/log"
	"github.com/Egor123qwe/logs-storage/internal/storage/repo"
	"github.com/Egor123qwe/logs-storage/pkg/sqlt"
)

type Storage interface {
	Log() repo.Log

	Close() error
}

type store struct {
	db  *sqlx.DB
	log repo.Log
}

func New(config Config) (Storage, error) {
	db, err := sqlx.Connect(config.db.Driver, config.db.URL)
	if err != nil {
		return nil, err
	}

	storage := store{
		db:  db,
		log: logrepo.New(sqlt.NewDB(db)),
	}

	return storage, nil
}

func (s store) Log() repo.Log {
	return s.log
}

func (s store) Close() error {
	return s.db.Close()
}
