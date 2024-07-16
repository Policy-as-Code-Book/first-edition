package main

import rego.v1

allowed_violations := 0
allowed_messages := ["world","Powhatan"]

default allow := false
allow if {
  count(violation) <= allowed_violations
}

violation contains msg if {
  input.kind == "message"  
  not input.message in allowed_messages
  msg := sprintf("invalid message: %q",[allowed_messages])
}


