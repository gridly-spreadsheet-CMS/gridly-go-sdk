# client\TransmemApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cleanup**](TransmemApi.md#Cleanup) | **Put** /v1/transmems/{tmId}/cleanup | Erases all the translation data of the provided tmId
[**Create**](TransmemApi.md#Create) | **Post** /v1/transmems | Create a new translation memory
[**CreateWithFile**](TransmemApi.md#CreateWithFile) | **Post** /v1/transmems/upload | Create a new translation memory by uploading tmx file
[**Delete**](TransmemApi.md#Delete) | **Delete** /v1/transmems/{tmId} | Delete a translation memory by id
[**Export**](TransmemApi.md#Export) | **Get** /v1/transmems/{tmId}/export | Export translation memory tmx file
[**Get**](TransmemApi.md#Get) | **Get** /v1/transmems/{tmId} | Get translation memory info by id
[**ImportTmx**](TransmemApi.md#ImportTmx) | **Post** /v1/transmems/{tmId}/import | Import a translation memory from tmx file
[**ListTM**](TransmemApi.md#ListTM) | **Get** /v1/transmems | List all available translation memories or create default one if there is no translation memory
[**Update**](TransmemApi.md#Update) | **Put** /v1/transmems/{tmId} | Update a translation memory



## Cleanup

> Cleanup(ctx, tmId).Execute()

Erases all the translation data of the provided tmId

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
    tmId := "tmId_example" // string | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Cleanup(context.Background(), tmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Cleanup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCleanupRequest struct via the builder pattern


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


## Create

> TransMem Create(ctx).CreateTransMem(createTransMem).Execute()

Create a new translation memory

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
    createTransMem := *gridly.NewCreateTransMem("Name_example") // CreateTransMem |  (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Create(context.Background()).CreateTransMem(createTransMem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransMem** | [**CreateTransMem**](CreateTransMem.md) |  | 

### Return type

[**TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWithFile

> TransMem CreateWithFile(ctx).File(file).Execute()

Create a new translation memory by uploading tmx file

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
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.CreateWithFile(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.CreateWithFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWithFile`: TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.CreateWithFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> TransMem Delete(ctx, tmId).Execute()

Delete a translation memory by id

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
    tmId := "tmId_example" // string | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Delete(context.Background(), tmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export

> *os.File Export(ctx, tmId).Format(format).SourceLang(sourceLang).TargetLangs(targetLangs).Execute()

Export translation memory tmx file

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
    tmId := "tmId_example" // string | 
    format := gridly.ExportFormat("tmx") // ExportFormat |  (optional)
    sourceLang := "sourceLang_example" // string |  (optional)
    targetLangs := []string{"Inner_example"} // []string |  (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Export(context.Background(), tmId).Format(format).SourceLang(sourceLang).TargetLangs(targetLangs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.Export`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**ExportFormat**](ExportFormat.md) |  | 
 **sourceLang** | **string** |  | 
 **targetLangs** | **[]string** |  | 

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

> TransMem Get(ctx, tmId).Execute()

Get translation memory info by id

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
    tmId := "tmId_example" // string | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Get(context.Background(), tmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTmx

> ImportTmx(ctx, tmId).File(file).Execute()

Import a translation memory from tmx file

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
    tmId := "tmId_example" // string | 
    file := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.ImportTmx(context.Background(), tmId).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.ImportTmx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportTmxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

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


## ListTM

> []TransMem ListTM(ctx).ProjectId(projectId).Execute()

List all available translation memories or create default one if there is no translation memory

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
    projectId := int64(789) // int64 |  (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.ListTM(context.Background()).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.ListTM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTM`: []TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.ListTM`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64** |  | 

### Return type

[**[]TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> TransMem Update(ctx, tmId).UpdateTransMem(updateTransMem).Execute()

Update a translation memory

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
    tmId := "tmId_example" // string | 
    updateTransMem := *gridly.NewUpdateTransMem() // UpdateTransMem |  (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.TransmemApi.Update(context.Background(), tmId).UpdateTransMem(updateTransMem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransmemApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: TransMem
    fmt.Fprintf(os.Stdout, "Response from `TransmemApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransMem** | [**UpdateTransMem**](UpdateTransMem.md) |  | 

### Return type

[**TransMem**](TransMem.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

