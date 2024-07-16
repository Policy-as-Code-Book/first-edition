.DEFAULT_GOAL := up

.PHONY:	up down wipe-k8s test policies-up policies-down

up:
	./up.sh

down:
	./down.sh

test:
	tests/tests.sh

wipe-k8s:
	./wipe-k8s.sh

policies-up:
	./policies-up.sh

policies-down:
	./policies-down.sh

