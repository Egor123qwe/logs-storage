package log

import (
	"context"

	"github.com/Egor123qwe/logs-storage/internal/model/log"
	"github.com/Egor123qwe/logs-storage/internal/storage"
)

type Service interface {
	Add(ctx context.Context, logs ...log.Log) error
}

type service struct {
	storage storage.Storage
}

func New(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s service) Add(ctx context.Context, logs ...log.Log) error {
	return nil
}
