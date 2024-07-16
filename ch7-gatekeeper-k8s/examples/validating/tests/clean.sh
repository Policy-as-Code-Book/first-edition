#!/usr/bin/env bash
#set -o errexit

CMD_KUBECTL="kubectl"

${CMD_KUBECTL} delete ns policy-test 2>&1
${CMD_KUBECTL} delete ns policy-test1 2>&1
