.PHONY: all
SHELL := /bin/bash
PKG = github.com/voidabhi/ms
PKGS := './..'

all: lint test

build: go build ./...

lint:
	golint ./...
	go fmt ./...
	go vet ./...

test:
	go test ./... -v
