// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"

	"fybrik.io/fybrik/pkg/logging"
	"fybrik.io/fybrik/pkg/model/datacatalog"
	fybrikTLS "fybrik.io/fybrik/pkg/tls"
)

type ConnectorController struct {
	TestServerURL string
	TestClient    *retryablehttp.Client
	Log           zerolog.Logger
}

func NewConnectorController(TestServerURL string) (*ConnectorController, error) {
	log := logging.LogInit(logging.CONNECTOR, "Test-connector")
	retryClient := retryablehttp.NewClient()
	if strings.HasPrefix(TestServerURL, "https") {
		config, err := fybrikTLS.GetClientTLSConfig(&log)
		if err != nil {
			log.Error().Err(err)
			return nil, err
		}
		if config != nil {
			log.Info().Msg("Set TLS config for Test connector as a client")
			retryClient.HTTPClient.Transport = &http.Transport{TLSClientConfig: config}
		}
	}

	return &ConnectorController{
		TestServerURL: TestServerURL,
		TestClient:    retryClient,
		Log:           log,
	}, nil
}

func (r *ConnectorController) CreateNewComponent(c *gin.Context) {
	// Parse request
	var request datacatalog.CreateNewComponentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logging.LogStructure("CreateNewComponentRequest object received:", request, &r.Log, zerolog.DebugLevel, false, false)

	// Unmarshal as GetPolicyDecisionsResponse for the sake of validation
	var response datacatalog.CreateNewComponentResponse
	response.Status = "Test!"

	r.Log.Info().Msg(
		"Sending response from Test connector : " + string(response.Status))

	c.JSON(http.StatusOK, response)
}
