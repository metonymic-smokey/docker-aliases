package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"strconv"
)

func main() {

	nc := flag.String("n", "", "number of containers")

	flag.Parse()
	n, _ := strconv.Atoi(*nc)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}

	if len(containers) < n {
		n = len(containers)
	}

	fmt.Println("Containers being stopped... ")
	for _, container := range containers[:n] {
		cli.ContainerStop(ctx, container.Names[0][1:], nil)
	}

	fmt.Println("Top ", n, "containers being removed. ")
	for _, container := range containers[:n] {
		cli.ContainerRemove(ctx, container.Names[0][1:], types.ContainerRemoveOptions{})
	}

}
