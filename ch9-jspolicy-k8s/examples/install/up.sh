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

helm install jspolicy jspolicy -n $NS \
--create-namespace --repo https://charts.loft.sh

LABEL=$(${KUBECTL} get ns $NS -oyaml | { grep jspolicy-ignore || true; })
if [[ "$LABEL" == "" ]]
then
  ${KUBECTL} label ns $NS jspolicy-ignore=ignore
  ${KUBECTL} get ns $NS -oyaml | grep jspolicy-ignore
fi

LABEL=$(${KUBECTL} get ns kube-system -oyaml | { grep jspolicy-ignore || true; })
if [[ "$LABEL" == "" ]]
then
  ${KUBECTL} label ns kube-system jspolicy-ignore=ignore
  ${KUBECTL} get ns kube-system -oyaml | grep jspolicy-ignore
fi




