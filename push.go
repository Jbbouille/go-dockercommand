package dockercommand

import (
	"os"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
)

type PushOptions struct {
	Image string
}

func (dock *Docker) Push(options *PushOptions) error {
	var image string
	var tag string
	if strings.Contains(options.Image, ":") {
		imageSplit := strings.Split(options.Image, ":")
		image = imageSplit[0]
		tag = imageSplit[1]
	} else {
		image = options.Image
		tag = "latest"
	}
	return dock.client.PushImage(docker.PushImageOptions{
		Name:         image,
		Tag:          tag,
		OutputStream: os.Stdout,
	}, docker.AuthConfiguration{})
}
