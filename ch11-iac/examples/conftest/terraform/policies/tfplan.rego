package examples.ch11.terraform

import input as tfplan
import rego.v1

disallowed_iam_types := {"aws_iam_policy", "aws_iam_role"}

scan_resource_types := {"aws_iam_policy", "aws_iam_role"}

# scan_resource_types := { x | some x in disallowed_iam_types }

# tp := tfplan

default tfplan_allow := false

tfplan_allow if {
	count(violation_detected_iam) == 0
}

violation_detected_iam contains msg if {
	detected_iam_resources

	msg := sprintf("IAM resources cannot be used in the plan:, %q", [disallowed_iam_types])
}

detected_iam_resources if {
	some x in disallowed_iam_types
	resources[x]
}

####################
# Terraform Library
####################

# list of all resources of a given type
# resources[in_resource_type] := all if {
# 	some resource_type in scan_resource_types
# 	in_resource_type == resource_type
# 	all := [name |
# 		some change in tfplan.resource_changes

# 		# change := tfplan.resource_changes[_]
# 		change.type == in_resource_type
# 	]
# }

# list of all resources of a given type
resources[resource_type] := all if {
    some resource_type
    scan_resource_types[resource_type]
    all := [name |
        name:= tfplan.resource_changes[_]
        name.type == resource_type
    ]
}
