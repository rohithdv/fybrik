// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	connectors "fybrik.io/fybrik/pkg/connectors/clients"
	pb "fybrik.io/fybrik/pkg/connectors/protobuf"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	return value
}

func constructInputParameters() *pb.ApplicationContext {
	timeoutinsecs := getEnv("CONNECTION_TIMEOUT")
	timeoutinseconds, err := strconv.Atoi(timeoutinsecs)
	if err != nil {
		log.Printf("Atoi conversion of timeoutinseconds failed: %v", err)
		return nil
	}

	fmt.Println("timeoutinseconds in MockupPilot: ", timeoutinseconds)

	catalogID := "<CATALOGID>"
	datasetID := "<ASSETID/DATASETID>"

	var datasetIDJson string
	if getEnv("CATALOG_PROVIDER_NAME") == "EGERIA" {
		datasetIDJson = "{\"ServerName\":\"mds1\", \"AssetGuid\":\"<ASSET_GUID>\"}"
	} else {
		datasetIDJson = "{\"catalog_id\":\"" + catalogID + "\",\"asset_id\":\"" + datasetID + "\"}"
	}

	// applicationDetails := &pb.ApplicationDetails{Purpose: "Fraud Detection", Role: "Data Scientist", ProcessingGeography: "Netherlands"}
	applicationDetails := &pb.ApplicationDetails{ProcessingGeography: "Netherlands", Properties: map[string]string{"intent": "Fraud Detection", "role": "Data Scientist"}}

	datasets := []*pb.DatasetContext{}
	datasets = append(datasets, createDatasetRead(datasetIDJson))
	datasets = append(datasets, createDatasetWrite(datasetIDJson))
	//datasets = append(datasets, createDatasetWrite(datasetIDJson))
	//datasets = append(datasets, createDatasetTransferFirst(datasetIDJson))
	// datasets = append(datasets, createDatasetTransferSecond(catalogID, datasetID))
	// datasets = append(datasets, createDatasetRead(catalogIDcos, datasetIDcos))
	// datasets = append(datasets, createDatasetRead(catalogIDDb2, datasetIDDb2))
	// datasets = append(datasets, createDatasetTransferFirst(catalogIDDb2, datasetIDDb2))
	// datasets = append(datasets, createDatasetTransferSecond(catalogIDDb2, datasetIDDb2))

	credential_path := "http://vault.fybrik-system:8200/v1/kubernetes-secrets/wkc-creds?namespace=cp4d"
	applicationContext := &pb.ApplicationContext{AppInfo: applicationDetails, Datasets: datasets, CredentialPath: credential_path}
	log.Printf("Sending Application Context: ")
	appContextStr, _ := json.MarshalIndent(applicationContext, "", "\t")
	log.Print(string(appContextStr))
	log.Println("1***************************************************************")

	return applicationContext
}

func main() {
	applicationContext := constructInputParameters()

	timeOutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeOut, err := strconv.Atoi(timeOutInSecs)
	connectionTimeout := time.Duration(timeOut) * time.Second
	mainPolicyManagerName := os.Getenv("MAIN_POLICY_MANAGER_NAME")
	//mainPolicyManagerURL := "opa-connector.m4d-system:80"
	mainPolicyManagerURL := os.Getenv("MAIN_POLICY_MANAGER_CONNECTOR_URL")
	policyManager, err := connectors.NewGrpcPolicyManager(mainPolicyManagerName, mainPolicyManagerURL, connectionTimeout)
	if err != nil {
		log.Println("returned with error ")
		log.Println("error in policyManager creation  %v", err)
		return
	}
	r, err := policyManager.GetPoliciesDecisions(context.Background(), applicationContext)

	if err != nil {
		errStatus, _ := status.FromError(err)
		fmt.Println("*********************************in error in  MockupPilot *****************************")
		fmt.Println("Message: ", errStatus.Message())
		fmt.Println("Code: ", errStatus.Code())

		// take specific action based on specific error?
		if codes.InvalidArgument == errStatus.Code() {
			fmt.Println("InvalidArgument in mockup pilot")
			return
		}
	} else {
		log.Printf("Response received from Policy Compiler below:")
		s, _ := json.MarshalIndent(r, "", "    ")
		log.Print(string(s))
		log.Println("2***************************************************************")
	}
}

func createDatasetRead(datasetIDJson string) *pb.DatasetContext {
	dataset := &pb.DatasetIdentifier{DatasetId: datasetIDJson}
	operation := &pb.AccessOperation{Type: pb.AccessOperation_READ}
	datasetContext := &pb.DatasetContext{Dataset: dataset, Operation: operation}
	return datasetContext
}

func createDatasetWrite(datasetIDJson string) *pb.DatasetContext {
	dataset := &pb.DatasetIdentifier{DatasetId: datasetIDJson}
	operation := &pb.AccessOperation{Type: pb.AccessOperation_WRITE}
	datasetContext := &pb.DatasetContext{Dataset: dataset, Operation: operation}
	return datasetContext
}

func createDatasetTransferFirst(datasetIDJson string) *pb.DatasetContext {
	dataset := &pb.DatasetIdentifier{DatasetId: datasetIDJson}
	operation := &pb.AccessOperation{Type: pb.AccessOperation_COPY, Destination: "US"}
	datasetContext := &pb.DatasetContext{Dataset: dataset, Operation: operation}
	return datasetContext
}

// func createDatasetTransferSecond(catalogID, datasetID string) *pb.DatasetContext {
// 	dataset := &pb.DatasetIdentifier{CatalogId: catalogID, DatasetId: datasetID}
// 	operation := &pb.AccessOperation{Type: pb.AccessOperation_COPY, Destination: "European Union"}
// 	datasetContext := &pb.DatasetContext{Dataset: dataset, Operation: operation}
// 	return datasetContext
// }
