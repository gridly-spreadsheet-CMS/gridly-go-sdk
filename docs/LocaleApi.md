# client\LocaleApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](LocaleApi.md#Create) | **Get** /v1/locales/all | create



## Create

> map[string]string Create(ctx).Execute()

create



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
	resp, r, err := apiClient.LocaleApi.Create(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaleApi.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `LocaleApi.Create`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


### Return type

**map[string]string**

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

