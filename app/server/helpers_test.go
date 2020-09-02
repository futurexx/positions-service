package server_test

import (
	"net/url"
	"testing"

	"github.com/futurexx/positions-service/app/server"
	"github.com/stretchr/testify/assert"
)

func TestGetOrDefault(t *testing.T) {
	values := url.Values{}
	values.Add("existed_key", "value")

	testCases := []struct {
		name         string
		vs           *url.Values
		key          string
		defaultValue string
		want         string
	}{
		{
			name:         "Test existed key",
			vs:           &values,
			key:          "existed_key",
			defaultValue: "default",
			want:         "value",
		},
		{
			name:         "Test default value",
			vs:           &values,
			key:          "nonexisted_key",
			defaultValue: "default",
			want:         "default",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, server.GetOrDefault(tc.vs, tc.key, tc.defaultValue))
		})
	}

}
