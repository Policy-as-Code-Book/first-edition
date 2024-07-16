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

helm install gatekeeper gatekeeper/gatekeeper \
    --namespace gatekeeper-system --create-namespace \
    --values values.yaml
    
    # --set logDenies=true --set replicas=1 --set version=v3.15.0 \
    # --set "controllerManager.exemptNamespaces={kube-system}" \
    # --set enableGeneratorResourceExpansion=true \
    # --set enableExternalData=true --set externaldataProviderResponseCacheTTL=0
    
    # --set emitAuditEvents=true 
    # --set logMutations=true --set mutationAnnotations=true \
    #  \
    # --set logLevel=DEBUG --set controllerManager.disableCertRotation=true \
    # --set logDenies=true --set replicas=1 --set version=v3.11.0 \
    # --set logDenies=true --set replicas=1 --set version=v3.15.0 \

LABEL=$(${KUBECTL} get ns kube-system -oyaml | { grep admission.gatekeeper.sh/ignore || true; })
if [[ "$LABEL" == "" ]]
then
  ${KUBECTL} label ns kube-system admission.gatekeeper.sh/ignore=jimmy
  ${KUBECTL} get ns kube-system -oyaml | grep admission.gatekeeper.sh/ignore
fi