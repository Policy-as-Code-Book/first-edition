package main

places_where_secrets_hide = [
    "secret",
    "apikey",
    "access"
]

# secrets bad
deny[msg] {    
    input[x].Cmd == "env"
    value := input[x].Value
    contains(lower(value[_]), places_where_secrets_hide[_])
    msg = sprintf("secrets found @ line %d",[x])
}

# latest bad
deny[msg] {
    input[x].Cmd == "from"
    value := split(input[x].Value[0], ":")
    contains(lower(value[1]), "latest")
    msg = sprintf("latest found @ line %d",[x])
}