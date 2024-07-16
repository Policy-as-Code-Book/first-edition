#!/usr/bin/env bash

# debugging
# set -x
# set -v

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  if [ "$1" != "0" ]; then
    echo "Error $1 occurred on $2"
  fi
}

KUBECTL="kubectl"
${KUBECTL} delete MutatingWebhookConfiguration -l app=opa --ignore-not-found
${KUBECTL} delete ValidatingWebhookConfiguration -l app=opa --ignore-not-found
${KUBECTL} delete ns opa --ignore-not-found
${KUBECTL} delete ns policy-test --ignore-not-found
# ${KUBECTL} delete clusterrole opa --ignore-not-found
# ${KUBECTL} delete clusterrolebinding opa --ignore-not-found
${KUBECTL} delete clusterrolebinding opa-kube-mgmt-viewer --ignore-not-found
# ${KUBECTL} delete clusterrolebinding kubelet-api-admin --ignore-not-found
${KUBECTL} label namespace kube-system openpolicyagent.org/webhook-
