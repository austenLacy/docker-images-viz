package main

import(
    "fmt"
    "github.com/austenLacy/docker-inspect/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

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

func apiPortToMap(ports []docker.APIPort) []map[string]interface{} {
	result := make([]map[string]interface{}, 2)
	for _, port := range ports {
		intPort := map[string]interface{}{
	        "IP":          port.IP,
			"Type":        port.Type,
			"PrivatePort": port.PrivatePort,
			"PublicPort":  port.PublicPort,
		}
		result = append(result, intPort)
	}
	return result
}