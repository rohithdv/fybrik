// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"

	v2opa "github.com/mesh-for-data/mesh-for-data/connectors/v2opa/openapiserver"
)

func main() {

	log.Printf("Server started")

	DefaultApiService := v2opa.NewDefaultApiService()
	DefaultApiController := v2opa.NewDefaultApiController(DefaultApiService)

	router := v2opa.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":50050", router))

}
