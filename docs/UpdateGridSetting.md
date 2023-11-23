# UpdateGridSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDateTimeFormat** | Pointer to [**DateTimeFormat**](DateTimeFormat.md) |  | [optional] 
**DefaultSourceLanguageStatus** | Pointer to **string** |  | [optional] 
**TranslatorCanViewAutomations** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewUpdateGridSetting

`func NewUpdateGridSetting() *UpdateGridSetting`

NewUpdateGridSetting instantiates a new UpdateGridSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGridSettingWithDefaults

`func NewUpdateGridSettingWithDefaults() *UpdateGridSetting`

NewUpdateGridSettingWithDefaults instantiates a new UpdateGridSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDateTimeFormat

`func (o *UpdateGridSetting) GetDefaultDateTimeFormat() DateTimeFormat`

GetDefaultDateTimeFormat returns the DefaultDateTimeFormat field if non-nil, zero value otherwise.

### GetDefaultDateTimeFormatOk

`func (o *UpdateGridSetting) GetDefaultDateTimeFormatOk() (*DateTimeFormat, bool)`

GetDefaultDateTimeFormatOk returns a tuple with the DefaultDateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDateTimeFormat

`func (o *UpdateGridSetting) SetDefaultDateTimeFormat(v DateTimeFormat)`

SetDefaultDateTimeFormat sets DefaultDateTimeFormat field to given value.

### HasDefaultDateTimeFormat

`func (o *UpdateGridSetting) HasDefaultDateTimeFormat() bool`

HasDefaultDateTimeFormat returns a boolean if a field has been set.

### GetDefaultSourceLanguageStatus

`func (o *UpdateGridSetting) GetDefaultSourceLanguageStatus() string`

GetDefaultSourceLanguageStatus returns the DefaultSourceLanguageStatus field if non-nil, zero value otherwise.

### GetDefaultSourceLanguageStatusOk

`func (o *UpdateGridSetting) GetDefaultSourceLanguageStatusOk() (*string, bool)`

GetDefaultSourceLanguageStatusOk returns a tuple with the DefaultSourceLanguageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSourceLanguageStatus

`func (o *UpdateGridSetting) SetDefaultSourceLanguageStatus(v string)`

SetDefaultSourceLanguageStatus sets DefaultSourceLanguageStatus field to given value.

### HasDefaultSourceLanguageStatus

`func (o *UpdateGridSetting) HasDefaultSourceLanguageStatus() bool`

HasDefaultSourceLanguageStatus returns a boolean if a field has been set.

### GetTranslatorCanViewAutomations

`func (o *UpdateGridSetting) GetTranslatorCanViewAutomations() bool`

GetTranslatorCanViewAutomations returns the TranslatorCanViewAutomations field if non-nil, zero value otherwise.

### GetTranslatorCanViewAutomationsOk

`func (o *UpdateGridSetting) GetTranslatorCanViewAutomationsOk() (*bool, bool)`

GetTranslatorCanViewAutomationsOk returns a tuple with the TranslatorCanViewAutomations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatorCanViewAutomations

`func (o *UpdateGridSetting) SetTranslatorCanViewAutomations(v bool)`

SetTranslatorCanViewAutomations sets TranslatorCanViewAutomations field to given value.

### HasTranslatorCanViewAutomations

`func (o *UpdateGridSetting) HasTranslatorCanViewAutomations() bool`

HasTranslatorCanViewAutomations returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateGridSetting) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateGridSetting) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateGridSetting) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateGridSetting) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


