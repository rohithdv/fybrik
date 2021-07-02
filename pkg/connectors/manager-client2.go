package main

import (
	"log"
	"time"

	connectors "github.com/mesh-for-data/mesh-for-data/pkg/connectors/clients"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

func main() {

	//mainPolicyManagerName := os.Getenv("MAIN_POLICY_MANAGER_NAME")
	mainPolicyManagerName := "OPEN API MANAGER"
	//mainPolicyManagerURL := os.Getenv("MAIN_POLICY_MANAGER_CONNECTOR_URL")
	mainPolicyManagerURL := "http://v2opaconnector.m4d-system:50050"
	//connectionTimeout, err := getConnectionTimeout()
	timeOutInSeconds := 120
	connectionTimeout := time.Duration(timeOutInSeconds) * time.Second

	policyManager, err := connectors.NewOpenApiPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		return
	}

	creds := "http://vault.m4d-system:8200/v1/kubernetes-secrets/wkc-creds?namespace=cp4d"

	// input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("{\"asset_id\": \"0bb3245e-e3ef-40b7-b639-c471bae4966c\", \"catalog_id\": \"503d683f-1d43-4257-a1a3-0ddf5e446ba5\"}", creds))} // []PolicymanagerRequest | input values that need to be considered for filter

	input := openapiclient.NewPolicymanagerRequestWithDefaults()
	input.SetAction(*openapiclient.NewAction(openapiclient.ActionType("read")))
	input.SetResource(*openapiclient.NewResource("{\"asset_id\": \"0bb3245e-e3ef-40b7-b639-c471bae4966c\", \"catalog_id\": \"503d683f-1d43-4257-a1a3-0ddf5e446ba5\"}"))
	//input.SetRequestContext(openapiclient.RequestContext{})

	// input := openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("{\"asset_id\": \"0bb3245e-e3ef-40b7-b639-c471bae4966c\", \"catalog_id\": \"503d683f-1d43-4257-a1a3-0ddf5e446ba5\"}", creds))} // []PolicymanagerRequest | input values that need to be considered for filter

	log.Println("in manager-client - policy manager request: ", input)
	log.Println("in manager-client - creds: ", creds)

	response, err := policyManager.GetPoliciesDecisions(input, creds)

	bytes, _ := response.MarshalJSON()
	log.Println("in manager-client - Response from `policyManager.GetPoliciesDecisions`: \n", string(bytes))
}
