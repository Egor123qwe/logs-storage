package grpc

import (
	"google.golang.org/grpc"

	api "github.com/Egor123qwe/logs-storage/internal/handler/grpc/generate"
	"github.com/Egor123qwe/logs-storage/internal/handler/grpc/resolver"
	srv "github.com/Egor123qwe/logs-storage/internal/service"
)

type Handler interface {
	Subscribe(server *grpc.Server)
}

type handler struct {
	api resolver.Handler
}

func New(srv srv.Service) Handler {
	return &handler{
		api: resolver.New(srv),
	}
}

func (h *handler) Subscribe(server *grpc.Server) {
	api.RegisterLogsServer(server, h.api)
}
