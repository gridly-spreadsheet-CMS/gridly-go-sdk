# CreateColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**LanguageCode** | Pointer to **string** |  | [optional] 
**LocalizationType** | Pointer to **string** |  | [optional] 
**NumberFormat** | Pointer to [**NumberFormat**](NumberFormat.md) |  | [optional] 
**SelectionOptions** | Pointer to **[]string** |  | [optional] 
**Reference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**Formula** | Pointer to [**Formula**](Formula.md) |  | [optional] 
**DateTimeFormat** | Pointer to [**DateTimeFormat**](DateTimeFormat.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateColumn

`func NewCreateColumn(name string, type_ string, ) *CreateColumn`

NewCreateColumn instantiates a new CreateColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateColumnWithDefaults

`func NewCreateColumnWithDefaults() *CreateColumn`

NewCreateColumnWithDefaults instantiates a new CreateColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateColumn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateColumn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateColumn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateColumn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateColumn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateColumn) SetType(v string)`

SetType sets Type field to given value.


### GetLanguageCode

`func (o *CreateColumn) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *CreateColumn) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *CreateColumn) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *CreateColumn) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetLocalizationType

`func (o *CreateColumn) GetLocalizationType() string`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *CreateColumn) GetLocalizationTypeOk() (*string, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *CreateColumn) SetLocalizationType(v string)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *CreateColumn) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### GetNumberFormat

`func (o *CreateColumn) GetNumberFormat() NumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *CreateColumn) GetNumberFormatOk() (*NumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *CreateColumn) SetNumberFormat(v NumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *CreateColumn) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetSelectionOptions

`func (o *CreateColumn) GetSelectionOptions() []string`

GetSelectionOptions returns the SelectionOptions field if non-nil, zero value otherwise.

### GetSelectionOptionsOk

`func (o *CreateColumn) GetSelectionOptionsOk() (*[]string, bool)`

GetSelectionOptionsOk returns a tuple with the SelectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionOptions

`func (o *CreateColumn) SetSelectionOptions(v []string)`

SetSelectionOptions sets SelectionOptions field to given value.

### HasSelectionOptions

`func (o *CreateColumn) HasSelectionOptions() bool`

HasSelectionOptions returns a boolean if a field has been set.

### GetReference

`func (o *CreateColumn) GetReference() Reference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateColumn) GetReferenceOk() (*Reference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateColumn) SetReference(v Reference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateColumn) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetFormula

`func (o *CreateColumn) GetFormula() Formula`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *CreateColumn) GetFormulaOk() (*Formula, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *CreateColumn) SetFormula(v Formula)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *CreateColumn) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *CreateColumn) GetDateTimeFormat() DateTimeFormat`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *CreateColumn) GetDateTimeFormatOk() (*DateTimeFormat, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *CreateColumn) SetDateTimeFormat(v DateTimeFormat)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *CreateColumn) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetId

`func (o *CreateColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateColumn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateColumn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateColumn) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


