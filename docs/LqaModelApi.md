# client\LqaModelApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLQAModel**](LqaModelApi.md#GetLQAModel) | **Get** /v1/lqa-models | Get default lqa model



## GetLQAModel

> LQAModelResponse GetLQAModel(ctx).ProjectId(projectId).Execute()

Get default lqa model

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
	projectId := int64(789) // int64 |  (optional)

	configuration := gridly.NewConfiguration()
	apiClient := gridly.NewAPIClient(configuration)
	resp, r, err := apiClient.LqaModelApi.GetLQAModel(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LqaModelApi.GetLQAModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLQAModel`: LQAModelResponse
	fmt.Fprintf(os.Stdout, "Response from `LqaModelApi.GetLQAModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLQAModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64** |  | 

### Return type

[**LQAModelResponse**](LQAModelResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

