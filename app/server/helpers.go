package server

import "net/url"

// GetOrDefault ...
func GetOrDefault(vs *url.Values, key string, defaultValue string) string {
	value := vs.Get(key)

	if value == "" {
		value = defaultValue
	}
	return value
}
