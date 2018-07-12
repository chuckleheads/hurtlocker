package drivers

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Docker struct {
	Client *client.Client
	Image  string
}

func (d Docker) Pull() {
	out, err := d.Client.ImagePull(context.Background(), d.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

func (d Docker) Create(cmd []string) string {
	resp, err := d.Client.ContainerCreate(context.Background(), &container.Config{
		Image: d.Image,
		Cmd:   cmd,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	return resp.ID
}

func (d Docker) Start(id string) {
	if err := d.Client.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
