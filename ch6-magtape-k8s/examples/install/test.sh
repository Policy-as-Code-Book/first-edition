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

${KUBECTL} -n policy-test apply -f \
https://raw.githubusercontent.com/tmobile/magtape/master/testing/deployments/test-deploy02.yaml

sleep 5

${KUBECTL} -n policy-test delete -f \
https://raw.githubusercontent.com/tmobile/magtape/master/testing/deployments/test-deploy02.yaml \
--ignore-not-found


