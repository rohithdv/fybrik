// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
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
	//time.Sleep(8 * time.Second)

	return in, nil
}

func (r *OpaReader) GetOPADecisions(in *openapiclient.PolicymanagerRequest, creds string, catalogReader *CatalogReader, policyToBeEvaluated string) (openapiclient.PolicymanagerResponse, error) {

	datasetsMetadata, err := catalogReader.GetDatasetsMetadataFromCatalog(in, creds)
	if err != nil {
		return openapiclient.PolicymanagerResponse{}, err
	}
	datasetID := in.GetResource().Name
	metadata := datasetsMetadata[datasetID]

	inputMap, ok := metadata.(map[string]interface{})
	if !ok {
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in unmarshalling dataset metadata (datasetID = %s): %v", datasetID, err)
	}

	in, _ = r.updatePolicyManagerRequestWithResourceInfo(in, inputMap)
	b, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err)
		return openapiclient.PolicymanagerResponse{}, err
	}
	inputJSON := "{ \"input\": " + string(b) + " }"
	fmt.Println("updated stringified policy manager request iun GetOPADecisions", inputJSON)

	opaEval, err := EvaluatePoliciesOnInput(inputJSON, r.opaServerURL, policyToBeEvaluated)
	if err != nil {
		log.Printf("error in EvaluatePoliciesOnInput : %v", err)
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in EvaluatePoliciesOnInput : %v", err)
	}
	log.Println("OPA Eval : " + opaEval)

	operation := in.GetAction().ActionType
	policyManagerResponse, err := ConvertOpaResponseToPolicymanagerResponse(opaEval, &operation)
	if err != nil {
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in ConvertOpaResponseToPolicymanagerResponse : %v", err)
	}

	log.Println("policyManagerResponse : ", policyManagerResponse)
	return policyManagerResponse, nil
}

// Translate the evaluation received from OPA for (dataset, operation) into pb.OperationDecision
func ConvertOpaResponseToPolicymanagerResponse(opaEval string, operation *openapiclient.ActionType) (openapiclient.PolicymanagerResponse, error) {
	resultInterface := make(map[string]interface{})
	err := json.Unmarshal([]byte(opaEval), &resultInterface)
	if err != nil {
		return openapiclient.PolicymanagerResponse{}, err
	}
	evaluationMap, ok := resultInterface["result"].(map[string]interface{})
	if !ok {
		return openapiclient.PolicymanagerResponse{}, errors.New("error in format of OPA evaluation (incorrect result map)")
	}

	var policyManagerResp = new(openapiclient.PolicymanagerResponse)
	policyManagerResp.Result = make([]openapiclient.PolicymanagerResponseResult, 0)

	if evaluationMap["deny"] != nil {
		lstDeny, ok := evaluationMap["deny"].([]interface{})
		if !ok {
			return openapiclient.PolicymanagerResponse{}, errors.New("unknown format of deny content")
		}
		if len(lstDeny) > 0 {
			//for i, reason := range lstDeny {
			for i := 0; i < len(lstDeny); i++ {
				var action1 = new(openapiclient.Action1)
				action1.ActionOnDatasets = new(openapiclient.ActionOnDatasets)
				action1.ActionOnDatasets.SetName(openapiclient.DENY_ACCESS)

				result := openapiclient.PolicymanagerResponseResult{
					Action: *action1,
					Policy: "",
				}

				log.Println("lstDeny[i]", lstDeny[i])
				log.Println("lstDeny[i]", lstDeny[i].(string))
				// Declared an empty map interface
				var result1 map[string]interface{}
				json.Unmarshal([]byte(lstDeny[i].(string)), &result1)
				log.Println("result1 type in ConvertOpaResponseToPolicymanagerResponse", reflect.TypeOf(result1))
				result.Policy = result1["policy"].(string)

				log.Println("lstDeny[i]", lstDeny[i])
				policyManagerResp.Result = append(policyManagerResp.Result, result)
			}
		}
	}

	if evaluationMap["transform"] != nil {
		lstTransformations, ok := evaluationMap["transform"].([]interface{})
		if !ok {
			return openapiclient.PolicymanagerResponse{}, errors.New("unknown format of transform content")
		}
		for i, transformAction := range lstTransformations {

			newEnforcementAction, ok := buildNewEnforcementAction(transformAction)
			if !ok {
				return openapiclient.PolicymanagerResponse{}, errors.New("unknown format of transform action")
			}
			var action1 = new(openapiclient.Action1)
			action1.ActionOnColumns = newEnforcementAction
			result := openapiclient.PolicymanagerResponseResult{
				Action: *action1,
				Policy: "",
			}

			log.Println("lstTransformations[i]", lstTransformations[i])
			log.Println("lstTransformations[i]", lstTransformations[i].(string))
			// Declared an empty map interface
			var result1 map[string]interface{}
			json.Unmarshal([]byte(lstTransformations[i].(string)), &result1)
			log.Println("result1 type in ConvertOpaResponseToPolicymanagerResponse transform part ", reflect.TypeOf(result1))
			result.Policy = result1["policy"].(string)

			// if newUsedPolicy == nil {
			// 	log.Printf("Warning: empty used policy field for transformation %d", i)
			// } else {
			// 	result.Policy = result.Policy + "; " + *newUsedPolicy
			// }
			policyManagerResp.Result = append(policyManagerResp.Result, result)
		}
	}

	if len(policyManagerResp.Result) == 0 { // allow action
		policyManagerResp.Result = append(policyManagerResp.Result, openapiclient.PolicymanagerResponseResult{})
	}

	log.Println("*policyManagerResp: ", *policyManagerResp)
	return *policyManagerResp, nil
}

