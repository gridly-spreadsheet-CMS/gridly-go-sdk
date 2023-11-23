# GridSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSourceLanguageStatus** | Pointer to **string** |  | [optional] 
**TranslatorCanViewAutomations** | Pointer to **bool** |  | [optional] 
**Categories** | Pointer to [**[]FileCategory**](FileCategory.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewGridSetting

`func NewGridSetting() *GridSetting`

NewGridSetting instantiates a new GridSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridSettingWithDefaults

`func NewGridSettingWithDefaults() *GridSetting`

NewGridSettingWithDefaults instantiates a new GridSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSourceLanguageStatus

`func (o *GridSetting) GetDefaultSourceLanguageStatus() string`

GetDefaultSourceLanguageStatus returns the DefaultSourceLanguageStatus field if non-nil, zero value otherwise.

### GetDefaultSourceLanguageStatusOk

`func (o *GridSetting) GetDefaultSourceLanguageStatusOk() (*string, bool)`

GetDefaultSourceLanguageStatusOk returns a tuple with the DefaultSourceLanguageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSourceLanguageStatus

`func (o *GridSetting) SetDefaultSourceLanguageStatus(v string)`

SetDefaultSourceLanguageStatus sets DefaultSourceLanguageStatus field to given value.

### HasDefaultSourceLanguageStatus

`func (o *GridSetting) HasDefaultSourceLanguageStatus() bool`

HasDefaultSourceLanguageStatus returns a boolean if a field has been set.

### GetTranslatorCanViewAutomations

`func (o *GridSetting) GetTranslatorCanViewAutomations() bool`

GetTranslatorCanViewAutomations returns the TranslatorCanViewAutomations field if non-nil, zero value otherwise.

### GetTranslatorCanViewAutomationsOk

`func (o *GridSetting) GetTranslatorCanViewAutomationsOk() (*bool, bool)`

GetTranslatorCanViewAutomationsOk returns a tuple with the TranslatorCanViewAutomations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatorCanViewAutomations

`func (o *GridSetting) SetTranslatorCanViewAutomations(v bool)`

SetTranslatorCanViewAutomations sets TranslatorCanViewAutomations field to given value.

### HasTranslatorCanViewAutomations

`func (o *GridSetting) HasTranslatorCanViewAutomations() bool`

HasTranslatorCanViewAutomations returns a boolean if a field has been set.

### GetCategories

`func (o *GridSetting) GetCategories() []FileCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GridSetting) GetCategoriesOk() (*[]FileCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GridSetting) SetCategories(v []FileCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GridSetting) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetMetadata

`func (o *GridSetting) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GridSetting) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GridSetting) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GridSetting) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedTime

`func (o *GridSetting) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GridSetting) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GridSetting) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *GridSetting) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *GridSetting) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *GridSetting) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *GridSetting) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *GridSetting) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GridSetting) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GridSetting) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GridSetting) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GridSetting) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *GridSetting) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *GridSetting) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *GridSetting) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *GridSetting) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


