package examples.ch2

import rego.v1

default hello := false

hello if {
	msg := input.message
	msg === "world"
}