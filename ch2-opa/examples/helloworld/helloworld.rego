package examples.ch2

import rego.v1

default hello := "goodbye"

hello := "world" if {
	msg := input.message
	msg == "world"
}

