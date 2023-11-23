# AddViewColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 
**LocalizationType** | Pointer to **string** |  | [optional] 
**NumberFormat** | Pointer to [**NumberFormat**](NumberFormat.md) |  | [optional] 
**SelectionOptions** | Pointer to **[]string** |  | [optional] 
**Reference** | Pointer to [**Reference**](Reference.md) |  | [optional] 
**Formula** | Pointer to [**Formula**](Formula.md) |  | [optional] 
**DateTimeFormat** | Pointer to [**DateTimeFormat**](DateTimeFormat.md) |  | [optional] 

## Methods

### NewAddViewColumn

`func NewAddViewColumn() *AddViewColumn`

NewAddViewColumn instantiates a new AddViewColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddViewColumnWithDefaults

`func NewAddViewColumnWithDefaults() *AddViewColumn`

NewAddViewColumnWithDefaults instantiates a new AddViewColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddViewColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddViewColumn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddViewColumn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddViewColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEditable

`func (o *AddViewColumn) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddViewColumn) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddViewColumn) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddViewColumn) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetName

`func (o *AddViewColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddViewColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddViewColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddViewColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddViewColumn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddViewColumn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddViewColumn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddViewColumn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AddViewColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddViewColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddViewColumn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddViewColumn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLanguageCode

`func (o *AddViewColumn) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *AddViewColumn) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *AddViewColumn) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *AddViewColumn) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetLocalizationType

`func (o *AddViewColumn) GetLocalizationType() string`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *AddViewColumn) GetLocalizationTypeOk() (*string, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *AddViewColumn) SetLocalizationType(v string)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *AddViewColumn) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### GetNumberFormat

`func (o *AddViewColumn) GetNumberFormat() NumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *AddViewColumn) GetNumberFormatOk() (*NumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *AddViewColumn) SetNumberFormat(v NumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *AddViewColumn) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetSelectionOptions

`func (o *AddViewColumn) GetSelectionOptions() []string`

GetSelectionOptions returns the SelectionOptions field if non-nil, zero value otherwise.

### GetSelectionOptionsOk

`func (o *AddViewColumn) GetSelectionOptionsOk() (*[]string, bool)`

GetSelectionOptionsOk returns a tuple with the SelectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionOptions

`func (o *AddViewColumn) SetSelectionOptions(v []string)`

SetSelectionOptions sets SelectionOptions field to given value.

### HasSelectionOptions

`func (o *AddViewColumn) HasSelectionOptions() bool`

HasSelectionOptions returns a boolean if a field has been set.

### GetReference

`func (o *AddViewColumn) GetReference() Reference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AddViewColumn) GetReferenceOk() (*Reference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AddViewColumn) SetReference(v Reference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AddViewColumn) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetFormula

`func (o *AddViewColumn) GetFormula() Formula`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *AddViewColumn) GetFormulaOk() (*Formula, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *AddViewColumn) SetFormula(v Formula)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *AddViewColumn) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *AddViewColumn) GetDateTimeFormat() DateTimeFormat`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *AddViewColumn) GetDateTimeFormatOk() (*DateTimeFormat, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *AddViewColumn) SetDateTimeFormat(v DateTimeFormat)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *AddViewColumn) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


