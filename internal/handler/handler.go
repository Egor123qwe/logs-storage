package handler

import (
	"github.com/Egor123qwe/logs-storage/internal/broker"
	"github.com/Egor123qwe/logs-storage/internal/handler/event"
	"github.com/Egor123qwe/logs-storage/internal/service"
	"github.com/Egor123qwe/logs-storage/pkg/msghandler"
)

type Handler struct {
	Event msghandler.MsgResolver
}

func New(srv service.Service, broker broker.Broker) Handler {
	handler := Handler{
		Event: event.New(srv, broker.Kafka),
	}

	return handler
}
