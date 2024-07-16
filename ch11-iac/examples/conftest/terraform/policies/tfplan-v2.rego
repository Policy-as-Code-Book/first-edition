package examples.ch11.terraform

import input as tfplan
import rego.v1

disallowed_iam_types := {"aws_iam_policy", "aws_iam_role"}

default tfplan_allow := false

tfplan_allow if {
	count(violation_detected_iam) == 0
}

violation_detected_iam contains msg if {
	detected_iam_resources

	msg := sprintf("IAM resources cannot be used in the plan:, %q", [disallowed_iam_types])
}

# contributed by srenatus (styra)
detected_iam_resources if {
	some x in disallowed_iam_types
	resources[x]
}

resources contains res.type if {
	some res in tfplan.resource_changes
	res.type in disallowed_iam_types
}
