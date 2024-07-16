# OPA AuthZ policy for debugging
package system.authz

import rego.v1

default allow := false

allow if {
	# ["body" "client_certificates" "headers" "identity" "method" "params" "path"]
	print("BODY: ", input.body)
	print("CLIENT CERTS: ", input.client_certificates)
	print("HEADERS: ", input.headers)
	print("IDENTITY: ", input.identity)
	print("METHOD: ", input.method)
	print("PARAMS: ", input.params)
	print("PATH: ", input.path)
	false
}
