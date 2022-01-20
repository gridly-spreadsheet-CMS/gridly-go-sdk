# client\GridMetadataApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GridMetadataApi.md#Create) | **Post** /v1/grids/{gridId}/metadata | create
[**Delete**](GridMetadataApi.md#Delete) | **Delete** /v1/grids/{gridId}/metadata/{metadataId} | delete
[**Get**](GridMetadataApi.md#Get) | **Get** /v1/grids/{gridId}/metadata/{metadataId} | get
[**List**](GridMetadataApi.md#List) | **Get** /v1/grids/{gridId}/metadata | list
[**Update**](GridMetadataApi.md#Update) | **Patch** /v1/grids/{gridId}/metadata/{metadataId} | update



## Create

> Metadata Create(ctx, gridId).CreateMetadataDTO(createMetadataDTO).Execute()

create

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
    createMetadataDTO := *gridly.NewCreateMetadata() // CreateMetadata | createMetadataDTO

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.GridMetadataApi.Create(context.Background(), gridId).CreateMetadataDTO(createMetadataDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridMetadataApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `GridMetadataApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMetadataDTO** | [**CreateMetadata**](CreateMetadata.md) | createMetadataDTO | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, gridId, metadataId).Execute()

delete

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
    metadataId := "metadataId_example" // string | metadataId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.GridMetadataApi.Delete(context.Background(), gridId, metadataId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridMetadataApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**metadataId** | **string** | metadataId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Metadata Get(ctx, gridId, metadataId).Execute()

get

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
    metadataId := "metadataId_example" // string | metadataId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.GridMetadataApi.Get(context.Background(), gridId, metadataId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridMetadataApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `GridMetadataApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**metadataId** | **string** | metadataId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Metadata**](Metadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Metadata List(ctx, gridId).Execute()

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

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.GridMetadataApi.List(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridMetadataApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Metadata
    fmt.Fprintf(os.Stdout, "Response from `GridMetadataApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Metadata**](Metadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Metadata Update(ctx, gridId, metadataId).UpdateMetadataDTO(updateMetadataDTO).Execute()

update

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
    metadataId := "metadataId_example" // string | metadataId
    updateMetadataDTO := *gridly.NewUpdateMetadata() // UpdateMetadata | updateMetadataDTO

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.GridMetadataApi.Update(context.Background(), gridId, metadataId).UpdateMetadataDTO(updateMetadataDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridMetadataApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Metadata
    fmt.Fprintf(os.Stdout, "Response from `GridMetadataApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**metadataId** | **string** | metadataId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMetadataDTO** | [**UpdateMetadata**](UpdateMetadata.md) | updateMetadataDTO | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

