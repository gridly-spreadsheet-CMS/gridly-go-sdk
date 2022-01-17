# api\ViewApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Export**](ViewApi.md#Export) | **Get** /v1/views/{viewId}/export | export
[**Get**](ViewApi.md#Get) | **Get** /v1/views/{viewId} | get
[**ImportView**](ViewApi.md#ImportView) | **Post** /v1/views/{viewId}/import | importView
[**List**](ViewApi.md#List) | **Get** /v1/views | list
[**Merge**](ViewApi.md#Merge) | **Post** /v1/views/{viewId}/merge | merge



## Export

> map[string]interface{} Export(ctx, viewId).ColumnIds(columnIds).Query(query).Sort(sort).Type_(type_).Execute()

export

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
    columnIds := []string{"Inner_example"} // []string | columnIds (optional)
    query := "query_example" // string | query (optional) (default to "{}")
    sort := "sort_example" // string | sort (optional) (default to "{}")
    type_ := "type__example" // string | type (optional) (default to "csv")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewApi.Export(context.Background(), viewId).ColumnIds(columnIds).Query(query).Sort(sort).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: map[string]interface{}
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

 **columnIds** | **[]string** | columnIds | 
 **query** | **string** | query | [default to &quot;{}&quot;]
 **sort** | **string** | sort | [default to &quot;{}&quot;]
 **type_** | **string** | type | [default to &quot;csv&quot;]

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "./openapi"
)

func main() {
    viewId := "viewId_example" // string | viewId
    columnIds := []string{"Inner_example"} // []string | columnIds (optional)
    include := []string{"Include_example"} // []string | include (optional)
    page := "page_example" // string | page (optional) (default to "{}")
    query := "query_example" // string | query (optional) (default to "{}")
    sort := "sort_example" // string | sort (optional) (default to "{}")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewApi.Get(context.Background(), viewId).ColumnIds(columnIds).Include(include).Page(page).Query(query).Sort(sort).Execute()
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

 **columnIds** | **[]string** | columnIds | 
 **include** | **[]string** | include | 
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
    openapiclient "./openapi"
)

func main() {
    viewId := "viewId_example" // string | viewId
    file := os.NewFile(1234, "some_file") // *os.File | The following file types are supported: csv, tsv, xls, xlsx and JSON
    importRequest := TODO // interface{} | importRequest (optional)
    type_ := TODO // interface{} | type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewApi.ImportView(context.Background(), viewId).File(file).ImportRequest(importRequest).Type_(type_).Execute()
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

 **file** | ***os.File** | The following file types are supported: csv, tsv, xls, xlsx and JSON | 
 **importRequest** | [**interface{}**](interface{}.md) | importRequest | 
 **type_** | [**interface{}**](interface{}.md) | type | 

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
    openapiclient "./openapi"
)

func main() {
    branchId := "branchId_example" // string | branchId (optional)
    gridId := "gridId_example" // string | gridId (optional)
    type_ := "type__example" // string | type (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewApi.List(context.Background()).BranchId(branchId).GridId(gridId).Type_(type_).Execute()
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

> Merge(ctx, viewId).DestinationViewId(destinationViewId).MergeRecordOptions(mergeRecordOptions).Execute()

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
    destinationViewId := "destinationViewId_example" // string | destinationViewId
    viewId := "viewId_example" // string | viewId
    mergeRecordOptions := []string{"MergeRecordOptions_example"} // []string | mergeRecordOptions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewApi.Merge(context.Background(), viewId).DestinationViewId(destinationViewId).MergeRecordOptions(mergeRecordOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewApi.Merge``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationViewId** | **string** | destinationViewId | 

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

