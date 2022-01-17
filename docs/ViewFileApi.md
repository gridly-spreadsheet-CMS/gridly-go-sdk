# api\ViewFileApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Download**](ViewFileApi.md#Download) | **Get** /v1/views/{viewId}/files/{fileId} | download
[**Upload**](ViewFileApi.md#Upload) | **Post** /v1/views/{viewId}/files | upload
[**UploadZip**](ViewFileApi.md#UploadZip) | **Post** /v1/views/{viewId}/files/zip | uploadZip



## Download

> map[string]interface{} Download(ctx, fileId, viewId).Execute()

download

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
    fileId := "fileId_example" // string | fileId
    viewId := "viewId_example" // string | viewId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewFileApi.Download(context.Background(), fileId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ViewFileApi.Download`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | fileId | 
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## Upload

> map[string]interface{} Upload(ctx, viewId).ColumnId(columnId).File(file).RecordId(recordId).Execute()

upload

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
    columnId := TODO // interface{} | columnId
    file := os.NewFile(1234, "some_file") // *os.File | file
    recordId := TODO // interface{} | recordId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewFileApi.Upload(context.Background(), viewId).ColumnId(columnId).File(file).RecordId(recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ViewFileApi.Upload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnId** | [**interface{}**](interface{}.md) | columnId | 
 **file** | ***os.File** | file | 
 **recordId** | [**interface{}**](interface{}.md) | recordId | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadZip

> map[string]interface{} UploadZip(ctx, viewId).ColumnId(columnId).File(file).FileMappings(fileMappings).Execute()

uploadZip

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
    columnId := TODO // interface{} | columnId
    file := os.NewFile(1234, "some_file") // *os.File | file
    fileMappings := TODO // interface{} | fileMappings

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ViewFileApi.UploadZip(context.Background(), viewId).ColumnId(columnId).File(file).FileMappings(fileMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.UploadZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadZip`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ViewFileApi.UploadZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** | viewId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **columnId** | [**interface{}**](interface{}.md) | columnId | 
 **file** | ***os.File** | file | 
 **fileMappings** | [**interface{}**](interface{}.md) | fileMappings | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

