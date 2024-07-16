#!/usr/bin/env bash

# error handling - DISABLED
# set -e
# trap 'catch $? $LINENO' ERR
# catch() {
#   if [ "$1" != "0" ]; then
#     echo "Error $1 occurred on $2"
#   fi
# }

echo ""
echo "----------------------Test policy with Kyverno test configuration"
kyverno test .

echo ""
echo "----------------------Apply policy against good/bad mocked resources"
kyverno apply --detailed-results registry-validate.yaml --resource resources.yaml -p

echo ""
echo "----------------------Apply policy against bad Pod manifest"
kyverno apply registry-validate.yaml --resource 2-test-pod.yaml -p

echo ""
echo "----------------------Apply policy against bad Deployment manifest"
kyverno apply registry-validate.yaml --resource 3-test-deploy.yaml -p

echo ""
echo "----------------------Apply policy against good Pod manifest"
kyverno apply registry-validate.yaml --resource 3-test-deploy.yaml -p

