package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

func main() {

	img_name := flag.String("name", "", "name of the image")

	img_id := flag.String("ID", "", "Image ID of the image")

	action := flag.String("op", "stop", "Stop or remove containers")
	flag.Parse()

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Containers being stopped: ")
	for _, container := range containers {
		if container.Image == *img_name || container.ImageID[7:19] == *img_id {
			fmt.Println(container.Image, " ", container.ImageID[7:19])
			cli.ContainerStop(ctx, container.Names[0][1:], nil)
		}

	}

	if *action == "remove" {
		fmt.Println("Containers being removed: ")
		for _, container := range containers {
			if container.Image == *img_name || container.ImageID[7:19] == *img_id {
				fmt.Println(container.Image, " ", container.ImageID[7:19])
				cli.ContainerRemove(ctx, container.Names[0][1:], types.ContainerRemoveOptions{})
			}

		}

	}
}
