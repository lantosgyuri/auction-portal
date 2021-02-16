package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

var ctx = context.Background()

type Bid struct {
	AuctionId int    `json:"AuctionBid"`
	UserId    int    `json:"UserId"`
	Value     int    `json:"Value"`
	Action    string `json:"Action"`
}

type CreateAuction struct {
	Name      string
	DueDate   int
	StartDate int
	Timestamp int
}

type Event struct {
	Event   string
	Payload []byte
}

func main() {
	publish()
}

func CreateRedisClient() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6364")
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}

func publish() {
	redisConn := CreateRedisClient()
	/*
					auction := CreateAuction{
						DueDate:   int(time.Now().AddDate(0, 0, 7).Unix()),
						StartDate: int(time.Now().AddDate(0, 0, 2).Unix()),
						Timestamp: int(time.Now().Unix()),
						Name:      "JOZSIKA",
					}

					userBytes, _ := json.Marshal(auction)

					event := Event{
						Event:   domain.AuctionRequested,
						Payload: userBytes,
					}

					messageBytes, _ := json.Marshal(event)

					redisConn.Publish(ctx, "Auction", messageBytes)

				user := domain.CreateUserRequested{
					Name:     "Johanna",
					Password: "Secret",
				}

				userCreateBytes, _ := json.Marshal(user)

				eventJohannaCreate := Event{
					Event:   domain.UserCreateRequested,
					Payload: userCreateBytes,
				}

				johannaBytes, _ := json.Marshal(eventJohannaCreate)

				redisConn.Publish(ctx, "User", johannaBytes)

				userIvan := domain.CreateUserRequested{
					Name:     "Ivan",
					Password: "Top Secret",
				}

				userIvanCreateBytes, _ := json.Marshal(userIvan)

				eventIvanCreate := Event{
					Event:   domain.UserCreateRequested,
					Payload: userIvanCreateBytes,
				}

				ivanBytes, _ := json.Marshal(eventIvanCreate)

				redisConn.Publish(ctx, "User", ivanBytes)

				userME := domain.CreateUserRequested{
					Name:     "Gyorgy",
					Password: "Top Secret",
				}

				userMECreateBytes, _ := json.Marshal(userME)

				eventMECreate := Event{
					Event:   domain.UserCreateRequested,
					Payload: userMECreateBytes,
				}

				meBytes, _ := json.Marshal(eventMECreate)

				redisConn.Publish(ctx, "User", meBytes)

			userME := domain.DeleteUserRequest{
				Name: "Gyorgy",
				Id:   3,
			}

			userMECreateBytes, _ := json.Marshal(userME)

			eventMECreate := Event{
				Event:   domain.UserDeleteRequested,
				Payload: userMECreateBytes,
			}

			meBytes, _ := json.Marshal(eventMECreate)

			redisConn.Publish(ctx, "User", meBytes)


		winner := domain.WinnerAnnounced{
			WinnerId:  2,
			AuctionId: "3b1b4a43-005f-4099-9ed8-68b3905ef2c9",
		}

		winnerBytes, _ := json.Marshal(winner)

		eventWinner := Event{
			Event:   domain.AuctionWinnerAnnounced,
			Payload: winnerBytes,
		}

		winnerMessageBytes, _ := json.Marshal(eventWinner)

		redisConn.Publish(ctx, "Auction", winnerMessageBytes)
	*/

	bidPLaced := domain.BidPlaced{
		Promoted:  false,
		Amount:    20,
		UserId:    2,
		AuctionId: "56bc900f-96fe-42ce-90d2-e5ef1b512157",
	}

	bidPLacedbytes, _ := json.Marshal(bidPLaced)

	bidPlacedEvent := Event{
		Event:   domain.BidPlaceRequested,
		Payload: bidPLacedbytes,
	}

	bidPLacedEvetnBytes, _ := json.Marshal(bidPlacedEvent)

	redisConn.Publish(ctx, "Bid", bidPLacedEvetnBytes)

	/*
		bidDeleted := domain.BidDeleted{
			BidId:     40,
			Amount:    20,
			UserId:    41,
			AuctionId: "test",
		}

		bidDeletedbytes, _ := json.Marshal(bidDeleted)

		bidDeletedEcent := Event{
			Event:   domain.BidDeleteRequested,
			Payload: bidDeletedbytes,
		}

		bidDeletedEvetnBytes, _ := json.Marshal(bidDeletedEcent)

		redisConn.Publish(ctx, "Bid", bidDeletedEvetnBytes)
	*/

}
