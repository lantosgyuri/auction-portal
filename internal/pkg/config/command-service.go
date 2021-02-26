package config

type CommandService struct {
	RedisConf csRedisConf `yaml:"redis"`
	SotDbConf csSotDbConf `yaml:"sotDb"`
}

type csRedisConf struct {
	WriteUrl string `yaml:"writeUrl"`
	QueryUrl string `yaml:"queryUrl"`
}

type csSotDbConf struct {
	Dsn string `yaml:"dsn"`
}
