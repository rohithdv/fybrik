// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
	"google.golang.org/grpc"
)

var _ PolicyManager = (*grpcPolicyManager)(nil)

type grpcPolicyManager struct {
	name       string
	connection *grpc.ClientConn
	client     pb.PolicyManagerServiceClient
}

// NewGrpcPolicyManager creates a PolicyManager facade that connects to a GRPC service
// You must call .Close() when you are done using the created instance
func NewGrpcPolicyManager(name string, connectionURL string, connectionTimeout time.Duration) (PolicyManager, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()
	connection, err := grpc.DialContext(ctx, connectionURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("NewGrpcPolicyManager failed when connecting to %s", connectionURL))
	}
	return &grpcPolicyManager{
		name:       name,
		client:     pb.NewPolicyManagerServiceClient(connection),
		connection: connection,
	}, nil
}

func (m *grpcPolicyManager) GetPoliciesDecisions(in *openapiclient.PolicymanagerRequest, creds string) (*openapiclient.PolicymanagerResponse, error) {

	credentialPath := creds
	processingGeo := ""
	properties := make(map[string]string)
	appInfo := &pb.ApplicationDetails{ProcessingGeography: processingGeo, Properties: properties}

	datasetContextList := []*pb.DatasetContext{}
	datasetId := ""
	dataset := &pb.DatasetIdentifier{DatasetId: datasetId}
	destination := ""
	operation := &pb.AccessOperation{Type: pb.AccessOperation_READ, Destination: destination}
	datasetContext := &pb.DatasetContext{Dataset: dataset, Operation: operation}
	datasetContextList = append(datasetContextList, datasetContext)

	appContext := &pb.ApplicationContext{CredentialPath: credentialPath, AppInfo: appInfo, Datasets: datasetContextList}

	result, err := m.client.GetPoliciesDecisions(context.Background(), appContext)

	respResult := []openapiclient.PolicymanagerResponseResult{}
	decisionId := ""
	policyManagerResp := &openapiclient.PolicymanagerResponse{DecisionId: &decisionId, Result: respResult}

	return policyManagerResp, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
	// return result, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))

}

// func (m *grpcPolicyManager) GetPoliciesDecisions(ctx context.Context, in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
// 	result, err := m.client.GetPoliciesDecisions(ctx, in)
// 	return result, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
// }

// func (m *grpcPolicyManager) Close() error {
// 	return m.connection.Close()
// }
