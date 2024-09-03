#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  if [ "$1" != "0" ]; then
    echo "Error $1 occurred on $2"
  fi
}

OWNER="jimmy"
ENV="dev"
BILLING="lob-cc"
KUBECTL="kubectl"
CONFIGS_DIRECTORY="generated/configs"
SECRETS_DIRECTORY="generated/secrets"
TEMPLATES_DIRECTORY="templates"

cat $TEMPLATES_DIRECTORY/crb-kubelet-api-admin-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g"  | sed -e "s/__ENV_VALUE__/${ENV}/g" | sed -e "s/__BILLING_VALUE__/${BILLING}/g" > "$CONFIGS_DIRECTORY/kubelet-api-admin.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/kubelet-api-admin.yaml
${KUBECTL} label --overwrite namespace kube-system openpolicyagent.org/webhook=ignore

# Install Styra Agents
curl -H "Authorization: Bearer <TOKEN>" "<STYRA_URL>/systems/<SYSTEM_ID>/assets/kubectl-all" | ${KUBECTL} apply -f -






