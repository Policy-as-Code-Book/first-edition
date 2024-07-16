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

helm install kyverno kyverno/kyverno -n kyverno --create-namespace \
--values values.yaml

# ${KUBECTL} get validatingwebhookconfiguration --watch

# helm install kyverno kyverno/kyverno -n kyverno --create-namespace \
# --values values.yaml --devel --set "extraArgs={-v=4,--dumpPayload=true,--enablePolicyException=true}"

# helm install kyverno kyverno/kyverno -n kyverno --create-namespace \
# --values values.yaml --devel --set "extraArgs={-v=4,--dumpPayload=true}"

# helm install kyverno kyverno/kyverno -n kyverno --create-namespace

# Grab the namespace ignore settings
# ${KUBECTL} -n kyverno get cm kyverno -o=jsonpath='{.data}'|jq -r '.webhooks'

# helm install kyverno kyverno/kyverno -n kyverno --create-namespace \
# --values values.yaml --devel --set "extraArgs={backgroundScanInterval=30m}"
