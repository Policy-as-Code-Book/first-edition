package main

#Service: {
    apiVersion: string
    kind:       string
    metadata: {
        name:      string
        namespace: string
    }
    spec: {
        selector: [string]: string
        ports: [{
            name:       string
            port:       int
            targetPort: int | string
        }]
    }
}
