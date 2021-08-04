ROOT_DIR := ../../../..
DOCKER_NAME = ctlg-cred-mock
DOCKER_CONTEXT = $(ROOT_DIR)
DOCKER_FILE = $(DOCKER_CONTEXT)/test/services/pilot/catalog-credential-mock/Dockerfile

include $(ROOT_DIR)/hack/make-rules/docker.mk
#include $(ROOT_DIR)/pkg/policy-compiler/policy-compiler.env

docker-all: docker-build docker-push

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o catalog-credential-mock mockup_pilot_catalog_con.go
	docker build . -t $(IMG) -f Dockerfile
	rm catalog-credential-mock

.PHONY: build
build:
	go build

.PHONY: run
run:
	nohup ./catalog-credential-mock > output-catalog-credential-mock.log &  2>&1 && sleep 10

.PHONY: terminate
terminate:

.PHONY: clean
clean:
	rm -f catalog-credential-mock && rm -f output-catalog-credential-mock.log

.PHONY: test
test:
