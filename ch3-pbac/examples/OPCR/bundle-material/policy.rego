package pacbook.things

import rego.v1

default allow := false

# allow if {
# 	p = data.normal.stuff[_]
# 	input.method == p.method
# 	input.name == p.lead
# 	input.project == p.project
# }

allow if {
	some stuff in data.normal.stuff
	input.method == stuff.method
	input.name == stuff.author
	input.project == stuff.project
}