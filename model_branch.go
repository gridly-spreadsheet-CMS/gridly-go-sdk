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

// Branch struct for Branch
type Branch struct {
	Columns *[]ViewColumn `json:"columns,omitempty"`
	DefaultAccessViewId *string `json:"defaultAccessViewId,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	IsMaster *bool `json:"isMaster,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewBranch instantiates a new Branch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranch() *Branch {
	this := Branch{}
	return &this
}

// NewBranchWithDefaults instantiates a new Branch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithDefaults() *Branch {
	this := Branch{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *Branch) GetColumns() []ViewColumn {
	if o == nil || o.Columns == nil {
		var ret []ViewColumn
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetColumnsOk() (*[]ViewColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *Branch) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ViewColumn and assigns it to the Columns field.
func (o *Branch) SetColumns(v []ViewColumn) {
	o.Columns = &v
}

// GetDefaultAccessViewId returns the DefaultAccessViewId field value if set, zero value otherwise.
func (o *Branch) GetDefaultAccessViewId() string {
	if o == nil || o.DefaultAccessViewId == nil {
		var ret string
		return ret
	}
	return *o.DefaultAccessViewId
}

// GetDefaultAccessViewIdOk returns a tuple with the DefaultAccessViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetDefaultAccessViewIdOk() (*string, bool) {
	if o == nil || o.DefaultAccessViewId == nil {
		return nil, false
	}
	return o.DefaultAccessViewId, true
}

// HasDefaultAccessViewId returns a boolean if a field has been set.
func (o *Branch) HasDefaultAccessViewId() bool {
	if o != nil && o.DefaultAccessViewId != nil {
		return true
	}

	return false
}

// SetDefaultAccessViewId gets a reference to the given string and assigns it to the DefaultAccessViewId field.
func (o *Branch) SetDefaultAccessViewId(v string) {
	o.DefaultAccessViewId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Branch) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Branch) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Branch) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Branch) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Branch) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Branch) SetId(v string) {
	o.Id = &v
}

// GetIsMaster returns the IsMaster field value if set, zero value otherwise.
func (o *Branch) GetIsMaster() bool {
	if o == nil || o.IsMaster == nil {
		var ret bool
		return ret
	}
	return *o.IsMaster
}

// GetIsMasterOk returns a tuple with the IsMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetIsMasterOk() (*bool, bool) {
	if o == nil || o.IsMaster == nil {
		return nil, false
	}
	return o.IsMaster, true
}

// HasIsMaster returns a boolean if a field has been set.
func (o *Branch) HasIsMaster() bool {
	if o != nil && o.IsMaster != nil {
		return true
	}

	return false
}

// SetIsMaster gets a reference to the given bool and assigns it to the IsMaster field.
func (o *Branch) SetIsMaster(v bool) {
	o.IsMaster = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Branch) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Branch) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Branch) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Branch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Branch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Branch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Branch) SetName(v string) {
	o.Name = &v
}

func (o Branch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.DefaultAccessViewId != nil {
		toSerialize["defaultAccessViewId"] = o.DefaultAccessViewId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsMaster != nil {
		toSerialize["isMaster"] = o.IsMaster
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBranch struct {
	value *Branch
	isSet bool
}

func (v NullableBranch) Get() *Branch {
	return v.value
}

func (v *NullableBranch) Set(val *Branch) {
	v.value = val
	v.isSet = true
}

func (v NullableBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranch(val *Branch) *NullableBranch {
	return &NullableBranch{value: val, isSet: true}
}

func (v NullableBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


