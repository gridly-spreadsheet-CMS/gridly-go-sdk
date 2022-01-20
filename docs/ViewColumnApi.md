# client\ViewColumnApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](ViewColumnApi.md#Add) | **Post** /v1/views/{viewId}/columns/{columnId}/add | add
[**Create**](ViewColumnApi.md#Create) | **Post** /v1/views/{viewId}/columns | create
[**Delete**](ViewColumnApi.md#Delete) | **Delete** /v1/views/{viewId}/columns/{columnId} | delete
[**Get**](ViewColumnApi.md#Get) | **Get** /v1/views/{viewId}/columns/{columnId} | get
[**Remove**](ViewColumnApi.md#Remove) | **Post** /v1/views/{viewId}/columns/{columnId}/remove | remove
[**Update**](ViewColumnApi.md#Update) | **Patch** /v1/views/{viewId}/columns/{columnId} | update



## Add

> ViewColumn Add(ctx, columnId, viewId).Execute()

add

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
    columnId := "columnId_example" // string | columnId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Add(context.Background(), columnId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Add``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Add`: ViewColumn
    fmt.Fprintf(os.Stdout, "Response from `ViewColumnApi.Add`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** | columnId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ViewColumn**](ViewColumn.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> ViewColumn Create(ctx, viewId).CreateColumn(createColumn).Execute()

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
    viewId := "viewId_example" // string | viewId
    createColumn := *gridly.NewCreateColumn("Type_example") // CreateColumn | createColumn

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Create(context.Background(), viewId).CreateColumn(createColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: ViewColumn
    fmt.Fprintf(os.Stdout, "Response from `ViewColumnApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createColumn** | [**CreateColumn**](CreateColumn.md) | createColumn | 

### Return type

[**ViewColumn**](ViewColumn.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, columnId, viewId).Execute()

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
    columnId := "columnId_example" // string | columnId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Delete(context.Background(), columnId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** | columnId | 
**viewId** | **string** | viewId | 

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

> ViewColumn Get(ctx, columnId, viewId).Execute()

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
    columnId := "columnId_example" // string | columnId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Get(context.Background(), columnId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: ViewColumn
    fmt.Fprintf(os.Stdout, "Response from `ViewColumnApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** | columnId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ViewColumn**](ViewColumn.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Remove

> Remove(ctx, columnId, viewId).Execute()

remove

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
    columnId := "columnId_example" // string | columnId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Remove(context.Background(), columnId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Remove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** | columnId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRequest struct via the builder pattern


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


## Update

> ViewColumn Update(ctx, columnId, viewId).UpdateColumn(updateColumn).Execute()

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
    columnId := "columnId_example" // string | columnId
    viewId := "viewId_example" // string | viewId
    updateColumn := *gridly.NewUpdateColumn() // UpdateColumn | updateColumn

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewColumnApi.Update(context.Background(), columnId, viewId).UpdateColumn(updateColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewColumnApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: ViewColumn
    fmt.Fprintf(os.Stdout, "Response from `ViewColumnApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**columnId** | **string** | columnId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateColumn** | [**UpdateColumn**](UpdateColumn.md) | updateColumn | 

### Return type

[**ViewColumn**](ViewColumn.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

