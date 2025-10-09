# LengthLimitSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** |  | [optional] 
**Unit** | **string** |  | 
**UseSourceTextLength** | Pointer to **bool** |  | [optional] 

## Methods

### NewLengthLimitSetting

`func NewLengthLimitSetting(unit string, ) *LengthLimitSetting`

NewLengthLimitSetting instantiates a new LengthLimitSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLengthLimitSettingWithDefaults

`func NewLengthLimitSettingWithDefaults() *LengthLimitSetting`

NewLengthLimitSettingWithDefaults instantiates a new LengthLimitSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *LengthLimitSetting) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LengthLimitSetting) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LengthLimitSetting) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *LengthLimitSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetUnit

`func (o *LengthLimitSetting) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *LengthLimitSetting) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *LengthLimitSetting) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUseSourceTextLength

`func (o *LengthLimitSetting) GetUseSourceTextLength() bool`

GetUseSourceTextLength returns the UseSourceTextLength field if non-nil, zero value otherwise.

### GetUseSourceTextLengthOk

`func (o *LengthLimitSetting) GetUseSourceTextLengthOk() (*bool, bool)`

GetUseSourceTextLengthOk returns a tuple with the UseSourceTextLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourceTextLength

`func (o *LengthLimitSetting) SetUseSourceTextLength(v bool)`

SetUseSourceTextLength sets UseSourceTextLength field to given value.

### HasUseSourceTextLength

`func (o *LengthLimitSetting) HasUseSourceTextLength() bool`

HasUseSourceTextLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


