# UpdateColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 
**LocalizationType** | Pointer to **string** |  | [optional] 
**SelectionOptions** | Pointer to **[]string** |  | [optional] 
**NumberFormat** | Pointer to [**NumberFormat**](NumberFormat.md) |  | [optional] 
**Reference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**Formula** | Pointer to [**Formula**](Formula.md) |  | [optional] 
**DateTimeFormat** | Pointer to [**DateTimeFormat**](DateTimeFormat.md) |  | [optional] 
**NewId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateColumn

`func NewUpdateColumn() *UpdateColumn`

NewUpdateColumn instantiates a new UpdateColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateColumnWithDefaults

`func NewUpdateColumnWithDefaults() *UpdateColumn`

NewUpdateColumnWithDefaults instantiates a new UpdateColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateColumn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateColumn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateColumn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateColumn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateColumn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateColumn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLanguageCode

`func (o *UpdateColumn) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *UpdateColumn) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *UpdateColumn) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *UpdateColumn) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetLocalizationType

`func (o *UpdateColumn) GetLocalizationType() string`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *UpdateColumn) GetLocalizationTypeOk() (*string, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *UpdateColumn) SetLocalizationType(v string)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *UpdateColumn) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### GetSelectionOptions

`func (o *UpdateColumn) GetSelectionOptions() []string`

GetSelectionOptions returns the SelectionOptions field if non-nil, zero value otherwise.

### GetSelectionOptionsOk

`func (o *UpdateColumn) GetSelectionOptionsOk() (*[]string, bool)`

GetSelectionOptionsOk returns a tuple with the SelectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionOptions

`func (o *UpdateColumn) SetSelectionOptions(v []string)`

SetSelectionOptions sets SelectionOptions field to given value.

### HasSelectionOptions

`func (o *UpdateColumn) HasSelectionOptions() bool`

HasSelectionOptions returns a boolean if a field has been set.

### GetNumberFormat

`func (o *UpdateColumn) GetNumberFormat() NumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *UpdateColumn) GetNumberFormatOk() (*NumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *UpdateColumn) SetNumberFormat(v NumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *UpdateColumn) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetReference

`func (o *UpdateColumn) GetReference() Reference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *UpdateColumn) GetReferenceOk() (*Reference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *UpdateColumn) SetReference(v Reference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *UpdateColumn) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetFormula

`func (o *UpdateColumn) GetFormula() Formula`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *UpdateColumn) GetFormulaOk() (*Formula, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *UpdateColumn) SetFormula(v Formula)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *UpdateColumn) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *UpdateColumn) GetDateTimeFormat() DateTimeFormat`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *UpdateColumn) GetDateTimeFormatOk() (*DateTimeFormat, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *UpdateColumn) SetDateTimeFormat(v DateTimeFormat)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *UpdateColumn) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetNewId

`func (o *UpdateColumn) GetNewId() string`

GetNewId returns the NewId field if non-nil, zero value otherwise.

### GetNewIdOk

`func (o *UpdateColumn) GetNewIdOk() (*string, bool)`

GetNewIdOk returns a tuple with the NewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewId

`func (o *UpdateColumn) SetNewId(v string)`

SetNewId sets NewId field to given value.

### HasNewId

`func (o *UpdateColumn) HasNewId() bool`

HasNewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


