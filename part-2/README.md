# Challenge Part-2

This second part of the challenge is responsible for reading a dump of product/image and generating another dump, grouping a maximum of 3 valid images for each product. A valid image is an image that returns 200 status code.

Input format:
```
{"productId":"pid1","image":"http://localhost:4567/images/167412.png"}
{"productId":"pid1","image":"http://localhost:4567/images/167414.png"}
```

Output format:
```
{"productId":"pid1","images":["http://localhost:4567/images/167412.png","http://localhost:4567/images/167414.png"]}
```

# Prerequisites

To run this project, you need to have the docker installed on the machine to run a mock API used to test the images. We also need Golang installed to run the project and unit tests.

* [Golang 1.14.5](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/engine/install/)

# Getting started

## Run the Mock API

To run the script to analyze the example dump, we need to start an API Mock to test the images locally. To do this, just run the following script:

```shell
./start-api.sh
```

To stop, you can still execute ctrl + c in the terminal that started or run the script

```shell
./stop-api.sh
```

## Example dump

The example dump is already at the root of that part of the challenge. You need to unzip it before using.

```shell
tar -xf input-dump.tar.gz
```

## Running

Before executing, we need to build the project. Run the command below to generate the executable `analyzer`.

```golang
go build -o analyzer
```

After the build is done, you can run the project by passing the file with the `dump` parameter. Here is an example using the existing dump in the project.

```shell
./analyzer -dump=input-dump
```

You can also run the project without making a build by running the following command:

```golang
go run . -dump=input-dump
```

# Run unit tests

To run unit tests, execute:

```golang
go test -coverprofile=coverage.out
```

To see test coverage in a more visual format, execute:

```golang
go tool cover -html=coverage.out
```
