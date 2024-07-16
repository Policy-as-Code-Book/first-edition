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

${KUBECTL} config set-context --current --namespace=policy-reporter

helm uninstall policy-reporter

${KUBECTL} delete ns policy-reporter --ignore-not-found

