// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"os"
	"time"

	"emperror.dev/errors"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

var _ PolicyManager2 = (*openApiPolicyManager)(nil)

type openApiPolicyManager struct {
	name   string
	client *openapiclient.APIClient
}

// NewopenApiPolicyManager creates a PolicyManager facade that connects to a openApi service
// You must call .Close() when you are done using the created instance
func NewOpenApiPolicyManager(name string, connectionURL string, connectionTimeout time.Duration) (PolicyManager2, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	// defer cancel()
	// connection, err := openApi.DialContext(ctx, connectionURL, openApi.WithInsecure(), openApi.WithBlock())
	// if err != nil {
	// 	return nil, errors.Wrap(err, fmt.Sprintf("NewopenApiPolicyManager failed when connecting to %s", connectionURL))
	// }

	configuration := &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: openapiclient.ServerConfigurations{
			{
				URL:         connectionURL,
				Description: "No description provided",
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}
	api_client := openapiclient.NewAPIClient(configuration)

	return &openApiPolicyManager{
		name:   name,
		client: api_client,
	}, nil
}

func (m *openApiPolicyManager) GetPoliciesDecisions(in *openapiclient.PolicymanagerRequest) (*openapiclient.PolicymanagerResponse, error) {
	//input := []openapiclient.PolicymanagerRequest{*openapiclient.NewPolicymanagerRequest(*openapiclient.NewAction(openapiclient.ActionType("read")), *openapiclient.NewResource("Name_example"))} // []PolicymanagerRequest | input values that need to be considered for filter

	//input := []openapiclient.PolicymanagerRequest{*in}

	resp, r, err := m.client.DefaultApi.GetPoliciesDecisions(context.Background()).Input(*in).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPoliciesDecisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
	}
	// response from `GetPoliciesDecisions`: []PolicymanagerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPoliciesDecisions`: %v\n", resp)
	return &resp[0], nil
}

func (m *openApiPolicyManager) Close() error {
	return nil
}
