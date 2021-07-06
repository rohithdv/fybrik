#!/usr/bin/env bash

set -e
set -x

# cd ../../config/taxonomy
# ./codegen_openapi_client_server.sh
# pwd
# cd ../../pkg/connectors
# rm -rf out_go_client 
# cp -r ../../config/taxonomy/out_go_client .
# pwd
# cd out_go_client
# rm go.mod 
# rm go.sum 
# cd ..

make docker-build 
make docker-push 
make undeploy 
make deploy 

# POD=$(kubectl get pod -l app=manager-client -o jsonpath="{.items[0].metadata.name}")
# while [[ $(kubectl get pod $POD -o jsonpath='{.status.conditions[?(@.type=="s")].status}') == "True" ]]; do echo "waiting" && sleep 1; done

# SECONDS=0

# until [[ $SECONDS -gt $end ]] || [[ $(kubectl get jobs $job_name -o jsonpath='{.status.conditions[?(@.type=="Failed")].status}') == "True" ]] || [[ $(kubectl get jobs $job_name -o jsonpath='{.status.conditions[?(@.type=="Complete")].status}') == "True" ]]; do


