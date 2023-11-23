# client\GridApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GridApi.md#Create) | **Post** /v1/grids | create
[**CreateCategory**](GridApi.md#CreateCategory) | **Post** /v1/grids/{gridId}/settings/categories | createCategory
[**Delete**](GridApi.md#Delete) | **Delete** /v1/grids/{gridId} | delete
[**DeleteCategory**](GridApi.md#DeleteCategory) | **Delete** /v1/grids/{gridId}/settings/categories/{categoryId} | deleteCategory
[**DeleteFile**](GridApi.md#DeleteFile) | **Delete** /v1/grids/{gridId}/settings/categories/{categoryId}/files/{fileId} | deleteFile
[**Get**](GridApi.md#Get) | **Get** /v1/grids/{gridId} | get
[**GetSetting**](GridApi.md#GetSetting) | **Get** /v1/grids/{gridId}/settings | getSetting
[**List**](GridApi.md#List) | **Get** /v1/grids | list
[**ListFiles**](GridApi.md#ListFiles) | **Get** /v1/grids/{gridId}/settings/files | listFiles
[**ListTemplateGrids**](GridApi.md#ListTemplateGrids) | **Get** /v1/template-grids | listTemplateGrids
[**Update**](GridApi.md#Update) | **Patch** /v1/grids/{gridId} | update
[**UpdateCategory**](GridApi.md#UpdateCategory) | **Put** /v1/grids/{gridId}/settings/categories/{categoryId} | updateCategory
[**UpdateSetting**](GridApi.md#UpdateSetting) | **Patch** /v1/grids/{gridId}/settings | updateSetting
[**UploadSettingFile**](GridApi.md#UploadSettingFile) | **Post** /v1/grids/{gridId}/settings/categories/{categoryId}/files | uploadSettingFile



## Create

> Grid Create(ctx).DbId(dbId).CreateGrid(createGrid).Execute()

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
    dbId := "dbId_example" // string | dbId
    createGrid := *gridly.NewCreateGrid("Name_example") // CreateGrid | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.Create(context.Background()).DbId(dbId).CreateGrid(createGrid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbId** | **string** | dbId | 
 **createGrid** | [**CreateGrid**](CreateGrid.md) |  | 

### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCategory

> FileCategory CreateCategory(ctx, gridId).CreateFileCategory(createFileCategory).Execute()

createCategory



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
    createFileCategory := *gridly.NewCreateFileCategory("Name_example") // CreateFileCategory | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.CreateCategory(context.Background(), gridId).CreateFileCategory(createFileCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: FileCategory
    fmt.Fprintf(os.Stdout, "Response from `GridApi.CreateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileCategory** | [**CreateFileCategory**](CreateFileCategory.md) |  | 

### Return type

[**FileCategory**](FileCategory.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, gridId).Execute()

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
    gridId := "gridId_example" // string | gridId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.Delete(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

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


## DeleteCategory

> DeleteCategory(ctx, gridId, categoryId).Execute()

deleteCategory



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
    categoryId := "categoryId_example" // string | categoryId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.DeleteCategory(context.Background(), gridId, categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**categoryId** | **string** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## DeleteFile

> DeleteFile(ctx, gridId, categoryId, fileId).Execute()

deleteFile



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
    categoryId := "categoryId_example" // string | categoryId
    fileId := "fileId_example" // string | fileId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.DeleteFile(context.Background(), gridId, categoryId, fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.DeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**categoryId** | **string** | categoryId | 
**fileId** | **string** | fileId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


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

> Grid Get(ctx, gridId).Execute()

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
    gridId := "gridId_example" // string | gridId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.Get(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> GridSetting GetSetting(ctx, gridId).Execute()

getSetting



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
    resp, r, err := apiClient.GridApi.GetSetting(context.Background(), gridId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.GetSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSetting`: GridSetting
    fmt.Fprintf(os.Stdout, "Response from `GridApi.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GridSetting**](GridSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Grid List(ctx).DbId(dbId).Execute()

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
    dbId := "dbId_example" // string | dbId

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.List(context.Background()).DbId(dbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbId** | **string** | dbId | 

### Return type

[**[]Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiles

> []SettingFile ListFiles(ctx, gridId).CategoryId(categoryId).Execute()

listFiles



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
    categoryId := []string{"Inner_example"} // []string | categoryId (optional)

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.ListFiles(context.Background(), gridId).CategoryId(categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.ListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiles`: []SettingFile
    fmt.Fprintf(os.Stdout, "Response from `GridApi.ListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryId** | **[]string** | categoryId | 

### Return type

[**[]SettingFile**](SettingFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTemplateGrids

> []Grid ListTemplateGrids(ctx).Execute()

listTemplateGrids



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

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.ListTemplateGrids(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.ListTemplateGrids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTemplateGrids`: []Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.ListTemplateGrids`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTemplateGridsRequest struct via the builder pattern


### Return type

[**[]Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Grid Update(ctx, gridId).UpdateGrid(updateGrid).Execute()

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
    gridId := "gridId_example" // string | gridId
    updateGrid := *gridly.NewUpdateGrid() // UpdateGrid | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.Update(context.Background(), gridId).UpdateGrid(updateGrid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Grid
    fmt.Fprintf(os.Stdout, "Response from `GridApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGrid** | [**UpdateGrid**](UpdateGrid.md) |  | 

### Return type

[**Grid**](Grid.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> FileCategory UpdateCategory(ctx, gridId, categoryId).UpdateCategory(updateCategory).Execute()

updateCategory



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
    categoryId := "categoryId_example" // string | categoryId
    updateCategory := *gridly.NewUpdateCategory("Name_example") // UpdateCategory | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.UpdateCategory(context.Background(), gridId, categoryId).UpdateCategory(updateCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: FileCategory
    fmt.Fprintf(os.Stdout, "Response from `GridApi.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**categoryId** | **string** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCategory** | [**UpdateCategory**](UpdateCategory.md) |  | 

### Return type

[**FileCategory**](FileCategory.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSetting

> GridSetting UpdateSetting(ctx, gridId).UpdateGridSetting(updateGridSetting).Execute()

updateSetting



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
    updateGridSetting := *gridly.NewUpdateGridSetting() // UpdateGridSetting | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.UpdateSetting(context.Background(), gridId).UpdateGridSetting(updateGridSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.UpdateSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSetting`: GridSetting
    fmt.Fprintf(os.Stdout, "Response from `GridApi.UpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGridSetting** | [**UpdateGridSetting**](UpdateGridSetting.md) |  | 

### Return type

[**GridSetting**](GridSetting.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSettingFile

> UploadedFile UploadSettingFile(ctx, gridId, categoryId).UploadSettingFileRequest(uploadSettingFileRequest).Execute()

uploadSettingFile



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
    categoryId := "categoryId_example" // string | categoryId
    uploadSettingFileRequest := *gridly.NewUploadSettingFileRequest() // UploadSettingFileRequest | 

    configuration := gridly.NewConfiguration()
    apiClient := gridly.NewAPIClient(configuration)
    resp, r, err := apiClient.GridApi.UploadSettingFile(context.Background(), gridId, categoryId).UploadSettingFileRequest(uploadSettingFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GridApi.UploadSettingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadSettingFile`: UploadedFile
    fmt.Fprintf(os.Stdout, "Response from `GridApi.UploadSettingFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gridId** | **string** | gridId | 
**categoryId** | **string** | categoryId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSettingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadSettingFileRequest** | [**UploadSettingFileRequest**](UploadSettingFileRequest.md) |  | 

### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

