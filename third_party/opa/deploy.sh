#!/usr/bin/env bash

set -e
set -x

NAMESPACE=m4d-system
TIMEOUT=8m
OPA_EXECUTABLE=../../hack/tools/bin/opa

check_valid_data_folder(){
    count=`ls -1 *.json 2>/dev/null | wc -l`
    count="$(echo -e "${count}" | tr -d '[:space:]')"
    local retVal=1
    if [ $count == 0 ]
    then
        retVal=0
    fi
    echo $retVal
}

check_valid_policy_folder(){
    count=`ls -1 *.rego 2>/dev/null | wc -l`
    count="$(echo -e "${count}" | tr -d '[:space:]')"
    local retVal=1
    if [ $count == 0 ]
    then
        retVal=0
    fi
    echo $retVal
}

unloadpolicy() {
    cd $1
    retVal=$(check_valid_policy_folder)
    if [ $retVal -eq 1 ];
    then
        policyfolder="${1##*/}"
        echo $policyfolder
        kubectl delete configmap $policyfolder --namespace=$NAMESPACE
    else
        echo "$1 is not a valid policy folder"
    fi
    cd -
}

unloaddata() {
    cd $1
    retVal=$(check_valid_data_folder)
    if [ $retVal -eq 1 ];
    then
        policydatafolder="${1##*/}"
        echo $policydatafolder
        kubectl delete configmap $policydatafolder --namespace=$NAMESPACE
    else
        echo "$1 is not a valid data folder"
    fi
    cd -
}

loadpolicy(){
    cd $1
    retVal=$(check_valid_policy_folder)
    if [ $retVal -eq 1 ];
    then
        policyfolder="${1##*/}"
        echo $policyfolder
        kubectl create configmap $policyfolder --from-file=./ --namespace=$NAMESPACE
        kubectl label configmap $policyfolder openpolicyagent.org/policy=rego --namespace=$NAMESPACE
    else
        echo "$1 is not a valid policy folder"
    fi
    cd -
}

loaddata(){
    cd $1
    retVal=$(check_valid_data_folder)
    if [ $retVal -eq 1 ];
    then
        policydatafolder="${1##*/}"
        echo $policydatafolder
        kubectl create configmap $policydatafolder --from-file=./ --namespace=$NAMESPACE
        kubectl label configmap $policydatafolder openpolicyagent.org/data=opa --namespace=$NAMESPACE
    else
        echo "$1 is not a valid data folder"
    fi
    cd -
}

validate_schema(){
   $OPA_EXECUTABLE eval  data.dataapi.authz.transform -i input-READ.json -d data-and-policies/user-created-policy-1/sample_policies.rego -d ../../charts/m4d/files/opa-server/policy-lib  -s data-and-policies/user-created-policy-1/schemas
}

case "$1" in
    validateschema)
        validate_schema
        ;;
    loadpolicy)
        loadpolicy "$2"
        ;;
    loaddata)
        loaddata "$2"
        ;;
    unloadpolicy)
        unloadpolicy "$2"
        ;;
    unloaddata)
        unloaddata "$2"
        ;;
    *)
        echo "usage: $0 [deploy|undeploy|loadpolicy <policydir>|loaddata <datadir>|unloadpolicy <policydir>|unloaddata <datadir>]"
        exit 1
        ;;
esac
