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
NS=${1:-jspolicy}

# ${KUBECTL} config set-context --current --namespace=$NS

helm -n $NS uninstall jspolicy

${KUBECTL} delete ns $NS --ignore-not-found

${KUBECTL} delete validatingwebhookconfiguration jspolicy --ignore-not-found
${KUBECTL} delete mutatingwebhookconfiguration jspolicy --ignore-not-found

${KUBECTL} api-resources --api-group='policy.jspolicy.com' \
-o name | xargs ${KUBECTL} delete --ignore-not-found crd

${KUBECTL} api-resources --api-group='wgpolicyk8s.io' \
-o name | xargs ${KUBECTL} delete  --ignore-not-found crd

${KUBECTL} label ns kube-system jspolicy-ignore-