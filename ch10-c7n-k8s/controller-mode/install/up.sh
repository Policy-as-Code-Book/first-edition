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
NS=${1:-c7n-kube}

read -p "Do you wish to install certmanager?" yn
yn=${yn:-n}
case $yn in
    [Yy]* ) ${KUBECTL} apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.0/cert-manager.yaml
    sleep 120
esac

read -p "Do you wish to install/update the namespace?" yn
yn=${yn:-n}
case $yn in
    [Yy]* ) ${KUBECTL} apply -f k8s/0-ns.yaml
esac

helm install c7n-kube c7n/c7n-kube -n $NS \
--create-namespace --values values.yaml




