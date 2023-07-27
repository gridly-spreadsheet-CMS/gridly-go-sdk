# client\PathApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PathApi.md#Create) | **Post** /v1/views/{viewId}/paths | create
[**Delete**](PathApi.md#Delete) | **Delete** /v1/views/{viewId}/paths | delete
[**List**](PathApi.md#List) | **Get** /v1/views/{viewId}/paths/tree | list
[**Move**](PathApi.md#Move) | **Post** /v1/views/{viewId}/paths/move | move
[**Update**](PathApi.md#Update) | **Put** /v1/views/{viewId}/paths/{path} | update



## Create

> PathList Create(ctx, viewId).CreatePath(createPath).Execute()

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
    createPath := *gridly.NewCreatePath() // CreatePath | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.PathApi.Create(context.Background(), viewId).CreatePath(createPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PathApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: PathList
    fmt.Fprintf(os.Stdout, "Response from `PathApi.Create`: %v\n", resp)
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

 **createPath** | [**CreatePath**](CreatePath.md) |  | 

### Return type

[**PathList**](PathList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, viewId).DeletePath(deletePath).Execute()

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
    deletePath := *gridly.NewDeletePath() // DeletePath | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.PathApi.Delete(context.Background(), viewId).DeletePath(deletePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PathApi.Delete``: %v\n", err)
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

 **deletePath** | [**DeletePath**](DeletePath.md) |  | 

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


## List

> []PathNode List(ctx, viewId).RootPath(rootPath).Execute()

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
    rootPath := "rootPath_example" // string | rootPath (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.PathApi.List(context.Background(), viewId).RootPath(rootPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PathApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []PathNode
    fmt.Fprintf(os.Stdout, "Response from `PathApi.List`: %v\n", resp)
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

 **rootPath** | **string** | rootPath | 

### Return type

[**[]PathNode**](PathNode.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Move

> PathList Move(ctx, viewId).MovePath(movePath).Execute()

move



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
    movePath := *gridly.NewMovePath() // MovePath | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.PathApi.Move(context.Background(), viewId).MovePath(movePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PathApi.Move``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Move`: PathList
    fmt.Fprintf(os.Stdout, "Response from `PathApi.Move`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **movePath** | [**MovePath**](MovePath.md) |  | 

### Return type

[**PathList**](PathList.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> PathSingle Update(ctx, viewId, path).UpdatePath(updatePath).Execute()

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
    viewId := "viewId_example" // string | viewId
    path := "path_example" // string | path
    updatePath := *gridly.NewUpdatePath("NewName_example") // UpdatePath | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.PathApi.Update(context.Background(), viewId, path).UpdatePath(updatePath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PathApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: PathSingle
    fmt.Fprintf(os.Stdout, "Response from `PathApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 
**path** | **string** | path | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePath** | [**UpdatePath**](UpdatePath.md) |  | 

### Return type

[**PathSingle**](PathSingle.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

