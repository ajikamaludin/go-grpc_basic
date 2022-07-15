package configs

import (
	"log"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/config"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/postgres"
	"github.com/sirupsen/logrus"
)

// bundle/wrap another service access here
type Configs struct {
	Config *config.Config
	Pg     *postgres.Conn
}

func New() (*Configs, *logrus.Logger, error) {
	config, err := config.New()

	if err != nil {
		return nil, nil, err
	}

	// force all writes to regular log to logger
	logger := logrus.New()
	log.SetOutput(logger.Writer()) // TODO: ?
	log.SetFlags(0)                // TODO: ?

	logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		ForceQuote:    true,
		FullTimestamp: true,
	}

	logger.Info("[CONFIG] Setup complete")

	// Test Pg Con
	pg, err := postgres.New(config)
	if err != nil {
		return nil, nil, err
	}

	return &Configs{
		Config: config,
		Pg:     pg,
	}, logger, nil
}
