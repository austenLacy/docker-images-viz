package main

import (
	"fmt"
	"github.com/austenLacy/docker-image-viz/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func main() {
    // use docker-machine
    client, _ := docker.NewClientFromEnv()
    imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})

	if len(imgs) == 0 {
		fmt.Println("No images available")
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
