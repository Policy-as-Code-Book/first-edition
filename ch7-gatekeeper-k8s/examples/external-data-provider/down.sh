#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  if [ "$1" != "0" ]; then
    echo "Error $1 occurred on $2"
  fi
}

CONFIGS_DIRECTORY="generated/configs"

KUBECTL="kubectl"

${KUBECTL} delete -f $CONFIGS_DIRECTORY/answers-app.yaml --ignore-not-found 

${KUBECTL} delete -n answers secret answers-server --ignore-not-found 

${KUBECTL} delete -f $CONFIGS_DIRECTORY/ns.yaml --ignore-not-found 

${KUBECTL} delete -f $CONFIGS_DIRECTORY/answers-provider.yaml --ignore-not-found 


