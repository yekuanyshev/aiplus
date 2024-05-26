package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/yekuanyshev/aiplus/config"
	"github.com/yekuanyshev/aiplus/internal/repository"
	"github.com/yekuanyshev/aiplus/internal/service"
	"github.com/yekuanyshev/aiplus/internal/transport/rest"
	"github.com/yekuanyshev/aiplus/pkg/postgres"
)

func main() {
	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	conf, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		return
	}

	pool, err := postgres.Connect(ctx, conf.PgDSN)
	if err != nil {
		slog.Error("failed to connect to postgres", slog.Any("error", err))
	}

	repository := repository.NewManager(pool)
	service := service.NewManager(repository)
	rest := rest.New(service)

	go func() {
		log.Println("starting...")
		err = rest.Start(conf.HTTPListen)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("stopping...")
	err = rest.Stop()
	if err != nil {
		log.Println("failed to stop rest", err)
	}

	log.Println("postgres pool closing...")
	pool.Close()
}
