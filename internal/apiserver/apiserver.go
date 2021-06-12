package apiserver

import (
	"github.com/adminoid/vuego/internal/config"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config config.Config
	logger *logrus.Logger
}

// New ...
func New(config config.Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}
