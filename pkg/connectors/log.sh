#!/usr/bin/env bash

set -e
set -x

POD=$(kubectl get pod -l app=manager-client -n m4d-system -o jsonpath="{.items[0].metadata.name}")
kubectl logs -n m4d-system $POD >out.log