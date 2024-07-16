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

OWNER="jimmy"
ENV="dev"
BILLING="lob-cc"
KUBECTL="kubectl"
CA_BUNDLE=""
CONFIGS_DIRECTORY="generated/configs"
SECRETS_DIRECTORY="generated/secrets"
TEMPLATES_DIRECTORY="templates"
POLICY_CM_DIRECTORY="../policy-configmaps/deploy"

if [ ! -d "$TEMPLATES_DIRECTORY" ]; then
  echo "$TEMPLATES_DIRECTORY not found, install aborted"
  exit 99
fi

if [ ! -d "$CONFIGS_DIRECTORY" ]; then
  mkdir -p $CONFIGS_DIRECTORY
fi

if [ ! -d "$SECRETS_DIRECTORY" ]; then
  mkdir -p $SECRETS_DIRECTORY
fi

rm -f $CONFIGS_DIRECTORY/*
rm -f $SECRETS_DIRECTORY/*

cat $TEMPLATES_DIRECTORY/ns-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" \
| sed -e "s/__ENV_VALUE__/${ENV}/g" \
| sed -e "s/__BILLING_VALUE__/${BILLING}/g" \
> "$CONFIGS_DIRECTORY/ns.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/ns.yaml

${KUBECTL} label --overwrite namespace kube-system \
openpolicyagent.org/webhook=ignore
${KUBECTL} label --overwrite namespace opa \
openpolicyagent.org/webhook=ignore

openssl genrsa -out $SECRETS_DIRECTORY/opa-ca.key 2048
openssl req -x509 -new -nodes -sha256 -key $SECRETS_DIRECTORY/opa-ca.key -days 100000 -out $SECRETS_DIRECTORY/opa-ca.crt -subj /CN=admission_ca 2>&1
openssl genrsa -out $SECRETS_DIRECTORY/opa-server.key 2048
openssl req -new -key $SECRETS_DIRECTORY/opa-server.key -sha256 -out $SECRETS_DIRECTORY/opa-server.csr -subj /CN=opa.opa.svc -config $TEMPLATES_DIRECTORY/opa-server.conf 2>&1
openssl x509 -req -in $SECRETS_DIRECTORY/opa-server.csr -sha256 -CA $SECRETS_DIRECTORY/opa-ca.crt -CAkey $SECRETS_DIRECTORY/opa-ca.key -CAcreateserial -out $SECRETS_DIRECTORY/opa-server.crt -days 100000 -extensions v3_ext -extfile $TEMPLATES_DIRECTORY/opa-server.conf

echo "Trying to delete \"opa-server\" secret..."
${KUBECTL} -n opa delete secret opa-server --ignore-not-found 2>&1
${KUBECTL} -n opa create secret tls opa-server --cert=$SECRETS_DIRECTORY/opa-server.crt --key=$SECRETS_DIRECTORY/opa-server.key

cat $TEMPLATES_DIRECTORY/admission-controller-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" | sed -e "s/__ENV_VALUE__/${ENV}/g" \
| sed -e "s/__BILLING_VALUE__/${BILLING}/g" > "$CONFIGS_DIRECTORY/opa-admission-controller.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/opa-admission-controller.yaml
  
CA_BUNDLE="$(base64 -i $SECRETS_DIRECTORY/opa-ca.crt)"

Sleep 10

${KUBECTL} apply -f $POLICY_CM_DIRECTORY/.

CA_BUNDLE="$(base64 -i $SECRETS_DIRECTORY/opa-ca.crt)"

cat $TEMPLATES_DIRECTORY/validating-webhook-configuration-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" | sed -e "s/__ENV_VALUE__/${ENV}/g" \
| sed -e "s/__BILLING_VALUE__/${BILLING}/g" | sed -e "s/__CA_BUNDLE_VALUE__/${CA_BUNDLE}/g" > "$CONFIGS_DIRECTORY/opa-validating-webhook-configuration.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/opa-validating-webhook-configuration.yaml

