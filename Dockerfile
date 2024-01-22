# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./

EXPOSE 8090

RUN GOOS=linux go build -o /pong

CMD ["/pong"]
