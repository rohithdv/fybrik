package main

// func main() {
// 	input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("Name_example"))} // []PolicymanagerRequest | input values that need to be considered for filter

// 	configuration := openapiclient.NewConfiguration()
// 	api_client := openapiclient.NewAPIClient(configuration)
// 	resp, r, err := api_client.DefaultApi.GetPoliciesDecisions(context.Background()).Input(input).Execute()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPoliciesDecisions``: %v\n", err)
// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
// 	}
// 	// response from `GetPoliciesDecisions`: []PolicymanagerResponse
// 	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPoliciesDecisions`: %v\n", resp)
// }
