# Challenge Part-1

In this first part of the challenge you are responsible for creating an API that rejects repeated requests in a period of 10 minutes.

# Prerequisites

To run this project, you need to have the docker-compose installed on the machine. We also need Golang installed to run unit tests.

* [Golang 1.14.5](https://golang.org/doc/install)
* [Docker Compose](https://docs.docker.com/compose/install/)

# Getting started

To run the application, just execute the following command.

```shell
docker-compose up
```

Using Docker Compose, we will upload a service with Redis and another service with the API proposed by the challenge.

To test, send a POST request to `http://localhost:8080/v1/products`. You can use a tool like Postman or run the sample CURL below.

```shell
curl -v -X POST http://localhost:8080/v1/products \
    -H "Content-Type: application/json"  \
    -d '[{"id": "123", "name": "mesa"}]'
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
