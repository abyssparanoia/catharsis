package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	port = ":50051"
)

func main() {

	server := newAuthenticationServer(port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for s := range c {
		log.Printf("Signal %s recieved\n", s.String())

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
				log.Println("Gracefully stop")
			case <-ctx.Done():
				log.Println("Force stop")
				server.Stop()
			}

			cancel()

			log.Println("Exit")
			return

		case syscall.SIGHUP:
		default:
			return

		}
	}

}
