# docker-inspect

A small Go app that displays useful info about any local docker images, containers, and env.

Currently uses:
* [go-dockerclient](https://github.com/fsouza/go-dockerclient)
* [cli.go](https://github.com/codegangsta/cli)


### TODO:
* <s>Make the tree view look better. Similar to the now deprecated `docker images --tree` output.</s>
* Add cli flags for truncating the image ID and showing cumulative image size vs individual image size
* Add subcommand/flag to show only labeled images as output (less verbose)