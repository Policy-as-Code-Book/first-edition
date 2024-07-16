.DEFAULT_GOAL := help

PLATFORM ?= linux/arm64
IMAGE_REGISTRY ?= jimmyraywv
IMAGE_REPO ?= debugger
DOCKERFILE ?= Dockerfile.debugger
VERSION ?= v0.1.0
NO_CACHE ?= true
DEBUGGER ?= debugger

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ OPA operations (install, test, etc.)

.PHONY: clean
clean:	## Clean locally generated configs and secrets
	./clean-local-config.sh 2>&1
	./clean-local-secrets.sh 2>&1

.PHONY: down
down:	## Uninstall OPA and associated resources
	./down.sh 2>&1

.PHONY: up
up:	## Install OPA and associated resources
	./up.sh 2>&1

.PHONY: test
test:	## Run policy tests in the cluster
	./tests/tests.sh 2>&1

.PHONY: test-clean
test-clean:	## Clean up after policy tests
	./tests/clean.sh 2>&1

.PHONY:
apply: apply	## Apply policies
	./policy-configmaps/apply.sh 2>&1

.PHONY:
delete: delete	## Delete policies
	./policy-configmaps/delete.sh 2>&1

.PHONY: test-rego
test-rego:	## Test Rego policies
	opa test policy-rego/*.rego -v

.PHONY: build-debugger
build-debugger:	## Build debugger pod
	docker buildx build -f $(DOCKERFILE) \
	    --push \
		--platform $(PLATFORM) \
		--tag $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION) . --no-cache=$(NO_CACHE)
	$(info	)

.PHONY: run-debugger-pod
run-debugger-pod:	## Run debugger pod
	kubectl run -n opa -it debugger \
	--image $(IMAGE_REGISTRY)/$(IMAGE_REPO):$(VERSION) --restart=Never

.PHONY: delete-debugger-pod
delete-debugger-pod:	## Delete debugger pod
	kubectl -n opa delete pod $(DEBUGGER)

