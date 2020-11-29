package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"os"
	"regexp"
)

func main() {

	x := os.Args[1:][0]

	re, _ := regexp.Compile(x)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Containers getting removed: ")
	for _, container := range containers {

		name := container.Names[0][1:]

		matched := re.MatchString(name)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			fmt.Println(name)

			//stop container with name matching regex
			err := cli.ContainerStop(ctx, name, nil)
			if err != nil {
				log.Fatal(err)
			}

			//remove containers after stopping
			err = cli.ContainerRemove(ctx, name, types.ContainerRemoveOptions{})
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}
