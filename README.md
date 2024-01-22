# pong

A little service which replies when you go to `/ping`, with `pong` (and a few
more helpful things).

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
