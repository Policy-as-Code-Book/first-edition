#!/usr/bin/env bash

trap 'catch $? $LINENO' ERR
catch() {
  echo "Error $1 occurred on $2"
}

KUBECTL="kubectl"

${KUBECTL} delete -f 1-c.yaml
sleep 10
${KUBECTL} delete -f 0-ct.yaml

