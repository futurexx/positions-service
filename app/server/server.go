package server

import (
	"net/http"
	"os"

	"github.com/futurexx/positions-service/app/storage"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// log - logger
var log = logrus.New()

// Server ...
type Server struct {
	config  *Config
	router  *mux.Router
	storage storage.IStorage
}

func new(config *Config) *Server {
	server := &Server{
		config: config,
		router: mux.NewRouter(),
	}

	return server
}

// ConfigureLogger ...
func configureLogger(logLevel string) error {

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	log.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.addRoute("/api/monitoring/ping", s.handlerPing(), "GET")
	s.addRoute("/api/positions/summary", s.handlerSummary(), "GET")
	s.addRoute("/api/positions", s.handlerPositions(), "GET")
}

func (s *Server) configureStorage() {
	st := storage.New(s.config.Storage)
	s.storage = st
}

// Start is a method of starting server
func Start(config *Config) error {
	server := new(config)

	configureLogger(config.LogLevel)
	if err := configureLogger(config.LogLevel); err != nil {
		return err
	}
	log.Info("Logger configureted")

	server.configureRouter()
	log.Info("Router configurated")

	server.configureStorage()
	log.Info("Storage configurated")

	if err := server.storage.Open(); err != nil {
		log.Error("Failed to create database connection")
		return err
	}
	log.Info("Database connected")
	defer func() {
		if err := server.storage.Close(); err != nil {
			log.Error("Failed to close database connection")
		}
	}()

	log.Info("Starting server...")
	return http.ListenAndServe(server.config.BindAddr, handlers.LoggingHandler(os.Stdout, server.router))
}
