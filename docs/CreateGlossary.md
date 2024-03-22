# CreateGlossary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Langs** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to [**[]GlossaryProject**](GlossaryProject.md) |  | [optional] 

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

### GetProjects

`func (o *CreateGlossary) GetProjects() []GlossaryProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateGlossary) GetProjectsOk() (*[]GlossaryProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateGlossary) SetProjects(v []GlossaryProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateGlossary) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


