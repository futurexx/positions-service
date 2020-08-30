package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) respond(w http.ResponseWriter, data interface{}, status int) {
	if data != nil {
		w.Header().Set("Context-Type", "application/json")
		payload, _ := json.Marshal(data)
		_, err := w.Write(payload)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "server",
				"function": "respond",
				"error":    err,
				"payload":  payload,
			}).Warning("Internal server error.")

			status = http.StatusInternalServerError
		}
	}

	if status != 200 {
		w.WriteHeader(status)
	}
}

func (s *Server) handlerPing() http.HandlerFunc {
	type response struct {
		OK bool `json:"ok"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response{OK: true}
		s.respond(w, resp, http.StatusOK)
	}
}
