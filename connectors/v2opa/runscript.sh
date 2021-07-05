#!/usr/bin/env bash

set -e
set -x

kubectl apply -f v2-opa-connector-config.yaml -n m4d-system
make docker-build 
make docker-push 
make undeploy 
make deploy 
