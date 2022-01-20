/*
Gridly API

Gridly API documentation

API version: 3.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// NumberFormat struct for NumberFormat
type NumberFormat struct {
	CurrencySymbol *string `json:"currencySymbol,omitempty"`
	DecimalPlaces *int32 `json:"decimalPlaces,omitempty"`
	Type *string `json:"type,omitempty"`
	Use1000Separator *bool `json:"use1000Separator,omitempty"`
}

// NewNumberFormat instantiates a new NumberFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberFormat() *NumberFormat {
	this := NumberFormat{}
	return &this
}

// NewNumberFormatWithDefaults instantiates a new NumberFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberFormatWithDefaults() *NumberFormat {
	this := NumberFormat{}
	return &this
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *NumberFormat) GetCurrencySymbol() string {
	if o == nil || o.CurrencySymbol == nil {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || o.CurrencySymbol == nil {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *NumberFormat) HasCurrencySymbol() bool {
	if o != nil && o.CurrencySymbol != nil {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *NumberFormat) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetDecimalPlaces returns the DecimalPlaces field value if set, zero value otherwise.
func (o *NumberFormat) GetDecimalPlaces() int32 {
	if o == nil || o.DecimalPlaces == nil {
		var ret int32
		return ret
	}
	return *o.DecimalPlaces
}

// GetDecimalPlacesOk returns a tuple with the DecimalPlaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetDecimalPlacesOk() (*int32, bool) {
	if o == nil || o.DecimalPlaces == nil {
		return nil, false
	}
	return o.DecimalPlaces, true
}

// HasDecimalPlaces returns a boolean if a field has been set.
func (o *NumberFormat) HasDecimalPlaces() bool {
	if o != nil && o.DecimalPlaces != nil {
		return true
	}

	return false
}

// SetDecimalPlaces gets a reference to the given int32 and assigns it to the DecimalPlaces field.
func (o *NumberFormat) SetDecimalPlaces(v int32) {
	o.DecimalPlaces = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NumberFormat) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NumberFormat) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NumberFormat) SetType(v string) {
	o.Type = &v
}

// GetUse1000Separator returns the Use1000Separator field value if set, zero value otherwise.
func (o *NumberFormat) GetUse1000Separator() bool {
	if o == nil || o.Use1000Separator == nil {
		var ret bool
		return ret
	}
	return *o.Use1000Separator
}

// GetUse1000SeparatorOk returns a tuple with the Use1000Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberFormat) GetUse1000SeparatorOk() (*bool, bool) {
	if o == nil || o.Use1000Separator == nil {
		return nil, false
	}
	return o.Use1000Separator, true
}

// HasUse1000Separator returns a boolean if a field has been set.
func (o *NumberFormat) HasUse1000Separator() bool {
	if o != nil && o.Use1000Separator != nil {
		return true
	}

	return false
}

// SetUse1000Separator gets a reference to the given bool and assigns it to the Use1000Separator field.
func (o *NumberFormat) SetUse1000Separator(v bool) {
	o.Use1000Separator = &v
}

func (o NumberFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencySymbol != nil {
		toSerialize["currencySymbol"] = o.CurrencySymbol
	}
	if o.DecimalPlaces != nil {
		toSerialize["decimalPlaces"] = o.DecimalPlaces
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Use1000Separator != nil {
		toSerialize["use1000Separator"] = o.Use1000Separator
	}
	return json.Marshal(toSerialize)
}

type NullableNumberFormat struct {
	value *NumberFormat
	isSet bool
}

func (v NullableNumberFormat) Get() *NumberFormat {
	return v.value
}

func (v *NullableNumberFormat) Set(val *NumberFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberFormat(val *NumberFormat) *NullableNumberFormat {
	return &NullableNumberFormat{value: val, isSet: true}
}

func (v NullableNumberFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


