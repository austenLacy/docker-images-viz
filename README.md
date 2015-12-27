# docker-inspect

A small Go app that displays useful info about any local docker images, containers, and env.

A lot of the motivation for the images tree view is because of the deprecation of the `docker images --tree` command and because I wanted to play around with the [go-dockerclient](https://github.com/fsouza/go-dockerclient)

## Steps to install

1. Make sure [go](https://golang.org/) is installed on your machine
2. Download `docker-inspect` source
3. Navigate to `docker-inspect` directory and run `go install` to put `docker-inspect` binary into your `$GOPATH/bin`
4. I suggest putting your `$GOPATH/bin` into your `$PATH` with `export PATH=$PATH:$GOPATH/bin`
5. If you followed step 4 then the command `docker-inspect` is now in your path and can be run anywhere with `docker-inspect`

>Note this currently only works with a `docker-machine` environment. See [docker-machine](https://docs.docker.com/machine/) for more info.

## Usage

### Commands

#### docker-inspect help

```no-highlight
# can also do `docker-inspect`, `docker-inspect -h`, or `docker-inspect help`
$ docker-inspect --help
NAME:
   docker-inspect - get some info on any docker images, containers, and env

USAGE:
   docker-inspect [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   images, i		view any docker images
   containers, c	view any docker containers running with docker-machine
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

#### docker-inspect --version

```no-highlight
# can also do `docker-inspect -v` or `docker-inspect --v`
$ docker-inspect --version
docker-inspect version 0.0.1
```

#### docker-inspect images

```no-highlight
$ docker-inspect images --help
NAME:
   docker-inspect images - view any docker images

USAGE:
   docker-inspect images [command options] [arguments...]

OPTIONS:
   --verbose, --verb 	if true show all (labeled and unlabeled) images, if false show only labeled images, true by default
   --truncate-id, --ti 	if true truncates the image id to just the first 12 characters, if false then shows entire id. true by default
   --accumulate, --acc 	if true accumulates the each image's size in tree view, if false then it shows each image's individual size, false by default
```

```no-highlight
$ docker-inspect images
├─ 9ee13ca3b908 -- Virtual Size: 125.1 MB
│ └─ 23cb15b0fcec -- Virtual Size: 0.0 B
│   └─ 5e5f21412e19 -- Virtual Size: 44.3 MB
│     └─ df82ac64861d -- Virtual Size: 122.2 MB
│       └─ 6a84d4eff4c4 -- Virtual Size: 134.0 MB
│         └─ f204abcb3569 -- Virtual Size: 0.0 B
│           └─ f2d7651c6d8a -- Virtual Size: 0.0 B
│             └─ 50386c55167d -- Virtual Size: 0.0 B
│               └─ feb3148492d4 -- Virtual Size: 278.2 MB
│                 └─ 6a6f1b05ca25 -- Virtual Size: 0.0 B
│                   └─ a81e9c53fc51 -- Virtual Size: 0.0 B
│                     └─ 58cc81e0c7ad -- Virtual Size: 0.0 B
│                       └─ 93f4c5ebe1a3 -- Virtual Size: 0.0 B
│                         └─ 08ff0c215f8f -- Virtual Size: 2.5 KB Tags: golang:latest
├─ a719479f5894 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d -- Virtual Size: 0.0 B
│   ├─ 3df5aff384fc -- Virtual Size: 0.0 B
│   │ └─ 4f3dc531a45a -- Virtual Size: 2.0 KB
│   │   └─ 687dd94c3fd3 -- Virtual Size: 221.0 B
│   │     └─ dd6bbcfbe827 -- Virtual Size: 0.0 B
│   │       └─ 5ba2077eefe2 -- Virtual Size: 7.7 MB
│   │         └─ f61f6b8f8a52 -- Virtual Size: 11.0 B
│   │           └─ 7eccb4b78817 -- Virtual Size: 11.0 B
│   │             └─ 22e963aa9f34 -- Virtual Size: 0.0 B
│   │               └─ ca1f5f48ef43 -- Virtual Size: 0.0 B
│   │                 └─ 8d5e6665a7a6 -- Virtual Size: 0.0 B Tags: nginx:latest
```

```no-highlight
$ docker-inspect images --verbose=false
├─ 9ee13ca3b908 -- Virtual Size: 125.1 MB
│ └─ 08ff0c215f8f -- Virtual Size: 2.5 KB Tags: golang:latest
├─ a719479f5894 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d -- Virtual Size: 0.0 B
│   ├─ 8d5e6665a7a6 -- Virtual Size: 0.0 B Tags: nginx:latest
```

```no-highlight
$ docker-inspect images --truncate-id=false
├─ 9ee13ca3b908aacceeb9eb6a3028d4566aa997ffc90915d55578a487f058e935 -- Virtual Size: 125.1 MB
│ └─ 23cb15b0fcece0623d706e9dbc3d9fce97937fb50ab3fbff8574206d258b6303 -- Virtual Size: 0.0 B
│   └─ 5e5f21412e197751f63ce12e40459cba6cdc2a1330c63f106f19a7292599985c -- Virtual Size: 44.3 MB
│     └─ df82ac64861d951d10f479113098db9e78f9778cc83bc5d87fb3cd04424045d2 -- Virtual Size: 122.2 MB
│       └─ 6a84d4eff4c4e1214f79b7b7cd3e006f36dab3a66067a67dc43fe120330f85b0 -- Virtual Size: 134.0 MB
│         └─ f204abcb356959ffaaac4d2b5c9e6000053d768cccab77f643f4c5d87f76f3a9 -- Virtual Size: 0.0 B
│           └─ f2d7651c6d8a15eef8f08c7874fa2fa37e3a7ef800b04b417ba8708eaa608ccf -- Virtual Size: 0.0 B
│             └─ 50386c55167d14312d0ed81cf9f0fa2588461be098b2cd360c15fe3f120e2f75 -- Virtual Size: 0.0 B
│               └─ feb3148492d4bd41caabfe87ba6733c07bf060f573a084b32c3acd59b89ab1c9 -- Virtual Size: 278.2 MB
│                 └─ 6a6f1b05ca254c23cd3b573180fe3c6d33c05281a1f3733bb35fffc6a6d36ecd -- Virtual Size: 0.0 B
│                   └─ a81e9c53fc519488bf6227529bb006cd1aa1ff1ccf2d4e0f4e06e5395e54ec44 -- Virtual Size: 0.0 B
│                     └─ 58cc81e0c7ad5e2ec98b17eb5357a18e63a9092c0fa6db07a95aea45629793bb -- Virtual Size: 0.0 B
│                       └─ 93f4c5ebe1a31b0b98e3adaaee51888f17fee08ab4b3e009143cb970217d562e -- Virtual Size: 0.0 B
│                         └─ 08ff0c215f8fec13c49ffe7034254f8c617c25ca5be8907a66f38f5401bc109a -- Virtual Size: 2.5 KB Tags: golang:latest
├─ a719479f5894e94befa7b0a678f52b0e65c4cfa055eb14c1d219d2b6d3acf574 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d483318e92036e26574e0c329d0d52299fe47462c12c5e554eb67 -- Virtual Size: 0.0 B
│   ├─ 3df5aff384fc7c76136fce7548e315bee24dac2cee42678f5b30d168e1c927a3 -- Virtual Size: 0.0 B
│   │ └─ 4f3dc531a45ab2f90e542340293e706a51ddfabae923f460b170fe42fd5a7d48 -- Virtual Size: 2.0 KB
│   │   └─ 687dd94c3fd349e8f1e16acdee0c5122317f4d931c3248ccdf5073926a7744fa -- Virtual Size: 221.0 B
│   │     └─ dd6bbcfbe827ea18495d53a2ec1d72be77be42a3a7b56111aa4d3c0164bbe313 -- Virtual Size: 0.0 B
│   │       └─ 5ba2077eefe21b926b35f83d6cd473366a3f5872781aa160ad281b6c7351da98 -- Virtual Size: 7.7 MB
│   │         └─ f61f6b8f8a5280508962d8417414b313463b0f12eeb4501839361aa5bbcecf02 -- Virtual Size: 11.0 B
│   │           └─ 7eccb4b788170aebbcb2ee4c28abee9fa10580e483747866164130a10ec07151 -- Virtual Size: 11.0 B
│   │             └─ 22e963aa9f34b18dfd5108ea14676ba41103e6f95aa005e298f0d4ecbf62b5be -- Virtual Size: 0.0 B
│   │               └─ ca1f5f48ef43d72726dda945ff6ade7b9c1c12ce6329a7be7a3c1a23f8703c97 -- Virtual Size: 0.0 B
│   │                 └─ 8d5e6665a7a6e3e38929d737206f6e4bf20574bfe696d1bc30bf572034bf81de -- Virtual Size: 0.0 B Tags: nginx:latest
```

```no-highlight
$ docker-inspect images --accumulate=true
├─ 9ee13ca3b908 -- Virtual Size: 125.1 MB
│ └─ 23cb15b0fcec -- Virtual Size: 125.1 MB
│   └─ 5e5f21412e19 -- Virtual Size: 169.4 MB
│     └─ df82ac64861d -- Virtual Size: 291.7 MB
│       └─ 6a84d4eff4c4 -- Virtual Size: 425.7 MB
│         └─ f204abcb3569 -- Virtual Size: 425.7 MB
│           └─ f2d7651c6d8a -- Virtual Size: 425.7 MB
│             └─ 50386c55167d -- Virtual Size: 425.7 MB
│               └─ feb3148492d4 -- Virtual Size: 703.8 MB
│                 └─ 6a6f1b05ca25 -- Virtual Size: 703.8 MB
│                   └─ a81e9c53fc51 -- Virtual Size: 703.8 MB
│                     └─ 58cc81e0c7ad -- Virtual Size: 703.8 MB
│                       └─ 93f4c5ebe1a3 -- Virtual Size: 703.8 MB
│                         └─ 08ff0c215f8f -- Virtual Size: 703.8 MB Tags: golang:latest
├─ a719479f5894 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d -- Virtual Size: 125.1 MB
│   ├─ 3df5aff384fc -- Virtual Size: 125.1 MB
│   │ └─ 4f3dc531a45a -- Virtual Size: 125.1 MB
│   │   └─ 687dd94c3fd3 -- Virtual Size: 125.1 MB
│   │     └─ dd6bbcfbe827 -- Virtual Size: 125.1 MB
│   │       └─ 5ba2077eefe2 -- Virtual Size: 132.8 MB
│   │         └─ f61f6b8f8a52 -- Virtual Size: 132.8 MB
│   │           └─ 7eccb4b78817 -- Virtual Size: 132.8 MB
│   │             └─ 22e963aa9f34 -- Virtual Size: 132.8 MB
│   │               └─ ca1f5f48ef43 -- Virtual Size: 132.8 MB
│   │                 └─ 8d5e6665a7a6 -- Virtual Size: 132.8 MB Tags: nginx:latest
```

```no-highlight
$ docker-inspect images --accumulate=true --verbose=false
├─ 9ee13ca3b908 -- Virtual Size: 125.1 MB
│ └─ 08ff0c215f8f -- Virtual Size: 703.8 MB Tags: golang:latest
├─ a719479f5894 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d -- Virtual Size: 125.1 MB
│   ├─ 8d5e6665a7a6 -- Virtual Size: 132.8 MB Tags: nginx:latest
```

```no-highlight
$ docker-inspect images --accumulate=true --verbose=false --truncate-id=false
├─ 9ee13ca3b908aacceeb9eb6a3028d4566aa997ffc90915d55578a487f058e935 -- Virtual Size: 125.1 MB
│ └─ 08ff0c215f8fec13c49ffe7034254f8c617c25ca5be8907a66f38f5401bc109a -- Virtual Size: 703.8 MB Tags: golang:latest
├─ a719479f5894e94befa7b0a678f52b0e65c4cfa055eb14c1d219d2b6d3acf574 -- Virtual Size: 125.1 MB
│ └─ 91bac885982d483318e92036e26574e0c329d0d52299fe47462c12c5e554eb67 -- Virtual Size: 125.1 MB
│   ├─ 8d5e6665a7a6e3e38929d737206f6e4bf20574bfe696d1bc30bf572034bf81de -- Virtual Size: 132.8 MB Tags: nginx:latest
```

```no-highlight
# view tree for a single image by its name. All previous cli flags work the same
$ docker-inspect images nginx
└─ a719479f5894 -- Virtual Size: 125.1 MB
  └─ 91bac885982d -- Virtual Size: 0.0 B
    └─ 3df5aff384fc -- Virtual Size: 0.0 B
      └─ 4f3dc531a45a -- Virtual Size: 2.0 KB
        └─ 687dd94c3fd3 -- Virtual Size: 221.0 B
          └─ dd6bbcfbe827 -- Virtual Size: 0.0 B
            └─ 5ba2077eefe2 -- Virtual Size: 7.7 MB
              └─ f61f6b8f8a52 -- Virtual Size: 11.0 B
                └─ 7eccb4b78817 -- Virtual Size: 11.0 B
                  └─ 22e963aa9f34 -- Virtual Size: 0.0 B
                    └─ ca1f5f48ef43 -- Virtual Size: 0.0 B
                      └─ 8d5e6665a7a6 -- Virtual Size: 0.0 B Tags: nginx:latest
```

#### docker-inspect containers

```no-highlight
$ docker-inspect containers
----------------------------------------------------------------------------------
ID:  7fb2cfbbd6dd
Image:  $IMAGE_NAME_HERE
Names:  [$CONTAINER_NAME_HERE]
Ports:

├───── IP:
├───── Type:  tcp
├───── PrivatePort:  443
├───── PublicPort:  0

├───── IP:  0.0.0.0
├───── Type:  tcp
├───── PrivatePort:  80
├───── PublicPort:  32887

Created:  1448893527
Status:  Up 22 hours
Command:  /usr/local/bin/run.sh
SizeRw:  0
SizeRootFs:  0
----------------------------------------------------------------------------------

----------------------------------------------------------------------------------
ID:  3810a29369ee
Image:  $IMAGE_NAME_HERE
Names:  [$CONTAINER_NAME_HERE]
Ports:

Created:  1448893496
Status:  Exited (137) 22 hours ago
Command:  /usr/local/bin/start.sh
SizeRw:  0
SizeRootFs:  0
----------------------------------------------------------------------------------
```

```no-highlight
$ docker-inspect containers --truncate-id=false
----------------------------------------------------------------------------------
ID:  7fb2cfbbd6dd1d69ee8002e1d746285bb4144ed58d44b1315603835459bfeb3f
Image:  $IMAGE_NAME_HERE
Names:  [$CONTAINER_NAME_HERE]
Ports:

├───── IP:
├───── Type:  tcp
├───── PrivatePort:  443
├───── PublicPort:  0

├───── IP:  0.0.0.0
├───── Type:  tcp
├───── PrivatePort:  80
├───── PublicPort:  32887

Created:  1448893527
Status:  Up 22 hours
Command:  /usr/local/bin/run.sh
SizeRw:  0
SizeRootFs:  0
----------------------------------------------------------------------------------

----------------------------------------------------------------------------------
ID:  3810a29369eee95ef1ea0aa02b9be66c818731c67e068278180fc46fe5bb716d
Image:  $IMAGE_NAME_HERE
Names:  [$CONTAINER_NAME_HERE]
Ports:

Created:  1448893496
Status:  Exited (137) 22 hours ago
Command:  /usr/local/bin/start.sh
SizeRw:  0
SizeRootFs:  0
----------------------------------------------------------------------------------
```

A lot of influence comes from the [dockviz](https://github.com/justone/dockviz) project.

Currently uses:
* [go-dockerclient](https://github.com/fsouza/go-dockerclient)
* [cli.go](https://github.com/codegangsta/cli)

### TODO:
* <s>Make the tree view look better. Similar to the now deprecated `docker images --tree` output.</s>
* <s>Break up code so it's not all in `main.go`</s>
* <s>Add ability to show images tree by image name</s>
* <s>Add ability to visualize any containers en masse (similar to running `docker ps -a`)</s>
* Add stats for containers (`docker stats $CONTAINER_ID`)
* Make it work for more envs than just `docker-machine`
* <s>Add cli flags for truncating the image ID and showing cumulative image size vs individual image size</s>
* <s>Add cli flag to show only labeled images as output (less verbose)</s>
* Any UI enhancements?
* COMMENTS COMMENTS COMMENTS
* TESTS TESTS TESTS
