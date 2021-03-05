package config

type DataTransformer struct {
	RedisConf  dtRedisConf  `yaml:"redis"`
	DynamoConf dtDynamoConf `yaml:"dynamo"`
}

type dtRedisConf struct {
	QueueUrl string `yaml:"queueUrl"`
}

type dtDynamoConf struct {
	DBUrl string `yaml:"dbUrl"`
}
