// Package server configures and starts servers for handling incoming requests.
package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"

	"github.com/Egor123qwe/logs-storage/internal/broker"
	"github.com/Egor123qwe/logs-storage/internal/handler"
	"github.com/Egor123qwe/logs-storage/internal/server/launcher"
	"github.com/Egor123qwe/logs-storage/internal/server/launcher/grpc"
	"github.com/Egor123qwe/logs-storage/internal/server/launcher/kafka"
	"github.com/Egor123qwe/logs-storage/internal/service"
)

var log = logging.MustGetLogger("server")

type server struct {
	servers []launcher.Server
}

func New(srv service.Service) (launcher.Server, error) {
	brk, err := broker.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create broker: %w", err)
	}

	h := handler.New(srv, brk)

	result := &server{
		servers: []launcher.Server{
			kafka.New(brk.Kafka, h.Event),
			grpc.New(grpc.NewConfig(), h.GRPC),
		},
	}

	return result, nil
}

func (s *server) Serve(ctx context.Context) error {
	gr, grCtx := errgroup.WithContext(ctx)

	// start server
	gr.Go(func() error {
		return s.serve(grCtx)
	})

	var err error

	if err = gr.Wait(); err != nil {
		log.Criticalf("app error: %v", err)
	}

	log.Infof("app: shutting down the server...")

	return err
}

func (s *server) serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, s := range s.servers {
		s := s

		gr.Go(func() error {
			defer wg.Done()

			return s.Serve(grCtx)
		})
	}

	wg.Wait()

	return gr.Wait()
}
