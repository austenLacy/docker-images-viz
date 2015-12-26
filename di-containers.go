package main

import(
    "fmt"
    "github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

type Container struct {
	Id         string                   `json: "id"`
	Image      string                   `json: "image"`
	Names      []string                 `json: "names"`
	Ports      []map[string]interface{} `json: "ports"`
	Created    int64                    `json: "created"`
	Status     string                   `json: "status"`
	Command    string                   `json: "command"`
    SizeRw     int64                    `json: "sizeRw"`
    SizeRootFs int64                    `json: "sizeRootFs"`
}


/********************************************************************************
 * PRINT CONTAINERS FUNCTIONS
 ********************************************************************************/

func printContainer(clientContainer docker.APIContainers, shouldTruncateId bool) {
    // TODO: use the `go-dockerclient Container type` here to show
    //       more info on the container than the basic `docker.APIContainers`
    //       See: client.InspectContainer(container.ID)
    container := Container{
        clientContainer.ID,
		clientContainer.Image,
		clientContainer.Names,
		apiPortToMap(clientContainer.Ports),
		clientContainer.Created,
		clientContainer.Status,
		clientContainer.Command,
        clientContainer.SizeRw,
        clientContainer.SizeRootFs,
    }

    fmt.Println("----------------------------------------------------------------------------------")
    if shouldTruncateId {
        fmt.Println("ID: ", truncateId(container.Id))
    } else {
        fmt.Println("ID: ", container.Id)
    }
    fmt.Println("Image: ", container.Image)
    fmt.Println("Names: ", container.Names)
    fmt.Println("Ports: \n")
    for _, port := range container.Ports {
        if len(port) < 1 {
            continue
        }
        fmt.Println("├───── IP: ", port["IP"])
        fmt.Println("├───── Type: ", port["Type"])
        fmt.Println("├───── PrivatePort: ", port["PrivatePort"])
        fmt.Println("├───── PublicPort: ", port["PublicPort"])
        fmt.Println("")
    }
    fmt.Println("Created: ", container.Created)
    fmt.Println("Status: ", container.Status)
    fmt.Println("Command: ", container.Command)
    fmt.Println("SizeRw: ", container.SizeRw)
    fmt.Println("SizeRootFs: ", container.SizeRootFs)
    fmt.Println("----------------------------------------------------------------------------------\n")
}

