# UpdateTransMem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProjectIds** | Pointer to **[]int64** |  | [optional] 
**FuzzyMatch** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsPausedConsuming** | Pointer to **bool** |  | [optional] 
**PopulateTranslationStatus** | Pointer to [**TranslationStatus**](TranslationStatus.md) |  | [optional] 

## Methods

### NewUpdateTransMem

`func NewUpdateTransMem() *UpdateTransMem`

NewUpdateTransMem instantiates a new UpdateTransMem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransMemWithDefaults

`func NewUpdateTransMemWithDefaults() *UpdateTransMem`

NewUpdateTransMemWithDefaults instantiates a new UpdateTransMem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTransMem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTransMem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTransMem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTransMem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTransMem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTransMem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTransMem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTransMem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectIds

`func (o *UpdateTransMem) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *UpdateTransMem) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *UpdateTransMem) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *UpdateTransMem) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetFuzzyMatch

`func (o *UpdateTransMem) GetFuzzyMatch() bool`

GetFuzzyMatch returns the FuzzyMatch field if non-nil, zero value otherwise.

### GetFuzzyMatchOk

`func (o *UpdateTransMem) GetFuzzyMatchOk() (*bool, bool)`

GetFuzzyMatchOk returns a tuple with the FuzzyMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzyMatch

`func (o *UpdateTransMem) SetFuzzyMatch(v bool)`

SetFuzzyMatch sets FuzzyMatch field to given value.

### HasFuzzyMatch

`func (o *UpdateTransMem) HasFuzzyMatch() bool`

HasFuzzyMatch returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UpdateTransMem) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UpdateTransMem) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UpdateTransMem) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UpdateTransMem) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsPausedConsuming

`func (o *UpdateTransMem) GetIsPausedConsuming() bool`

GetIsPausedConsuming returns the IsPausedConsuming field if non-nil, zero value otherwise.

### GetIsPausedConsumingOk

`func (o *UpdateTransMem) GetIsPausedConsumingOk() (*bool, bool)`

GetIsPausedConsumingOk returns a tuple with the IsPausedConsuming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausedConsuming

`func (o *UpdateTransMem) SetIsPausedConsuming(v bool)`

SetIsPausedConsuming sets IsPausedConsuming field to given value.

### HasIsPausedConsuming

`func (o *UpdateTransMem) HasIsPausedConsuming() bool`

HasIsPausedConsuming returns a boolean if a field has been set.

### GetPopulateTranslationStatus

`func (o *UpdateTransMem) GetPopulateTranslationStatus() TranslationStatus`

GetPopulateTranslationStatus returns the PopulateTranslationStatus field if non-nil, zero value otherwise.

### GetPopulateTranslationStatusOk

`func (o *UpdateTransMem) GetPopulateTranslationStatusOk() (*TranslationStatus, bool)`

GetPopulateTranslationStatusOk returns a tuple with the PopulateTranslationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateTranslationStatus

`func (o *UpdateTransMem) SetPopulateTranslationStatus(v TranslationStatus)`

SetPopulateTranslationStatus sets PopulateTranslationStatus field to given value.

### HasPopulateTranslationStatus

`func (o *UpdateTransMem) HasPopulateTranslationStatus() bool`

HasPopulateTranslationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


