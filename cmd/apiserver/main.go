package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/astrokir/http_rest_api/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to api server cofig file")
}

func main() {
	flag.Parse()

	config := apiserver.GetConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server := apiserver.NewAPIServer(config)

	if err := server.StartAPIServer(); err != nil {
		log.Fatal(err)
	}
}
