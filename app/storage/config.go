package storage

// Config is a view of database configuration
type Config struct {
	DatabaseFile string `toml:"db_file"`
}

// InitConfig ...
func InitConfig() *Config {
	return &Config{}
}
