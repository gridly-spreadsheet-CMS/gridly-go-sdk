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

// Privilege struct for Privilege
type Privilege struct {
	Description *string `json:"description,omitempty"`
	Extra *string `json:"extra,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Level *string `json:"level,omitempty"`
	Name *string `json:"name,omitempty"`
	Order *int32 `json:"order,omitempty"`
	ServiceId *string `json:"serviceId,omitempty"`
}

// NewPrivilege instantiates a new Privilege object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilege() *Privilege {
	this := Privilege{}
	return &this
}

// NewPrivilegeWithDefaults instantiates a new Privilege object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegeWithDefaults() *Privilege {
	this := Privilege{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Privilege) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Privilege) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Privilege) SetDescription(v string) {
	o.Description = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Privilege) GetExtra() string {
	if o == nil || o.Extra == nil {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetExtraOk() (*string, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Privilege) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *Privilege) SetExtra(v string) {
	o.Extra = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Privilege) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Privilege) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Privilege) SetId(v int64) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Privilege) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Privilege) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Privilege) SetLabel(v string) {
	o.Label = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Privilege) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Privilege) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *Privilege) SetLevel(v string) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Privilege) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Privilege) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Privilege) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Privilege) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Privilege) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *Privilege) SetOrder(v int32) {
	o.Order = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *Privilege) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Privilege) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *Privilege) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *Privilege) SetServiceId(v string) {
	o.ServiceId = &v
}

func (o Privilege) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ServiceId != nil {
		toSerialize["serviceId"] = o.ServiceId
	}
	return json.Marshal(toSerialize)
}

type NullablePrivilege struct {
	value *Privilege
	isSet bool
}

func (v NullablePrivilege) Get() *Privilege {
	return v.value
}

func (v *NullablePrivilege) Set(val *Privilege) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilege) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilege) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilege(val *Privilege) *NullablePrivilege {
	return &NullablePrivilege{value: val, isSet: true}
}

func (v NullablePrivilege) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilege) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


