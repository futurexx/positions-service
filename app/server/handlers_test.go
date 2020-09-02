package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerPing(t *testing.T) {
	s := Server{}
	testCases := []struct {
		name string
		want []byte
	}{
		{
			name: "Test response OK",
			want: []byte(`{"ok":true}`),
		},
	}

	handler := http.HandlerFunc(s.handlerPing())
	for _, tc := range testCases {
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
