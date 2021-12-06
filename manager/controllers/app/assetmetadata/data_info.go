// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package assetmetadata

import (
	"errors"

	"fybrik.io/fybrik/pkg/serde"

	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/utils"
	pb "fybrik.io/fybrik/pkg/connectors/protobuf"
	datacatalogTaxonomyModels "fybrik.io/fybrik/pkg/taxonomy/model/datacatalog/base"
)

// DataDetails is the asset metadata received from the catalog connector
// This structure is in use by the manager and other components, such as policy manager and config policies evaluator
type DataDetails struct {
	// Name of the asset
	Name string
	// Interface is the protocol and format
	Interface app.InterfaceDetails
	// Geography is the geo-location of the asset
	Geography string
	// Connection is the connection details in raw format as received from the connector
	Connection serde.Arbitrary
	// Metadata
	TagMetadata *pb.DatasetMetadata
}

// Transforms a CatalogDatasetInfo into a DataDetails struct
// TODO Think about getting rid of one or the other and reuse
// func CatalogDatasetToDataDetails(response *pb.CatalogDatasetInfo) (*DataDetails, error) {
func CatalogDatasetToDataDetails(response *datacatalogTaxonomyModels.DataCatalogResponse) (*DataDetails, error) {
	details := response.GetDetails()
	if details == nil {
		return nil, errors.New("no metadata found for " + response.DatasetId)
	}
	//format := details.DataFormat
	format := response.Details.DataFormat
	if format == nil {
		format = new(string)
	}
	//connection := serde.NewArbitrary(details.DataStore)
	connection := serde.NewArbitrary(response.Details.Connection)
	protocol, err := utils.GetProtocol(&response.Details.Connection)

	return &DataDetails{
		Name: response.ResourceMetadata.Name,
		Interface: app.InterfaceDetails{
			Protocol:   protocol,
			DataFormat: *format,
		},
		Geography:   *response.ResourceMetadata.Geography,
		Connection:  *connection,
		TagMetadata: details.Metadata,
	}, err
}
