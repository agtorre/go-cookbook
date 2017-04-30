FROM golang:alpine

ENV GOPATH /code/
ADD . /code/src/github.com/agtorre/go-cookbook/chapter10/monitoring
WORKDIR /code/src/github.com/agtorre/go-cookbook/chapter10/monitoring
RUN go build

ENTRYPOINT /code/src/github.com/agtorre/go-cookbook/chapter10/monitoring/monitoring
