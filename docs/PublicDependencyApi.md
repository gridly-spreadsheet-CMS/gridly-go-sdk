# api\PublicDependencyApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PublicDependencyApi.md#Create) | **Post** /v1/views/{viewId}/dependencies | create
[**Delete**](PublicDependencyApi.md#Delete) | **Delete** /v1/views/{viewId}/dependencies | delete
[**Delete_0**](PublicDependencyApi.md#Delete_0) | **Delete** /v1/views/{viewId}/dependencies/{dependencyId} | delete
[**Get**](PublicDependencyApi.md#Get) | **Get** /v1/views/{viewId}/dependencies/{dependencyId} | get
[**List**](PublicDependencyApi.md#List) | **Get** /v1/views/{viewId}/dependencies | list
[**Update**](PublicDependencyApi.md#Update) | **Put** /v1/views/{viewId}/dependencies/{dependencyId} | update



## Create

> PublicDependency Create(ctx, viewId).CreateDependency(createDependency).Execute()

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
    viewId := "viewId_example" // string | viewId
    createDependency := *openapiclient.NewCreateDependency() // CreateDependency | createDependency

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.Create(context.Background(), viewId).CreateDependency(createDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: PublicDependency
    fmt.Fprintf(os.Stdout, "Response from `PublicDependencyApi.Create`: %v\n", resp)
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

 **createDependency** | [**CreateDependency**](CreateDependency.md) | createDependency | 

### Return type

[**PublicDependency**](PublicDependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, viewId).DeleteDependency(deleteDependency).Execute()

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
    viewId := "viewId_example" // string | viewId
    deleteDependency := *openapiclient.NewDeleteDependency() // DeleteDependency | deleteDependency

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.Delete(context.Background(), viewId).DeleteDependency(deleteDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteDependency** | [**DeleteDependency**](DeleteDependency.md) | deleteDependency | 

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


## Delete_0

> Delete_0(ctx, dependencyId, viewId).Execute()

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
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.Delete_0(context.Background(), dependencyId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.Delete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dependencyId** | **string** | dependencyId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete_1Request struct via the builder pattern


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

> PublicDependency Get(ctx, dependencyId, viewId).Execute()

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
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.Get(context.Background(), dependencyId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: PublicDependency
    fmt.Fprintf(os.Stdout, "Response from `PublicDependencyApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dependencyId** | **string** | dependencyId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PublicDependency**](PublicDependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []PublicDependency List(ctx, viewId).Execute()

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
    viewId := "viewId_example" // string | viewId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.List(context.Background(), viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []PublicDependency
    fmt.Fprintf(os.Stdout, "Response from `PublicDependencyApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicDependency**](PublicDependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> PublicDependency Update(ctx, dependencyId, viewId).UpdateDependency(updateDependency).Execute()

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
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId
    updateDependency := *openapiclient.NewUpdateDependency() // UpdateDependency | updateDependency

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PublicDependencyApi.Update(context.Background(), dependencyId, viewId).UpdateDependency(updateDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicDependencyApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: PublicDependency
    fmt.Fprintf(os.Stdout, "Response from `PublicDependencyApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dependencyId** | **string** | dependencyId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDependency** | [**UpdateDependency**](UpdateDependency.md) | updateDependency | 

### Return type

[**PublicDependency**](PublicDependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

