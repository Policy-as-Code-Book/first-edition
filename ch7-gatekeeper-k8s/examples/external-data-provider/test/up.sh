#!/usr/bin/env bash

trap 'catch $? $LINENO' ERR
catch() {
  echo "Error $1 occurred on $2"
}

KUBECTL="kubectl"

${KUBECTL} apply -f 0-ct.yaml
sleep 30
${KUBECTL} apply -f 1-c.yaml

