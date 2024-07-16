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

# ${KUBECTL} config set-context --current --namespace=$NS

helm -n $NS uninstall c7n-kube

# ${KUBECTL} delete mutatingwebhookconfiguration c7n-admission

read -p "Do you wish to delete the namespace? " yn
yn=${yn:-n}
case $yn in
    [Yy]* ) ${KUBECTL} delete -f k8s/0-ns.yaml --ignore-not-found
esac

read -p "Do you wish to uninstall cert manager? " yn
yn=${yn:-n}
case $yn in
    [Yy]* ) ${KUBECTL} delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.0/cert-manager.yaml --ignore-not-found
esac

