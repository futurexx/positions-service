package server

import (
	"github.com/futurexx/semrush_test_task/app/storage"
)

// Config is view of toml config file
type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LogLevel      string `toml:"log_level"`
	StorageConfig *storage.Config
}

// InitConfig ...
func InitConfig() *Config {
	return &Config{}
}
