# \DefaultApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetInfo**](DefaultApi.md#GetAssetInfo) | **Post** /getAssetInfo | This REST API gets data asset information from the data catalog configured in fybrik for the data sets indicated in FybrikApplication yaml
[**WriteAssetInfo**](DefaultApi.md#WriteAssetInfo) | **Post** /writeAssetInfo | This REST API write data asset information to the data catalog configured in fybrik



## GetAssetInfo

> GetAssetResponse GetAssetInfo(ctx).XRequestDatacatalogCred(xRequestDatacatalogCred).GetAssetRequest(getAssetRequest).Execute()

This REST API gets data asset information from the data catalog configured in fybrik for the data sets indicated in FybrikApplication yaml

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestDatacatalogCred := "xRequestDatacatalogCred_example" // string | 
    getAssetRequest := TODO // GetAssetRequest | Data Catalog Request Object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAssetInfo(context.Background()).XRequestDatacatalogCred(xRequestDatacatalogCred).GetAssetRequest(getAssetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAssetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetInfo`: GetAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAssetInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestDatacatalogCred** | **string** |  | 
 **getAssetRequest** | [**GetAssetRequest**](GetAssetRequest.md) | Data Catalog Request Object. | 

### Return type

[**GetAssetResponse**](GetAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteAssetInfo

> WriteAssetResponse WriteAssetInfo(ctx).XRequestDatacatalogWriteCred(xRequestDatacatalogWriteCred).WriteAssetRequest(writeAssetRequest).Execute()

This REST API write data asset information to the data catalog configured in fybrik

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xRequestDatacatalogWriteCred := "xRequestDatacatalogWriteCred_example" // string | 
    writeAssetRequest := TODO // WriteAssetRequest | Write Asset Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.WriteAssetInfo(context.Background()).XRequestDatacatalogWriteCred(xRequestDatacatalogWriteCred).WriteAssetRequest(writeAssetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WriteAssetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WriteAssetInfo`: WriteAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.WriteAssetInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWriteAssetInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestDatacatalogWriteCred** | **string** |  | 
 **writeAssetRequest** | [**WriteAssetRequest**](WriteAssetRequest.md) | Write Asset Request | 

### Return type

[**WriteAssetResponse**](WriteAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

