package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerApi struct {
	dockerClient *client.Client
}

func (d *DockerApi) GetDockerContainer() []types.Container {
	containers, err := d.dockerClient.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	// for _, container := range containers {
	// 	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	// }

	fmt.Println("docker-->")
	return containers
}

func (d *DockerApi) RestartContainerID(id string) {

	opt := container.StopOptions{}

	err := d.dockerClient.ContainerRestart(context.TODO(), id, opt)

	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(id))

}
