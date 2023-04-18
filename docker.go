package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/swarm"
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

func (d *DockerApi) GetSwarmNode() []swarm.Node {

	res, err := d.dockerClient.NodeList(context.Background(), types.NodeListOptions{})

	if err != nil {
		fmt.Println(err)
	}
	// log.Println(res)
	return res

}
func (d *DockerApi) GetDockerServices() []swarm.Service {

	res, err := d.dockerClient.ServiceList(context.Background(), types.ServiceListOptions{})

	if err != nil {
		fmt.Println(err)
	}
	// log.Println(res)
	return res

}

type Msg struct {
	PreviousSpec []byte
}

func (d *DockerApi) DockerServicesUpdate(id string) {

	// inspect, _, err := d.dockerClient.ServiceInspectWithRaw(context.Background(), id, types.ServiceInspectOptions{})
	// fmt.Println(inspect.ID)
	// fmt.Println("------")

	// // obj := Msg{f}
	// // json, _ := json.Marshal(obj)
	// // fmt.Println(string(json))
	// // myString := string(f[:])
	// // fmt.Println(myString)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // s := inspect.PreviousSpec
	// // fmt.Println(*s)

	// res, err := d.dockerClient.ServiceUpdate(context.Background(), inspect.ID, swarm.Version{Index: inspect.Meta.Version.Index}, inspect.Spec, types.ServiceUpdateOptions{RegistryAuthFrom: "previous-spec"})

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("------")

	d.dockerClient.ServiceRemove(context.Background(), id)
	// log.Println(res.Warnings)

}
