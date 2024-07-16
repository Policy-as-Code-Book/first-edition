#!/usr/bin/env bash

# error handling
set -e
trap 'catch $? $LINENO' ERR
catch() {
  echo "Error $1 occurred on $2"
}

CONFIGS_DIRECTORY="generated/configs"

rm $CONFIGS_DIRECTORY/*
