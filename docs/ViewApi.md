# client\ViewApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ViewApi.md#Create) | **Post** /v1/views | create
[**Export**](ViewApi.md#Export) | **Get** /v1/views/{viewId}/export | export
[**Get**](ViewApi.md#Get) | **Get** /v1/views/{viewId} | get
[**GetStatistic**](ViewApi.md#GetStatistic) | **Get** /v1/views/{viewId}/statistic | getStatistic
[**ImportView**](ViewApi.md#ImportView) | **Post** /v1/views/{viewId}/import | importView
[**List**](ViewApi.md#List) | **Get** /v1/views | list
[**Merge**](ViewApi.md#Merge) | **Post** /v1/views/{viewId}/merge | merge



## Create

> View Create(ctx).CreateView(createView).Execute()

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
    createView := *gridly.NewCreateView("Name_example", "GridId_example") // CreateView | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.Create(context.Background()).CreateView(createView).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createView** | [**CreateView**](CreateView.md) |  | 

### Return type

[**View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> *os.File Export(ctx, viewId).ColumnIds(columnIds).FileHeader(fileHeader).Query(query).Sort(sort).Type_(type_).Execute()

export



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
    columnIds := []string{"Inner_example"} // []string | columnIds (optional) (default to [])
    fileHeader := gridly.ExportFileHeader("none") // ExportFileHeader | fileHeader (optional)
    query := "query_example" // string | query (optional) (default to "{}")
    sort := "sort_example" // string | sort (optional) (default to "{}")
    type_ := gridly.FileType("csv") // FileType | type (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.Export(context.Background(), viewId).ColumnIds(columnIds).FileHeader(fileHeader).Query(query).Sort(sort).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnIds** | **[]string** | columnIds | [default to []]
 **fileHeader** | [**ExportFileHeader**](ExportFileHeader.md) | fileHeader | 
 **query** | **string** | query | [default to &quot;{}&quot;]
 **sort** | **string** | sort | [default to &quot;{}&quot;]
 **type_** | [**FileType**](FileType.md) | type | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> View Get(ctx, viewId).ColumnIds(columnIds).Include(include).Page(page).Query(query).Sort(sort).Execute()

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
    viewId := "viewId_example" // string | viewId
    columnIds := []string{"Inner_example"} // []string | columnIds (optional) (default to [])
    include := []string{"Include_example"} // []string | include (optional) (default to [])
    page := "page_example" // string | page (optional) (default to "{}")
    query := "query_example" // string | query (optional) (default to "{}")
    sort := "sort_example" // string | sort (optional) (default to "{}")

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.Get(context.Background(), viewId).ColumnIds(columnIds).Include(include).Page(page).Query(query).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: View
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnIds** | **[]string** | columnIds | [default to []]
 **include** | **[]string** | include | [default to []]
 **page** | **string** | page | [default to &quot;{}&quot;]
 **query** | **string** | query | [default to &quot;{}&quot;]
 **sort** | **string** | sort | [default to &quot;{}&quot;]

### Return type

[**View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatistic

> ViewStatistic GetStatistic(ctx, viewId).ColumnIds(columnIds).Execute()

getStatistic



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
    columnIds := []string{"Inner_example"} // []string | columnIds (optional) (default to [])

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.GetStatistic(context.Background(), viewId).ColumnIds(columnIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.GetStatistic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatistic`: ViewStatistic
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.GetStatistic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatisticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnIds** | **[]string** | columnIds | [default to []]

### Return type

[**ViewStatistic**](ViewStatistic.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportView

> ImportView(ctx, viewId).File(file).ImportRequest(importRequest).Type_(type_).Execute()

importView



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
    file := os.NewFile(1234, "some_file") // *os.File | The following file types are supported: csv, tsv, xls, xlsx and json
    importRequest := "importRequest_example" // string | importRequest (optional) (default to "{}")
    type_ := gridly.FileType("csv") // FileType | type (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.ImportView(context.Background(), viewId).File(file).ImportRequest(importRequest).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.ImportView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The following file types are supported: csv, tsv, xls, xlsx and json | 
 **importRequest** | **string** | importRequest | [default to &quot;{}&quot;]
 **type_** | [**FileType**](FileType.md) | type | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []View List(ctx).BranchId(branchId).GridId(gridId).Type_(type_).Execute()

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
    branchId := "branchId_example" // string | branchId (optional)
    gridId := "gridId_example" // string | gridId (optional)
    type_ := "type__example" // string | type (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.List(context.Background()).BranchId(branchId).GridId(gridId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []View
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **branchId** | **string** | branchId | 
 **gridId** | **string** | gridId | 
 **type_** | **string** | type | 

### Return type

[**[]View**](View.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Merge

> Task Merge(ctx, viewId).DestinationViewId(destinationViewId).MergeBranchRequest(mergeBranchRequest).MergeRecordOptions(mergeRecordOptions).Execute()

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
    destinationViewId := "destinationViewId_example" // string | destinationViewId
    viewId := "viewId_example" // string | viewId
    mergeBranchRequest := *gridly.NewMergeBranchRequest() // MergeBranchRequest | 
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional) (default to [])

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewApi.Merge(context.Background(), viewId).DestinationViewId(destinationViewId).MergeBranchRequest(mergeBranchRequest).MergeRecordOptions(mergeRecordOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Merge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Merge`: Task
    fmt.Fprintf(os.Stdout, "Response from `ViewApi.Merge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationViewId** | **string** | destinationViewId | 

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

