package server

import (
	"net/http"
	"os"

	"github.com/futurexx/semrush_test_task/app/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func new(config *Config) *Server {
	server := &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

	return server
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

// handlers registration
func (s *Server) configureRouter() {
	s.addRoute("/api/monitoring/ping", s.handlerPing(), "GET")
}

// Start is a method of starting server
func Start(config *Config) error {
	server := new(config)

	err := server.configureLogger()
	if err != nil {
		return err
	}
	server.logger.Info("Logger configureted")

	server.configureRouter()
	server.logger.Info("Router configurated")

	server.logger.Info("Starting server...")
	return http.ListenAndServe(server.config.BindAddr, handlers.LoggingHandler(os.Stdout, server.router))
}
