OUT := endpoint
VERSION := $(shell git describe --always --long)

all: test server

server:
	go build -o ${OUT} -ldflags="-X main.version=${VERSION}"

test:
	go test

run:
	./${OUT}

clean:
	-@rm ${OUT} ${OUT}-v*