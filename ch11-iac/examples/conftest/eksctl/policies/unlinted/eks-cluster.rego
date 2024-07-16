package examples.ch11

import future.keywords.in

allowed_violations := 0
allowed_versions := ["1.26","1.27"]
allowed_regions := ["us-east-1","us-west-2"]

default allow := false
allow {
  count(violation) <= allowed_violations
}

violation[msg] {
  input.kind == "ClusterConfig"
  not input.metadata.version in allowed_versions

  msg := sprintf("invalid EKS version, version must be in %q", [allowed_versions])
}

violation[msg] {
  input.kind == "ClusterConfig"
  not input.metadata.region in allowed_regions

  msg := sprintf("invalid region, region must be in %q", [allowed_regions])
}
