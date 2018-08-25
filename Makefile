.DEFAULT_GOAL := build

build:
	go build -o z -ldflags="-s -w" ./cmd/z

.PHONY: build
