# client\ProjectApi

All URIs are relative to *https://api.gridly.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ProjectApi.md#Create) | **Post** /v1/projects | create
[**Delete**](ProjectApi.md#Delete) | **Delete** /v1/projects/{projectId} | delete
[**FindOneDetail**](ProjectApi.md#FindOneDetail) | **Get** /v1/projects/{projectId} | findOneDetail
[**List**](ProjectApi.md#List) | **Get** /v1/projects | list
[**Update**](ProjectApi.md#Update) | **Put** /v1/projects/{projectId} | update



## Create

> Project Create(ctx).CreateProject(createProject).Execute()

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
    createProject := *gridly.NewCreateProject("Name_example") // CreateProject | createProject

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.Create(context.Background()).CreateProject(createProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProject** | [**CreateProject**](CreateProject.md) | createProject | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, projectId).Execute()

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
    projectId := int64(789) // int64 | projectId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.Delete(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | projectId | 

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


## FindOneDetail

> ProjectDetail FindOneDetail(ctx, projectId).Expand(expand).Execute()

findOneDetail

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
    expand := "expand_example" // string | expand
    projectId := int64(789) // int64 | projectId

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.FindOneDetail(context.Background(), projectId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.FindOneDetail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindOneDetail`: ProjectDetail
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.FindOneDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindOneDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | expand | 


### Return type

[**ProjectDetail**](ProjectDetail.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []Project List(ctx).Execute()

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

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.List(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.List`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


### Return type

[**[]Project**](Project.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Project Update(ctx, projectId).UpdateProject(updateProject).Execute()

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
    projectId := int64(789) // int64 | projectId
    updateProject := *gridly.NewUpdateProject("Name_example") // UpdateProject | updateProject

    configuration := gridly.NewConfiguration()
    api_client := gridly.NewAPIClient(configuration)
    resp, r, err := api_client.ProjectApi.Update(context.Background(), projectId).UpdateProject(updateProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProject** | [**UpdateProject**](UpdateProject.md) | updateProject | 

### Return type

[**Project**](Project.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

