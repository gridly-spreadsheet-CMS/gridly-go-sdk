# Glossary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Langs** | Pointer to **[]string** |  | [optional] 
**Projects** | Pointer to [**[]GlossaryProject**](GlossaryProject.md) |  | [optional] 

## Methods

### NewGlossary

`func NewGlossary() *Glossary`

NewGlossary instantiates a new Glossary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlossaryWithDefaults

`func NewGlossaryWithDefaults() *Glossary`

NewGlossaryWithDefaults instantiates a new Glossary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Glossary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Glossary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Glossary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Glossary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Glossary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Glossary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Glossary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Glossary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Glossary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Glossary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Glossary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Glossary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLangs

`func (o *Glossary) GetLangs() []string`

GetLangs returns the Langs field if non-nil, zero value otherwise.

### GetLangsOk

`func (o *Glossary) GetLangsOk() (*[]string, bool)`

GetLangsOk returns a tuple with the Langs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangs

`func (o *Glossary) SetLangs(v []string)`

SetLangs sets Langs field to given value.

### HasLangs

`func (o *Glossary) HasLangs() bool`

HasLangs returns a boolean if a field has been set.

### GetProjects

`func (o *Glossary) GetProjects() []GlossaryProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Glossary) GetProjectsOk() (*[]GlossaryProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Glossary) SetProjects(v []GlossaryProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Glossary) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


