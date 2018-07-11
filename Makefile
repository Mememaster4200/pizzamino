# Generic Golang Makefile
# `make get` for dependencies, `make` to build

GOPATH=$(shell pwd)
GOFILES=$(wildcard src/*.go)

build: # compile
	@GOPATH=$(GOPATH) GOBIN=$(GOPATH)/bin go build -o bin/pizzamino $(GOFILES)

clean: # clean mess
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

fmt: # clean syntax
	gofmt -l -w src

get: # get dependencies
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get github.com/veandco/go-sdl2/{sdl,img,mix,ttf}

install: # install build
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)
