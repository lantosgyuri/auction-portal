package main

import (
	command_service "github.com/lantosgyuri/auction-portal/internal/pkg/command-service"
	"github.com/lantosgyuri/auction-portal/internal/pkg/input"
	"gopkg.in/yaml.v3"
	"log"
)

type Config struct {
	RedisConf redisConf `yaml:"redis"`
}

type redisConf struct {
	Url string `yaml:"url"`
}

func main() {
	quit := make(chan bool)
	var conf Config

	confBytes, err := input.ReadFile("config.yaml")

	if err != nil {
		log.Fatal("Can not read config file")
	}

	err = yaml.Unmarshal(confBytes, &conf)

	if err != nil {
		log.Fatal("Can not unmarshal config file")
	}

	command_service.StartSubscriber(conf.RedisConf.Url)
	<-quit
}
