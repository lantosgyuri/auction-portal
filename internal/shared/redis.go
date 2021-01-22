package shared

import "github.com/go-redis/redis/v8"

func SetUpRedis(url string) *redis.Client {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}
