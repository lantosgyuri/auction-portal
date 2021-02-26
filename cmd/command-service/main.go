package main

import (
	command_service "github.com/lantosgyuri/auction-portal/internal/command-service"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"github.com/lantosgyuri/auction-portal/internal/pkg/input"
	"gopkg.in/yaml.v3"
	"log"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var conf config.CommandService

	confBytes, err := input.ReadFile("config.yaml")

	if err != nil {
		log.Fatal("Can not read config file")
	}

	err = yaml.Unmarshal(confBytes, &conf)

	if err != nil {
		log.Fatal("Can not unmarshal config file")
	}

	connection.InitializeMariaDb(conf.SotDbConf.Dsn)
	defer connection.CloseMariaDb()
	adapter.MigrateSotDb(&wg)

	command_service.StartSubscriber(conf, &wg)
	wg.Wait()
}
