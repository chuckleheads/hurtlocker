package config

type RabbitMQConfig struct {
	Host     string   `mapstructure:"host"`
	Port     int      `mapstructure:"port"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	Exchange string   `mapstructure:"exchange"`
	Queue    string   `mapstructure:"queue"`
	Topic    []string `mapstructure:"topic"`
}
