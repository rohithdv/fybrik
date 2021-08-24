// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

/*
 * Policy Manager Service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	opabl "fybrik.io/fybrik/connectors/open_policy_agent/lib"
	openapiclientmodels "fybrik.io/fybrik/pkg/taxonomy/model/base"
	"github.com/gin-gonic/gin"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	log.Printf("Env. variable extracted: %s - %s\n", key, value)
	return value
}

func constructPolicyManagerRequest(inputString string) *openapiclientmodels.PolicyManagerRequest {
	log.Println("inconstructPolicymanagerRequest")
	log.Println("inputString")
	log.Println(inputString)
	var input openapiclientmodels.PolicyManagerRequest
	json.Unmarshal([]byte(inputString), &input)
	log.Printf("input: %v", input)
	//resource := (&bird).GetResource()
	//fmt.Println(fmt.Sprintf("bird creds: %v", (&resource).GetCreds()))

	return &input
}

func displayReq(input2 openapiclientmodels.PolicyManagerRequest) error {
	bytes, err := input2.MarshalJSON()
	if err != nil {
		log.Println("err: ", err)
		return nil
	}
	log.Println("marshalled request:", string(bytes))
	return nil
}

func contactOPA(input openapiclientmodels.PolicyManagerRequest, creds string) (openapiclientmodels.PolicyManagerResponse, error) {

	catalogConnectorAddress := getEnv("CATALOG_CONNECTOR_URL")
	policyToBeEvaluated := "dataapi/authz/verdict"

	timeOutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeOut, err := strconv.Atoi(timeOutInSecs)

	if err != nil {
		return openapiclientmodels.PolicyManagerResponse{}, fmt.Errorf("conversion of timeOutinseconds failed: %v", err)
	}

	opaServerURL := getEnv("OPA_SERVER_URL")
	opaReader := opabl.NewOpaReader(opaServerURL)

	catalogReader := opabl.NewCatalogReader(catalogConnectorAddress, timeOut)
	eval, err := opaReader.GetOPADecisions(&input, creds, catalogReader, policyToBeEvaluated)
	if err != nil {
		log.Println("GetOPADecisions err:", err)
		return openapiclientmodels.PolicyManagerResponse{}, err
	}

	jsonOutput, err := json.MarshalIndent(eval, "", "\t")
	if err != nil {
		return openapiclientmodels.PolicyManagerResponse{}, fmt.Errorf("error during MarshalIndent of OPA decisions: %v", err)
	}
	log.Println("contactOPA: Received evaluation after execution of GetOPADecisions : " + string(jsonOutput))

	return eval, nil
}

// GetPoliciesDecisions - getPoliciesDecisions
func GetPoliciesDecisions(c *gin.Context) {
	log.Println("in GetPoliciesDecisions of V2 OPA Connector!")
	data, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("ctx.Request.body: %v", string(data))
	log.Println("creds value is", c.Request.Header["X-Request-Cred"][0])

	input2 := constructPolicyManagerRequest(string(data))
	log.Printf("input2:")
	displayReq(*input2)
	resp, err := contactOPA(*input2, c.Request.Header["X-Request-Cred"][0])
	if err != nil {
		log.Println("err: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, resp)
}
