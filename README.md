# Auction portal
##### _A weekend project to implement cool things from cool articles_

### Domain
- Create auction
- Create user
- Delete user
- Place bid
- Delete bid
- Double bids through a promoter
- Announce winner

User can place bid what is smaller than the current highest bid, but has to be higher than the latest bid by the user.
User can delete bids.
The Promoter picks a random bid in a predefined intervall and doubles it for the lucky users.

## Architecture

I'm using event sourcing with CQRS in this project (_articles_ [1], [2]). The async response-request pattern is used to communicate with the client (_article_ [3]).
The services communicate with each other only with message queues using Redis PubSub. I chose this because the main benefit of the message queue: Fast response to the client and independent scaling of the response and consumer layer.
The event database is an SQL database because of the ACIDity. The SQL database saving the events, auction snapshots, users, bids and auctions. This database is the "source of truth".
The query database is a NoSql database where the data is already saved in the form as the client needs it.
The packages follows ports and adapters architecture (_talk_ [4])

### Services

##### Rest api(TBD)
- Routes:
  GET /auctions: returns all active auctions
  PUT /auction/:id/bid: adds a new bid to an auction
  GET /auction/:id/bids: get a user’s bid for an auction
  DELETE /auction/:id/bid/:id: delete a user’s bid for an auction

##### Write command service (DONE)
- Handling the writing part to write in the SQL DB. It consumes events from the `WriteQueue`. Builds the aggregate if needed and validates the write requests.
- Subscribed to 3 channels: `User`, `Bid`, `Auction`. Currently the channels are handled by one queue but I chose to split them to provide a consistent time for Auction and User creation as these messages will be not stacked up on bid messages.
- For handling the event reaction I'm using the _Command Pattern_ (_article_ [5]).

##### Data transformer (TBD)
- Subscribed to the events coming out from the `Write Command Service`. Transforms data to be consume ready.

##### Notifier (TBD)
- Subscribed to the events coming out from the `Write Command Service`. Users can subscribe to notifications: won auctions, max bid updates for auctions they participated in, and doubled bids.

##### Auction creator (TBD)
- CLI tool to create Auctions.

##### Promoter (TBD)
- Randomly, in each auction a bid is doubled by the auction house.
- How a bid is promoted is most likely something that the user wants to configure later on. To have the promoter as flexible as possible I'm using the _Functional Options_ pattern (_article_ [6]). This allows me to later easily attach for example a dashboard where the user can configure which bids should be promoted with higher chance.

Architecture as an image (_sorry I left some notes there :|_):
![alt text](https://i.ibb.co/PC0GbGC/arch.png "Architecture as an image")


## CAP Theorem fit
The source of truth part (the event store) is more consistent than available. The command service computes a lot before saving one event.
The Query part is more available as it does not compute. But it is not consistent as it has a delay before getting the real data.
I think this is the good solution as the source of the truth part should be as consistent as possible. And the Query part as available as possible.
(_article_ [7])

### Things to be implemented after the MVP:
- Auth
- Cache
- Pessimistic lock
- Exponential backoff retry(if the message can not be consumed right away)
- Distributed transactions (slice the write service, and add more functionality)
- Force Sync from user (to have better consistency at the client side)


[1]: https://victoramartinez.com/posts/event-sourcing-in-go/
[2]: https://threedots.tech/post/basic-cqrs-in-go/
[3]: https://docs.microsoft.com/en-us/azure/architecture/patterns/async-request-reply
[4]: https://www.youtube.com/watch?v=vKbVrsMnhDc
[5]: https://levelup.gitconnected.com/the-command-pattern-with-go-fd5dabc84c7
[6]: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
[7]: https://towardsdatascience.com/cap-theorem-and-distributed-database-management-systems-5c2be977950e

