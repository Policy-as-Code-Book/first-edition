#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  echo "Error $1 occurred on $2"
}

SECRETS_DIRECTORY="generated/secrets"

rm $SECRETS_DIRECTORY/*
