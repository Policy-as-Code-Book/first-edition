#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  if [ "$1" != "0" ]; then
    echo "Error $1 occurred on $2"
  fi
}

KUBECTL="kubectl"

${KUBECTL} config set-context --current --namespace=kyverno

helm uninstall kyverno

# No longer needed
# ${KUBECTL} delete mutatingwebhookconfigurations -l webhook.kyverno.io/managed-by=kyverno

# ${KUBECTL} delete validatingwebhookconfigurations -l webhook.kyverno.io/managed-by=kyverno

${KUBECTL} delete ns kyverno --ignore-not-found

