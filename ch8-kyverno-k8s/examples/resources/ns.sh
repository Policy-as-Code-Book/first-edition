kubectl apply -f - <<EOF
apiVersion: v1
kind: Namespace
metadata:
  name: policy-test
  labels:    
    pod-security.kubernetes.io/enforce: privileged
    pod-security.kubernetes.io/audit: privileged
    pod-security.kubernetes.io/warn: privileged
EOF