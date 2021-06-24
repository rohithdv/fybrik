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
	mainPolicyManagerURL := "http://localhost:8080"
	//connectionTimeout, err := getConnectionTimeout()
	timeOutInSeconds := 120
	connectionTimeout := time.Duration(timeOutInSeconds) * time.Second

	policyManager, err := connectors.NewOpenApiPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		return
	}

	input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("Name_example"))} // []PolicymanagerRequest | input values that need to be considered for filter
	response, err := policyManager.GetPoliciesDecisions(&input[0])

	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPoliciesDecisions`: %v\n", response)
}
