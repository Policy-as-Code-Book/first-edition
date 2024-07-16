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

${KUBECTL} config set-context --current --namespace=gatekeeper-system

helm uninstall gatekeeper

${KUBECTL} delete crd -l gatekeeper.sh/system=yes --ignore-not-found

${KUBECTL} delete ns gatekeeper-system --ignore-not-found

${KUBECTL} label ns kube-system admission.gatekeeper.sh/ignore-

