package examples.ch11

import rego.v1

allowed_violations := 0

allowed_versions := {"1.26", "1.27"}

allowed_regions := {"us-east-1", "us-west-2"}

default allow := false

allow if {
	count(violation_versions) + count(violation_region) <= allowed_violations
}

violation_versions contains msg if {
	input.kind == "ClusterConfig"
	not input.metadata.version in allowed_versions

	msg := sprintf("invalid EKS version, version must be in %q", [allowed_versions])
}

violation_region contains msg if {
	input.kind == "ClusterConfig"
	not input.metadata.region in allowed_regions

	msg := sprintf("invalid region, region must be in %q", [allowed_regions])
}
