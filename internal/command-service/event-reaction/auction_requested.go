package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/port"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

func CreateAuctionRequestedCommand(conf config.CommandService) EventReactor {
	handler := command.CreateAuctionHandler{
		Repo: adapter.CreateMariaDbAuctionRepository(),
	}
	preserver := command.SaveAuctionEventHandler{Repo: adapter.CreateMariaDbAuctionRepository()}
	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.AuctionChannel)
	return EventReactor{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}
