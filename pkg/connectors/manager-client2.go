package main

import (
	"fmt"
	"os"
	"time"

	connectors "github.com/mesh-for-data/mesh-for-data/pkg/connectors/clients"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

func main() {

	//mainPolicyManagerName := os.Getenv("MAIN_POLICY_MANAGER_NAME")
	mainPolicyManagerName := "OPEN API MANAGER"
	//mainPolicyManagerURL := os.Getenv("MAIN_POLICY_MANAGER_CONNECTOR_URL")
	mainPolicyManagerURL := "https://v2opaconnector:50050"
	//connectionTimeout, err := getConnectionTimeout()
	timeOutInSeconds := 120
	connectionTimeout := time.Duration(timeOutInSeconds) * time.Second

	policyManager, err := connectors.NewOpenApiPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		return
	}

	input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("{\"asset_id\": \"0bb3245e-e3ef-40b7-b639-c471bae4966c\", \"catalog_id\": \"503d683f-1d43-4257-a1a3-0ddf5e446ba5\"}"))} // []PolicymanagerRequest | input values that need to be considered for filter
	creds := "http://vault.m4d-system:8200/v1/kubernetes-secrets/wkc-creds?namespace=cp4d"
	input[0].Resource.Creds = &creds
	response, err := policyManager.GetPoliciesDecisions(&input[0])

	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPoliciesDecisions`: %v\n", response)
}
