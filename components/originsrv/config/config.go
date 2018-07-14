package config

type Config struct {
	Port      int `mapstructure:"port"`
	HttpPort  int `mapstructure:"http_port"`
	Datastore DBConfig
}
