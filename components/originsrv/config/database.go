package config

type DBConfig struct {
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	SSLMode    string `mapstructure:"ssl-mode"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Migrations string `mapstructure:"migrations"`
}
