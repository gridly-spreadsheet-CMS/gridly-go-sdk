# client\BranchApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BranchApi.md#Create) | **Post** /v1/branches | create
[**CreateDiffCheck**](BranchApi.md#CreateDiffCheck) | **Post** /v1/branches/diffcheck | createDiffCheck
[**Delete**](BranchApi.md#Delete) | **Delete** /v1/branches/{branchId} | delete
[**Get**](BranchApi.md#Get) | **Get** /v1/branches/{branchId} | get
[**GetDiffCheck**](BranchApi.md#GetDiffCheck) | **Get** /v1/branches/diffcheck/{taskId} | getDiffCheck
[**List**](BranchApi.md#List) | **Get** /v1/branches | list
[**Merge**](BranchApi.md#Merge) | **Post** /v1/branches/{branchId}/merge | merge



## Create

> Branch Create(ctx).CreateBranch(createBranch).GridId(gridId).BranchId(branchId).Execute()

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
    createBranch := *gridly.NewCreateBranch("Name_example") // CreateBranch | 
    gridId := "gridId_example" // string | gridId (optional)
    branchId := "branchId_example" // string | branchId (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.Create(context.Background()).CreateBranch(createBranch).GridId(gridId).BranchId(branchId).Execute()
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
 **createBranch** | [**CreateBranch**](CreateBranch.md) |  | 
 **gridId** | **string** | gridId | 
 **branchId** | **string** | branchId | 

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


## CreateDiffCheck

> Task CreateDiffCheck(ctx).SourceViewId(sourceViewId).DestinationViewId(destinationViewId).Execute()

createDiffCheck



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
    sourceViewId := "sourceViewId_example" // string | sourceViewId
    destinationViewId := "destinationViewId_example" // string | destinationViewId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.CreateDiffCheck(context.Background()).SourceViewId(sourceViewId).DestinationViewId(destinationViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.CreateDiffCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiffCheck`: Task
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.CreateDiffCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiffCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceViewId** | **string** | sourceViewId | 
 **destinationViewId** | **string** | destinationViewId | 

### Return type

[**Task**](Task.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.Delete(context.Background(), branchId).Execute()
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
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.Get(context.Background(), branchId).Execute()
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


## GetDiffCheck

> []BranchDiffRecord GetDiffCheck(ctx, taskId).MergeRecordOptions(mergeRecordOptions).Query(query).Page(page).Execute()

getDiffCheck



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
    taskId := "taskId_example" // string | taskId
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional) (default to ["add","update","delete"])
    query := "query_example" // string | query (optional) (default to "{}")
    page := "page_example" // string | page (optional) (default to "{}")

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.GetDiffCheck(context.Background(), taskId).MergeRecordOptions(mergeRecordOptions).Query(query).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.GetDiffCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiffCheck`: []BranchDiffRecord
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.GetDiffCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | taskId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiffCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mergeRecordOptions** | **[]string** | mergeRecordOptions | [default to [&quot;add&quot;,&quot;update&quot;,&quot;delete&quot;]]
 **query** | **string** | query | [default to &quot;{}&quot;]
 **page** | **string** | page | [default to &quot;{}&quot;]

### Return type

[**[]BranchDiffRecord**](BranchDiffRecord.md)

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
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.List(context.Background()).GridId(gridId).Execute()
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

> Task Merge(ctx, branchId).DestinationBranchId(destinationBranchId).MergeBranchRequest(mergeBranchRequest).MergeRecordOptions(mergeRecordOptions).Execute()

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
    mergeBranchRequest := *gridly.NewMergeBranchRequest() // MergeBranchRequest | 
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional) (default to [])

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.BranchApi.Merge(context.Background(), branchId).DestinationBranchId(destinationBranchId).MergeBranchRequest(mergeBranchRequest).MergeRecordOptions(mergeRecordOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BranchApi.Merge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Merge`: Task
    fmt.Fprintf(os.Stdout, "Response from `BranchApi.Merge`: %v\n", resp)
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
 **mergeBranchRequest** | [**MergeBranchRequest**](MergeBranchRequest.md) |  | 
 **mergeRecordOptions** | **[]string** | mergeRecordOptions | [default to []]

### Return type

[**Task**](Task.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

