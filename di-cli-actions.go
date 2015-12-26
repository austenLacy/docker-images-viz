package main

import(
    "fmt"
    "github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func imagesAction(client *docker.Client, isVerbose bool, shouldTruncateId bool, shouldAccumulate bool) {
    var images *[]Image
    var imgs []Image

	clientImages, err := client.ListImages(docker.ListImagesOptions{All: true})

	if err != nil || len(clientImages) < 1 {
		fmt.Println("No image info available. Double check you have a docker-machine machine running.")
		return
	}

	for _, image := range clientImages {
		imgs = append(imgs, Image{
				image.ID,
				image.ParentID,
				image.RepoTags,
				image.VirtualSize,
				image.Size,
				image.Created,
			})
	}
	images = &imgs

	imagesByParent := collectChildren(images)

	roots := collectRoots(images)

    if isVerbose {
        fmt.Print(printImages(roots, imagesByParent, shouldTruncateId, shouldAccumulate))
    } else {
        *images, imagesByParent = filterOnlyLabeledImages(images, &imagesByParent)
        fmt.Print(printImages(roots, imagesByParent, shouldTruncateId, shouldAccumulate))
    }
}

func containersAction(client *docker.Client, shouldTruncateId bool) {
    clientContainers, err := client.ListContainers(docker.ListContainersOptions{All: true})

    if err != nil || len(clientContainers) < 1 {
        fmt.Println("No containers info available. Double check you have a docker-machine machine running")
        return
    }

    for _, container := range clientContainers {
        printContainer(container, shouldTruncateId)
    }
}

