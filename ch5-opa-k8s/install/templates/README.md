# Templates Directory

Templates in this directory are used to install OPA the hard way.

## Validating Admission Webhook Settings

The following are valid and invalid rule combinations for `ValidatingWebhookConfiguration` resource rules.

### Valid Settings

All these settings work.

```
    rules:
      - operations: ["CREATE", "UPDATE"]
        apiGroups: ["*"]
        apiVersions: ["v1"]
        resources: ["pods","deployments","services"]
```

```
    rules:
      - operations: ["CREATE", "UPDATE"]
        apiGroups: ["*"]
        apiVersions: ["*"]
        resources: ["pods"]
```

```
    rules:
      - operations: ["CREATE", "UPDATE"]
        apiGroups: ["*"]
        apiVersions: ["*"]
        resources: ["*"]
```

### Invalid Settings

> Note: For multiple resources, explicitly called-out, you should use explicit versions, like `v1`.

```
    rules:
      - operations: ["CREATE", "UPDATE"]
        apiGroups: ["*"]
        apiVersions: ["*"]
        resources: ["pods","deployments","services"]
```
