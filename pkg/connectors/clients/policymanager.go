// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"io"
	"log"
	"strings"

	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

// PolicyManager is an interface of a facade to connect to a policy manager.
type PolicyManager3 interface {
	pb.PolicyManagerServiceServer
	io.Closer
}

// PolicyManager is an interface of a facade to connect to a policy manager.
type PolicyManager interface {
	//pb.PolicyManagerServiceServer
	GetPoliciesDecisions(in *openapiclient.PolicymanagerRequest, creds string) (*openapiclient.PolicymanagerResponse, error)
	io.Closer
}

func MergePoliciesDecisions(in ...*pb.PoliciesDecisions) *pb.PoliciesDecisions {
	result := &pb.PoliciesDecisions{}

	for _, decisions := range in {
		result.ComponentVersions = append(result.ComponentVersions, decisions.ComponentVersions...)
		result.GeneralDecisions = append(result.GeneralDecisions, decisions.GeneralDecisions...)
		result.DatasetDecisions = append(result.DatasetDecisions, decisions.DatasetDecisions...)
	}

	result = compactPolicyDecisions(result)
	return result
}

// compactPolicyDecisions compacts policy decisions by merging decisions of same dataset identifier and same operation.
func compactPolicyDecisions(in *pb.PoliciesDecisions) *pb.PoliciesDecisions {
	if in == nil {
		return nil
	}

	result := &pb.PoliciesDecisions{
		ComponentVersions: in.ComponentVersions,
		DatasetDecisions:  []*pb.DatasetDecision{},
		GeneralDecisions:  compactOperationDecisions(in.GeneralDecisions),
	}

	// Group and flatten decisions by dataset id
	decisionsByIDKeys := []string{} // for determitistric results
	decisionsByID := map[string]*pb.DatasetDecision{}
	for _, datasetDecision := range in.DatasetDecisions {
		datasetID := datasetDecision.Dataset.DatasetId
		if _, exists := decisionsByID[datasetID]; !exists {
			decisionsByIDKeys = append(decisionsByIDKeys, datasetID)
			decisionsByID[datasetID] = &pb.DatasetDecision{
				Dataset: datasetDecision.Dataset,
			}
		}
		decisionsByID[datasetID].Decisions = append(decisionsByID[datasetID].Decisions, datasetDecision.Decisions...)
	}

	// Compact DatasetDecisions
	for _, key := range decisionsByIDKeys {
		datasetDecision := decisionsByID[key]
		result.DatasetDecisions = append(result.DatasetDecisions, &pb.DatasetDecision{
			Dataset:   datasetDecision.Dataset,
			Decisions: compactOperationDecisions(datasetDecision.Decisions),
		})
	}

	return result
}

func compactOperationDecisions(in []*pb.OperationDecision) []*pb.OperationDecision {
	if len(in) == 0 {
		return nil
	}

	type operationKeyType [2]interface{}

	// Group and flatten decisions for a specific dataset id by operation
	decisionsByOperationKeys := []operationKeyType{} // for determitistric results
	decisionsByOperation := map[operationKeyType]*pb.OperationDecision{}
	for _, operationDecision := range in {
		key := operationKeyType{operationDecision.Operation.Type, operationDecision.Operation.Destination}
		if _, exists := decisionsByOperation[key]; !exists {
			decisionsByOperationKeys = append(decisionsByOperationKeys, key)
			decisionsByOperation[key] = &pb.OperationDecision{
				Operation: operationDecision.Operation,
			}
		}
		decisionsByOperation[key].EnforcementActions = append(decisionsByOperation[key].EnforcementActions, operationDecision.EnforcementActions...)
		decisionsByOperation[key].UsedPolicies = append(decisionsByOperation[key].UsedPolicies, operationDecision.UsedPolicies...)
	}

	decisions := make([]*pb.OperationDecision, 0, len(decisionsByOperation))
	for _, key := range decisionsByOperationKeys {
		decision := decisionsByOperation[key]
		decisions = append(decisions, decision)
	}

	return decisions
}

func convertOpenApiReqToGrpcReq(in *openapiclient.PolicymanagerRequest, creds string) *pb.ApplicationContext {

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

	return appContext
}

func convGrpcRespToOpenApiResp(result *pb.PoliciesDecisions) *openapiclient.PolicymanagerResponse {

	// convert GRPC response to Open Api Response - start
	//decisionId := "3ffb47c7-c3c7-4fe7-b244-b38dc8951b87"
	// we dont get decision id returned from OPA from GRPC response. So we generate random hex string
	decisionId, _ := randomHex(20)
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
			//operation := decision.GetOperation()

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

	return policyManagerResp
}

func MergePoliciesDecisions2(in ...*openapiclient.PolicymanagerResponse) *openapiclient.PolicymanagerResponse {
	result := &openapiclient.PolicymanagerResponse{}
	decisionIdList := make([]string, 0)

	policyManagerRespResultArr := make([]openapiclient.PolicymanagerResponseResult, 0)

	for _, response := range in {
		decisionIdList = append(decisionIdList, response.GetDecisionId())
		policyManagerRespResult := response.GetResult()

		for i := 0; i < len(policyManagerRespResult); i++ {
			policyManagerRespResultArr = append(policyManagerRespResultArr, policyManagerRespResult[i])
		}
	}
	result.SetDecisionId(strings.Join(decisionIdList, ";"))
	result.SetResult(policyManagerRespResultArr)

	return result
}
