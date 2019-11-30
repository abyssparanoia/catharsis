package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abyssparanoia/catharsis/pkg/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"net/http"

	env "github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
)

func main() {
	// initilize environment variables
	envObj := &environment{}
	if err := env.Parse(envObj); err != nil {
		panic(err)
	}

	logger, err := log.New(envObj.Envrionment)
	if err != nil {
		panic(err)
	}

	// Dependency
	d := dependency{}
	d.Inject(envObj)

	r := chi.NewRouter()
	routing(r, d)

	baseCtx := ctxzap.ToContext(context.Background(), logger)

	//server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", envObj.Port),
		Handler: chi.ServerBaseContext(baseCtx, r),
	}

	logger.Info(fmt.Sprintf("[START] server. port: %s\n", envObj.Port))

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Info(fmt.Sprintf("[CLOSED] server closed with error: %s\n", err))
		}
	}()

	// graceful shuttdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	logger.Info(fmt.Sprintf("SIGNAL %d received, so server shutting down now...\n", <-quit))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.Info(fmt.Sprintf("failed to gracefully shutdown: %s\n", err))
	}

	logger.Info(fmt.Sprintf("server shutdown completed\n"))
}
