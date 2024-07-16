package main

// Set the schema
nginxSvc: #Service

// Set the concrete values
nginxSvc: {
    apiVersion: "v1"
    kind:       "Service"
    metadata: {
        name:      "nginx"
        namespace: "default"
    }
    spec: {
        selector: "app.kubernetes.io/name": "nginx"
        ports: [{
            name:       "http"
            port:       80
            targetPort: 80
        }]
    }
}