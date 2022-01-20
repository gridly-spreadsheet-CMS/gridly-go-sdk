# client\ViewDependencyApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ViewDependencyApi.md#Create) | **Post** /v1/views/{viewId}/dependencies | create
[**Delete**](ViewDependencyApi.md#Delete) | **Delete** /v1/views/{viewId}/dependencies | delete
[**Delete_0**](ViewDependencyApi.md#Delete_0) | **Delete** /v1/views/{viewId}/dependencies/{dependencyId} | delete
[**Get**](ViewDependencyApi.md#Get) | **Get** /v1/views/{viewId}/dependencies/{dependencyId} | get
[**List**](ViewDependencyApi.md#List) | **Get** /v1/views/{viewId}/dependencies | list
[**Update**](ViewDependencyApi.md#Update) | **Put** /v1/views/{viewId}/dependencies/{dependencyId} | update



## Create

> Dependency Create(ctx, viewId).CreateDependency(createDependency).Execute()

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
    createDependency := *gridly.NewCreateDependency() // CreateDependency | createDependency

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.Create(context.Background(), viewId).CreateDependency(createDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Dependency
    fmt.Fprintf(os.Stdout, "Response from `ViewDependencyApi.Create`: %v\n", resp)
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

[**Dependency**](Dependency.md)

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
    gridly "./openapi"
)

func main() {
    viewId := "viewId_example" // string | viewId
    deleteDependency := *gridly.NewDeleteDependency() // DeleteDependency | deleteDependency

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.Delete(context.Background(), viewId).DeleteDependency(deleteDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.Delete``: %v\n", err)
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

- **Content-Type**: application/json
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
    gridly "./openapi"
)

func main() {
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.Delete_0(context.Background(), dependencyId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.Delete_0``: %v\n", err)
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

> Dependency Get(ctx, dependencyId, viewId).Execute()

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
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.Get(context.Background(), dependencyId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Dependency
    fmt.Fprintf(os.Stdout, "Response from `ViewDependencyApi.Get`: %v\n", resp)
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

[**Dependency**](Dependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Dependency List(ctx, viewId).Execute()

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
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.List(context.Background(), viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Dependency
    fmt.Fprintf(os.Stdout, "Response from `ViewDependencyApi.List`: %v\n", resp)
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

[**[]Dependency**](Dependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Dependency Update(ctx, dependencyId, viewId).UpdateDependency(updateDependency).Execute()

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
    dependencyId := "dependencyId_example" // string | dependencyId
    viewId := "viewId_example" // string | viewId
    updateDependency := *gridly.NewUpdateDependency() // UpdateDependency | updateDependency

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ViewDependencyApi.Update(context.Background(), dependencyId, viewId).UpdateDependency(updateDependency).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewDependencyApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Dependency
    fmt.Fprintf(os.Stdout, "Response from `ViewDependencyApi.Update`: %v\n", resp)
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

[**Dependency**](Dependency.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

