// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0​
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	//"github.com/gogo/protobuf/proto"

	pb "fybrik.io/fybrik/pkg/connectors/protobuf"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Env Variable %v not defined", key)
	}
	return value
}

func GetMetadata(datasetID string) error {
	catalogConnectorURL := getEnv("CATALOG_CONNECTOR_URL")
	catalogProviderName := getEnv("CATALOG_PROVIDER_NAME")

	timeoutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeoutInSeconds, err := strconv.Atoi(timeoutInSecs)
	if err != nil {
		log.Printf("Atoi conversion of timeoutinseconds failed: %v", err)
		return errors.Wrap(err, "Atoi conversion of timeoutinseconds failed")
	}

	fmt.Println("timeoutInSeconds: ", timeoutInSeconds)
	fmt.Println("catalogConnectorURL: ", catalogConnectorURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInSeconds)*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, catalogConnectorURL, grpc.WithInsecure())

	if err != nil {
		log.Printf("Connection to "+catalogProviderName+" Catalog Connector failed: %v", err)
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
		return errors.Wrap(err, "Connection to "+catalogProviderName+" Catalog Connector failed")
	}
	defer conn.Close()

	c := pb.NewDataCatalogServiceClient(conn)

	credential_path := "http://vault.fybrik-system:8200/v1/kubernetes-secrets/wkc-creds?namespace=cp4d"
	objToSend := &pb.CatalogDatasetRequest{DatasetId: datasetID, CredentialPath: credential_path}

	log.Printf("Sending CatalogDatasetRequest: ")
	catDataReqStr, _ := json.MarshalIndent(objToSend, "", "\t")
	log.Print(string(catDataReqStr))
	log.Println("1***************************************************************")

	log.Printf("Sending request to " + catalogProviderName + " Catalog Connector Server")
	r, err := c.GetDatasetInfo(ctx, objToSend)

	// updated for better exception handling using standard GRPC codes
	if err != nil {
		log.Printf("Error sending data to %s Catalog Connector: %v", catalogProviderName, err)
		errStatus, _ := status.FromError(err)
		fmt.Println("Message:", errStatus.Message())
		// lets print the error code which is `INVALID_ARGUMENT`
		fmt.Println("Code:", errStatus.Code())
		return errors.Wrap(err, "Error sending data to Catalog Connector")
	}

	fmt.Println("***************************************************************")
	log.Printf("Received Response for GetDatasetInfo with  datasetID: %s\n", r.GetDatasetId())
	fmt.Println("***************************************************************")
	log.Printf("Response received from %s is given below:", catalogProviderName)
	s, _ := json.MarshalIndent(r, "", "\t")
	fmt.Print(string(s))
	fmt.Println("***************************************************************")
	fmt.Println("###############################################################")
	fmt.Println(r)
	fmt.Println("###############################################################")
	// t := proto.TextMarshaler{}
	// t.Marshal(os.Stdout, r)
	// fmt.Println("*********************eeee************************")
	// s, _ = json.MarshalIndent(r, "", "    ")
	// fmt.Print(string(s))
	// fmt.Println("***********************eee********************************")
	// fmt.Println("********@@@@@@@@@@@@@@***********************")
	// enc := json.NewEncoder(os.Stdout)
	// if err := enc.Encode(r); err != nil {
	// 	panic(err)
	// }
	// enc.SetIndent("", "  ")
	// if err := enc.Encode(r); err != nil {
	// 	panic(err)
	// }
	fmt.Println("********@@@@@@@@@@@@@@***********************")
	log.Printf("Received Response for GetDatasetInfo with  details.name : %s\n", r.GetDetails().GetName())
	log.Printf("Received Response for GetDatasetInfo with  details.dataowner : %s\n", r.GetDetails().GetDataOwner())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.name : %s\n", r.GetDetails().GetDataStore().GetName())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.type : %s\n", r.GetDetails().GetDataStore().GetType())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.s3.endpoint : %s\n", r.GetDetails().GetDataStore().GetS3().GetEndpoint())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.s3.bucket : %s\n", r.GetDetails().GetDataStore().GetS3().GetBucket())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.s3.objectkey : %s\n", r.GetDetails().GetDataStore().GetS3().GetObjectKey())
	log.Printf("Received Response for GetDatasetInfo with  details.datastore.s3.region : %s\n", r.GetDetails().GetDataStore().GetS3().GetRegion())
	return nil
}

