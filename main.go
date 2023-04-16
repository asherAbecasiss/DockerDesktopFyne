package main

import (
	"github.com/docker/docker/client"
	//"fyne.io/fyne/v2/theme"
)

var data = []string{"aefewewwfwfwfwefwefwefwefwefwfewefwfwfwfwfewfwfwfwefwfwadfasdfweef", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list", "a", "string", "list"}

func NewApp(client *client.Client) *DockerApi {
	return &DockerApi{
		dockerClient: client,
	}
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	docker := NewApp(cli)
	docker.Run()

}
