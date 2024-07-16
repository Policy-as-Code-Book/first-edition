.DEFAULT_GOAL := build

GIT_URL ?= $(shell git remote get-url --push origin)
GIT_COMMIT ?= $(shell git rev-parse HEAD)
TIMESTAMP := $(shell date '+%Y-%m-%d_%I:%M:%S%p')
REGION ?= us-east-2
IMAGE_REGISTRY ?= jimmyraywv
IMAGE_REPO ?= answers
DOCKERFILE ?= Dockerfile
NO_CACHE ?= true
GIT_COMMIT_IN ?=
GIT_URL_IN ?=
GO_MOD_PATH ?= jimmyray.io/magic-8-ball
# PLATFORM ?= linux/amd64,linux/arm64
PLATFORM ?= linux/arm64
# PLATFORM ?= linux/amd64

ifeq ($(strip $(GIT_COMMIT)),)
GIT_COMMIT := $(GIT_COMMIT_IN)
endif

ifeq ($(strip $(GIT_URL)),)
GIT_URL := $(GIT_URL_IN)
endif

VERSION_HASH := $(shell echo $(GIT_COMMIT)|cut -c 1-10)
# $(info [$(VERSION_HASH)])
VERSION_FROM_FILE ?= $(shell head -n 1 version)
VERSION ?=

ifeq ($(strip $(VERSION_HASH)),)
VERSION := $(VERSION_FROM_FILE)
else
VERSION := $(VERSION_FROM_FILE)-$(VERSION_HASH)
endif

##@ General

phony: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

phony: meta
meta:	## Provides metadata for other commands; good for DevOps logging. Can be called as a target, but is mostly used by other targets as a dependency.
	$(info    [METADATA])
	$(info    timestamp: [$(TIMESTAMP)])
	$(info    git commit: [$(GIT_COMMIT)])
	$(info    git URL: [$(GIT_URL)])
	$(info    Container image version: [$(VERSION)])
	$(info	)

##@ Build and deploy

phony: build
build:	meta ## Build container with Docker buildx, based on PLATFORM argument (default linux/amd64)
	$(info    [BUILD_CONTAINER_IMAGE])
	# docker buildx build \
	# 	--load \
	# 	--platform $(PLATFORM) \
	# 	--tag $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION) . --no-cache=$(NO_CACHE)
	docker buildx build \
	    --push \
		--platform $(PLATFORM) \
		--tag $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION) . --no-cache=$(NO_CACHE)
	$(info	)

phony: login
login:	## Login to remote image registry
	$(info    [REMOTE_REGISTRY_LOGIN])
	aws ecr get-login-password --region $(REGION) | docker login --username AWS --password-stdin $(IMAGE_REGISTRY)
	$(info	)

phony: push
push:	meta	## Push to remote image registry
	$(info    [PUSH_CONTAINER_IMAGE])
	docker push $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION)
	$(info	)

phony: push
pull:	meta	## Pull from remote image registry
	$(info    [PULL_CONTAINER_IMAGE])
	docker pull $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION)
	$(info	)

##@ Local Development

phony: compile
compile:	clean	meta	## Compile for local MacOS
	$(info   [COMPILE])
	go env -w GOPROXY=direct && GOOS=darwin GOARCH=arm64 go build -a -o main.bin .
	$(info	)

# go env -w GOPROXY=direct && CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -o main.bin .

phony: clean
clean:	## Remove compile binary
	-@rm main.bin

phony: init
init:	## Initialize Go project
	-@rm go.mod
	-@rm go.sum
	go mod init $(GO_MOD_PATH)
	go mod tidy

phony: check
check:	## Vet and Lint Go codebase
	-go generate ./...
	-go vet ./...
	-go fmt ./...
	-golangci-lint run

phony: test
test:	## Run tests
	go test $(GO_MOD_PATH) -test.v

phony: run
run:	## Run local binary
	./main.bin --config server-config.json
