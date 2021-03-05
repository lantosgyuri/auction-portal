package main

import (
	data_transformer "github.com/lantosgyuri/auction-portal/internal/data-transformer"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/lantosgyuri/auction-portal/internal/pkg/input"
	"gopkg.in/yaml.v3"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var conf config.DataTransformer

	confBytes, err := input.ReadFile("config.yaml")

	if err != nil {
		log.Fatal("Can not read config file")
	}

	err = yaml.Unmarshal(confBytes, &conf)

	data_transformer.StartSubscriber(conf, &wg)
	wg.Wait()
}
