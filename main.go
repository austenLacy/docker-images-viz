package main

import (
	"fmt"
	"github.com/austenLacy/docker-image-viz/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func main() {
    // endpoint := "unix:///var/run/docker.sock"
    
    // use docker-machine
    client, _ := docker.NewClientFromEnv()
    imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
    for _, img := range imgs {
        fmt.Println("ID: ", img.ID)
        fmt.Println("RepoTags: ", img.RepoTags)
        fmt.Println("Created: ", img.Created)
        fmt.Println("Size: ", img.Size)
        fmt.Println("VirtualSize: ", img.VirtualSize)
        fmt.Println("ParentId: ", img.ParentID)
    }
}
