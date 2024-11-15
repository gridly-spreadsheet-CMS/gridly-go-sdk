# client\CdnApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](CdnApi.md#List) | **Get** /v1/cdns | list
[**Publish**](CdnApi.md#Publish) | **Put** /v1/cdns/{cdnId}/publish | publish
[**UnPublish**](CdnApi.md#UnPublish) | **Put** /v1/cdns/{cdnId}/unpublish | unPublish



## List

> []CDN List(ctx).GridId(gridId).Offset(offset).Limit(limit).Ids(ids).Published(published).Execute()

list



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    gridly "./openapi"
)

func main() {
    gridId := "gridId_example" // string | gridId
    offset := int32(56) // int32 | offset (optional) (default to 0)
    limit := int32(56) // int32 | limit (optional) (default to 128)
    ids := []string{"Inner_example"} // []string | ids (optional)
    published := true // bool | published (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnApi.List(context.Background()).GridId(gridId).Offset(offset).Limit(limit).Ids(ids).Published(published).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []CDN
    fmt.Fprintf(os.Stdout, "Response from `CdnApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridId** | **string** | gridId | 
 **offset** | **int32** | offset | [default to 0]
 **limit** | **int32** | limit | [default to 128]
 **ids** | **[]string** | ids | 
 **published** | **bool** | published | 

### Return type

[**[]CDN**](CDN.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Publish

> CDN Publish(ctx, cdnId).Execute()

publish



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    gridly "./openapi"
)

func main() {
    cdnId := "cdnId_example" // string | cdnId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnApi.Publish(context.Background(), cdnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnApi.Publish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Publish`: CDN
    fmt.Fprintf(os.Stdout, "Response from `CdnApi.Publish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdnId** | **string** | cdnId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CDN**](CDN.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnPublish

> CDN UnPublish(ctx, cdnId).Execute()

unPublish



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    gridly "./openapi"
)

func main() {
    cdnId := "cdnId_example" // string | cdnId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnApi.UnPublish(context.Background(), cdnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnApi.UnPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnPublish`: CDN
    fmt.Fprintf(os.Stdout, "Response from `CdnApi.UnPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdnId** | **string** | cdnId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CDN**](CDN.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

