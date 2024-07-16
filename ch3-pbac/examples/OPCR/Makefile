.DEFAULT_GOAL := build

GIT_URL ?= $(shell git remote get-url --push origin)
GIT_COMMIT ?= $(shell git rev-parse HEAD)
TIMESTAMP := $(shell date '+%Y-%m-%d_%I:%M:%S%p')
REGION ?= us-east-2
# IMAGE_REGISTRY ?= opcr.io
IMAGE_REGISTRY ?= ghcr.io
IMAGE_REPO ?= jimmyraywv/pacbook
GIT_COMMIT_IN ?=
GIT_URL_IN ?=

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

.PHONY: build push pull meta clean help	run bundle

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY:	meta
meta:	## Provides metadata for other commands; good for DevOps logging. Can be called as a target, but is mostly used by other targets as a dependency.
	$(info    [METADATA])
	$(info    timestamp: [$(TIMESTAMP)])
	$(info    git commit: [$(GIT_COMMIT)])
	$(info    git URL: [$(GIT_URL)])
	$(info    Image version: [$(VERSION)])
	$(info	)

##@ Build and deploy

.PHONY:	build
build:	meta ## Build container with Docker buildx, based on PLATFORM argument (default linux/amd64)
	policy build ./bundle-material -t $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION)

.PHONY:	login
login:	## Login to remote image registry
	# GH_PAT=`cat <PATH_TO_GH_PAT_FILE>` && policy login -u jimmyraywv -p "$GH_PAT"
	GH_PAT=`cat <PATH_TO_GH_PAT_FILE>` && echo $GH_PAT| policy login -s ghcr.io -u jimmyraywv --password-stdin

.PHONY:	push
push:	meta	## Push to remote image registry
	policy push $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION)

.PHONY:	pull
pull:	meta	## Pull from remote image registry
	policy pull $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION)

.PHONY:	clean
clean:	## Clean by removing image and tarball
	policy rm -f $(IMAGE_REPO):$(VERSION)
	yes "yes" | rm -vrf ~/.policy/policies-root/blobs/sha256/*

.PHONY:	run
run:	## Run the OPA repl with the bundle image
	-@policy repl $(IMAGE_REPO):$(VERSION)

.PHONY:	bundle
bundle:	## Create bundle
	opa build ./bundle-material

