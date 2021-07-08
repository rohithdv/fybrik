#!/usr/bin/env bash

set -e
set -x

POD=$(kubectl get pod  -n m4d-system -l app.kubernetes.io/component=opa-connector -o jsonpath="{.items[0].metadata.name}")
kubectl logs -n m4d-system $POD >out.log