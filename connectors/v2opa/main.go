// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	v2opa "github.com/mesh-for-data/mesh-for-data/connectors/v2opa/http"
	openapiclient "github.com/mesh-for-data/mesh-for-data/pkg/connectors/out_go_client"
)

// var opaServerURL = ""

// const defaultPort = "50082" // synced with opa_connector.yaml

// type server struct {
// 	pb.UnimplementedPolicyManagerServiceServer
// 	opaReader *opabl.OpaReader
// }

// func getEnv(key string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		log.Fatalf("Env Variable %v not defined", key)
// 	}
// 	log.Printf("Env. variable extracted: %s - %s\n", key, value)
// 	return value
// }

// func getEnvWithDefault(key string, defaultValue string) string {
// 	value, exists := os.LookupEnv(key)
// 	if !exists {
// 		log.Printf("Env. variable not found, default value used: %s - %s\n", key, defaultValue)
// 		return defaultValue
// 	}

// 	log.Printf("Env. variable extracted: %s - %s\n", key, value)
// 	return value
// }

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() v2opa.DefaultApiServicer {
	return &DefaultApiService{}
}

// GetPoliciesDecisions - getPoliciesDecisions
func (s *DefaultApiService) GetPoliciesDecisions(ctx context.Context, input []openapiclient.PolicymanagerRequest) (v2opa.ImplResponse, error) {
	fmt.Println("Reached Server!")
	fmt.Println("input")
	fmt.Println(input)
	// TODO - update GetPoliciesDecisions with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []PolicymanagerResponse{}) or use other options such as http.Ok ...
	//return Response(200, []PolicymanagerResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return v2opa.Response(http.StatusNotImplemented, nil), errors.New("GetPoliciesDecisions method not implemented")
}

func main() {
	// port := getEnvWithDefault("PORT_OPA_CONNECTOR", defaultPort)
	// opaServerURL = getEnv("OPA_SERVER_URL") // set global variable

	// log.Println("OPA_SERVER_URL env variable in OPAConnector: ", opaServerURL)
	// log.Println("Using port to start go opa connector : ", port)

	// log.Printf("Server starts listening on port %v", port)
	// lis, err := net.Listen("tcp", ":"+port)
	// if err != nil {
	// 	log.Fatalf("Error in listening: %v", err)
	// }
	// s := grpc.NewServer()
	// srv := &server{opaReader: opabl.NewOpaReader(opaServerURL)}
	// pb.RegisterPolicyManagerServiceServer(s, srv)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Error in service: %v", err)
	// }

	log.Printf("Server started")

	DefaultApiService := NewDefaultApiService()
	DefaultApiController := v2opa.NewDefaultApiController(DefaultApiService)

	router := v2opa.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))

}
