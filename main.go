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
	// use docker-machine
	client, err := docker.NewClientFromEnv()

	if err != nil {
		fmt.Println("Unable to get docker client from docker-machine env")
	}

	app := cli.NewApp()
	app.Name = "docker-inspect"
	app.Usage = "get some stats on your docker images, containers, and env"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "verbose",
			Usage: "if used with images it shows all (labeled and unlabeled) images, if with containers shows running and non-running. true by default",
			Destination: &isVerbose,
		},
		cli.StringFlag{
			Name: "truncate-id",
			Usage: "if true truncates the image/container ids to just the first 12 characters, if false then show entire id. true by default",
			Destination: &shouldTruncateId,
		},
		cli.StringFlag{
			Name: "should-accumulate",
			Usage: "if true accumulates the each image's size in tree view, if false then it shows each image's individual size, false by default",
			Destination: &shouldAccumulate,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "images",
			Aliases: []string{"i"},
			Usage:   "view any images running with docker-machine",
			Action: func(c *cli.Context) {
				shouldTruncate := true
				verbose := true
				accumulate := false

				if shouldTruncateId == "false" {
					shouldTruncate = false
				}

				if isVerbose == "false" {
					verbose = false
				}

				if shouldAccumulate == "true" {
					accumulate = true
				}

				viewImagesAction(client, verbose, shouldTruncate, accumulate)
			},
		},
	}

	app.Run(os.Args)
}