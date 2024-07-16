.DEFAULT_GOAL := help

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Kyverno operations (install, test, etc.)

.PHONY:	up
up:	## Install Kyverno
	./up.sh

.PHONY:	down
down:	## Uninstall Kyverno
	./down.sh

.PHONY:	up-reporter
up-reporter:	## Install Kyverno Policy Reporter UI
	./up-reporter.sh

.PHONY:	down-reporter
down-reporter:	## Uninstall Kyverno Policy Reporter UI
	./down-reporter.sh

.PHONY: up-minikube
up-minikube:	## Launch minikube
	minikube start -p kyverno \
	--extra-config=apiserver.runtime-config=\
	admissionregistration.k8s.io/v1alpha1,admissionregistration.k8s.io/v1beta1  \
	--feature-gates='ValidatingAdmissionPolicy=true'

.PHONY: down-minikube
down-minikube:	## Down minikube
	minikube stop -p kyverno




