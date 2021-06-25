ROOT_DIR := ../..
include $(ROOT_DIR)/Makefile.env
include $(ROOT_DIR)/hack/make-rules/tools.mk

DOCKER_NAME = manager-client
KUBE_NAMESPACE ?= m4d-system

docker-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o manager-client manager-client2.go
	docker build . -t $(IMG)
	rm manager-client

.PHONY: build
build: $(TOOLBIN)/protoc $(TOOLBIN)/protoc-gen-go
	rm -f protobuf/*
	PATH=$(ABSTOOLBIN) protoc -I protos/ protos/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:$(ROOT_DIR)/pkg/connectors/protobuf
