package examples.ch2

import rego.v1

test_hello_world if {
	hello with input as {"message": "world"}
}

test_hello_goodbye if {
	hello with input as {"message": "planet"}
}

