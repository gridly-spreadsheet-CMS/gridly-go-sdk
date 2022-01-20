# client\BranchApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BranchApi.md#Create) | **Post** /v1/branches | create
[**Delete**](BranchApi.md#Delete) | **Delete** /v1/branches/{branchId} | delete
[**Get**](BranchApi.md#Get) | **Get** /v1/branches/{branchId} | get
[**List**](BranchApi.md#List) | **Get** /v1/branches | list
[**Merge**](BranchApi.md#Merge) | **Post** /v1/branches/{branchId}/merge | merge



## Create

> Branch Create(ctx).CreateBranch(createBranch).BranchId(branchId).GridId(gridId).Execute()

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
    createBranch := *gridly.NewCreateBranch() // CreateBranch | createBranch
    branchId := "branchId_example" // string | branchId (optional)
    gridId := "gridId_example" // string | gridId (optional)

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.BranchApi.Create(context.Background()).CreateBranch(createBranch).BranchId(branchId).GridId(gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Branch
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBranch** | [**CreateBranch**](CreateBranch.md) | createBranch | 
 **branchId** | **string** | branchId | 
 **gridId** | **string** | gridId | 

### Return type

[**Branch**](Branch.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, branchId).Execute()

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
    branchId := "branchId_example" // string | branchId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.BranchApi.Delete(context.Background(), branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** | branchId | 

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

> Branch Get(ctx, branchId).Execute()

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
    branchId := "branchId_example" // string | branchId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.BranchApi.Get(context.Background(), branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Branch
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** | branchId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Branch**](Branch.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Branch List(ctx).GridId(gridId).Execute()

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
    resp, r, err := api_client.BranchApi.List(context.Background()).GridId(gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Branch
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gridId** | **string** | gridId | 

### Return type

[**[]Branch**](Branch.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Merge

> Merge(ctx, branchId).DestinationBranchId(destinationBranchId).MergeRecordOptions(mergeRecordOptions).Execute()

merge

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
    branchId := "branchId_example" // string | branchId
    destinationBranchId := "destinationBranchId_example" // string | destinationBranchId
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional)

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.BranchApi.Merge(context.Background(), branchId).DestinationBranchId(destinationBranchId).MergeRecordOptions(mergeRecordOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.Merge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**branchId** | **string** | branchId | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationBranchId** | **string** | destinationBranchId | 
 **mergeRecordOptions** | **[]string** | mergeRecordOptions | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

