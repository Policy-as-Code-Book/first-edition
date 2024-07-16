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

# Delete resources
curl -H "Authorization: Bearer <TOKEN>" "<STYRA_URL>/systems/<SYSTEM_ID>/assets/kubectl-all" | ${KUBECTL} delete -f - 

# Release SLP Persistant Volume Claims
${KUBECTL} delete --namespace=styra-system pvc -l slp-pvc=slp-kubernetes-app-pvc --ignore-not-found

${KUBECTL} delete clusterrolebinding kubelet-api-admin --ignore-not-found
${KUBECTL} label namespace kube-system openpolicyagent.org/webhook-
