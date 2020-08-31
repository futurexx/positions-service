package server

import (
	"github.com/futurexx/positions-service/app/storage"
)

// Config is view of toml config file
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

// InitConfig ...
func InitConfig() *Config {
	return &Config{Storage: storage.InitConfig()}
}
