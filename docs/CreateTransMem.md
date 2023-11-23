# CreateTransMem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ProjectIds** | Pointer to **[]int64** |  | [optional] 
**FuzzyMatch** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsPausedConsuming** | Pointer to **bool** |  | [optional] 
**ContextLookup** | Pointer to **bool** |  | [optional] 
**PopulateTranslationStatus** | Pointer to [**TranslationStatus**](TranslationStatus.md) |  | [optional] 

## Methods

### NewCreateTransMem

`func NewCreateTransMem(name string, ) *CreateTransMem`

NewCreateTransMem instantiates a new CreateTransMem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransMemWithDefaults

`func NewCreateTransMemWithDefaults() *CreateTransMem`

NewCreateTransMemWithDefaults instantiates a new CreateTransMem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTransMem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransMem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransMem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTransMem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTransMem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTransMem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTransMem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectIds

`func (o *CreateTransMem) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *CreateTransMem) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *CreateTransMem) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *CreateTransMem) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetFuzzyMatch

`func (o *CreateTransMem) GetFuzzyMatch() bool`

GetFuzzyMatch returns the FuzzyMatch field if non-nil, zero value otherwise.

### GetFuzzyMatchOk

`func (o *CreateTransMem) GetFuzzyMatchOk() (*bool, bool)`

GetFuzzyMatchOk returns a tuple with the FuzzyMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzyMatch

`func (o *CreateTransMem) SetFuzzyMatch(v bool)`

SetFuzzyMatch sets FuzzyMatch field to given value.

### HasFuzzyMatch

`func (o *CreateTransMem) HasFuzzyMatch() bool`

HasFuzzyMatch returns a boolean if a field has been set.

### GetIsDisabled

`func (o *CreateTransMem) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateTransMem) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateTransMem) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateTransMem) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsPausedConsuming

`func (o *CreateTransMem) GetIsPausedConsuming() bool`

GetIsPausedConsuming returns the IsPausedConsuming field if non-nil, zero value otherwise.

### GetIsPausedConsumingOk

`func (o *CreateTransMem) GetIsPausedConsumingOk() (*bool, bool)`

GetIsPausedConsumingOk returns a tuple with the IsPausedConsuming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausedConsuming

`func (o *CreateTransMem) SetIsPausedConsuming(v bool)`

SetIsPausedConsuming sets IsPausedConsuming field to given value.

### HasIsPausedConsuming

`func (o *CreateTransMem) HasIsPausedConsuming() bool`

HasIsPausedConsuming returns a boolean if a field has been set.

### GetContextLookup

`func (o *CreateTransMem) GetContextLookup() bool`

GetContextLookup returns the ContextLookup field if non-nil, zero value otherwise.

### GetContextLookupOk

`func (o *CreateTransMem) GetContextLookupOk() (*bool, bool)`

GetContextLookupOk returns a tuple with the ContextLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLookup

`func (o *CreateTransMem) SetContextLookup(v bool)`

SetContextLookup sets ContextLookup field to given value.

### HasContextLookup

`func (o *CreateTransMem) HasContextLookup() bool`

HasContextLookup returns a boolean if a field has been set.

### GetPopulateTranslationStatus

`func (o *CreateTransMem) GetPopulateTranslationStatus() TranslationStatus`

GetPopulateTranslationStatus returns the PopulateTranslationStatus field if non-nil, zero value otherwise.

### GetPopulateTranslationStatusOk

`func (o *CreateTransMem) GetPopulateTranslationStatusOk() (*TranslationStatus, bool)`

GetPopulateTranslationStatusOk returns a tuple with the PopulateTranslationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateTranslationStatus

`func (o *CreateTransMem) SetPopulateTranslationStatus(v TranslationStatus)`

SetPopulateTranslationStatus sets PopulateTranslationStatus field to given value.

### HasPopulateTranslationStatus

`func (o *CreateTransMem) HasPopulateTranslationStatus() bool`

HasPopulateTranslationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


