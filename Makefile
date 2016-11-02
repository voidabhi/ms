.PHONY: all
SHELL := /bin/bash
PKG = github.com/voidabhi/ms
PKGS := './..'

all: lint test

lint:
	golint ./...
	go fmt ./...
	go vet ./...

test:
	go test ./... -v
