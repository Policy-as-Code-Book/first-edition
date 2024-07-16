package examples.ch2

import rego.v1

default hello := ""

hello := output if {
    input.method == "POST"
    input.message in {"world", "planet"}
    output := build_return_msg("Hello",input.from)
}

build_return_msg (msg, from) := result if {
    result := sprintf("%s, %s", [msg, from])	
} 
