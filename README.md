# docker-inspect

A small Go app that displays useful info about any local docker images, containers, and env.

A lot of influence from the [dockviz](https://github.com/justone/dockviz) project.

Currently uses:
* [go-dockerclient](https://github.com/fsouza/go-dockerclient)
* [cli.go](https://github.com/codegangsta/cli)

### TODO:
* <s>Make the tree view look better. Similar to the now deprecated `docker images --tree` output.</s>
* <s>Break up code so it's not all in `main.go`</s>
* Add the ability to visualize any containers
* <s>Add cli flags for truncating the image ID and showing cumulative image size vs individual image size</s>
* <s>Add subcommand/flag to show only labeled images as output (less verbose)</s>
* Any UI enhancements?
* COMMENTS COMMENTS COMMENTS
* TESTS TESTS TESTS
