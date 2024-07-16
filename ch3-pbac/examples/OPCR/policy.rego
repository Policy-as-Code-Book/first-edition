package project

import rego.v1

default allow := false

allow if {
	some project in data.projects
	# p = data.projects[_]
	input.method == project.method
	input.name == project.author
	input.project == project.project
}

