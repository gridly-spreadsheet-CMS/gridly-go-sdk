# api\VersionControlApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](VersionControlApi.md#Create) | **Post** /v1/branches | create
[**Delete**](VersionControlApi.md#Delete) | **Delete** /v1/branches/{branchId} | delete
[**Get**](VersionControlApi.md#Get) | **Get** /v1/branches/{branchId} | get
[**List**](VersionControlApi.md#List) | **Get** /v1/branches | list
[**Merge**](VersionControlApi.md#Merge) | **Post** /v1/branches/{branchId}/merge | merge



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
    openapiclient "./openapi"
)

func main() {
    createBranch := *openapiclient.NewCreateBranch() // CreateBranch | createBranch
    branchId := "branchId_example" // string | branchId (optional)
    gridId := "gridId_example" // string | gridId (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionControlApi.Create(context.Background()).CreateBranch(createBranch).BranchId(branchId).GridId(gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionControlApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Branch
    fmt.Fprintf(os.Stdout, "Response from `VersionControlApi.Create`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    branchId := "branchId_example" // string | branchId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionControlApi.Delete(context.Background(), branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionControlApi.Delete``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    branchId := "branchId_example" // string | branchId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionControlApi.Get(context.Background(), branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionControlApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Branch
    fmt.Fprintf(os.Stdout, "Response from `VersionControlApi.Get`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    gridId := "gridId_example" // string | gridId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionControlApi.List(context.Background()).GridId(gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionControlApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Branch
    fmt.Fprintf(os.Stdout, "Response from `VersionControlApi.List`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    branchId := "branchId_example" // string | branchId
    destinationBranchId := "destinationBranchId_example" // string | destinationBranchId
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VersionControlApi.Merge(context.Background(), branchId).DestinationBranchId(destinationBranchId).MergeRecordOptions(mergeRecordOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionControlApi.Merge``: %v\n", err)
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