func WriteMetadata(datasetID string) error {
	catalogConnectorURL := getEnv("CATALOG_CONNECTOR_URL")
	catalogProviderName := getEnv("CATALOG_PROVIDER_NAME")

	timeoutInSecs := getEnv("CONNECTION_TIMEOUT")
	timeoutInSeconds, err := strconv.Atoi(timeoutInSecs)
	if err != nil {
		log.Printf("Atoi conversion of timeoutinseconds failed: %v", err)
		return errors.Wrap(err, "Atoi conversion of timeoutinseconds failed")
	}

	fmt.Println("timeoutInSeconds: ", timeoutInSeconds)
	fmt.Println("catalogConnectorURL: ", catalogConnectorURL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInSeconds)*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, catalogConnectorURL, grpc.WithInsecure())

	if err != nil {
		log.Printf("Connection to "+catalogProviderName+" Catalog Connector failed: %v", err)
		errStatus, _ := status.FromError(err)
		fmt.Println(errStatus.Message())
		fmt.Println(errStatus.Code())
		return errors.Wrap(err, "Connection to "+catalogProviderName+" Catalog Connector failed")
	}
	defer conn.Close()

	c := pb.NewDataCatalogServiceClient(conn)

	// catalogCredentials := "v1/m4d/user-creds/datauser1/notebook-with-kafka/WKC"
	//appID := "datauser1/notebook-with-kafka"
	//appID := getEnv("APPID")
	//objToSend := &pb.CatalogDatasetWriteRequest{AppId: appID, DatasetId: datasetID}

	datasetName := "ingest data"
	dataOwner := ""
	endpoint := "s3.eu-de.cloud-object-storage.appdomain.cloud"
	bucket := "datamesh-objectstorage-secret-provider-test"
	objectKey := "transactions.csv"
	region := "Turkey"
	s3datastore := &pb.S3DataStore{Endpoint: endpoint, Bucket: bucket, ObjectKey: objectKey, Region: region}
	datastoreType := pb.DataStore_S3
	datastoreName := ""
	dataStore := &pb.DataStore{Type: datastoreType, Name: datastoreName, S3: s3datastore}
	dataFormat := "csv"
	geo := "ne"
	var tags = []string{"confidential1", "confidential2"}

	var datasetTags = []string{"resideny = Tukrey"}
	namedMetadata := make(map[string]string)
	namedMetadata["testkey"] = "testvalue"
	namedMetadata["business_term_id"] = ",businesstermid2"
	namedMetadata["business_term_name"] = "Customer Name,‎‪GRID Identifier‬"
	datasetNamedMetadata := make(map[string]string)
	datasetNamedMetadata["testkey"] = "testvalue"

	componentType := "column"
	componentsMetadata1 := &pb.DataComponentMetadata{ComponentType: componentType, NamedMetadata: namedMetadata, Tags: tags}
	componentsMetadata := make(map[string]*pb.DataComponentMetadata)
	componentsMetadata["amount"] = componentsMetadata1
	componentsMetadata["oldbalanceOrg"] = componentsMetadata1
	datasetmetadata := &pb.DatasetMetadata{DatasetNamedMetadata: datasetNamedMetadata, DatasetTags: datasetTags, ComponentsMetadata: componentsMetadata}
	dataset_details := &pb.DatasetDetails{Name: datasetName, DataOwner: dataOwner, DataStore: dataStore, DataFormat: dataFormat, Geo: geo, Metadata: datasetmetadata}

	accessKey := "testaccessKey"
	datasetId := "testdatasetId"
	username := "testusername"
	password := "testpassword"
	apiKey := "testapikey"
	resourceInstanceId := "testresourceinstanceid"
	destinationCatalogId := "<DEST_CATALOG_ID>"
	creds := &pb.Credentials{AccessKey: accessKey, SecretKey: datasetId, Username: username, Password: password, ApiKey: apiKey, ResourceInstanceId: resourceInstanceId}
	objToSend := &pb.RegisterAssetRequest{Creds: creds, DatasetDetails: dataset_details, DestinationCatalogId: destinationCatalogId}

	log.Printf("Sending CatalogDatasetWriteRequest: ")
	catDataReqStr, _ := json.MarshalIndent(objToSend, "", "\t")
	log.Print(string(catDataReqStr))
	log.Println("1***************************************************************")

	log.Printf("Sending request to " + catalogProviderName + " Catalog Connector Server")
	r, err := c.RegisterDatasetInfo(ctx, objToSend)

	// updated for better exception handling using standard GRPC codes
	if err != nil {
		log.Printf("Error sending data to %s Catalog Connector: %v", catalogProviderName, err)
		errStatus, _ := status.FromError(err)
		fmt.Println("Message:", errStatus.Message())
		// lets print the error code which is `INVALID_ARGUMENT`
		fmt.Println("Code:", errStatus.Code())
		return errors.Wrap(err, "Error sending data to Catalog Connector")
		fmt.Println("***************************************************************")
	}

	log.Printf("Received Response for WriteDatasetInfo with  datasetID: %s\n", r.GetAssetId())
	fmt.Println("***************************************************************")
	log.Printf("Response received from %s is given below:", catalogProviderName)
	s, _ := json.MarshalIndent(r, "", "\t")
	fmt.Print(string(s))
	fmt.Println("***************************************************************")
	return nil
}

func main() {

	catalogID := "<CATALOGID>"
	datasetID := "<ASSETID/DATASETID>"

	var datasetIDJson string
	if getEnv("CATALOG_PROVIDER_NAME") == "EGERIA" {
		datasetIDJson = "{\"ServerName\":\"mds1\",\"AssetGuid\":\"<AssetGUIID>\"}"
	} else {
		datasetIDJson = "{\"catalog_id\":\"" + catalogID + "\",\"asset_id\":\"" + datasetID + "\"}"
		log.Println("datasetIDJson used: ", datasetIDJson)
	}

	err := GetMetadata(datasetIDJson)
	if err != nil {
		fmt.Printf("Error in GetCredentials:\n %v\n\n", err)
		fmt.Printf("Error in GetCredentials Details:\n%+v\n\n", err)
		// errors.Cause() provides access to original error.
		fmt.Printf("Error in GetCredentials Cause: %v\n", errors.Cause(err))
		fmt.Printf("Error in GetCredentials Extended Cause:\n%+v\n", errors.Cause(err))
	}

	// err = WriteMetadata(datasetIDJson)
	// if err != nil {
	// 	fmt.Printf("Error in WriteMetadata:\n %v\n\n", err)
	// 	fmt.Printf("Error in WriteMetadata Details:\n%+v\n\n", err)
	// 	// errors.Cause() provides access to original error.
	// 	fmt.Printf("Error in WriteMetadata Cause: %v\n", errors.Cause(err))
	// 	fmt.Printf("Error in WriteMetadata Extended Cause:\n%+v\n", errors.Cause(err))
	// }
}