func buildNewEnforcementAction(transformAction interface{}) (*openapiclient.ActionOnColumns, bool) {
	log.Println("transformAction", transformAction)
	log.Println("transformAction", transformAction.(string))
	// Declared an empty map interface
	var result1 map[string]interface{}
	json.Unmarshal([]byte(transformAction.(string)), &result1)
	log.Println("transformAction type :", reflect.TypeOf(result1))
	log.Println("result1[\"action\"].(string) :", result1["action"].(map[string]interface{}))

	//if action, ok := transformAction.(map[string]interface{}); ok {
	// newUsedPolicy, ok := buildNewPolicy(action["used_policy"])
	// if !ok {
	// 	log.Println("Warning: unknown format of used policy information. Skipping policy", action)
	// }

	var actionOnColumns = new(openapiclient.ActionOnColumns)

	if result, ok := result1["action"].(map[string]interface{}); ok {
		res1 := result["name"].(string)
		switch res1 {
		case string(openapiclient.REMOVE_COLUMN):
			actionOnColumns.SetName(openapiclient.REMOVE_COLUMN)
			log.Println("Name:", openapiclient.REMOVE_COLUMN)

			resCols := result["columns"].([]interface{})
			log.Println("resCols", resCols)
			lstOfCols := []string{}
			for i := 0; i < len(resCols); i++ {
				lstOfCols = append(lstOfCols, resCols[i].(string))
			}
			log.Println("lstOfCols", lstOfCols)
			actionOnColumns.SetColumns(lstOfCols)

			//if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
			return actionOnColumns, true

		case string(openapiclient.ENCRYPT_COLUMN):
			//if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
			actionOnColumns.SetName(openapiclient.ENCRYPT_COLUMN)
			log.Println("Name:", openapiclient.ENCRYPT_COLUMN)

			resCols := result["columns"].([]interface{})
			log.Println("resCols", resCols)
			lstOfCols := []string{}
			for i := 0; i < len(resCols); i++ {
				lstOfCols = append(lstOfCols, resCols[i].(string))
			}
			log.Println("lstOfCols", lstOfCols)
			actionOnColumns.SetColumns(lstOfCols)

			return actionOnColumns, true
			//}

		case string(openapiclient.REDACT_COLUMN):
			//if columnName, ok := extractArgument(action["arguments"], "column_name"); ok {
			actionOnColumns.SetName(openapiclient.REDACT_COLUMN)
			log.Println("Name:", openapiclient.REDACT_COLUMN)

			resCols := result["columns"].([]interface{})
			log.Println("resCols", resCols)
			lstOfCols := []string{}
			for i := 0; i < len(resCols); i++ {
				lstOfCols = append(lstOfCols, resCols[i].(string))
			}
			log.Println("lstOfCols", lstOfCols)
			actionOnColumns.SetColumns(lstOfCols)

			return actionOnColumns, true
			//}

		case string(openapiclient.PERIODIC_BLACKOUT):
			//if monthlyDaysNum, ok := extractArgument(action["arguments"], "monthly_days_end"); ok {
			actionOnColumns.SetName(openapiclient.PERIODIC_BLACKOUT)
			log.Println("Name:", openapiclient.PERIODIC_BLACKOUT)

			resCols := result["columns"].([]interface{})
			log.Println("resCols", resCols)
			lstOfCols := []string{}
			for i := 0; i < len(resCols); i++ {
				lstOfCols = append(lstOfCols, resCols[i].(string))
			}
			log.Println("lstOfCols", lstOfCols)
			actionOnColumns.SetColumns(lstOfCols)

			return actionOnColumns, true
			//}
			//else if yearlyDaysNum, ok := extractArgument(action["arguments"], "yearly_days_end"); ok {
			// actionOnColumns.SetName(openapiclient.PERIODIC_BLACKOUT)
			// actionOnColumns.SetColumns(result["columns"].([]string))
			// return actionOnColumns, true
			//}

		default:
			log.Printf("Unknown Enforcement Action receieved from OPA")
		}
	}
	return nil, false
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
	log.Println("in buildNewPolicy")
	if policy, ok := usedPolicy.(map[string]interface{}); ok {
		//todo: add other fields that can be returned as part of the policy struct
		if description, ok := policy["policy"].(string); ok {
			//newUsedPolicy := &pb.Policy{Description: description}
			newUsedPolicy := description
			return &newUsedPolicy, true
		}
	}

	return nil, false
}
