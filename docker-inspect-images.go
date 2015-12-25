package main

import(
    "fmt"
	"bytes"
    "strings"
)

type Image struct {
	Id          string
	ParentId    string
	RepoTags    []string
	VirtualSize int64
	Size        int64
	Created     int64
}

/********************************************************************************
 * PRINT FUNCTIONS
 ********************************************************************************/
func printImages(images []Image, byParent map[string][]Image, noTrunc bool, incremental bool) string {
	var buffer bytes.Buffer
	initialPrefix := ""

	printAsTree(&buffer, images, byParent, noTrunc, incremental, initialPrefix)

	return buffer.String()
}


func printAsTree(buffer *bytes.Buffer, images []Image, byParent map[string][]Image, shouldTruncate bool, incremental bool, prefix string) {
	var length = len(images)
	if length > 1 {
		for idx, image := range images {
			var nextPrefix string = ""
			if (idx + 1) == length {
				printTreeNode(buffer, image, shouldTruncate, incremental, prefix + "└─")
				nextPrefix = "  "
			} else {
				printTreeNode(buffer, image, shouldTruncate, incremental, prefix + "├─")
				nextPrefix = "│ "
			}
			if subimages, exists := byParent[image.Id]; exists {
				printAsTree(buffer, subimages, byParent, shouldTruncate, incremental, prefix + nextPrefix)
			}
		}
	} else {
		for _, image := range images {
			printTreeNode(buffer, image, shouldTruncate, incremental, prefix + "└─")
			if subimages, exists := byParent[image.Id]; exists {
				printAsTree(buffer, subimages, byParent, shouldTruncate, incremental, prefix + "  ")
			}
		}
	}
}

func printTreeNode(buffer *bytes.Buffer, image Image, shouldTruncate bool, incremental bool, prefix string) {
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

	buffer.WriteString(fmt.Sprintf("%s(%s) -- Virtual Size: %s", prefix, imageID, convertToHumanReadableSize(size)))

	if image.RepoTags[0] != "<none>:<none>" {
		buffer.WriteString(fmt.Sprintf(" Tags: %s\n", strings.Join(image.RepoTags, ", ")))
	} else {
		buffer.WriteString(fmt.Sprintf("\n"))
	}
}


/********************************************************************************
 * UTILITY FUNCTIONS
 ********************************************************************************/

func collectRoots(images *[]Image) []Image {
	var roots []Image
	for _, image := range *images {
		if image.ParentId == "" {
			roots = append(roots, image)
		}
	}

	return roots
}

func filterOnlyLabeledImages(images *[]Image, byParent *map[string][]Image) (filteredImages []Image, filteredChildren map[string][]Image) {
	for i := 0; i < len(*images); i++ {
		// image should visible
		//   1. it has a label
		//   2. it is root
		//   3. it is a node
		visible := ((*images)[i].RepoTags[0] != "<none>:<none>") ||
                   ((*images)[i].ParentId == "") ||
                   (len((*byParent)[(*images)[i].Id]) > 1)
		if visible {
			filteredImages = append(filteredImages, (*images)[i])
		} else {
			// change childs parent id
			// if items are filtered with only one child
			for j := 0; j < len(filteredImages); j++ {
				if (filteredImages[j].ParentId == (*images)[i].Id) {
					filteredImages[j].ParentId = (*images)[i].ParentId
				}
			}
			for j := 0; j < len(*images); j++ {
				if ((*images)[j].ParentId == (*images)[i].Id) {
					(*images)[j].ParentId = (*images)[i].ParentId
				}
			}
		}
	}

	filteredChildren = collectChildren(&filteredImages)

	return filteredImages, filteredChildren
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