package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sso/internal/grpc/auth"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) run() error {
	const op = "grpcapp.run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting gRPC server", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() error {
	const op = "grpcapp.Stop"
	log := a.log.With(slog.String("op", op))
	log.Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()

	return nil
}
