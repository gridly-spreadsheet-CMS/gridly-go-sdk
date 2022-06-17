# client\ViewFileApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ViewFileApi.md#Delete) | **Delete** /v1/views/{viewId}/files | delete
[**Download**](ViewFileApi.md#Download) | **Get** /v1/views/{viewId}/files/{fileId} | download
[**Upload**](ViewFileApi.md#Upload) | **Post** /v1/views/{viewId}/files | upload
[**UploadZip**](ViewFileApi.md#UploadZip) | **Post** /v1/views/{viewId}/files/zip | uploadZip



## Delete

> Delete(ctx, viewId).ColumnId(columnId).RecordId(recordId).DeleteFile(deleteFile).Execute()

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
    columnId := "columnId_example" // string | columnId
    recordId := "recordId_example" // string | recordId
    viewId := "viewId_example" // string | viewId
    deleteFile := *gridly.NewDeleteFile([]string{"Ids_example"}) // DeleteFile | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewFileApi.Delete(context.Background(), viewId).ColumnId(columnId).RecordId(recordId).DeleteFile(deleteFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.Delete``: %v\n", err)
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
 **columnId** | **string** | columnId | 
 **recordId** | **string** | recordId | 

 **deleteFile** | [**DeleteFile**](DeleteFile.md) |  | 

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


## Download

> *os.File Download(ctx, fileId, viewId).Execute()

download



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
    fileId := "fileId_example" // string | fileId
    viewId := "viewId_example" // string | viewId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewFileApi.Download(context.Background(), fileId, viewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.Download``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Download`: *os.File
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

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> UploadedFile Upload(ctx, viewId).ColumnId(columnId).RecordId(recordId).File(file).Execute()

upload



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
    columnId := "columnId_example" // string | columnId
    recordId := "recordId_example" // string | recordId
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewFileApi.Upload(context.Background(), viewId).ColumnId(columnId).RecordId(recordId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: UploadedFile
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

 **columnId** | **string** | columnId | 
 **recordId** | **string** | recordId | 
 **file** | ***os.File** |  | 

### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadZip

> []Record UploadZip(ctx, viewId).ColumnId(columnId).FileMappings(fileMappings).File(file).Execute()

uploadZip



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
    columnId := "columnId_example" // string | columnId
    fileMappings := "fileMappings_example" // string | fileMappings
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewFileApi.UploadZip(context.Background(), viewId).ColumnId(columnId).FileMappings(fileMappings).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewFileApi.UploadZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadZip`: []Record
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

 **columnId** | **string** | columnId | 
 **fileMappings** | **string** | fileMappings | 
 **file** | ***os.File** |  | 

### Return type

[**[]Record**](Record.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

