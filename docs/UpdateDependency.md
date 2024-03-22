# UpdateDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewId** | Pointer to **string** |  | [optional] 
**TargetColumnId** | **string** |  | 
**SourceColumnId** | **string** |  | 

## Methods

### NewUpdateDependency

`func NewUpdateDependency(targetColumnId string, sourceColumnId string, ) *UpdateDependency`

NewUpdateDependency instantiates a new UpdateDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDependencyWithDefaults

`func NewUpdateDependencyWithDefaults() *UpdateDependency`

NewUpdateDependencyWithDefaults instantiates a new UpdateDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewId

`func (o *UpdateDependency) GetNewId() string`

GetNewId returns the NewId field if non-nil, zero value otherwise.

### GetNewIdOk

`func (o *UpdateDependency) GetNewIdOk() (*string, bool)`

GetNewIdOk returns a tuple with the NewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewId

`func (o *UpdateDependency) SetNewId(v string)`

SetNewId sets NewId field to given value.

### HasNewId

`func (o *UpdateDependency) HasNewId() bool`

HasNewId returns a boolean if a field has been set.

### GetTargetColumnId

`func (o *UpdateDependency) GetTargetColumnId() string`

GetTargetColumnId returns the TargetColumnId field if non-nil, zero value otherwise.

### GetTargetColumnIdOk

`func (o *UpdateDependency) GetTargetColumnIdOk() (*string, bool)`

GetTargetColumnIdOk returns a tuple with the TargetColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetColumnId

`func (o *UpdateDependency) SetTargetColumnId(v string)`

SetTargetColumnId sets TargetColumnId field to given value.


### GetSourceColumnId

`func (o *UpdateDependency) GetSourceColumnId() string`

GetSourceColumnId returns the SourceColumnId field if non-nil, zero value otherwise.

### GetSourceColumnIdOk

`func (o *UpdateDependency) GetSourceColumnIdOk() (*string, bool)`

GetSourceColumnIdOk returns a tuple with the SourceColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColumnId

`func (o *UpdateDependency) SetSourceColumnId(v string)`

SetSourceColumnId sets SourceColumnId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


