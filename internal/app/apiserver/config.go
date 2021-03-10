package apiserver

import (
	"encoding/json"
	"github.com/HumoFX/http-rest-api/internal/app/store"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BindAddr string `json:"bind_addr"`
	LogLevel string `json:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}
func (c *Config) SetConfig(configPath string) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		log.Print(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(byteValue, c)
	if err != nil {
		log.Fatal(err)
	}
}