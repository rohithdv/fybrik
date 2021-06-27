#!/bin/bash
# Copyright 2020 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

set -x
set -e

: ${KUBE_NAMESPACE:=m4d-system}
: ${WITHOUT_VAULT=true}
: ${ROOT_DIR=../..}

v2_opa_connector_delete() {
        printf "\nRemoving kubectl resources on active cluster"
        $ROOT_DIR/hack/tools/bin/kustomize build . | kubectl delete -f - || true
}

v2_opa_connector_create() {
        $ROOT_DIR/hack/tools/bin/kustomize build . | kubectl apply -f -
}

config() {
        kubectl create secret docker-registry cloud-registry  \
                --docker-server="$DOCKER_HOSTNAME" \
                --docker-username="$DOCKER_USERNAME" \
                --docker-password="$DOCKER_PASSWORD" \
                -n $KUBE_NAMESPACE || true
        kubectl patch serviceaccount default -p \
                "{\"imagePullSecrets\": [{\"name\": \"cloud-registry\"}]}" \
                -n $KUBE_NAMESPACE || true
}

undeploy() {
        v2_opa_connector_delete
}

deploy() {
	config || true
        v2_opa_connector_create
}

case "$1" in
undeploy)
    undeploy
    ;;
*)
    deploy
    ;;
esac
