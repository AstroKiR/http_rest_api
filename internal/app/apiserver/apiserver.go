package apiserver

import "github.com/sirupsen/logrus"

type APIServer struct {
	config *Config
	logger *logrus.Logger
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

func (server *APIServer) configeLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) StartAPIServer() error {
	server.configeLogger()

	server.logger.Infof("Starting api server on port %s\n", server.config.BindAddr)
	return nil
}
