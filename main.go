package main

import (
	"fmt"
	"os"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

var isVerbose string
var shouldTruncateId string
var shouldAccumulate string

func main() {
	// defaults
	verbose := true
	truncate := true
	accumulate := false

	// TODO: support more env's than just docker-machine
	// use docker-machine
	client, err := docker.NewClientFromEnv()

	if err != nil {
		fmt.Println("Unable to get docker client from docker-machine env")
	}

	app := cli.NewApp()
	app.Name = "docker-inspect"
	app.Usage = "get some info on any docker images, containers, and env"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name: "images",
			Aliases: []string{"i"},
			Usage: "view any docker images",
			Action: func(c *cli.Context) {
				// check if user has provided image ID to build tree for
				var argImageId string

				if len(c.Args()) > 0 {
					argImageId = c.Args()[0]
				}

				if shouldTruncateId == "false" {
					truncate = false
				}

				if isVerbose == "false" {
					verbose = false
				}

				if shouldAccumulate == "true" {
					accumulate = true
				}

				imagesAction(client, argImageId, verbose, truncate, accumulate)
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "verbose, verb",
					Usage: "if true show all (labeled and unlabeled) images, if false show only labeled images, true by default",
					Destination: &isVerbose,
				},
				cli.StringFlag{
					Name: "truncate-id, ti",
					Usage: "if true truncates the image id to just the first 12 characters, if false then shows entire id. true by default",
					Destination: &shouldTruncateId,
				},
				cli.StringFlag{
					Name: "accumulate, acc",
					Usage: "if true accumulates the each image's size in tree view, if false then it shows each image's individual size, false by default",
					Destination: &shouldAccumulate,
				},
			},
		},
		{
			Name: "containers",
			Aliases: []string{"c"},
			Usage: "view any docker containers running with docker-machine",
			Action: func(c *cli.Context) {
				if shouldTruncateId == "false" {
					truncate = false
				}

				containersAction(client, truncate);
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "truncate-id, ti",
					Usage: "if true truncates the container id to just the first 12 characters, if false then shows entire id. true by default",
					Destination: &shouldTruncateId,
				},
			},
		},
	}

	app.Run(os.Args)
}