/*
Gridly API

Gridly API documentation

API version: 5.9.0
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// UpdatePath struct for UpdatePath
type UpdatePath struct {
	NewName string `json:"newName"`
}

// NewUpdatePath instantiates a new UpdatePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePath(newName string) *UpdatePath {
	this := UpdatePath{}
	this.NewName = newName
	return &this
}

// NewUpdatePathWithDefaults instantiates a new UpdatePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePathWithDefaults() *UpdatePath {
	this := UpdatePath{}
	return &this
}

// GetNewName returns the NewName field value
func (o *UpdatePath) GetNewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value
// and a boolean to check if the value has been set.
func (o *UpdatePath) GetNewNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NewName, true
}

// SetNewName sets field value
func (o *UpdatePath) SetNewName(v string) {
	o.NewName = v
}

func (o UpdatePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["newName"] = o.NewName
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePath struct {
	value *UpdatePath
	isSet bool
}

func (v NullableUpdatePath) Get() *UpdatePath {
	return v.value
}

func (v *NullableUpdatePath) Set(val *UpdatePath) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePath) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePath(val *UpdatePath) *NullableUpdatePath {
	return &NullableUpdatePath{value: val, isSet: true}
}

func (v NullableUpdatePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


