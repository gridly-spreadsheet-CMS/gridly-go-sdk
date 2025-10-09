# LQAModelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsSystem** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**LQAModelStatus**](LQAModelStatus.md) |  | [optional] 
**Categories** | Pointer to [**[]LQAModelCategoryResponse**](LQAModelCategoryResponse.md) |  | [optional] 
**Severities** | Pointer to [**[]LQAModelSeverityResponse**](LQAModelSeverityResponse.md) |  | [optional] 
**Projects** | Pointer to [**[]LQAModelProjectResponse**](LQAModelProjectResponse.md) |  | [optional] 

## Methods

### NewLQAModelResponse

`func NewLQAModelResponse() *LQAModelResponse`

NewLQAModelResponse instantiates a new LQAModelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLQAModelResponseWithDefaults

`func NewLQAModelResponseWithDefaults() *LQAModelResponse`

NewLQAModelResponseWithDefaults instantiates a new LQAModelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LQAModelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LQAModelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LQAModelResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LQAModelResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LQAModelResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LQAModelResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LQAModelResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LQAModelResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDefault

`func (o *LQAModelResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *LQAModelResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *LQAModelResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *LQAModelResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsSystem

`func (o *LQAModelResponse) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *LQAModelResponse) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *LQAModelResponse) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *LQAModelResponse) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetStatus

`func (o *LQAModelResponse) GetStatus() LQAModelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LQAModelResponse) GetStatusOk() (*LQAModelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LQAModelResponse) SetStatus(v LQAModelStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LQAModelResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategories

`func (o *LQAModelResponse) GetCategories() []LQAModelCategoryResponse`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *LQAModelResponse) GetCategoriesOk() (*[]LQAModelCategoryResponse, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *LQAModelResponse) SetCategories(v []LQAModelCategoryResponse)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *LQAModelResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeverities

`func (o *LQAModelResponse) GetSeverities() []LQAModelSeverityResponse`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *LQAModelResponse) GetSeveritiesOk() (*[]LQAModelSeverityResponse, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *LQAModelResponse) SetSeverities(v []LQAModelSeverityResponse)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *LQAModelResponse) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetProjects

`func (o *LQAModelResponse) GetProjects() []LQAModelProjectResponse`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *LQAModelResponse) GetProjectsOk() (*[]LQAModelProjectResponse, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *LQAModelResponse) SetProjects(v []LQAModelProjectResponse)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *LQAModelResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


