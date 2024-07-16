.DEFAULT_GOAL := help

PLATFORM ?= linux/arm64
IMAGE_REGISTRY ?= jimmyraywv
IMAGE_REPO ?= debugger
VERSION ?= v0.1.0
DEBUGGER ?= debugger

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ MagTape operations (install, test, debug, etc.)

.PHONY: down
down:	## Uninstall MagTape and associated resources
	./down.sh 2>&1

.PHONY: up
up:	## Install Magtape and associated resources
	./up.sh 2>&1

.PHONY: test
test:	## Test MagTape
	./test.sh 2>&1

.PHONY: run-debugger-pod
run-debugger-pod:	## Run debugger pod
	kubectl run -n opa -it debugger \
	--image $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION) --restart=Never

.PHONY: delete-debugger-pod
delete-debugger-pod:	## Delete debugger pod
	kubectl -n opa delete pod $(DEBUGGER)

