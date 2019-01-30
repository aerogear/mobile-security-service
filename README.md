# Mobile Security Service

[![CircleCI](https://circleci.com/gh/aerogear/mobile-security-service.svg?style=svg)](https://circleci.com/gh/aerogear/mobile-security-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/aerogear/mobile-security-service)](https://goreportcard.com/report/github.com/aerogear/mobile-security-service)
[![Coverage Status](https://coveralls.io/repos/github/aerogear/mobile-security-service/badge.svg?branch=master)](https://coveralls.io/github/aerogear/mobile-security-service?branch=master)

This is the server component of the AeroGear Mobile Security Service. It is a RESTful API that allows developers to view, enable and disable specific versions of applications on demand, with the information stored in a PostgreSQL database. The service is written in [Golang](https://golang.org/).

## Prerequisites

* [Install Golang](https://golang.org/doc/install)
* [Ensure the $GOPATH environment variable is set](https://github.com/golang/go/wiki/SettingGOPATH)
* [Install the dep package manager](https://golang.github.io/dep/docs/installation.html)
* [Install Docker and Docker Compose](https://docs.docker.com/compose/install/)

See the [Contributing Guide](./CONTRIBUTING.md) for more information.

## Getting Started

If you'd like to simply run the entire application in `docker-compose`, check out the relevant section.

Golang projects are kept in a [workspace](https://golang.org/doc/code.html#Workspaces) that follows a very specific architecture. Before cloning this repo, be sure you have a `$GOPATH` environment variable set up.

### Clone the Repository

```sh
git clone git@github.com:aerogear/mobile-security-service.git $GOPATH/src/github.com/aerogear/mobile-security-service
```

### Install Dependencies

```sh
make setup
```

Note this is using the `dep` package manager under the hood. You will see the dependencies installed in the `vendor` folder.

### Start the Server

```sh
go run cmd/mobile-security-service/main.go
```

### Run Entire Application with Docker Compose

This section shows how to start the entire application with `docker-compose`. This is useful for doing some quick tests (using the SDKs) for example.

First, compile a Linux compatible binary:

```bash
go build -o mobile-security-service cmd/mobile-security-service/main.go
```

This binary will be used to build the Docker image. Now start the entire application.

```bash
docker-compose up
```

## Environment Variables

The **mobile-security-service** is configured using environment variables.

1. By default, the application will look for system environment variables to use.
2. If a system environment variable cannot be found, the application will then check the `.env` file in the application root.
3. If the `.env` file does not exist, or if the variable is not defined in the file, the application will use the default value defined in [config.go](./pkg/config/config.go).

### Add your own .env file

Make a copy of the example file `.env.example`.

```sh
cp .env.example .env
```

Now the application will use the values defined in `.env`.

### Server Configuration

| Variable                         | Default | Description                                                                                                                        |
|----------------------------------|---------|------------------------------------------------------------------------------------------------------------------------------------|
| PORT                             | 3000    | The port the server will listen on                                                                                                 |
| LOG_LEVEL                        | info    | Can be one of `[debug, info, warning, error, fatal, panic]`                                                                        |
| LOG_FORMAT                       | text    | Can be one of `[text, json]`                                                                                                       |
| ACCESS_CONTROL_ALLOW_ORIGIN      | *       | Can be multiple URL values separated with commas. Example: `ACCESS_CONTROL_ALLOW_ORIGIN=http://www.example.com,http://example.com` |
| ACCESS_CONTROL_ALLOW_CREDENTIALS | false   | Can be one of `[true, false]`                                                                                                      |

### Using Swagger UI

#### Method 1

A [Swagger](https://swagger.io/) UI can be used for testing the mobile-security-service service.

```bash
docker run -p 8080:8080 -e API_URL=https://raw.githubusercontent.com/aerogear/mobile-security-service/master/apispec.yaml swaggerapi/swagger-ui
```

The Swagger UI is available at [localhost:8080](http://localhost:8080).

#### Method 2

There is also a [Chrome extension](https://chrome.google.com/webstore/detail/swagger-ui-console/ljlmonadebogfjabhkppkoohjkjclfai?hl=en) you can use instead of running a Docker container.

Paste [https://raw.githubusercontent.com/aerogear/mobile-security-service/master/apispec.yaml](https://raw.githubusercontent.com/aerogear/mobile-security-service/master/apispec.yaml) and press **Explore**.

## Building & Testing

The `Makefile` provides commands for building and testing the code. Some dependencies are required to run these commands.

### Dependencies

Dependencies may be required to run some of the `Make` commands. Below are instructions on how to install them.

#### errcheck

[errcheck](https://github.com/kisielk/errcheck) is required to run the `make errcheck` command.

Install:

```sh
go get -u github.com/kisielk/errcheck
```

| Command                       | Description                                                                                     |
|-------------------------------|-------------------------------------------------------------------------------------------------|
| `make setup`                  | Downloads dependencies into `vendor`                                                            |
| `make build`                  | Compile a binary compatible with your current system into `./mobile-security-service`    |
| `make build_linux`            | Compile a Linux binary into `./dist/linux_amd64/mobile-security-service`                 |
| `make docker_build`           | Compile a binary and create a Docker image from it.                                             |
| `make docker_build_release`   | Compile a binary and create a Docker image with a release tag                                   |
| `make docker_build_master`    | Compile a binary and create a Docker image tagged `master`                                      |
| `make test`                   | Runs unit tests                                                                                 |
| `make test-integration`       | Runs integration tests                                                                          |
| `make test-integration-cover` | Runs integration tests and outputs results to a log file                                        |
| `make errcheck`               | Checks for unchecked errors using [errcheck](https://github.com/kisielk/errcheck)               |
| `make vet`                    | Examines source code and reports suspicious constructs using [vet](https://golang.org/cmd/vet/) |
| `make fmt`                    | Formats code using [gofmt](https://golang.org/cmd/gofmt/)                                       |
| `make clean`                  | Removes binary compiled using `make build`                                                      |
| `make docker_push_release`    | Pushes release image to Docker image hosting repository                                         |
| `make docker_push_master`     | Pushes master image to Docker image hosting repository                                          |

## Built With

* [Golang](https://golang.org/) - Programming language used
* [Echo](https://echo.labstack.com/) - Web framework used

## Contributing

All contributions are hugely appreciated. Please see our [Contributing Guide](./CONTRIBUTING.md) for guidelines on how to open issues and pull requests. Please check out our [Code of Conduct](./.github/CODE_OF_CONDUCT) too.

## Questions

There are a number of ways you can get in in touch with us:

* Open a [GitHub issue](https://github.com/aerogear/mobile-security-service/issues/new).
* Start a thread in the [Aerogear Mailing List](https://groups.google.com/forum/#!forum/aerogear) (Open to anyone).
* Reach out to us on IRC. The Aerogear team can be found at the #Aerogear channel on [freenode.net](https://freenode.net/) (Open to anyone).