# Go API client for gridly

Gridly API documentation

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.31.0
- Package version: 1.2.7
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit []()

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import gridly "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), gridly.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), gridly.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), gridly.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gridly.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.gridly.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BranchApi* | [**Create**](docs/BranchApi.md#create) | **Post** /v1/branches | create
*BranchApi* | [**Delete**](docs/BranchApi.md#delete) | **Delete** /v1/branches/{branchId} | delete
*BranchApi* | [**Get**](docs/BranchApi.md#get) | **Get** /v1/branches/{branchId} | get
*BranchApi* | [**List**](docs/BranchApi.md#list) | **Get** /v1/branches | list
*BranchApi* | [**Merge**](docs/BranchApi.md#merge) | **Post** /v1/branches/{branchId}/merge | merge
*DatabaseApi* | [**Create**](docs/DatabaseApi.md#create) | **Post** /v1/databases | create
*DatabaseApi* | [**Delete**](docs/DatabaseApi.md#delete) | **Delete** /v1/databases/{dbId} | delete
*DatabaseApi* | [**Duplicate**](docs/DatabaseApi.md#duplicate) | **Post** /v1/databases/{dbId}/duplicate | duplicate
*DatabaseApi* | [**Get**](docs/DatabaseApi.md#get) | **Get** /v1/databases/{dbId} | get
*DatabaseApi* | [**List**](docs/DatabaseApi.md#list) | **Get** /v1/databases | list
*DatabaseApi* | [**Update**](docs/DatabaseApi.md#update) | **Put** /v1/databases/{dbId} | update
*GridApi* | [**Create**](docs/GridApi.md#create) | **Post** /v1/grids | create
*GridApi* | [**Delete**](docs/GridApi.md#delete) | **Delete** /v1/grids/{gridId} | delete
*GridApi* | [**Get**](docs/GridApi.md#get) | **Get** /v1/grids/{gridId} | get
*GridApi* | [**List**](docs/GridApi.md#list) | **Get** /v1/grids | list
*GridApi* | [**ListTemplateGrids**](docs/GridApi.md#listtemplategrids) | **Get** /v1/template-grids | listTemplateGrids
*GridApi* | [**Update**](docs/GridApi.md#update) | **Patch** /v1/grids/{gridId} | update
*ProjectApi* | [**Create**](docs/ProjectApi.md#create) | **Post** /v1/projects | create
*ProjectApi* | [**Delete**](docs/ProjectApi.md#delete) | **Delete** /v1/projects/{projectId} | delete
*ProjectApi* | [**FindOne**](docs/ProjectApi.md#findone) | **Get** /v1/projects/{projectId} | findOne
*ProjectApi* | [**List**](docs/ProjectApi.md#list) | **Get** /v1/projects | list
*ProjectApi* | [**Update**](docs/ProjectApi.md#update) | **Put** /v1/projects/{projectId} | update
*RecordApi* | [**Create**](docs/RecordApi.md#create) | **Post** /v1/views/{viewId}/records | create
*RecordApi* | [**Delete**](docs/RecordApi.md#delete) | **Delete** /v1/views/{viewId}/records | delete
*RecordApi* | [**Fetch**](docs/RecordApi.md#fetch) | **Get** /v1/views/{viewId}/records | fetch
*RecordApi* | [**Update**](docs/RecordApi.md#update) | **Patch** /v1/views/{viewId}/records | update
*RecordApi* | [**UpdateRecord**](docs/RecordApi.md#updaterecord) | **Patch** /v1/views/{viewId}/records/{id} | updateRecord
*TaskApi* | [**Get**](docs/TaskApi.md#get) | **Get** /v1/tasks/{taskId} | get
*ViewApi* | [**Create**](docs/ViewApi.md#create) | **Post** /v1/views | create
*ViewApi* | [**Export**](docs/ViewApi.md#export) | **Get** /v1/views/{viewId}/export | export
*ViewApi* | [**Get**](docs/ViewApi.md#get) | **Get** /v1/views/{viewId} | get
*ViewApi* | [**ImportView**](docs/ViewApi.md#importview) | **Post** /v1/views/{viewId}/import | importView
*ViewApi* | [**List**](docs/ViewApi.md#list) | **Get** /v1/views | list
*ViewApi* | [**Merge**](docs/ViewApi.md#merge) | **Post** /v1/views/{viewId}/merge | merge
*ViewColumnApi* | [**Add**](docs/ViewColumnApi.md#add) | **Post** /v1/views/{viewId}/columns/{columnId}/add | add
*ViewColumnApi* | [**Create**](docs/ViewColumnApi.md#create) | **Post** /v1/views/{viewId}/columns | create
*ViewColumnApi* | [**Delete**](docs/ViewColumnApi.md#delete) | **Delete** /v1/views/{viewId}/columns/{columnId} | delete
*ViewColumnApi* | [**Get**](docs/ViewColumnApi.md#get) | **Get** /v1/views/{viewId}/columns/{columnId} | get
*ViewColumnApi* | [**Remove**](docs/ViewColumnApi.md#remove) | **Post** /v1/views/{viewId}/columns/{columnId}/remove | remove
*ViewColumnApi* | [**Update**](docs/ViewColumnApi.md#update) | **Patch** /v1/views/{viewId}/columns/{columnId} | update
*ViewDependencyApi* | [**Create**](docs/ViewDependencyApi.md#create) | **Post** /v1/views/{viewId}/dependencies | create
*ViewDependencyApi* | [**Delete**](docs/ViewDependencyApi.md#delete) | **Delete** /v1/views/{viewId}/dependencies | delete
*ViewDependencyApi* | [**DeleteById**](docs/ViewDependencyApi.md#deletebyid) | **Delete** /v1/views/{viewId}/dependencies/{dependencyId} | deleteById
*ViewDependencyApi* | [**Get**](docs/ViewDependencyApi.md#get) | **Get** /v1/views/{viewId}/dependencies/{dependencyId} | get
*ViewDependencyApi* | [**List**](docs/ViewDependencyApi.md#list) | **Get** /v1/views/{viewId}/dependencies | list
*ViewDependencyApi* | [**Update**](docs/ViewDependencyApi.md#update) | **Put** /v1/views/{viewId}/dependencies/{dependencyId} | update
*ViewFileApi* | [**Delete**](docs/ViewFileApi.md#delete) | **Delete** /v1/views/{viewId}/files | delete
*ViewFileApi* | [**Download**](docs/ViewFileApi.md#download) | **Get** /v1/views/{viewId}/files/{fileId} | download
*ViewFileApi* | [**Upload**](docs/ViewFileApi.md#upload) | **Post** /v1/views/{viewId}/files | upload
*ViewFileApi* | [**UploadZip**](docs/ViewFileApi.md#uploadzip) | **Post** /v1/views/{viewId}/files/zip | uploadZip


## Documentation For Models

 - [AddViewColumn](docs/AddViewColumn.md)
 - [Branch](docs/Branch.md)
 - [Cell](docs/Cell.md)
 - [ColumnReference](docs/ColumnReference.md)
 - [CreateBranch](docs/CreateBranch.md)
 - [CreateColumn](docs/CreateColumn.md)
 - [CreateDatabase](docs/CreateDatabase.md)
 - [CreateDependency](docs/CreateDependency.md)
 - [CreateGrid](docs/CreateGrid.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateView](docs/CreateView.md)
 - [Database](docs/Database.md)
 - [DeleteDependency](docs/DeleteDependency.md)
 - [DeleteFile](docs/DeleteFile.md)
 - [DeleteRecord](docs/DeleteRecord.md)
 - [Dependency](docs/Dependency.md)
 - [ExportFileHeader](docs/ExportFileHeader.md)
 - [FetchFileOption](docs/FetchFileOption.md)
 - [Grid](docs/Grid.md)
 - [Group](docs/Group.md)
 - [NumberFormat](docs/NumberFormat.md)
 - [Privilege](docs/Privilege.md)
 - [Project](docs/Project.md)
 - [ProjectDetail](docs/ProjectDetail.md)
 - [Record](docs/Record.md)
 - [RecordIdentifierWrapper](docs/RecordIdentifierWrapper.md)
 - [Reference](docs/Reference.md)
 - [ReferencedColumn](docs/ReferencedColumn.md)
 - [ReferencedGrid](docs/ReferencedGrid.md)
 - [Role](docs/Role.md)
 - [SetCell](docs/SetCell.md)
 - [SetRecord](docs/SetRecord.md)
 - [Task](docs/Task.md)
 - [UpdateColumn](docs/UpdateColumn.md)
 - [UpdateDatabase](docs/UpdateDatabase.md)
 - [UpdateDependency](docs/UpdateDependency.md)
 - [UpdateGrid](docs/UpdateGrid.md)
 - [UpdateProject](docs/UpdateProject.md)
 - [UploadZipRequest](docs/UploadZipRequest.md)
 - [UploadedFile](docs/UploadedFile.md)
 - [View](docs/View.md)
 - [ViewColumn](docs/ViewColumn.md)


## Documentation For Authorization



### ApiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



