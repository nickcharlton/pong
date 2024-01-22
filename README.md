# pong

A little service which replies when you go to `/ping`, with `pong` (and a few
more helpful things).

The main purpose of this project is to have a simple, stateless application for
testing. It's [published as a basic Docker image][1].

## Usage

```sh
$ go run main.go
pong
Listening on 8090...
```

```go
$ curl localhost:8090/ping
pong
```

In addition:

* `/` → `I am pong`,
* `/params` → any query params you provide,
* `/headers` → any headers provided to the request

## Docker

### Building

You can build and tag the image like so:

```sh
docker build --tag pong .
```

Then tag the image with a new version:

```sh
docker image tag pong:latest pong:v1.0
```

```sh
$ docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
pong         latest    0abf2f9e56fc   5 minutes ago   892MB
pong         v1.0      0abf2f9e56fc   5 minutes ago   892MB
```

You can run it locally with:

```sh
$ docker run --publish 8090:8090 pong
pong
Listening on 8090...
```

### Publishing

```sh
docker tag pong nickcharlton/pong:v1
docker push nickcharlton/pong:v1
```

[1]: https://hub.docker.com/repository/docker/nickcharlton/pong
