/*
Gridly API

Gridly API documentation

API version: 3.30.0
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// CreateProject struct for CreateProject
type CreateProject struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewCreateProject instantiates a new CreateProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProject(name string) *CreateProject {
	this := CreateProject{}
	this.Name = name
	return &this
}

// NewCreateProjectWithDefaults instantiates a new CreateProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectWithDefaults() *CreateProject {
	this := CreateProject{}
	return &this
}

// GetName returns the Name field value
func (o *CreateProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProject) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateProject) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProject) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateProject) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateProject) SetDescription(v string) {
	o.Description = &v
}

func (o CreateProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCreateProject struct {
	value *CreateProject
	isSet bool
}

func (v NullableCreateProject) Get() *CreateProject {
	return v.value
}

func (v *NullableCreateProject) Set(val *CreateProject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProject(val *CreateProject) *NullableCreateProject {
	return &NullableCreateProject{value: val, isSet: true}
}

func (v NullableCreateProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


