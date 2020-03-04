
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o greeter-web *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t greeter-web:latest
