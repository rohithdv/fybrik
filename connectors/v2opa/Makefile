ROOT_DIR := ../..
DOCKER_NAME = v2opaconnector

include $(ROOT_DIR)/Makefile.env
include $(ROOT_DIR)/hack/make-rules/docker.mk

KUBE_NAMESPACE ?= m4d-system

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o v2opaconnector main.go
	docker build . -t $(IMG)
	rm v2opaconnector


.PHONY: deploy
deploy: $(TOOLBIN)/kustomize
	./deploy.sh

.PHONY: undeploy
undeploy:
	./deploy.sh undeploy


-include ${ROOT_DIR}/.env
.EXPORT_ALL_VARIABLES:
