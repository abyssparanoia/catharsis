package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/abyssparanoia/catharsis/pkg/log"
)

const (
	port = ":50051"
)

func main() {

	// TODO: change logger mode by env
	logger, err := log.NewDevelopment()
	if err != nil {
		panic(err)
	}

	server := newAuthenticationServer(logger, port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for s := range c {
		logger.Info("Signal recieved", zap.String("signal", s.String()))

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			shutdown := make(chan bool, 1)

			go func() {

				server.GracefulStop()

				shutdown <- true
			}()

			select {
			case <-shutdown:
				logger.Info("Gracefully stop")
			case <-ctx.Done():
				logger.Info("Force stop")
				server.Stop()
			}

			cancel()

			logger.Info("Exit")
			return

		case syscall.SIGHUP:
		default:
			return

		}
	}

}
