package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajikamaludin/go-grpc_basic/configs"
	"github.com/ajikamaludin/go-grpc_basic/router"
	"golang.org/x/sync/errgroup"
)

func main() {
	var err error

	configs, logger, err := configs.New()
	if err != nil {
		log.Printf("[SERVER] ERROR %v", err)
		log.Fatal(err)
	}

	logger.Infof("[SERVER] Environment %s is ready", configs.Config.Env)

	g, _ := errgroup.WithContext(context.Background())

	var servers []*http.Server

	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			logger.Infof("receive signal: %s\n", sig)
			for i, s := range servers {
				if err := s.Shutdown(context.Background()); err != nil {
					if err == nil {
						logger.Infof("error shutting down server %d: %v", i, err)
						return err
					}
				}
			}
			os.Exit(1)
		}
		return nil
	})

	g.Go(func() error { return router.NewGRPCServer(configs, logger) })
	g.Go(func() error { return router.NewHTTPServer(configs, logger) })

	if err := g.Wait(); !router.IgnoreErr(err) {
		log.Printf("[SERVER] ERROR %v", err)
		logger.Fatal(err)
	}

	logger.Infoln("[APP] DONE.")
}
