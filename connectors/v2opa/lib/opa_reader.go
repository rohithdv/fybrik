// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

type OpaReader struct {
	opaServerURL string
}

func NewOpaReader(opasrvurl string) *OpaReader {
	return &OpaReader{opaServerURL: opasrvurl}
}

func (r *OpaReader) updatePolicyManagerRequestWithResourceInfo(in *openapiclient.PolicymanagerRequest, wkcMetadata map[string]interface{}) (*openapiclient.PolicymanagerRequest, error) {

	responseBytes, errJSON := json.MarshalIndent(wkcMetadata, "", "\t")
	if errJSON != nil {
		return nil, fmt.Errorf("error Marshalling External Catalog Connector Response: %v", errJSON)
	}

	var wkcJson interface{}
	err := json.Unmarshal(responseBytes, &wkcJson)
	if err != nil {
		return nil, fmt.Errorf("error UnMarshalling WKC Catalog Connector Response: %v", err)
	}
	if main, ok := wkcJson.(map[string]interface{}); ok {
		if details, ok := main["details"].(map[string]interface{}); ok {
			if metadata, ok := details["metadata"].(map[string]interface{}); ok {
				if datasetTags, ok := metadata["dataset_tags"].([]interface{}); ok {
					tagArr := make([]string, 0)
					for i := 0; i < len(datasetTags); i++ {
						tagVal := datasetTags[i].(string)
						tagArr = append(tagArr, tagVal)
					}
					log.Println("tagArr: ", tagArr)

					tagInReq := make(map[string]map[string]interface{})
					tagVal := make(map[string]interface{})
					for i := 0; i < len(tagArr); i++ {
						splitStr := strings.Split(tagArr[i], " = ")
						// residency = Turkey
						tagVal[splitStr[0]] = splitStr[1]
					}
					tagInReq["tags"] = tagVal
					resource := in.GetResource()
					(&resource).SetTags(tagInReq)
					in.SetResource(resource)
					log.Println("in.GetResource().GetTags(): ", (&resource).GetTags())
				}
				if componentsMetadata, ok := metadata["components_metadata"].(map[string]interface{}); ok {
					listofcols := []string{}
					listoftags := [][]string{}
					lstOfValueTags := []string{}
					for key, val := range componentsMetadata {
						log.Println("key :", key)
						log.Println("val :", val)
						listofcols = append(listofcols, key)

						if columnsMetadata, ok := val.(map[string]interface{}); ok {
							if tagsList, ok := columnsMetadata["tags"].([]interface{}); ok {
								lstOfValueTags = []string{}
								for _, tagElem := range tagsList {
									lstOfValueTags = append(lstOfValueTags, tagElem.(string))
								}
								listoftags = append(listoftags, lstOfValueTags)
							} else {
								lstOfValueTags = []string{}
								listoftags = append(listoftags, lstOfValueTags)
							}
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
							for j := 0; j < len(listoftags[i]); j++ {
								q[listoftags[i][j]] = "true"
							}
							p["tags"] = q
							newcol.SetTags(p)
						}
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
				}
			}
		}
	}
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

	catalogProviderName := getEnv("CATALOG_PROVIDER_NAME")
	if catalogProviderName == "WKC" {
		in, _ = r.updatePolicyManagerRequestWithResourceInfo(in, inputMap)
	} else {
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("processing metadata from non WKC catalog not supported yet (datasetID = %s): %v", datasetID, err)
	}
	b, err := json.Marshal(*in)
	if err != nil {
		fmt.Println(err)
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error during marshal in GetOPADecisions: %v", err)
	}
	inputJSON := "{ \"input\": " + string(b) + " }"
	fmt.Println("updated stringified policy manager request iun GetOPADecisions", inputJSON)

	opaEval, err := EvaluatePoliciesOnInput(inputJSON, r.opaServerURL, policyToBeEvaluated)
	if err != nil {
		log.Printf("error in EvaluatePoliciesOnInput : %v", err)
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in EvaluatePoliciesOnInput : %v", err)
	}
	log.Println("OPA Eval : " + opaEval)

	//operation := in.GetAction().ActionType
	policyManagerResponse := new(openapiclient.PolicymanagerResponse)
	err = json.Unmarshal([]byte(opaEval), &policyManagerResponse)
	if err != nil {
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in GetOPADecisions during unmarshalling OPA response to Policy Manager Response : %v", err)
	}
	log.Println("unmarshalled policyManagerResp in GetOPADecisions:", policyManagerResponse)

	res, err := json.MarshalIndent(policyManagerResponse, "", "\t")
	if err != nil {
		return openapiclient.PolicymanagerResponse{}, fmt.Errorf("error in GetOPADecisions during MarshalIndent Policy Manager Response : %v", err)
	}
	log.Println("Marshalled PolicyManagerResponse from OPA response in GetOPADecisions:", string(res))

	return *policyManagerResponse, nil
}

func buildNewEnforcementAction(transformAction interface{}) (*openapiclient.ActionOnColumns, bool) {
	log.Println("transformAction", transformAction)
	var actionOnColumns = new(openapiclient.ActionOnColumns)
	if result1, ok := transformAction.(map[string]interface{}); ok {
		log.Println("transformAction type :", reflect.TypeOf(result1))
		log.Println("result1[\"action\"].(string) :", result1["action"].(map[string]interface{}))
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

				return actionOnColumns, true

			case string(openapiclient.ENCRYPT_COLUMN):
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

			case string(openapiclient.REDACT_COLUMN):
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
			newUsedPolicy := description
			return &newUsedPolicy, true
		}
	}

	return nil, false
}
