package main

import rego.v1

deny[msg] if {
	input[i].Cmd == "from"
	val := input[i].Value
	contains(val[i], data.denylist[_])

	msg = sprintf("unallowed image found %s", [val])
}