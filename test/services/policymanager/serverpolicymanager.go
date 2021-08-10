// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	mockup "fybrik.io/fybrik/manager/controllers/mockup"
	openapiclientmodels "fybrik.io/fybrik/pkg/taxonomy/model/base"
	"github.com/gin-gonic/gin"
)

const (
	PORT = 50082
)

var router *gin.Engine

func constructPolicyManagerRequest(inputString string) *openapiclientmodels.PolicyManagerRequest {
	fmt.Println("inconstructPolicymanagerRequest")
	fmt.Println("inputString")
	fmt.Println(inputString)
	var input openapiclientmodels.PolicyManagerRequest
	err := json.Unmarshal([]byte(inputString), &input)
	if err != nil {
		return nil
	}
	fmt.Println("input:", input)
	return &input
}

func main() {
	router = gin.Default()

	router.POST("/getPoliciesDecisions", func(c *gin.Context) {
		creds := ""
		if values := c.Request.Header["X-Request-Cred"]; len(values) > 0 {
			creds = values[0]
		}
		fmt.Println("creds extracted from POST request:", creds)
		input, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println("input extracted from POST request body:", string(input))

		policyManagerReq := constructPolicyManagerRequest(string(input))
		policyManager := &mockup.MockPolicyManager{}
		policyManagerResp, err := policyManager.GetPoliciesDecisions(policyManagerReq, creds)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error in GetPoliciesDecisions!")
			return
		}
		c.JSON(http.StatusOK, policyManagerResp)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	log.Fatal(router.Run(":" + strconv.Itoa(PORT)))
}
