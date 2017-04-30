FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/github.com/agtorre/go-cookbook/chapter10/docker
WORKDIR /code/src/github.com/agtorre/go-cookbook/chapter10/docker/example
RUN go build

ENTRYPOINT /code/src/github.com/agtorre/go-cookbook/chapter10/docker/example/example


