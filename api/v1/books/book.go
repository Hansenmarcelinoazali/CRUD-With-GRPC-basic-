package books

import (
	"latihan/configs"

	"github.com/sirupsen/logrus"
)

type Server struct {
	config *configs.Configs
	logger *logrus.Logger
}

func New(config *configs.Configs, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}
