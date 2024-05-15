# TransMem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjectIds** | Pointer to **[]int64** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**IsPausedConsuming** | Pointer to **bool** |  | [optional] 
**PopulateTranslationStatus** | Pointer to [**TranslationStatus**](TranslationStatus.md) |  | [optional] 
**ContextLookup** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**FuzzyMatch** | Pointer to **bool** |  | [optional] 
**AllowAlternative** | Pointer to **bool** |  | [optional] 
**AllowAlternativeHasSameRecordId** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransMem

`func NewTransMem(name string, ) *TransMem`

NewTransMem instantiates a new TransMem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransMemWithDefaults

`func NewTransMemWithDefaults() *TransMem`

NewTransMemWithDefaults instantiates a new TransMem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransMem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransMem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransMem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransMem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectIds

`func (o *TransMem) GetProjectIds() []int64`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *TransMem) GetProjectIdsOk() (*[]int64, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *TransMem) SetProjectIds(v []int64)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *TransMem) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetIsDisabled

`func (o *TransMem) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *TransMem) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *TransMem) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *TransMem) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsPausedConsuming

`func (o *TransMem) GetIsPausedConsuming() bool`

GetIsPausedConsuming returns the IsPausedConsuming field if non-nil, zero value otherwise.

### GetIsPausedConsumingOk

`func (o *TransMem) GetIsPausedConsumingOk() (*bool, bool)`

GetIsPausedConsumingOk returns a tuple with the IsPausedConsuming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPausedConsuming

`func (o *TransMem) SetIsPausedConsuming(v bool)`

SetIsPausedConsuming sets IsPausedConsuming field to given value.

### HasIsPausedConsuming

`func (o *TransMem) HasIsPausedConsuming() bool`

HasIsPausedConsuming returns a boolean if a field has been set.

### GetPopulateTranslationStatus

`func (o *TransMem) GetPopulateTranslationStatus() TranslationStatus`

GetPopulateTranslationStatus returns the PopulateTranslationStatus field if non-nil, zero value otherwise.

### GetPopulateTranslationStatusOk

`func (o *TransMem) GetPopulateTranslationStatusOk() (*TranslationStatus, bool)`

GetPopulateTranslationStatusOk returns a tuple with the PopulateTranslationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateTranslationStatus

`func (o *TransMem) SetPopulateTranslationStatus(v TranslationStatus)`

SetPopulateTranslationStatus sets PopulateTranslationStatus field to given value.

### HasPopulateTranslationStatus

`func (o *TransMem) HasPopulateTranslationStatus() bool`

HasPopulateTranslationStatus returns a boolean if a field has been set.

### GetContextLookup

`func (o *TransMem) GetContextLookup() bool`

GetContextLookup returns the ContextLookup field if non-nil, zero value otherwise.

### GetContextLookupOk

`func (o *TransMem) GetContextLookupOk() (*bool, bool)`

GetContextLookupOk returns a tuple with the ContextLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextLookup

`func (o *TransMem) SetContextLookup(v bool)`

SetContextLookup sets ContextLookup field to given value.

### HasContextLookup

`func (o *TransMem) HasContextLookup() bool`

HasContextLookup returns a boolean if a field has been set.

### GetName

`func (o *TransMem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransMem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransMem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TransMem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransMem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransMem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransMem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFuzzyMatch

`func (o *TransMem) GetFuzzyMatch() bool`

GetFuzzyMatch returns the FuzzyMatch field if non-nil, zero value otherwise.

### GetFuzzyMatchOk

`func (o *TransMem) GetFuzzyMatchOk() (*bool, bool)`

GetFuzzyMatchOk returns a tuple with the FuzzyMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuzzyMatch

`func (o *TransMem) SetFuzzyMatch(v bool)`

SetFuzzyMatch sets FuzzyMatch field to given value.

### HasFuzzyMatch

`func (o *TransMem) HasFuzzyMatch() bool`

HasFuzzyMatch returns a boolean if a field has been set.

### GetAllowAlternative

`func (o *TransMem) GetAllowAlternative() bool`

GetAllowAlternative returns the AllowAlternative field if non-nil, zero value otherwise.

### GetAllowAlternativeOk

`func (o *TransMem) GetAllowAlternativeOk() (*bool, bool)`

GetAllowAlternativeOk returns a tuple with the AllowAlternative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlternative

`func (o *TransMem) SetAllowAlternative(v bool)`

SetAllowAlternative sets AllowAlternative field to given value.

### HasAllowAlternative

`func (o *TransMem) HasAllowAlternative() bool`

HasAllowAlternative returns a boolean if a field has been set.

### GetAllowAlternativeHasSameRecordId

`func (o *TransMem) GetAllowAlternativeHasSameRecordId() bool`

GetAllowAlternativeHasSameRecordId returns the AllowAlternativeHasSameRecordId field if non-nil, zero value otherwise.

### GetAllowAlternativeHasSameRecordIdOk

`func (o *TransMem) GetAllowAlternativeHasSameRecordIdOk() (*bool, bool)`

GetAllowAlternativeHasSameRecordIdOk returns a tuple with the AllowAlternativeHasSameRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlternativeHasSameRecordId

`func (o *TransMem) SetAllowAlternativeHasSameRecordId(v bool)`

SetAllowAlternativeHasSameRecordId sets AllowAlternativeHasSameRecordId field to given value.

### HasAllowAlternativeHasSameRecordId

`func (o *TransMem) HasAllowAlternativeHasSameRecordId() bool`

HasAllowAlternativeHasSameRecordId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


