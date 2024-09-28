package resolver

import (
	"github.com/op/go-logging"

	srv "github.com/Egor123qwe/logs-storage/internal/service"
	model "github.com/Egor123qwe/logs-storage/pkg/proto"
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
