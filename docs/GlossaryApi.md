# client\GlossaryApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GlossaryApi.md#Create) | **Post** /v1/glossaries | Create a new glossary
[**Delete**](GlossaryApi.md#Delete) | **Delete** /v1/glossaries/{id} | Delete a glossary
[**ExportFile**](GlossaryApi.md#ExportFile) | **Get** /v1/glossaries/{id}/export | Export a glossary
[**Get**](GlossaryApi.md#Get) | **Get** /v1/glossaries/{id} | Get glossary info
[**GetAll**](GlossaryApi.md#GetAll) | **Get** /v1/glossaries | List all glossaries
[**ImportFile**](GlossaryApi.md#ImportFile) | **Post** /v1/glossaries/{id}/import | Import a glossary from file
[**Update**](GlossaryApi.md#Update) | **Put** /v1/glossaries/{id} | Update glossary info



## Create

> Glossary Create(ctx).CreateGlossary(createGlossary).Execute()

Create a new glossary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	createGlossary := *gridly.NewCreateGlossary("Name_example") // CreateGlossary | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.GlossaryApi.Create(context.Background()).CreateGlossary(createGlossary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: Glossary
	fmt.Fprintf(os.Stdout, "Response from `GlossaryApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGlossary** | [**CreateGlossary**](CreateGlossary.md) |  | 

### Return type

[**Glossary**](Glossary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()

Delete a glossary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	r, err := apiClient.GlossaryApi.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

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


## ExportFile

> *os.File ExportFile(ctx, id).Fields(fields).Format(format).Langs(langs).Execute()

Export a glossary

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	id := "id_example" // string | 
	fields := []string{"Inner_example"} // []string |  (optional)
	format := gridly.ExportGlossaryFormat("csv") // ExportGlossaryFormat |  (optional) (default to "tbx")
	langs := []string{"Inner_example"} // []string |  (optional)

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.GlossaryApi.ExportFile(context.Background(), id).Fields(fields).Format(format).Langs(langs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.ExportFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `GlossaryApi.ExportFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** |  | 
 **format** | [**ExportGlossaryFormat**](ExportGlossaryFormat.md) |  | [default to &quot;tbx&quot;]
 **langs** | **[]string** |  | 

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

> Glossary Get(ctx, id).Execute()

Get glossary info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	id := "id_example" // string | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.GlossaryApi.Get(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: Glossary
	fmt.Fprintf(os.Stdout, "Response from `GlossaryApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Glossary**](Glossary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> []Glossary GetAll(ctx).Execute()

List all glossaries

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.GlossaryApi.GetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: []Glossary
	fmt.Fprintf(os.Stdout, "Response from `GlossaryApi.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**[]Glossary**](Glossary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFile

> ImportFile(ctx, id).ImportOption(importOption).File(file).Execute()

Import a glossary from file

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	id := "id_example" // string | 
	importOption := gridly.ImportGlossaryOption("ADD") // ImportGlossaryOption |  (optional) (default to "UPDATE")
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	r, err := apiClient.GlossaryApi.ImportFile(context.Background(), id).ImportOption(importOption).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.ImportFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importOption** | [**ImportGlossaryOption**](ImportGlossaryOption.md) |  | [default to &quot;UPDATE&quot;]
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Glossary Update(ctx, id).UpdateGlossary(updateGlossary).Execute()

Update glossary info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	gridly "github.com/gridly-spreadsheet-CMS/gridly-go-sdk"
)

func main() {
	id := "id_example" // string | 
	updateGlossary := *gridly.NewUpdateGlossary() // UpdateGlossary | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.GlossaryApi.Update(context.Background(), id).UpdateGlossary(updateGlossary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlossaryApi.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: Glossary
	fmt.Fprintf(os.Stdout, "Response from `GlossaryApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGlossary** | [**UpdateGlossary**](UpdateGlossary.md) |  | 

### Return type

[**Glossary**](Glossary.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

