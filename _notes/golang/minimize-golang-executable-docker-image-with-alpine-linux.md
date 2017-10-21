# Building Minimize Golang Executable Docker Image with Alpine Linux

## Introduction

Building a minimize golang docker image to save image upload time and repository space.

**The final image is only about 10 MB**.

## Step by Step

### Docker version above 17.05

1. Use multi-stage builds, Notice:
    * naming stage on FROM by `as`
    * copy stage file by `COPY --from`

    ```dockerfile
    # Start from Alpine Linux image with the latest version of Golang
    # Naming build stage as builder
    FROM golang:alpine as builder

    # Install Git for go get
    RUN set -eux; \
        apk add --no-cache --virtual git

    # Set ENV
    ENV GOPATH /go/
    ENV GO_WORKDIR $GOPATH/src/github.com/dinos80152/golang-docker-alpine/

    # Set WORKDIR to go source code directory
    WORKDIR $GO_WORKDIR

    # Add files to image
    ADD . $GO_WORKDIR

    # Fetch Golang Dependency and Build Binary
    RUN go get
    RUN go install


    # Start from a raw Alpine Linux image
    FROM alpine:latest

    # Install ca-certificates for ssl
    RUN set -eux; \
        apk add --no-cache --virtual ca-certificates

    # Copy binary from builder stage into image
    COPY --from=builder /go/bin/golang-docker-alpine .

    # Execute binary when docker container starts
    CMD /golang-docker-alpine

    # Expose port 8080 to be connected from outside
    EXPOSE 8080
    ```

1. Build image

    ```bash
    docker build -t golang-docker-alpine .
    ```

1. Test

    ```bash
    docker run -p 80:8080 -d golang-docker-alpine
    ```

### Docker version lower than 17.05

1. Write Dockerfile for compiling source code to binary on alpine linux

    * DockerfileSrc

    ```dockerfile
    # Start from an Alpine image with the latest version of Go installed
    FROM golang:alpine

    # Install Git for go get
    RUN set -eux; \
        apk add --no-cache --virtual git

    # Set ENV
    ENV GOPATH /go/
    ENV GO_WORKDIR $GOPATH/src/github.com/dinos80152/golang-docker-alpine/

    # Set WORKDIR to go source code directory
    WORKDIR $GO_WORKDIR

    # Add files to image
    ADD . $GO_WORKDIR

    # Fetch Golang Dependency and Build Binary
    RUN go get
    RUN go install
    ```

1. Build image from DockerfileSrc

    ```bash
    # build app binary on Alpine OS with golang package
    docker build -t dinos80152/golang-docker-alpine-src -f DockerfileSrc .
    ```

1. Copy binary from container to host

    ```bash
    # Copy app binary from container to host
    docker cp $(docker create dinos80152/golang-docker-alpine-src):/go/bin/golang-docker-alpine ./
    ```

1. Write Dockerfile for building image include app binary and pure alpine linux

    * DockerfileBin

    ```dockerfile
    # Start from a raw Alpine image
    FROM alpine:latest

    # Install ca-certificates for ssl
    RUN set -eux; \
        apk add --no-cache --virtual ca-certificates

    # Copy binary into image
    ADD golang-docker-alpine /

    # Execute binary when docker container starts
    CMD /golang-docker-alpine

    # Expose port 8080 to be connected from outside
    EXPOSE 8080
    ```

1. Build docker image from DockerfileBin

    ```bash
    # Build docker image with pure alpine os and app binary
    docker build -t golang-docker-alpine -f DockerfileBin .
    ```

1. Test

    ```bash
    docker run -p 80:8080 -d golang-docker-alpine
    ```

## Warning

Function **time.LoadLocation() will panic in binary image**, because of

> LoadLocation looks in the directory or uncompressed zip file named by the ZONEINFO environment variable, if any, then looks in known installation locations on Unix systems, and finally looks in $GOROOT/lib/time/zoneinfo.zip.

You have to **set $GOROOT and copy zoneinfo.zip manually**.

## Template

[Github Repo online demo with Travis CI and Heroku](https://github.com/dinos80152/golang-docker-alpine)

## Reference

* [Alpine Linux](https://alpinelinux.org/)
* [Use multi-stage builds@docker docs](https://docs.docker.com/engine/userguide/eng-image/multistage-build/)
* [Docker golang:alpine](https://github.com/docker-library/golang/tree/64b88dc3e9d83e71eafc000fed1f0d5e289b3e65/1.8/alpine)
* [Docker - how can I copy a file from an image to a host?](https://stackoverflow.com/questions/25292198/)
* [Docker CLI@docker docs](https://docs.docker.com/engine/reference/commandline/docker/)
* [Dockerfile reference@docker docs](https://docs.docker.com/engine/reference/builder/)
* [x509: failed to load system roots and no roots provided](https://github.com/zenazn/goji/issues/126)