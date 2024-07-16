package system.authz

import rego.v1

default allow := false

allow if {                      
    # some permission
    # lookup_permissions[permission] 
    # permission.path == "*"  
	claims.username == "jimmy"
	claims.owner == true 
	# signature == "cc8c68b50fb6fcc4b7e4024dae4716d9eefae49bc257d8fbc10bfc6e61761d40"         
}

bearer_token := t if {
	v := input.identity
	startswith(v, "Bearer ")
	t := substring(v, count("Bearer "), -1)
}

# lookup_permissions[permission] {             
# 	token := tokens[signature]  
#     role := token.roles[_]           
#     permission := permissions[role]         
# }

claims := payload if {
	io.jwt.verify_hs256(bearer_token, "7904c2ee-9e31-47df-aae8-ed6ae4d16502")
	[_, payload, _] := io.jwt.decode(bearer_token)
}

signature := sig if {
	io.jwt.verify_hs256(bearer_token, "7904c2ee-9e31-47df-aae8-ed6ae4d16502")
	[_, _, sig] := io.jwt.decode(bearer_token)
}

