ROOT_DIR := ../../../..
DOCKER_NAME = plcy-mgr-mock
DOCKER_CONTEXT = $(ROOT_DIR)
DOCKER_FILE = $(DOCKER_CONTEXT)/test/services/pilot/policy-manager-mock/Dockerfile

include $(ROOT_DIR)/hack/make-rules/docker.mk
#include $(ROOT_DIR)/pkg/policy-compiler/policy-compiler.env

docker-all: docker-build docker-push

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o policy-manager-mock mockup_pilot.go
	docker build . -t $(IMG) -f Dockerfile
	rm policy-manager-mock

.PHONY: build
build:
	go build

.PHONY: run
run:
	nohup ./policy-manager-mock > output-policy-manager-mock.log &  2>&1 && sleep 10

.PHONY: terminate
terminate:

.PHONY: clean
clean:
	rm -f policy-manager-mock && rm -f output-policy-manager-mock.log

.PHONY: test
test:
