// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
	pb "github.com/mesh-for-data/mesh-for-data/pkg/connectors/protobuf"
)

type OpaReader struct {
	opaServerURL string
}

func NewOpaReader(opasrvurl string) *OpaReader {
	return &OpaReader{opaServerURL: opasrvurl}
}

func (r *OpaReader) updatePolicyManagerRequestWithResourceInfo(in *openapiclient.PolicymanagerRequest, metadata map[string]interface{}) (*openapiclient.PolicymanagerRequest, error) {

	responseBytes, errJSON := json.MarshalIndent(metadata, "", "\t")
	if errJSON != nil {
		return nil, fmt.Errorf("error Marshalling External Catalog Connector Response: %v", errJSON)
	}

	// https://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects
	var f interface{}
	json.Unmarshal(responseBytes, &f)

	m := f.(map[string]interface{})
	log.Println("dataset_id", m["dataset_id"])

	f = m["details"]
	m = f.(map[string]interface{})

	f = m["metadata"]
	m = f.(map[string]interface{})

	f = m["components_metadata"]
	m = f.(map[string]interface{})

	listofcols := []string{}
	listoftags := [][]string{}
	lstOfValueTags := []string{}
	for key, val := range m {
		log.Println("key :", key)
		log.Println("val :", val)
		listofcols = append(listofcols, key)

		m = val.(map[string]interface{})
		if v, ok := m["tags"]; ok {
			l := v.([]interface{})

			lstOfValueTags = []string{}
			for _, l1 := range l {
				lstOfValueTags = append(lstOfValueTags, l1.(string))
			}
			listoftags = append(listoftags, lstOfValueTags)
		} else {
			lstOfValueTags = []string{}
			listoftags = append(listoftags, lstOfValueTags)
		}
	}
	log.Println("******** listofcols : *******", listofcols)
	log.Println("******** listoftags: *******", listoftags)

	cols := []openapiclient.ResourceColumns{}

	var newcol *openapiclient.ResourceColumns
	numOfCols := len(listofcols)
	numOfTags := 0
	for i := 0; i < numOfCols; i++ {
		newcol = new(openapiclient.ResourceColumns)
		newcol.SetName(listofcols[i])
		numOfTags = len(listoftags[i])
		if numOfTags > 0 {
			p := make(map[string]map[string]interface{})
			q := make(map[string]interface{})
			q[listofcols[i]] = listoftags[i]
			p["tags"] = q
			newcol.SetTags(p)
		}
		// if val, ok := inputMap2["tags"]; ok {
		// 	newcol.SetTags(val.(map[string]map[string]interface{}))
		// }
		cols = append(cols, *newcol)
	}
	log.Println("******** cols : *******")
	log.Println("cols=", cols)
	for i := 0; i < numOfCols; i++ {
		log.Println("cols=", cols[i].GetName())
		log.Println("cols=", cols[i].GetTags())
	}
	log.Println("******** in before: *******", *in)
	log.Println("******** res before: *******", in.Resource)
	res := in.Resource
	(&res).SetColumns(cols)
	in.SetResource(res)
	log.Println("******** res after: *******", res)
	log.Println("******** in after: *******", *in)

	log.Println("******** udpated policy manager resp object : *******")
	b, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("stringified policy manager request", string(b))
	log.Println("******** udpated policy manager resp object end : *******")
	time.Sleep(8 * time.Second)

	return in, nil
}

func (r *OpaReader) GetOPADecisions(in *openapiclient.PolicymanagerRequest, catalogReader *CatalogReader, policyToBeEvaluated string) (*openapiclient.PolicymanagerResponse, error) {

	datasetsMetadata, err := catalogReader.GetDatasetsMetadataFromCatalog(in)
	if err != nil {
		return nil, err
	}
	datasetID := in.GetResource().Name
	metadata := datasetsMetadata[datasetID]

	inputMap, ok := metadata.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error in unmarshalling dataset metadata (datasetID = %s): %v", datasetID, err)
	}

	in, _ = r.updatePolicyManagerRequestWithResourceInfo(in, inputMap)

	b, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("stringified policy manager request", string(b))
	inputJSON := "{ \"input\": " + string(b) + " }"

	opaEval, err := EvaluatePoliciesOnInput(inputJSON, r.opaServerURL, policyToBeEvaluated)
	if err != nil {
		log.Printf("error in EvaluatePoliciesOnInput : %v", err)
		return nil, fmt.Errorf("error in EvaluatePoliciesOnInput : %v", err)
	}
	log.Println("OPA Eval : " + opaEval)
	operation := in.GetAction().ActionType
	opaOperationDecision, err := GetOPAOperationDecision(opaEval, &operation)
	if err != nil {
		return nil, fmt.Errorf("error in GetOPAOperationDecision : %v", err)
	}
	return opaOperationDecision, nil
}

// Translate the evaluation received from OPA for (dataset, operation) into pb.OperationDecision
func GetOPAOperationDecision(opaEval string, operation *openapiclient.ActionType) (*openapiclient.PolicymanagerResponse, error) {
	resultInterface := make(map[string]interface{})
	err := json.Unmarshal([]byte(opaEval), &resultInterface)
	if err != nil {
		return nil, err
	}
	evaluationMap, ok := resultInterface["result"].(map[string]interface{})
	if !ok {
		return nil, errors.New("error in format of OPA evaluation (incorrect result map)")
	}

	// Now iterate over
	enforcementActions := make([]*pb.EnforcementAction, 0)
	usedPolicies := make([]*pb.Policy, 0)

	var policyManagerResp = new(openapiclient.PolicymanagerResponse)
	policyManagerResp.Result = make([]openapiclient.PolicymanagerResponseResult, 0)

	if evaluationMap["deny"] != nil {
		lstDeny, ok := evaluationMap["deny"].([]interface{})
		if !ok {
			return nil, errors.New("unknown format of deny content")
		}
		if len(lstDeny) > 0 {

			var action1 = new(openapiclient.Action1)
			action1.ActionOnDatasets = new(openapiclient.ActionOnDatasets)
			action1.ActionOnDatasets.SetName(openapiclient.DENY_ACCESS)

			result := openapiclient.PolicymanagerResponseResult{
				Action: *action1,
				Policy: "",
			}
			//newEnforcementAction := &pb.EnforcementAction{Name: "Deny", Id: "Deny-ID", Level: pb.EnforcementAction_DATASET, Args: map[string]string{}}
			//enforcementActions = append(enforcementActions, newEnforcementAction)

			for i, reason := range lstDeny {
				if reasonMap, ok := reason.(map[string]interface{}); ok {
					if newUsedPolicy, ok := buildNewPolicy(reasonMap["used_policy"]); ok {
						result.Policy = result.Policy + "; " + *newUsedPolicy
						continue
					}
				}
				log.Printf("Warning: unknown format of argument %d of lstDeny list. Skipping", i)
				continue
			}

			policyManagerResp.Result = append(policyManagerResp.Result, result)
		}
	}

	if evaluationMap["transform"] != nil {
		lstTransformations, ok := evaluationMap["transform"].([]interface{})
		if !ok {
			return nil, errors.New("unknown format of transform content")
		}
		for i, transformAction := range lstTransformations {

			newEnforcementAction, newUsedPolicy, ok := buildNewEnfrocementAction(transformAction)
			if !ok {
				return nil, errors.New("unknown format of transform action")
			}
			var action1 = new(openapiclient.Action1)
			action1.ActionOnColumns = newEnforcementAction
			result := openapiclient.PolicymanagerResponseResult{
				Action: *action1,
				Policy: "",
			}

			// enforcementActions = append(enforcementActions, newEnforcementAction)
			if newUsedPolicy == nil {
				log.Printf("Warning: empty used policy field for transformation %d", i)
			} else {
				//usedPolicies = append(usedPolicies, newUsedPolicy)
				result.Policy = result.Policy + "; " + *newUsedPolicy
			}
			policyManagerResp.Result = append(policyManagerResp.Result, result)
		}
	}

	if len(enforcementActions) == 0 { // allow action
		// newEnforcementAction := &pb.EnforcementAction{Name: "Allow", Id: "Allow-ID", Level: pb.EnforcementAction_DATASET, Args: map[string]string{}}
		// enforcementActions = append(enforcementActions, newEnforcementAction)

		var action1 = new(openapiclient.Action1)
		action1.ActionOnDatasets = new(openapiclient.ActionOnDatasets)
		action1.ActionOnDatasets.SetName(openapiclient.ALLOW_ACCESS)

		result := openapiclient.PolicymanagerResponseResult{
			Action: *action1,
			Policy: "",
		}
		policyManagerResp.Result = append(policyManagerResp.Result, result)

	}

	log.Println("enforcementActions: ", enforcementActions)
	log.Println("usedPolicies: ", usedPolicies)

	return policyManagerResp, nil
}

func buildNewEnfrocementAction(transformAction interface{}) (*openapiclient.ActionOnColumns, *string, bool) {
	if action, ok := transformAction.(map[string]interface{}); ok {
		newUsedPolicy, ok := buildNewPolicy(action["used_policy"])
		if !ok {
			log.Println("Warning: unknown format of used policy information. Skipping policy", action)
		}

		var actionOnColumns = new(openapiclient.ActionOnColumns)

		if result, ok := action["action_name"].(string); ok {
			switch result {
			case string(openapiclient.REMOVE_COLUMN):
				if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
					//newEnforcementAction := &pb.EnforcementAction{Name: "removed", Id: "removed-ID",
					//Level: pb.EnforcementAction_COLUMN, Args: map[string]string{"column_name": columnName}}
					actionOnColumns.SetName(openapiclient.REMOVE_COLUMN)
					actionOnColumns.SetColumns([]string{columnName})

					return actionOnColumns, newUsedPolicy, true
				}
			case string(openapiclient.ENCRYPT_COLUMN):
				if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
					// newEnforcementAction := &pb.EnforcementAction{Name: "encrypted", Id: "encrypted-ID",
					// 	Level: pb.EnforcementAction_COLUMN, Args: map[string]string{"column_name": columnName}}
					actionOnColumns.SetName(openapiclient.ENCRYPT_COLUMN)
					actionOnColumns.SetColumns([]string{columnName})
					return actionOnColumns, newUsedPolicy, true
				}
			case string(openapiclient.REDACT_COLUMN):
				if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
					// newEnforcementAction := &pb.EnforcementAction{Name: "redact", Id: "redact-ID",
					// 	Level: pb.EnforcementAction_COLUMN, Args: map[string]string{"column_name": columnName}}
					actionOnColumns.SetName(openapiclient.REDACT_COLUMN)
					actionOnColumns.SetColumns([]string{columnName})
					return actionOnColumns, newUsedPolicy, true
				}
			case string(openapiclient.PERIODIC_BLACKOUT):
				if monthlyDaysNum, ok := extractArgument(action["arguments"], "monthly_days_end"); ok {
					// newEnforcementAction := &pb.EnforcementAction{Name: "periodic_blackout", Id: "periodic_blackout-ID",
					// 	Level: pb.EnforcementAction_DATASET, Args: map[string]string{"monthly_days_end": monthlyDaysNum}}
					//return newEnforcementAction, newUsedPolicy, true
					actionOnColumns.SetName(openapiclient.PERIODIC_BLACKOUT)
					actionOnColumns.SetColumns([]string{monthlyDaysNum})
					return actionOnColumns, newUsedPolicy, true
				} else if yearlyDaysNum, ok := extractArgument(action["arguments"], "yearly_days_end"); ok {
					// newEnforcementAction := &pb.EnforcementAction{Name: "periodic_blackout", Id: "periodic_blackout-ID",
					// 	Level: pb.EnforcementAction_DATASET, Args: map[string]string{"yearly_days_end": yearlyDaysNum}}
					// return newEnforcementAction, newUsedPolicy, true
					actionOnColumns.SetName(openapiclient.PERIODIC_BLACKOUT)
					actionOnColumns.SetColumns([]string{yearlyDaysNum})
					return actionOnColumns, newUsedPolicy, true
				}
			default:
				log.Printf("Unknown Enforcement Action receieved from OPA")
			}
		}
	}
	return nil, nil, false
}

func extractArgument(arguments interface{}, argName string) (string, bool) {
	if argsMap, ok := arguments.(map[string]interface{}); ok {
		if value, ok := argsMap[argName].(string); ok {
			return value, true
		}
	}
	return "", false
}

func buildNewPolicy(usedPolicy interface{}) (*string, bool) {
	if policy, ok := usedPolicy.(map[string]interface{}); ok {
		//todo: add other fields that can be returned as part of the policy struct
		if description, ok := policy["description"].(string); ok {
			//newUsedPolicy := &pb.Policy{Description: description}
			newUsedPolicy := description
			return &newUsedPolicy, true
		}
	}

	return nil, false
}
