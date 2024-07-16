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

${KUBECTL} apply -f - <<EOF
apiVersion: v1
kind: Namespace
metadata:
  name: policy-test
EOF