package service

import (
	"github.com/Egor123qwe/logs-storage/internal/service/log"
	"github.com/Egor123qwe/logs-storage/internal/storage"
)

type Service interface {
	Logs() log.Service
}

type service struct {
	session log.Service
}

func New(storage storage.Storage) Service {
	return &service{
		session: log.New(storage),
	}
}

func (s *service) Logs() log.Service {
	return s.session
}
