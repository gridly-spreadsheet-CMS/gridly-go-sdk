# client\RecordApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RecordApi.md#Create) | **Post** /v1/views/{viewId}/records | create
[**Delete**](RecordApi.md#Delete) | **Delete** /v1/views/{viewId}/records | delete
[**Fetch**](RecordApi.md#Fetch) | **Get** /v1/views/{viewId}/records | fetch
[**Update**](RecordApi.md#Update) | **Patch** /v1/views/{viewId}/records | update
[**UpdateRecord**](RecordApi.md#UpdateRecord) | **Patch** /v1/views/{viewId}/records/{id} | updateRecord



## Create

> []Record Create(ctx, viewId).SetRecord(setRecord).Execute()

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
    setRecord := []gridly.SetRecord{*gridly.NewSetRecord()} // []SetRecord | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordApi.Create(context.Background(), viewId).SetRecord(setRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: []Record
    fmt.Fprintf(os.Stdout, "Response from `RecordApi.Create`: %v\n", resp)
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

 **setRecord** | [**[]SetRecord**](SetRecord.md) |  | 

### Return type

[**[]Record**](Record.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, viewId).DeleteRecord(deleteRecord).Execute()

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
    deleteRecord := *gridly.NewDeleteRecord() // DeleteRecord | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordApi.Delete(context.Background(), viewId).DeleteRecord(deleteRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordApi.Delete``: %v\n", err)
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

 **deleteRecord** | [**DeleteRecord**](DeleteRecord.md) |  | 

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


## Fetch

> []Record Fetch(ctx, viewId).ColumnIds(columnIds).Page(page).Query(query).Sort(sort).FetchFileOption(fetchFileOption).Execute()

fetch



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
    page := "page_example" // string | page (optional) (default to "{}")
    query := "query_example" // string | query (optional) (default to "{}")
    sort := "sort_example" // string | sort (optional) (default to "{}")
    fetchFileOption := gridly.FetchFileOption("all") // FetchFileOption | fetchFileOption (optional) (default to "id")

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordApi.Fetch(context.Background(), viewId).ColumnIds(columnIds).Page(page).Query(query).Sort(sort).FetchFileOption(fetchFileOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordApi.Fetch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Fetch`: []Record
    fmt.Fprintf(os.Stdout, "Response from `RecordApi.Fetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnIds** | **[]string** | columnIds | [default to []]
 **page** | **string** | page | [default to &quot;{}&quot;]
 **query** | **string** | query | [default to &quot;{}&quot;]
 **sort** | **string** | sort | [default to &quot;{}&quot;]
 **fetchFileOption** | [**FetchFileOption**](FetchFileOption.md) | fetchFileOption | [default to &quot;id&quot;]

### Return type

[**[]Record**](Record.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> []Record Update(ctx, viewId).SetRecord(setRecord).Execute()

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
    setRecord := []gridly.SetRecord{*gridly.NewSetRecord()} // []SetRecord | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordApi.Update(context.Background(), viewId).SetRecord(setRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: []Record
    fmt.Fprintf(os.Stdout, "Response from `RecordApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setRecord** | [**[]SetRecord**](SetRecord.md) |  | 

### Return type

[**[]Record**](Record.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecord

> Record UpdateRecord(ctx, id, viewId).SetRecord(setRecord).Path(path).Execute()

updateRecord



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
    id := "id_example" // string | id
    viewId := "viewId_example" // string | viewId
    setRecord := *gridly.NewSetRecord() // SetRecord | 
    path := "path_example" // string | path (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordApi.UpdateRecord(context.Background(), id, viewId).SetRecord(setRecord).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordApi.UpdateRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecord`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordApi.UpdateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setRecord** | [**SetRecord**](SetRecord.md) |  | 
 **path** | **string** | path | 

### Return type

[**Record**](Record.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

