package main

import rego.v1

places_where_secrets_hide := [
    "secret",
    "apikey",
    "access"
]

bad_adds := [
    "example.com"
]

bad_args := [
    "example.com"
]

bad_runs := [
    "wget",
    "curl"
]

maybe_bad_runs := [
    "apk",
    "yum"
]

# secrets bad
deny[msg] if {    
    input[x].Cmd == "env"
    value := input[x].Value
    contains(lower(value[_]), places_where_secrets_hide[_])
    msg = sprintf("secrets found @ line %d",[x])
}

# latest bad
deny[msg] if {
    input[x].Cmd == "from"
    value := split(input[x].Value[0], ":")
    contains(lower(value[1]), "latest")
    msg = sprintf("FROM with latest found @ line %d",[x])
}

# add potentially bad
warn[msg] if {
    input[x].Cmd == "add"
    # value := split(input[x].Value[0], ":")
    # contains(lower(value[1]), "latest")
    msg = sprintf("ADD found @ line %d",[x])
}

# run apk potentially bad
warn[msg] if {
    input[x].Cmd == "run"
    value := input[x].Value[0]
    contains(lower(value), "apk")
    msg = sprintf("RUN with potentially bad commands, %q, found @ line %d",[maybe_bad_runs,x])
}

# run with curl/wget bad
deny[msg] if {
    input[x].Cmd == "run"
    value := input[x].Value[0]
    contains(lower(value), bad_runs[_])
    msg = sprintf("RUN with bad commands, %q, found @ line %d",[bad_runs,x])
}

# add with example.com bad
deny[msg] if {
    input[x].Cmd == "add"
    value := input[x].Value[0]
    contains(lower(value), bad_adds[_])
    msg = sprintf("ADD with bad source, %q, found @ line %d",[bad_adds,x])
}

# example.com bad
deny[msg] if {
    input[x].Cmd == "arg"
    value := input[x].Value[0]
    contains(lower(value), bad_args[_])
    msg = sprintf("ARG with bad args, %q, found @ line %d",[bad_args,x])
}