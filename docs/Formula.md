# Formula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormulaText** | **string** |  | 
**AlwaysFormatResultValueAsList** | Pointer to **bool** |  | [optional] 
**DetectResultValueType** | Pointer to **string** |  | [optional] 

## Methods

### NewFormula

`func NewFormula(formulaText string, ) *Formula`

NewFormula instantiates a new Formula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormulaWithDefaults

`func NewFormulaWithDefaults() *Formula`

NewFormulaWithDefaults instantiates a new Formula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormulaText

`func (o *Formula) GetFormulaText() string`

GetFormulaText returns the FormulaText field if non-nil, zero value otherwise.

### GetFormulaTextOk

`func (o *Formula) GetFormulaTextOk() (*string, bool)`

GetFormulaTextOk returns a tuple with the FormulaText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormulaText

`func (o *Formula) SetFormulaText(v string)`

SetFormulaText sets FormulaText field to given value.


### GetAlwaysFormatResultValueAsList

`func (o *Formula) GetAlwaysFormatResultValueAsList() bool`

GetAlwaysFormatResultValueAsList returns the AlwaysFormatResultValueAsList field if non-nil, zero value otherwise.

### GetAlwaysFormatResultValueAsListOk

`func (o *Formula) GetAlwaysFormatResultValueAsListOk() (*bool, bool)`

GetAlwaysFormatResultValueAsListOk returns a tuple with the AlwaysFormatResultValueAsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysFormatResultValueAsList

`func (o *Formula) SetAlwaysFormatResultValueAsList(v bool)`

SetAlwaysFormatResultValueAsList sets AlwaysFormatResultValueAsList field to given value.

### HasAlwaysFormatResultValueAsList

`func (o *Formula) HasAlwaysFormatResultValueAsList() bool`

HasAlwaysFormatResultValueAsList returns a boolean if a field has been set.

### GetDetectResultValueType

`func (o *Formula) GetDetectResultValueType() string`

GetDetectResultValueType returns the DetectResultValueType field if non-nil, zero value otherwise.

### GetDetectResultValueTypeOk

`func (o *Formula) GetDetectResultValueTypeOk() (*string, bool)`

GetDetectResultValueTypeOk returns a tuple with the DetectResultValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectResultValueType

`func (o *Formula) SetDetectResultValueType(v string)`

SetDetectResultValueType sets DetectResultValueType field to given value.

### HasDetectResultValueType

`func (o *Formula) HasDetectResultValueType() bool`

HasDetectResultValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


