#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  if [ "$1" != "0" ]; then
    echo "Error $1 occurred on $2"
  fi
}

CMD_KUBECTL="kubectl"

${CMD_KUBECTL} delete -f 0-ns.yaml --ignore-not-found 2>&1

