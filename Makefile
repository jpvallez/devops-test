# OUT is the name of the binary
OUT := endpoint
# VERSION grabs the latest tag from repo
VERSION := $(shell git describe --always --long)

all: test server

server:
	go build -o ${OUT} -ldflags="-X main.version=${VERSION}"

test:
	go test
