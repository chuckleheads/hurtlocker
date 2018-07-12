package drivers

import (
	"github.com/docker/docker/client"
)

type Driver interface {
	Pull()
	Create([]string) string
	Start(string)
}

func New() Driver {
	// TED: uncomment this when we have more drivers
	// if viper.GetString("driver") == "docker" {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return Docker{
		Client: cli,
		Image:  "habitat/default-studio:0.58.0",
	}
}
