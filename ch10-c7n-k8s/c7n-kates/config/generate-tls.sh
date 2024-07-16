#!/usr/bin/env bash

trap 'catch $? $LINENO' ERR
catch() {
  echo "Error $1 occurred on $2"
}

SERVER_CONF_DIR=${1:-"server"}
SECRETS_DIRECTORY="secrets"

if [ ! -d "$SERVER_CONF_DIR" ]; then
  echo "$SERVER_CONF_DIR not found, process aborted"
  exit 99
fi

if [ ! -d "$SECRETS_DIRECTORY" ]; then
  mkdir -p $SECRETS_DIRECTORY
fi

rm -f $SECRETS_DIRECTORY/*

openssl genrsa -out $SECRETS_DIRECTORY/ca.key 2048
openssl req -x509 -new -nodes -sha256 -key $SECRETS_DIRECTORY/ca.key -days 365 \
-out $SECRETS_DIRECTORY/ca.crt -subj /CN=admission_ca 2>&1
openssl genrsa -out $SECRETS_DIRECTORY/server.key 2048
openssl req -new -key $SECRETS_DIRECTORY/server.key -sha256 -out $SECRETS_DIRECTORY/server.csr \
-subj /CN=c7n-admission -config $SERVER_CONF_DIR/server.conf 2>&1
openssl x509 -req -days 365 -in $SECRETS_DIRECTORY/server.csr -sha256 -CA $SECRETS_DIRECTORY/ca.crt \
-CAkey $SECRETS_DIRECTORY/ca.key -CAcreateserial -out $SECRETS_DIRECTORY/server.crt -days 100000 -extensions \
v3_ext -extfile $SERVER_CONF_DIR/server.conf

# base64 -i $SECRETS_DIRECTORY/ca.crt > $SECRETS_DIRECTORY/ca-bundle.out
# base64 -i $SECRETS_DIRECTORY/server.crt > $SECRETS_DIRECTORY/tls-crt.out
# base64 -i $SECRETS_DIRECTORY/server.key > $SECRETS_DIRECTORY/tls-key.out
