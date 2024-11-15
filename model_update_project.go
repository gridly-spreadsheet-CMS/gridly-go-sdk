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

// UpdateProject struct for UpdateProject
type UpdateProject struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewUpdateProject instantiates a new UpdateProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProject(name string) *UpdateProject {
	this := UpdateProject{}
	this.Name = name
	return &this
}

// NewUpdateProjectWithDefaults instantiates a new UpdateProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectWithDefaults() *UpdateProject {
	this := UpdateProject{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProject) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProject) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProject) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProject) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProject) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateProject struct {
	value *UpdateProject
	isSet bool
}

func (v NullableUpdateProject) Get() *UpdateProject {
	return v.value
}

func (v *NullableUpdateProject) Set(val *UpdateProject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProject(val *UpdateProject) *NullableUpdateProject {
	return &NullableUpdateProject{value: val, isSet: true}
}

func (v NullableUpdateProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


