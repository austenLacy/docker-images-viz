package main

import (
	"fmt"
	"os"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

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
	app.Action = func(c *cli.Context) {
		fmt.Println("Hello friend!")
	}

	app.Commands = []cli.Command{
		{
			Name:    "images",
			Aliases: []string{"i"},
			Usage:   "view any images running with docker-machine",
			Action: func(c *cli.Context) {
				viewImages(client)
			},
		},
	}

	app.Run(os.Args)
}

func viewImages(client *docker.Client) {
	imgs, err := client.ListImages(docker.ListImagesOptions{All: false})

	if err != nil || len(imgs) < 1 {
		fmt.Println("No image info available. Double check you have a machine running.")
		return
	}

	for _, img := range imgs {
		fmt.Println("Image ID: ", img.ID)
		fmt.Println("Image RepoTags: ", img.RepoTags)
		fmt.Println("Image Created: ", img.Created)
		fmt.Println("Image Size: ", img.Size)

		// IMAGE HISTORY INFO
		//
		// Shows each of the layers that make up this image
		//
		// for more granular info on the image see -->
		//     client.InspectImage(img.ID)
		//
		imgHistory, _ := client.ImageHistory(img.ID)

		for _, imgHist := range imgHistory {
			fmt.Println("Img History ID: ", imgHist.ID)
			fmt.Println("Img History Tags: ", imgHist.Tags)
			fmt.Println("Img History Created: ", imgHist.Created)
			fmt.Println("Img History Created By: ", imgHist.CreatedBy)
			fmt.Println("Img History Size: ", imgHist.Size)
			fmt.Println("")
		}
		fmt.Println("===============================================")
		fmt.Println("")
	}
}
