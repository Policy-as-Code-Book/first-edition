package pacbook.things

default allow = true

allow {
	p := data.pacbook.stuff[_]
	input.method == p.method
	input.name == p.lead
	input.project == p.project
}
