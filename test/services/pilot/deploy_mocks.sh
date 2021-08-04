#!/bin/bash
# Copyright 2020 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

set -x
set -e

: ${KUBE_NAMESPACE:=default}
: ${WITHOUT_VAULT=false}
: ${ROOT_DIR=../../../}

POLICY_DIR=$ROOT_DIR/pkg/policy-compiler
#source $POLICY_DIR/policy-compiler.env

kustomize_build() {
        local operation=$1
        local pass=$2
        local TEMP=$(mktemp -d)
        cp -aR deploy/base/* $TEMP
        cd $TEMP

        local images="ctlg-cred-mock plcy-mgr-mock"
        for image in ${images}; do \
                kustomize edit set image ${image}=${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${image}:${DOCKER_TAGNAME}; \
        done
        kustomize build . | kubectl $operation -f -  || $pass

        cd -
}

mocks_delete() {
        printf "\nRemoving kubectl mock resources on active cluster"
        #kubectl delete --ignore-not-found -f $ROOT_DIR/manager/config/prod/deployment_configmap.yaml
        #cd $ROOT_DIR/test/services/pilot/deploy
        #kustomize build patch/$REGISTRY | kubectl delete -f - || true
        kustomize_build delete true
        #cd -
        kubectl delete secret cloud-registry -n "${KUBE_NAMESPACE}" || true
        # kubectl delete -f  m4d-config.yaml -n "${KUBE_NAMESPACE}"
}

mocks_create() {
        #kubectl apply -f $ROOT_DIR/manager/config/prod/deployment_configmap.yaml
        #pwd
        #cd $ROOT_DIR/test/services/pilot/deploy
        #kustomize build patch/$REGISTRY | kubectl apply -f -
        kustomize_build apply false
}

kube_cluster_info() {
        printf "\nSleep ...\n"
        sleep 10

        printf "\nGetting current status of active cluster. Please wait...\n"
        kubectl cluster-info
        kubectl get all -n "${KUBE_NAMESPACE}" 

        printf "\n(TIP:) You can use the command \'watch kubectl get all\' to continuously monitor the cluster resources!\n"
        printf "\nThe deployment script has completed successfully!\n"
}

config(){
        kubectl create secret docker-registry cloud-registry  \
                  --docker-server="${DOCKER_HOSTNAME}" \
                  --docker-username="${DOCKER_USERNAME}" \
                  --docker-password="${DOCKER_PASSWORD}" \
                  -n "${KUBE_NAMESPACE}" || true

        kubectl patch serviceaccount default -p \
                  "{\"imagePullSecrets\": [{\"name\": \"cloud-registry\"}]}" \
                  -n "${KUBE_NAMESPACE}" || true

        # helm template --debug ../../../charts/m4d/ > out.txt
        # sed -n "/^\# Source: m4d\/templates\/m4d-config.yaml/,/---/p" out.txt | sed '1d;$d' > m4d-config.yaml
        # kubectl apply -f  m4d-config.yaml -n "${KUBE_NAMESPACE}"
        # rm out.txt
        # rm m4d-config.yaml
}

undeploy() {
        mocks_delete
        kube_cluster_info
}

deploy() {
        mocks_delete
        config
        mocks_create
        kube_cluster_info
}

case "$1" in
undeploy)
    undeploy
    ;;
*)
    deploy
    ;;
esac
