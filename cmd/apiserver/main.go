package main

import (
	"flag"
	"github.com/HumoFX/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.json", "path to config file")
}

func main()  {
	flag.Parse()

	config := apiserver.NewConfig()
	config.SetConfig(configPath)

	s := apiserver.New(config)
	if err := s.Start(); err !=nil {
		log.Fatal(err)
	}

}