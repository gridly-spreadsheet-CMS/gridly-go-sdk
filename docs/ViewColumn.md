# ViewColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DateTimeFormat** | Pointer to [**DateTimeFormat**](DateTimeFormat.md) |  | [optional] 
**DependsOn** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Formula** | Pointer to [**Formula**](Formula.md) |  | [optional] 
**IsSource** | Pointer to **bool** |  | [optional] 
**IsTarget** | Pointer to **bool** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 
**LocalizationType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberFormat** | Pointer to [**NumberFormat**](NumberFormat.md) |  | [optional] 
**Reference** | Pointer to [**ColumnReference**](ColumnReference.md) |  | [optional] 
**SelectionOptions** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewViewColumn

`func NewViewColumn() *ViewColumn`

NewViewColumn instantiates a new ViewColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewColumnWithDefaults

`func NewViewColumnWithDefaults() *ViewColumn`

NewViewColumnWithDefaults instantiates a new ViewColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewColumn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewColumn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewColumn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ViewColumn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateTimeFormat

`func (o *ViewColumn) GetDateTimeFormat() DateTimeFormat`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *ViewColumn) GetDateTimeFormatOk() (*DateTimeFormat, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *ViewColumn) SetDateTimeFormat(v DateTimeFormat)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *ViewColumn) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### GetDependsOn

`func (o *ViewColumn) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ViewColumn) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ViewColumn) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ViewColumn) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetDescription

`func (o *ViewColumn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViewColumn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViewColumn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ViewColumn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormula

`func (o *ViewColumn) GetFormula() Formula`

GetFormula returns the Formula field if non-nil, zero value otherwise.

### GetFormulaOk

`func (o *ViewColumn) GetFormulaOk() (*Formula, bool)`

GetFormulaOk returns a tuple with the Formula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormula

`func (o *ViewColumn) SetFormula(v Formula)`

SetFormula sets Formula field to given value.

### HasFormula

`func (o *ViewColumn) HasFormula() bool`

HasFormula returns a boolean if a field has been set.

### GetIsSource

`func (o *ViewColumn) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *ViewColumn) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *ViewColumn) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *ViewColumn) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsTarget

`func (o *ViewColumn) GetIsTarget() bool`

GetIsTarget returns the IsTarget field if non-nil, zero value otherwise.

### GetIsTargetOk

`func (o *ViewColumn) GetIsTargetOk() (*bool, bool)`

GetIsTargetOk returns a tuple with the IsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTarget

`func (o *ViewColumn) SetIsTarget(v bool)`

SetIsTarget sets IsTarget field to given value.

### HasIsTarget

`func (o *ViewColumn) HasIsTarget() bool`

HasIsTarget returns a boolean if a field has been set.

### GetLanguageCode

`func (o *ViewColumn) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *ViewColumn) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *ViewColumn) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *ViewColumn) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetLocalizationType

`func (o *ViewColumn) GetLocalizationType() string`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *ViewColumn) GetLocalizationTypeOk() (*string, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *ViewColumn) SetLocalizationType(v string)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *ViewColumn) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### GetName

`func (o *ViewColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberFormat

`func (o *ViewColumn) GetNumberFormat() NumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *ViewColumn) GetNumberFormatOk() (*NumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *ViewColumn) SetNumberFormat(v NumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *ViewColumn) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetReference

`func (o *ViewColumn) GetReference() ColumnReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ViewColumn) GetReferenceOk() (*ColumnReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ViewColumn) SetReference(v ColumnReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ViewColumn) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSelectionOptions

`func (o *ViewColumn) GetSelectionOptions() []string`

GetSelectionOptions returns the SelectionOptions field if non-nil, zero value otherwise.

### GetSelectionOptionsOk

`func (o *ViewColumn) GetSelectionOptionsOk() (*[]string, bool)`

GetSelectionOptionsOk returns a tuple with the SelectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionOptions

`func (o *ViewColumn) SetSelectionOptions(v []string)`

SetSelectionOptions sets SelectionOptions field to given value.

### HasSelectionOptions

`func (o *ViewColumn) HasSelectionOptions() bool`

HasSelectionOptions returns a boolean if a field has been set.

### GetType

`func (o *ViewColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewColumn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewColumn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


