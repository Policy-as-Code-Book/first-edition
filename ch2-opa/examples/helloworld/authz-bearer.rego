package system.authz

import rego.v1

permissions := {
    "admin": {
        "path": "*"
    }
}

tokens := {
    "21ad4323-f187-4237-9b88-1e0aa6a4599d": {
        "roles": ["admin"]
    }
}

default allow := false

allow if {                      
    some permission
    lookup_permissions[permission] 
    permission.path == "*"            
}

# allow {                    
#     some permission
#     lookup_permissions[permission]        
#     permission.path == input.path      
# }

lookup_permissions[permission] if {             
    token := tokens[input.identity]  
    role := token.roles[_]           
    permission := permissions[role]         
}