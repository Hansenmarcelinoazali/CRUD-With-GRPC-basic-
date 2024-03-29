package configs

import (
	"latihan/pkg/v1/config"
	"latihan/pkg/v1/postgres"
	"log"

	"github.com/sirupsen/logrus"
)

type Configs struct {
	Config *config.Config
	Pg     *postgres.Conn
}

// Setup sets up the environment.
func New() (*Configs, *logrus.Logger, error) {
	config, err := config.New()
	if err != nil {
		return nil, nil, err
	}

	logger := logrus.New()

	// force all writes to regular log to logger
	log.SetOutput(logger.Writer())
	log.SetFlags(0)

	// configure logging for environment
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		ForceQuote:    true,
		FullTimestamp: true,
	}

	logger.Info("[CONFIG] Setup complete")

	pg, err := postgres.New(config)
	if err != nil {
		return nil, nil, err
	}

	logger.Infof("[CLIENT] setup complete")

	return &Configs{
		Config: config,
		Pg:     pg,
	}, logger, nil
}
