// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"io"

	"fybrik.io/fybrik/pkg/model/datacatalog"
)

// DataCatalog is an interface of a facade to a data catalog.
type DataCatalog interface {
	GetAssetInfo(in *datacatalog.GetAssetRequest, creds string) (*datacatalog.GetAssetResponse, error)
	CreateAsset(in *datacatalog.CreateAssetRequest, creds string) (*datacatalog.CreateAssetResponse, error)
	DeleteAsset(in *datacatalog.DeleteAssetRequest, creds string) (*datacatalog.DeleteAssetResponse, error)
	UpdateAsset(in *datacatalog.UpdateAssetRequest, creds string) (*datacatalog.UpdateAssetResponse, error)
	CreateNewComponent(in *datacatalog.CreateNewComponentRequest, creds string) (*datacatalog.CreateNewComponentResponse, error)
	io.Closer
}

func NewDataCatalog(catalogProviderName, catalogConnectorAddress string) (DataCatalog, error) {
	return NewOpenAPIDataCatalog(catalogProviderName, catalogConnectorAddress), nil
}
