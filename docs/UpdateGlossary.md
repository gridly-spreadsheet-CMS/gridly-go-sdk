# UpdateGlossary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Langs** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to [**[]GlossaryProject**](GlossaryProject.md) |  | [optional] 

## Methods

### NewUpdateGlossary

`func NewUpdateGlossary() *UpdateGlossary`

NewUpdateGlossary instantiates a new UpdateGlossary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGlossaryWithDefaults

`func NewUpdateGlossaryWithDefaults() *UpdateGlossary`

NewUpdateGlossaryWithDefaults instantiates a new UpdateGlossary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGlossary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGlossary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGlossary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGlossary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGlossary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGlossary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGlossary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGlossary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLangs

`func (o *UpdateGlossary) GetLangs() []string`

GetLangs returns the Langs field if non-nil, zero value otherwise.

### GetLangsOk

`func (o *UpdateGlossary) GetLangsOk() (*[]string, bool)`

GetLangsOk returns a tuple with the Langs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangs

`func (o *UpdateGlossary) SetLangs(v []string)`

SetLangs sets Langs field to given value.

### HasLangs

`func (o *UpdateGlossary) HasLangs() bool`

HasLangs returns a boolean if a field has been set.

### GetProjects

`func (o *UpdateGlossary) GetProjects() []GlossaryProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UpdateGlossary) GetProjectsOk() (*[]GlossaryProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UpdateGlossary) SetProjects(v []GlossaryProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UpdateGlossary) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


