// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"emperror.dev/errors"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
	"google.golang.org/grpc"
)

var _ PolicyManager2 = (*grpcPolicyManager)(nil)

type grpcPolicyManager struct {
	name       string
	connection *grpc.ClientConn
	client     pb.PolicyManagerServiceClient
}

// ref: https://sosedoff.com/2014/12/15/generate-random-hex-string-in-go.html
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// NewGrpcPolicyManager creates a PolicyManager facade that connects to a GRPC service
// You must call .Close() when you are done using the created instance
func NewGrpcPolicyManager(name string, connectionURL string, connectionTimeout time.Duration) (PolicyManager2, error) {
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

	// convert GRPC response to Open Api Response - start
	//decisionId := "3ffb47c7-c3c7-4fe7-b244-b38dc8951b87"
	// we dont get decision id returned from OPA from GRPC response. So we generate random hex string
	decisionId, err := randomHex(20)
	log.Println("decision id generated", decisionId)

	var datasetDecisions []*pb.DatasetDecision
	var decisions []*pb.OperationDecision
	datasetDecisions = result.GetDatasetDecisions()
	respResult := []openapiclient.PolicymanagerResponseResult{}

	// we assume only one dataset decision is passed
	for i := 0; i < len(datasetDecisions); i++ {
		datasetDecision := datasetDecisions[i]
		//datasetID := datasetDecision.GetDataset()
		decisions = datasetDecision.GetDecisions()

		for j := 0; j < len(decisions); i++ {
			decision := decisions[j]
			operation = decision.GetOperation()

			var enfActionList []*pb.EnforcementAction
			var usedPoliciesList []*pb.Policy
			enfActionList = decision.GetEnforcementActions()
			usedPoliciesList = decision.GetUsedPolicies()

			for k := 0; k < len(enfActionList); k++ {
				enfAction := enfActionList[k]
				name := enfAction.GetName()
				level := enfAction.GetLevel()
				args := enfAction.GetArgs()
				policyManagerResult := openapiclient.PolicymanagerResponseResult{}

				if level == pb.EnforcementAction_COLUMN {
					actionOnCols := openapiclient.ActionOnColumns{}
					if name == "redact" {
						actionOnCols.SetName(openapiclient.REDACT_COLUMN)
						actionOnCols.SetColumns([]string{args["columns_name"]})
					}
					if name == "encrypted" {
						actionOnCols.SetName(openapiclient.ENCRYPT_COLUMN)
						actionOnCols.SetColumns([]string{args["columns_name"]})
					}
					if name == "removed" {
						actionOnCols.SetName(openapiclient.REMOVE_COLUMN)
						actionOnCols.SetColumns([]string{args["columns_name"]})
					}
					policyManagerResult.SetAction(
						openapiclient.ActionOnColumnsAsAction1(&actionOnCols))
				}

				if level == pb.EnforcementAction_DATASET {
					actionOnDataset := openapiclient.ActionOnDatasets{}
					if name == "Deny" {
						actionOnDataset.SetName(openapiclient.DENY_ACCESS)
					}
					policyManagerResult.SetAction(
						openapiclient.ActionOnDatasetsAsAction1(&actionOnDataset))
				}
				policy := usedPoliciesList[k].GetDescription()
				log.Println("usedPoliciesList[k].GetDescription()", policy)
				policyManagerResult.SetPolicy(policy)

				respResult = append(respResult, policyManagerResult)
			}
		}
	}
	// convert GRPC response to Open Api Response - end
	policyManagerResp := &openapiclient.PolicymanagerResponse{DecisionId: &decisionId, Result: respResult}

	return policyManagerResp, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
	// return result, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
}

// func (m *grpcPolicyManager) GetPoliciesDecisions(ctx context.Context, in *pb.ApplicationContext) (*pb.PoliciesDecisions, error) {
// 	result, err := m.client.GetPoliciesDecisions(ctx, in)
// 	return result, errors.Wrap(err, fmt.Sprintf("get policies decisions from %s failed", m.name))
// }

func (m *grpcPolicyManager) Close() error {
	return m.connection.Close()
}
