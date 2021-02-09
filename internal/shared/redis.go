package shared

import "github.com/go-redis/redis/v8"

var clients map[string]*redis.Client

func SetUpRedis(url string) *redis.Client {
	if clients == nil {
		clients = make(map[string]*redis.Client)
	}
	if _, ok := clients[url]; !ok {
		opt, err := redis.ParseURL(url)
		if err != nil {
			panic(err)
		}

		clients[url] = redis.NewClient(opt)
	}

	return clients[url]
}
