package main

import (
    "fmt"
    "github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

func imagesAction(client *docker.Client, argImageId string, isVerbose bool, shouldTruncateId bool, shouldAccumulate bool) {
    var images *[]Image
    var imgs []Image
    var clientImages []docker.APIImages
    var clientImageHistory []docker.ImageHistory
    var err1 error = nil
    var err2 error = nil
    var roots []Image

    if argImageId != "" {
        clientImages, err1 = client.ListImages(docker.ListImagesOptions{All: false, Filter: argImageId})

        /*
           We use the ImageHistory here to get the parent images starting with
           the image ID passed in and returned by ListImages(...) back to its root.
        */
        clientImageHistory, err2 = client.ImageHistory(argImageId)
    } else {
        clientImages, err1 = client.ListImages(docker.ListImagesOptions{All: true})
    }

    if err1 != nil || err2 != nil || len(clientImages) < 1 {
        fmt.Println("No image info available. Double check you have a docker-machine machine running or if an image name was provided, that it is valid.")
        return
    }

    if len(clientImageHistory) < 1 {
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
    } else {
        for _, imageHist := range clientImageHistory {
            var repoTags string
            img, err := client.InspectImage(imageHist.ID)

            if err != nil {
                fmt.Println("Problem inspecting image with ID: ", imageHist.ID)
                return
            }

            if img.ID == clientImages[0].ID {
                repoTags = clientImages[0].RepoTags[0]
            } else {
                repoTags = "<none>:<none>"
            }

            imgs = append(imgs, Image{
                img.ID,
                img.Parent,
                []string{repoTags},
                img.VirtualSize,
                img.Size,
                img.Created.Unix(),
            })
        }
    }

    images = &imgs

    imagesByParent := collectChildren(images)

    roots = collectRoots(images)

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
