# NumberFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencySymbol** | Pointer to **string** |  | [optional] 
**DecimalPlaces** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Use1000Separator** | Pointer to **bool** |  | [optional] 

## Methods

### NewNumberFormat

`func NewNumberFormat() *NumberFormat`

NewNumberFormat instantiates a new NumberFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberFormatWithDefaults

`func NewNumberFormatWithDefaults() *NumberFormat`

NewNumberFormatWithDefaults instantiates a new NumberFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencySymbol

`func (o *NumberFormat) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *NumberFormat) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *NumberFormat) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *NumberFormat) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### GetDecimalPlaces

`func (o *NumberFormat) GetDecimalPlaces() int32`

GetDecimalPlaces returns the DecimalPlaces field if non-nil, zero value otherwise.

### GetDecimalPlacesOk

`func (o *NumberFormat) GetDecimalPlacesOk() (*int32, bool)`

GetDecimalPlacesOk returns a tuple with the DecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPlaces

`func (o *NumberFormat) SetDecimalPlaces(v int32)`

SetDecimalPlaces sets DecimalPlaces field to given value.

### HasDecimalPlaces

`func (o *NumberFormat) HasDecimalPlaces() bool`

HasDecimalPlaces returns a boolean if a field has been set.

### GetType

`func (o *NumberFormat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumberFormat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumberFormat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NumberFormat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUse1000Separator

`func (o *NumberFormat) GetUse1000Separator() bool`

GetUse1000Separator returns the Use1000Separator field if non-nil, zero value otherwise.

### GetUse1000SeparatorOk

`func (o *NumberFormat) GetUse1000SeparatorOk() (*bool, bool)`

GetUse1000SeparatorOk returns a tuple with the Use1000Separator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse1000Separator

`func (o *NumberFormat) SetUse1000Separator(v bool)`

SetUse1000Separator sets Use1000Separator field to given value.

### HasUse1000Separator

`func (o *NumberFormat) HasUse1000Separator() bool`

HasUse1000Separator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


