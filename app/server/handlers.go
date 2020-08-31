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
			}).Warning("Failed to write payload.")

			status = http.StatusInternalServerError
		}
	}

	if status != 200 {
		w.WriteHeader(status)
	}
}

func (s *Server) handlerPing() http.HandlerFunc {
	type response struct {
		Ok bool `json:"ok"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response{Ok: true}
		s.respond(w, resp, http.StatusOK)
	}
}

func (s *Server) handlerSummary() http.HandlerFunc {
	type response struct {
		Domain         string `json:"domain"`
		PositionsCount uint   `json:"positions_count"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			s.logger.Warning("`domain` param is required")
			s.respond(w, nil, http.StatusBadRequest)
		}

		positionsCount, err := s.storage.Positions().Summary(domain)
		if err != nil {
			s.logger.WithFields(logrus.Fields{
				"package":  "server",
				"function": "handlerSummary",
				"error":    err,
				"domain":   domain,
			}).Warning("Failed to get summary for domain")

			s.respond(w, nil, http.StatusInternalServerError)
		}

		resp := response{
			Domain:         domain,
			PositionsCount: positionsCount}
		s.respond(w, resp, http.StatusOK)
	}
}
