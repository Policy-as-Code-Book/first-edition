package model

const JsonProviderErrorMessage = `{
    "apiVersion": "externaldata.gatekeeper.sh/v1alpha1",
    "kind": "ProviderResponse",
    "response": {
        "idempotent": true,
        "items": [
            {
                "key": "message",
                "value": "Cannot predict now."
            },
            {
                "key": "code",
                "value": "M"
            }
        ],"systemError":"internal error"
    }
}`
