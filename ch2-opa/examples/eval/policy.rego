package example

import rego.v1

default hello := "goodbye"

hello := "world" if {
	msg := input.message
	msg in  data.msgs
}