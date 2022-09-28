// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"emperror.dev/errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"fybrik.io/fybrik/pkg/environment"
	fybrikTLS "fybrik.io/fybrik/pkg/tls"
)

const (
	envTESTServerURL = "TEST_SERVER_URL"
	envServicePort   = "SERVICE_PORT"
)

var (
	gitCommit string
	gitTag    string
)

// NewRouter returns a new router.
func NewRouter(controller *ConnectorController) *gin.Engine {
	router := gin.Default()
	router.POST("/createNewComponent", controller.CreateNewComponent)
	return router
}

// RootCmd defines the root cli command
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "TEST-connector",
		Short: "Kubernetes based policy manager connector for Fybrik",
	}
	cmd.AddCommand(RunCmd())
	return cmd
}

// RunCmd defines the command for running the connector
func RunCmd() *cobra.Command {
	ip := ""
	portStr, err := environment.MustGetEnv(envServicePort)
	if err != nil {
		log.Err(err).Msg(envServicePort + " env var is not defined")
		return nil
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Err(err).Msg(fmt.Sprintf("error in converting %s = [%s] to integer", envServicePort, portStr))
		return nil
	}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run TEST connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			gin.SetMode(gin.ReleaseMode)
			// Parse environment variables
			TESTServerURL, err := environment.MustGetEnv(envTESTServerURL)
			if err != nil {
				return errors.Wrap(err, "failed to retrieve URL for communicating with TEST server")
			}

			if !strings.HasPrefix(TESTServerURL, "https://") && !strings.HasPrefix(TESTServerURL, "http://") {
				return errors.New("server URL for TEST server must have http or https schema")
			}

			// Create and start connector
			controller, err := NewConnectorController(TESTServerURL)
			if err != nil {
				return errors.Wrap(err, "failed to set connection to TEST server")
			}
			controller.Log.Info().Msg("based on: gitTag=" + gitTag + ", latest gitCommit=" + gitCommit)
			router := NewRouter(controller)
			router.Use(gin.Logger())

			bindAddress := fmt.Sprintf("%s:%d", ip, port)
			if environment.IsUsingTLS() {
				tlsConfig, err := fybrikTLS.GetServerConfig(&controller.Log)
				if err != nil {
					return errors.Wrap(err, "failed to get tls config")
				}
				server := http.Server{Addr: bindAddress, Handler: router, TLSConfig: tlsConfig}
				return server.ListenAndServeTLS("", "")
			}
			controller.Log.Info().Msg(fybrikTLS.TLSDisabledMsg)
			return router.Run(bindAddress)
		},
	}
	cmd.Flags().StringVar(&ip, "ip", ip, "IP address")
	cmd.Flags().IntVar(&port, "port", port, "Listening port")
	return cmd
}

func main() {
	// Run the cli
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
