package main

import rego.v1

# latest bad
deny_latest contains msg if {
	some container in input.spec.template.spec.containers
	image := container.image
	endswith(lower(image), "latest")
	msg = sprintf("Image with latest found: %q", [image])
}

# no version bad
deny_no_version contains msg if {
	some container in input.spec.template.spec.containers
	image := container.image
	count(split(image, ":")) == 1
	msg = sprintf("Image with no version found: %q", [image])
}
