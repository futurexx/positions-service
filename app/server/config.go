package server

// Config is view of toml config file
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// InitConfig ...
func InitConfig() *Config {
	return &Config{}
}
