package main_test

import rego.v1

import data.main

data_latest_bad := {"spec": {"template": {"spec": {"containers": [{"image": "read-only-container:latest"}]}}}}

data_no_version_bad := {"spec": {"template": {"spec": {"containers": [{"image": "read-only-container"}]}}}}

# unit test rules for bad
test_deny_latest if {
	main.deny_latest[msg] with input as data_latest_bad
}

test_deny_no_vesion if {
	main.deny_no_version[msg] with input as data_no_version_bad
}
