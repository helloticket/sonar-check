# General
GIT_BRANCH			:= $(shell git symbolic-ref HEAD | sed -e 's/^refs\/heads\///')
GIT_LAST_COMMIT		:= $(shell git rev-parse --short HEAD)

# Version
VMAJOR_MINOR 		:= $(or ${VBRANCH}, ${VTAG}, ${GIT_BRANCH})
VBUILD 				:= $(or ${VBUILD}, 0)
VREV 				:= $(or ${VREV}, ${GIT_LAST_COMMIT})
VERSION				:= "$(VMAJOR_MINOR).$(VBUILD).$(shell echo ${VREV} | cut -c 1-8)"

all: help

tests: unit-test

unit-test:
	@go test -cover $(shell go list ./... | grep -v /vendor | grep -v /tests)

build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o release/linux/sonarcheck cmd/main.go
	
release:
	@tar -czvf release/linux/sonarcheck_$(VERSION)_linux_amd64.tar.gz release/linux

.PHONY: all tests build release
