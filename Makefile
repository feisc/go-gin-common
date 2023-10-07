## Harbor server
HARBOR_ADDR := harbor.zvos.zoomlion.com

## Harbor project name
PROJECT_NAME := edge

## Name of the service/application
SERVICE_NAME := zvos-edge-command-control

## Docker image name for the project
IMAGE_NAME := $(PROJECT_NAME)/$(SERVICE_NAME)

## Repository url for this project
REPOSITORY := $(HARBOR_ADDR)/$(IMAGE_NAME)

## Shell to use for running scripts
SHELL := $(shell which bash)

## Get docker path or an empty string
DOCKER := $(shell command -v docker)

## Get the main unix group for the user running make (to be used by docker-compose later)
GID := $(shell id -g)

## Get the unix user id for the user running make (to be used by docker-compose later)
UID := $(shell id -u)

## Commit hash from git
COMMIT=$(shell git rev-parse --verify --short HEAD)

## Branch from git
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

## image name
IMG ?= $(REPOSITORY):$(BRANCH)_$(COMMIT)

## env name 
ENV ?= dev

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Run go fmt against code.
.PHONY: fmt
fmt:
	go fmt ./...

## Run go vet against code.
.PHONY: vet
vet: 
	go vet ./...

## Build image
.PHONY: docker-build
docker-build: fmt vet
	docker build -t ${IMG} .

## Push image
.PHONY: docker-push
docker-push:
	docker push ${IMG}
