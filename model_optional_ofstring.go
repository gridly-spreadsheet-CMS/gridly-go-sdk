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

// OptionalOfstring struct for OptionalOfstring
type OptionalOfstring struct {
	Empty *bool `json:"empty,omitempty"`
	Present *bool `json:"present,omitempty"`
}

// NewOptionalOfstring instantiates a new OptionalOfstring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionalOfstring() *OptionalOfstring {
	this := OptionalOfstring{}
	return &this
}

// NewOptionalOfstringWithDefaults instantiates a new OptionalOfstring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionalOfstringWithDefaults() *OptionalOfstring {
	this := OptionalOfstring{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *OptionalOfstring) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalOfstring) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *OptionalOfstring) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *OptionalOfstring) SetEmpty(v bool) {
	o.Empty = &v
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *OptionalOfstring) GetPresent() bool {
	if o == nil || o.Present == nil {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalOfstring) GetPresentOk() (*bool, bool) {
	if o == nil || o.Present == nil {
		return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *OptionalOfstring) HasPresent() bool {
	if o != nil && o.Present != nil {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *OptionalOfstring) SetPresent(v bool) {
	o.Present = &v
}

func (o OptionalOfstring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.Present != nil {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableOptionalOfstring struct {
	value *OptionalOfstring
	isSet bool
}

func (v NullableOptionalOfstring) Get() *OptionalOfstring {
	return v.value
}

func (v *NullableOptionalOfstring) Set(val *OptionalOfstring) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionalOfstring) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionalOfstring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionalOfstring(val *OptionalOfstring) *NullableOptionalOfstring {
	return &NullableOptionalOfstring{value: val, isSet: true}
}

func (v NullableOptionalOfstring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionalOfstring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

