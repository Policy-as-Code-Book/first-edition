package main

import rego.v1

bad_answer := ":N"

# deny_bad_answer
deny_bad_answer contains msg if {
	some x in input.response.items
	a := x.value
	endswith(a, bad_answer)
	msg = sprintf("Bad Answer: %q", [a])
}

