package server

import (
	"net/http"
	"os"

	"github.com/futurexx/positions-service/app/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
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

func (s *Server) configureRouter() {
	s.addRoute("/api/monitoring/ping", s.handlerPing(), "GET")
	s.addRoute("/api/positions/summary", s.handlerSummary(), "GET")
}

func (s *Server) configureStorage() {
	st := storage.New(s.config.Storage)
	s.storage = st
}

// Start is a method of starting server
func Start(config *Config) error {
	server := new(config)

	err := server.configureLogger()
	if err != nil {
		server.logger.Error("Failed to configure logger")
		return err
	}
	server.logger.Info("Logger configureted")

	server.configureRouter()
	server.logger.Info("Router configurated")

	server.configureStorage()
	server.logger.Info("Storage configurated")

	if err := server.storage.Open(); err != nil {
		server.logger.Error("Failed to create database connection")
		return err
	}
	server.logger.Info("Database connected")
	defer func() {
		if err := server.storage.Close(); err != nil {
			server.logger.Error("Failed to close database connection")
		}
	}()

	server.logger.Info("Starting server...")
	return http.ListenAndServe(server.config.BindAddr, handlers.LoggingHandler(os.Stdout, server.router))
}
