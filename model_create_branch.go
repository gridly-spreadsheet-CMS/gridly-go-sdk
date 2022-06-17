/*
Gridly API

Gridly API documentation

API version: 3.29.0
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// CreateBranch struct for CreateBranch
type CreateBranch struct {
	Name string `json:"name"`
}

// NewCreateBranch instantiates a new CreateBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBranch(name string) *CreateBranch {
	this := CreateBranch{}
	this.Name = name
	return &this
}

// NewCreateBranchWithDefaults instantiates a new CreateBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBranchWithDefaults() *CreateBranch {
	this := CreateBranch{}
	return &this
}

// GetName returns the Name field value
func (o *CreateBranch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBranch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBranch) SetName(v string) {
	o.Name = v
}

func (o CreateBranch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateBranch struct {
	value *CreateBranch
	isSet bool
}

func (v NullableCreateBranch) Get() *CreateBranch {
	return v.value
}

func (v *NullableCreateBranch) Set(val *CreateBranch) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBranch(val *CreateBranch) *NullableCreateBranch {
	return &NullableCreateBranch{value: val, isSet: true}
}

func (v NullableCreateBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


