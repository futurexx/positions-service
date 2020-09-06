package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server = Server{storage: &MockStorage{}}
var brokenServer = Server{storage: &MockBrokenStorage{}}

func TestHandlerPing(t *testing.T) {
	testCases := []struct {
		name     string
		isBroken bool
		want     []byte
	}{
		{
			name:     "Test response OK",
			isBroken: false,
			want:     []byte(`{"ok":true}`),
		},
		{
			name:     "Test response broken still OK",
			isBroken: true,
			want:     []byte(`{"ok":true}`),
		},
	}

	for _, tc := range testCases {
		s := server
		if tc.isBroken {
			s = brokenServer
		}

		handler := http.HandlerFunc(s.handlerPing())
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/api/monitoring/ping", nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, tc.want, rec.Body.Bytes())

		})
	}
}

func TestHandlerSummary(t *testing.T) {
	testCases := []struct {
		name     string
		domain   string
		status   int
		isBroken bool
		want     []byte
	}{
		{
			name:     "Test response existed domain",
			domain:   "existed.domain",
			status:   http.StatusOK,
			isBroken: false,
			want:     []byte(`{"domain":"existed.domain","positions_count":5}`),
		},
		{
			name:     "Test response not existed domain",
			domain:   "not.existed.domain",
			status:   http.StatusOK,
			isBroken: false,
			want:     []byte(`{"domain":"not.existed.domain","positions_count":0}`),
		},
		{
			name:     "Test response bad domain",
			domain:   "",
			status:   http.StatusBadRequest,
			isBroken: false,
			want:     []byte(`{"domain":"","positions_count":0}`),
		},
		{
			name:     "Test response bad domain",
			domain:   "",
			status:   http.StatusBadRequest,
			isBroken: false,
			want:     []byte(`{"domain":"","positions_count":0}`),
		},
		{
			name:     "Test response broken storage",
			domain:   "existed.domain",
			isBroken: true,
			status:   http.StatusInternalServerError,
			want:     []byte(`{"domain":"existed.domain","positions_count":0}`),
		},
	}

	for _, tc := range testCases {
		s := server
		if tc.isBroken {
			s = brokenServer
		}

		handler := http.HandlerFunc(s.handlerSummary())
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			url := fmt.Sprintf("/api/pasitions/summary?domain=%s", tc.domain)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.status, rec.Code)
			assert.Equal(t, tc.want, rec.Body.Bytes())

		})
	}
}

func TestHandlerPositions(t *testing.T) {
	testCases := []struct {
		name     string
		domain   string
		isBroken bool
		status   int
	}{
		{
			name:     "Test response existed domain",
			domain:   "existed.domain",
			isBroken: false,
			status:   http.StatusOK,
		},
		{
			name:     "Test response not existed domain",
			domain:   "not.existed.domain",
			isBroken: false,
			status:   http.StatusOK,
		},
		{
			name:     "Test response bad domain",
			domain:   "",
			isBroken: false,
			status:   http.StatusBadRequest,
		},
		{
			name:     "Test response broken storage",
			domain:   "existed.domain",
			isBroken: true,
			status:   http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		s := server
		if tc.isBroken {
			s = brokenServer
		}

		handler := http.HandlerFunc(s.handlerPositions())
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			url := fmt.Sprintf("/api/pasitions?domain=%s", tc.domain)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler.ServeHTTP(rec, req)

			assert.Equal(t, tc.status, rec.Code)

		})
	}
}
