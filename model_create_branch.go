/*
Gridly API

Gridly API documentation

API version: 3.21.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// CreateBranch struct for CreateBranch
type CreateBranch struct {
	Name *string `json:"name,omitempty"`
}

// NewCreateBranch instantiates a new CreateBranch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBranch() *CreateBranch {
	this := CreateBranch{}
	return &this
}

// NewCreateBranchWithDefaults instantiates a new CreateBranch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBranchWithDefaults() *CreateBranch {
	this := CreateBranch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateBranch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateBranch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateBranch) SetName(v string) {
	o.Name = &v
}

func (o CreateBranch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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


