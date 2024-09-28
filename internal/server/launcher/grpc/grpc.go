package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	handler "github.com/Egor123qwe/logs-storage/internal/handler/grpc"
	"github.com/Egor123qwe/logs-storage/internal/server/launcher"
)

var log = logging.MustGetLogger("grpc")

type server struct {
	config Config
	srv    *grpc.Server
}

func New(config Config, handler handler.Handler) launcher.Server {
	srv := grpc.NewServer()
	reflection.Register(srv)

	handler.Subscribe(srv)

	return &server{
		config: config,
		srv:    srv,
	}
}

func (s *server) Serve(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		log.Fatalf("Failed to listen grpc server: %v", err)
	}

	log.Infof("serving gRPC on http://localhost:%d\n", s.config.Port)

	errCh := make(chan error)

	go func() {
		errCh <- s.srv.Serve(lis)

		close(errCh)
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("grpc-server: %w", err)

	case <-ctx.Done():
		s.srv.GracefulStop()
		log.Infof("grpc-server: server stopped successfully.")
	}

	return nil
}
