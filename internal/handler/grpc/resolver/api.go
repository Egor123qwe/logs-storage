package resolver

import (
	"github.com/op/go-logging"

	model "github.com/Egor123qwe/logs-storage/internal/handler/grpc/generate"
	srv "github.com/Egor123qwe/logs-storage/internal/service"
)

var log = logging.MustGetLogger("worker handler")

type Handler struct {
	model.UnimplementedLogsServer
	srv srv.Service
}

func New(srv srv.Service) Handler {
	return Handler{
		srv: srv,
	}
}
