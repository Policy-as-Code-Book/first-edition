package examples.ch11.terraform

import rego.v1

allowed_eks_versions := {"1.26", "1.27"}

allowed_regions := {"us-east-1", "us-west-2"}

default allow := false

allow if {
	count(violation) == 0
}

violation contains msg if {
	not input.cluster_version in allowed_eks_versions

	msg := sprintf("invalid EKS version, version must be in %q", [allowed_eks_versions])
}

violation contains msg if {
	not input.region in allowed_regions

	msg := sprintf("invalid region, region must be in %q", [allowed_regions])
}
