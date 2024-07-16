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

# helm install policy-reporter policy-reporter/policy-reporter -n policy-reporter \
#   --set metrics.enabled=true --set rest.enabled=true --create-namespace

# helm upgrade --install policy-reporter policy-reporter/policy-reporter \
#   --create-namespace -n policy-reporter --set metrics.enabled=true --set api.enabled=true

# helm install policy-reporter policy-reporter/policy-reporter -n policy-reporter \
#   --set metrics.enabled=true --set rest.enabled=true --create-namespace

helm install policy-reporter policy-reporter/policy-reporter \
  --set kyvernoPlugin.enabled=true --set ui.enabled=true \
  --set ui.plugins.kyverno=true -n policy-reporter --create-namespace

