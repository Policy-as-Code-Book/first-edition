.DEFAULT_GOAL := validate

.PHONY:	validate
validate:
	docker run --rm -v $(pwd):/ssc openpolicyagent/conftest:v0.47.0 test \
	-p /ssc/policy/docker/docker.rego /ssc/configs/DockerFile -o github
