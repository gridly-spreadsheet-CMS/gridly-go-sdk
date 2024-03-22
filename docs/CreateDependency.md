# CreateDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TargetColumnId** | **string** |  | 
**SourceColumnId** | **string** |  | 

## Methods

### NewCreateDependency

`func NewCreateDependency(targetColumnId string, sourceColumnId string, ) *CreateDependency`

NewCreateDependency instantiates a new CreateDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDependencyWithDefaults

`func NewCreateDependencyWithDefaults() *CreateDependency`

NewCreateDependencyWithDefaults instantiates a new CreateDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateDependency) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateDependency) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateDependency) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateDependency) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetColumnId

`func (o *CreateDependency) GetTargetColumnId() string`

GetTargetColumnId returns the TargetColumnId field if non-nil, zero value otherwise.

### GetTargetColumnIdOk

`func (o *CreateDependency) GetTargetColumnIdOk() (*string, bool)`

GetTargetColumnIdOk returns a tuple with the TargetColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetColumnId

`func (o *CreateDependency) SetTargetColumnId(v string)`

SetTargetColumnId sets TargetColumnId field to given value.


### GetSourceColumnId

`func (o *CreateDependency) GetSourceColumnId() string`

GetSourceColumnId returns the SourceColumnId field if non-nil, zero value otherwise.

### GetSourceColumnIdOk

`func (o *CreateDependency) GetSourceColumnIdOk() (*string, bool)`

GetSourceColumnIdOk returns a tuple with the SourceColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceColumnId

`func (o *CreateDependency) SetSourceColumnId(v string)`

SetSourceColumnId sets SourceColumnId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


