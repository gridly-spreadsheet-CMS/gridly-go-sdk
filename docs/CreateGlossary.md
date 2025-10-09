# CreateGlossary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Langs** | Pointer to **[]string** |  | [optional] 
**ProjectIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCreateGlossary

`func NewCreateGlossary(name string, ) *CreateGlossary`

NewCreateGlossary instantiates a new CreateGlossary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGlossaryWithDefaults

`func NewCreateGlossaryWithDefaults() *CreateGlossary`

NewCreateGlossaryWithDefaults instantiates a new CreateGlossary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGlossary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGlossary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGlossary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGlossary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGlossary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGlossary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGlossary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLangs

`func (o *CreateGlossary) GetLangs() []string`

GetLangs returns the Langs field if non-nil, zero value otherwise.

### GetLangsOk

`func (o *CreateGlossary) GetLangsOk() (*[]string, bool)`

GetLangsOk returns a tuple with the Langs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangs

`func (o *CreateGlossary) SetLangs(v []string)`

SetLangs sets Langs field to given value.

### HasLangs

`func (o *CreateGlossary) HasLangs() bool`

HasLangs returns a boolean if a field has been set.

### GetProjectIds

`func (o *CreateGlossary) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateGlossary) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateGlossary) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateGlossary) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


