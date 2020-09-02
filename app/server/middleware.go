package server

import "net/http"

func (s *Server) addRoute(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	methods = append(methods, "OPTIONS")
	s.router.HandleFunc(path, handler).Methods(methods...)
}
