package event

import (
	"github.com/op/go-logging"

	"github.com/Egor123qwe/logs-storage/internal/broker/kafka"
	"github.com/Egor123qwe/logs-storage/internal/handler/model/msg"
	"github.com/Egor123qwe/logs-storage/internal/handler/model/msg/event"
	"github.com/Egor123qwe/logs-storage/internal/service"
	"github.com/Egor123qwe/logs-storage/pkg/msghandler"
)

var log = logging.MustGetLogger("content handler")

type handler struct {
	srv    service.Service
	router msghandler.MsgHandler
}

func New(srv service.Service, broker kafka.Service) msghandler.MsgResolver {
	// type parser for messages of our contract type
	eventParser := func(m []byte) (string, error) {
		msg, err := msg.New(m).Parse()

		return msg.Type, err
	}

	handler := &handler{
		srv:    srv,
		router: msghandler.New(eventParser),
	}

	handler.initEvents()

	return handler.router
}

func (h handler) initEvents() {
	h.router.Add(string(event.AddLogs), h.AddLogs)
}
