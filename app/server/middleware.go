package server

import "net/http"

func (s *Server) addRoute(path string, handler func(http.ResponseWriter, *http.Request), methods ...string) {
	methods = append(methods, "OPTIONS")
	// TODO: CORPS, auth and claims with username
	s.router.HandleFunc(path, handler).Methods(methods...)
}
