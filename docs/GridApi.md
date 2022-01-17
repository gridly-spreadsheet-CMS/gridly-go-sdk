# api\GridApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GridApi.md#Create) | **Post** /v1/grids | create
[**Delete**](GridApi.md#Delete) | **Delete** /v1/grids/{gridId} | delete
[**Get**](GridApi.md#Get) | **Get** /v1/grids/{gridId} | get
[**List**](GridApi.md#List) | **Get** /v1/grids | list
[**ListTemplateGrids**](GridApi.md#ListTemplateGrids) | **Get** /v1/template-grids | listTemplateGrids
[**Update**](GridApi.md#Update) | **Patch** /v1/grids/{gridId} | update



## Create

> Grid Create(ctx).DbId(dbId).CreateCompositeGrid(createCompositeGrid).Execute()

create

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
    dbId := "dbId_example" // string | dbId
    createCompositeGrid := *openapiclient.NewCreateGrid() // CreateGrid | createCompositeGrid

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.Create(context.Background()).DbId(dbId).CreateCompositeGrid(createCompositeGrid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbId** | **string** | dbId | 
 **createCompositeGrid** | [**CreateGrid**](CreateGrid.md) | createCompositeGrid | 

### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, gridId).Execute()

delete

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
    gridId := "gridId_example" // string | gridId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.Delete(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

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

> Grid Get(ctx, gridId).Execute()

get

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
    gridId := "gridId_example" // string | gridId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.Get(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Grid List(ctx).DbId(dbId).Execute()

list

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
    dbId := "dbId_example" // string | dbId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.List(context.Background()).DbId(dbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbId** | **string** | dbId | 

### Return type

[**[]Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateGrids

> []Grid ListTemplateGrids(ctx).Execute()

listTemplateGrids

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.ListTemplateGrids(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.ListTemplateGrids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplateGrids`: []Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.ListTemplateGrids`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateGridsRequest struct via the builder pattern


### Return type

[**[]Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Grid Update(ctx, gridId).UpdateCompositeGrid(updateCompositeGrid).Execute()

update

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
    gridId := "gridId_example" // string | gridId
    updateCompositeGrid := *openapiclient.NewUpdateGrid() // UpdateGrid | updateCompositeGrid

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GridApi.Update(context.Background(), gridId).UpdateCompositeGrid(updateCompositeGrid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCompositeGrid** | [**UpdateGrid**](UpdateGrid.md) | updateCompositeGrid | 

### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

