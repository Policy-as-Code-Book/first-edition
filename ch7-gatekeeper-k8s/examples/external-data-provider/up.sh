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
CA_BUNDLE=""
CONFIGS_DIRECTORY="generated/configs"
SECRETS_DIRECTORY="generated/secrets"
TEMPLATES_DIRECTORY="templates"

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

openssl genrsa -out $SECRETS_DIRECTORY/answers-ca.key 2048
openssl req -x509 -new -nodes -sha256 -key $SECRETS_DIRECTORY/answers-ca.key -days 365 -out $SECRETS_DIRECTORY/answers-ca.crt -subj /CN=admission_ca 2>&1
openssl genrsa -out $SECRETS_DIRECTORY/answers-server.key 2048
openssl req -new -key $SECRETS_DIRECTORY/answers-server.key -sha256 -out $SECRETS_DIRECTORY/answers-server.csr -subj /CN=answers.answers.svc -config $TEMPLATES_DIRECTORY/server.conf 2>&1
openssl x509 -req -in $SECRETS_DIRECTORY/answers-server.csr -sha256 -CA $SECRETS_DIRECTORY/answers-ca.crt -CAkey $SECRETS_DIRECTORY/answers-ca.key -CAcreateserial -out $SECRETS_DIRECTORY/answers-server.crt -days 100000 -extensions v3_ext -extfile $TEMPLATES_DIRECTORY/server.conf

cat $TEMPLATES_DIRECTORY/ns-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" | sed -e "s/__ENV_VALUE__/${ENV}/g" | sed -e "s/__BILLING_VALUE__/${BILLING}/g" > "$CONFIGS_DIRECTORY/ns.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/ns.yaml

echo "Trying to delete \"answers-server\" secret..."
${KUBECTL} -n answers delete secret answers-server --ignore-not-found 2>&1
${KUBECTL} -n answers create secret tls answers-server --cert=$SECRETS_DIRECTORY/answers-server.crt --key=$SECRETS_DIRECTORY/answers-server.key

cat $TEMPLATES_DIRECTORY/answers-app-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" | sed -e "s/__ENV_VALUE__/${ENV}/g" | sed -e "s/__BILLING_VALUE__/${BILLING}/g" > "$CONFIGS_DIRECTORY/answers-app.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/answers-app.yaml

CA_BUNDLE="$(base64 -i $SECRETS_DIRECTORY/answers-ca.crt)"
cat $TEMPLATES_DIRECTORY/provider-template.yaml | sed -e "s/__OWNER_VALUE__/${OWNER}/g" | sed -e "s/__ENV_VALUE__/${ENV}/g" | sed -e "s/__BILLING_VALUE__/${BILLING}/g" | sed -e "s/__CA_BUNDLE_VALUE__/${CA_BUNDLE}/g" > "$CONFIGS_DIRECTORY/answers-provider.yaml"
${KUBECTL} apply -f $CONFIGS_DIRECTORY/answers-provider.yaml


