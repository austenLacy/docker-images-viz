package main

import (
	"fmt"
	"bytes"
	"os"
	"strings"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

type Image struct {
	Id          string   `json: "id,omitempty"`
	ParentId    string   `json: "parentId,omitempty"`
	RepoTags    []string `json: "repoTags,omitempty"`
	VirtualSize int64    `json: "virtualSize,omitempty"`
	Size        int64    `json: "size,omitempty"`
	Created     int64    `json: "created,omitempty"`
}

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
		defaultAction()
	}

	app.Commands = []cli.Command{
		{
			Name:    "images",
			Aliases: []string{"i"},
			Usage:   "view any images running with docker-machine",
			Action: func(c *cli.Context) {
				viewImagesAction(client)
			},
		},
	}

	app.Run(os.Args)
}

func defaultAction() {
	fmt.Println("Hello, friend!")
}

func viewImagesAction(client *docker.Client) {
	var images *[]Image
	var imgs []Image

	clientImages, err := client.ListImages(docker.ListImagesOptions{All: true})

	if err != nil || len(clientImages) < 1 {
		fmt.Println("No image info available. Double check you have a machine running.")
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

	////////////////////////////////////////////////////////
	// TODO: probably want to make these cli flags
	shouldTruncate := true
	shouldShowCumulativeSize := false
	////////////////////////////////////////////////////////

	fmt.Print(printImages(roots, imagesByParent, shouldTruncate, shouldShowCumulativeSize))
}

func printImages(images []Image, byParent map[string][]Image, noTrunc bool, incremental bool) string {
	var buffer bytes.Buffer
	defaultPrefix := ""

	printAsTree(&buffer, images, byParent, noTrunc, incremental, defaultPrefix)

	return buffer.String()
}


func collectRoots(images *[]Image) []Image {
	var roots []Image
	for _, image := range *images {
		if image.ParentId == "" {
			roots = append(roots, image)
		}
	}

	return roots
}

func collectChildren(images *[]Image) map[string][]Image {
	var imagesByParent = make(map[string][]Image)
	for _, image := range *images {
		if children, exists := imagesByParent[image.ParentId]; exists {
			imagesByParent[image.ParentId] = append(children, image)
		} else {
			imagesByParent[image.ParentId] = []Image{image}
		}
	}

	return imagesByParent
}

func printAsTree(buffer *bytes.Buffer, images []Image, byParent map[string][]Image, shouldTruncate bool, incremental bool, prefix string) {
	var length = len(images)
	if length > 1 {
		for idx, image := range images {
			var nextPrefix string = ""
			if (idx + 1) == length {
				PrintTreeNode(buffer, image, shouldTruncate, incremental, prefix + "└─")
				nextPrefix = "  "
			} else {
				PrintTreeNode(buffer, image, shouldTruncate, incremental, prefix + "├─")
				nextPrefix = "│ "
			}
			if subimages, exists := byParent[image.Id]; exists {
				printAsTree(buffer, subimages, byParent, shouldTruncate, incremental, prefix + nextPrefix)
			}
		}
	} else {
		for _, image := range images {
			PrintTreeNode(buffer, image, shouldTruncate, incremental, prefix + "└─")
			if subimages, exists := byParent[image.Id]; exists {
				printAsTree(buffer, subimages, byParent, shouldTruncate, incremental, prefix + "  ")
			}
		}
	}
}

func PrintTreeNode(buffer *bytes.Buffer, image Image, shouldTruncate bool, incremental bool, prefix string) {
	var imageID string
	if shouldTruncate {
		imageID = truncateId(image.Id)
	} else {
		imageID = image.Id
	}

	var size int64
	if incremental {
		size = image.VirtualSize
	} else {
		size = image.Size
	}

	buffer.WriteString(fmt.Sprintf("%s%s Virtual Size: %s", prefix, imageID, convertToHumanReadableSize(size)))

	if image.RepoTags[0] != "<none>:<none>" {
		buffer.WriteString(fmt.Sprintf(" Tags: %s\n", strings.Join(image.RepoTags, ", ")))
	} else {
		buffer.WriteString(fmt.Sprintf("\n"))
	}
}

func convertToHumanReadableSize(raw int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB"}

	rawFloat := float64(raw)
	ind := 0

	for {
		if rawFloat < 1000 {
			break
		} else {
			rawFloat = rawFloat / 1000
			ind = ind + 1
		}
	}

	return fmt.Sprintf("%.01f %s", rawFloat, sizes[ind])
}

func truncateId(id string) string {
	return id[0:12]
}
