# client\AutomationApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkTrigger**](AutomationApi.md#BulkTrigger) | **Post** /v1/automations/trigger | 
[**Get**](AutomationApi.md#Get) | **Get** /v1/automations/{automationId} | 
[**GetExecution**](AutomationApi.md#GetExecution) | **Get** /v1/automations/{automationId}/executions/{executionId} | 
[**GetExecutions**](AutomationApi.md#GetExecutions) | **Get** /v1/automations/{automationId}/executions | 
[**List**](AutomationApi.md#List) | **Get** /v1/automations | 
[**Trigger**](AutomationApi.md#Trigger) | **Post** /v1/automations/{automationId}/trigger | 



## BulkTrigger

> []ExecutionResponse BulkTrigger(ctx).BulkExecutionTriggerRequest(bulkExecutionTriggerRequest).Execute()



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
	bulkExecutionTriggerRequest := *gridly.NewBulkExecutionTriggerRequest() // BulkExecutionTriggerRequest | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.BulkTrigger(context.Background()).BulkExecutionTriggerRequest(bulkExecutionTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.BulkTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkTrigger`: []ExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.BulkTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkExecutionTriggerRequest** | [**BulkExecutionTriggerRequest**](BulkExecutionTriggerRequest.md) |  | 

### Return type

[**[]ExecutionResponse**](ExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Automation Get(ctx, automationId).Execute()



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
	automationId := "automationId_example" // string | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.Get(context.Background(), automationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: Automation
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Automation**](Automation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> ExecutionResponse GetExecution(ctx, automationId, executionId).Execute()



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
	automationId := "automationId_example" // string | 
	executionId := "executionId_example" // string | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.GetExecution(context.Background(), automationId, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecution`: ExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExecutionResponse**](ExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutions

> []ExecutionResponse GetExecutions(ctx, automationId).Page(page).Size(size).Execute()



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
	automationId := "automationId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 0)
	size := int32(56) // int32 |  (optional) (default to 10)

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.GetExecutions(context.Background(), automationId).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutions`: []ExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **size** | **int32** |  | [default to 10]

### Return type

[**[]ExecutionResponse**](ExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Automation List(ctx).ViewId(viewId).Execute()



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
	viewId := "viewId_example" // string | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.List(context.Background()).ViewId(viewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: []Automation
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **string** |  | 

### Return type

[**[]Automation**](Automation.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Trigger

> ExecutionResponse Trigger(ctx, automationId).ExecutionTriggerRequest(executionTriggerRequest).Execute()



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
	automationId := "automationId_example" // string | 
	executionTriggerRequest := *gridly.NewExecutionTriggerRequest() // ExecutionTriggerRequest | 

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomationApi.Trigger(context.Background(), automationId).ExecutionTriggerRequest(executionTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.Trigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Trigger`: ExecutionResponse
	fmt.Fprintf(os.Stdout, "Response from `AutomationApi.Trigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**automationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **executionTriggerRequest** | [**ExecutionTriggerRequest**](ExecutionTriggerRequest.md) |  | 

### Return type

[**ExecutionResponse**](ExecutionResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

