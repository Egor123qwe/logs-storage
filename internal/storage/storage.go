package storage

import (
	"github.com/Egor123qwe/logs-storage/internal/storage/db/psql"
	"github.com/Egor123qwe/logs-storage/internal/storage/repo"
)

type Storage interface {
	Log() repo.Log

	Close() error
}

type storage struct {
	psql psql.Storage
}

func New() (Storage, error) {
	var err error
	var storage storage

	storage.psql, err = psql.New(psql.NewConfig())
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s storage) Close() error {
	if err := s.psql.Close(); err != nil {
		return err
	}

	return nil
}
