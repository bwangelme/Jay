.PHONY: build

QAE_APPROOT=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
export QAE_APPROOT

build:
	go build -o bin/server main.go

st:
	@go test -v -run ${func} ./... || echo