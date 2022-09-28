ROOT_DIR := ../..
DOCKER_NAME = test-connector

include $(ROOT_DIR)/Makefile.env
include $(ROOT_DIR)/hack/make-rules/docker.mk
include $(ROOT_DIR)/hack/make-rules/tools.mk
include $(ROOT_DIR)/hack/make-rules/version.mk

KUBE_NAMESPACE ?= fybrik-system
POLICY_FILES ?=

.PHONY: all
all: docker-build docker-push

.PHONY: docker-build
docker-build: source-build
ifdef FROM_IMAGE
	docker build . -t ${IMG} --build-arg image=${FROM_IMAGE}
	rm test-connector
else
	docker build . -t ${IMG}
	rm test-connector
endif

.PHONY: source-build
source-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build $(LDFLAGS) -o test-connector .

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -v ./...

.PHONY: opaserver
opaserver: $(TOOLBIN)/opa
	${TOOLBIN}/opa run --server ${POLICY_FILES}

.PHONY: deploy
deploy:
	helm install fybrik-testconnector $(ROOT_DIR)/charts/fybrik \
		--set manager.enabled=false --set coordinator.enabled=false --set worker.enabled=false \
		--set testConnector.enabled=true \
		--set testConnector.image=${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/$(DOCKER_NAME):${DOCKER_TAGNAME} \
		-n $(KUBE_NAMESPACE) \
		--wait

.PHONY: undeploy
undeploy:
	helm uninstall fybrik-testconnector -n $(KUBE_NAMESPACE)

-include ${ROOT_DIR}/.env
.EXPORT_ALL_VARIABLES:
