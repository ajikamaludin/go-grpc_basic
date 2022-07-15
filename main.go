package main

import (
	"log"

	"github.com/ajikamaludin/go-grpc_basic/configs"
)

func main() {
	var err error

	configs, logger, err := configs.New()
	if err != nil {
		log.Printf("[SERVER] ERROR %v", err)
		log.Fatal(err)
	}

	logger.Infof("[SERVER] Environment %s is ready", configs.Config.Env)
}
