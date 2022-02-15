ROOT_DIR := ../../../..
DOCKER_NAME = write-asset-client

include $(ROOT_DIR)/hack/make-rules/docker.mk

include $(ROOT_DIR)/Makefile.env

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o write-asset-client
	docker build . -t $(IMG)
	rm write-asset-client

.PHONY: build
build:
	go build

.PHONY: run
run:
	nohup ./write-asset-client > output-write-asset-client.log &  2>&1 && sleep 10

.PHONY: terminate
terminate:

.PHONY: clean
clean:
	rm -f write-asset-client && rm -f output-write-asset-client.log

.PHONY: test
test:

-include ${ROOT_DIR}/.env
.EXPORT_ALL_VARIABLES:


.PHONY: deploy
deploy: $(TOOLBIN)/kustomize
	./deploy.sh

.PHONY: undeploy
undeploy:
	./deploy.sh undeploy
