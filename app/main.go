package main

import (
	"flag"

	"github.com/BurntSushi/toml"

	"github.com/futurexx/semrush_test_task/app/server"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/config.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := server.InitConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		panic(err)
	}

	if err := server.Start(config); err != nil {
		panic(err)
	}
}
