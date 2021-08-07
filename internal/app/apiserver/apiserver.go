package apiserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/test1", test1)
	server.router.HandleFunc("/test2", test2)
}

func (server *APIServer) StartAPIServer() error {
	if err := server.configureLogger(); err != nil {
		log.Fatal(err)
	}
	server.configureRouter()
	server.logger.Infof("Starting api server on port %s\n", server.config.BindAddr)
	return http.ListenAndServe(server.config.BindAddr, server.router)
}
