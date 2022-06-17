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

// Cell struct for Cell
type Cell struct {
	ColumnId *string `json:"columnId,omitempty"`
	DependencyStatus *string `json:"dependencyStatus,omitempty"`
	ReferencedIds []string `json:"referencedIds,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewCell instantiates a new Cell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCell() *Cell {
	this := Cell{}
	return &this
}

// NewCellWithDefaults instantiates a new Cell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellWithDefaults() *Cell {
	this := Cell{}
	return &this
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *Cell) GetColumnId() string {
	if o == nil || o.ColumnId == nil {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cell) GetColumnIdOk() (*string, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *Cell) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *Cell) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetDependencyStatus returns the DependencyStatus field value if set, zero value otherwise.
func (o *Cell) GetDependencyStatus() string {
	if o == nil || o.DependencyStatus == nil {
		var ret string
		return ret
	}
	return *o.DependencyStatus
}

// GetDependencyStatusOk returns a tuple with the DependencyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cell) GetDependencyStatusOk() (*string, bool) {
	if o == nil || o.DependencyStatus == nil {
		return nil, false
	}
	return o.DependencyStatus, true
}

// HasDependencyStatus returns a boolean if a field has been set.
func (o *Cell) HasDependencyStatus() bool {
	if o != nil && o.DependencyStatus != nil {
		return true
	}

	return false
}

// SetDependencyStatus gets a reference to the given string and assigns it to the DependencyStatus field.
func (o *Cell) SetDependencyStatus(v string) {
	o.DependencyStatus = &v
}

// GetReferencedIds returns the ReferencedIds field value if set, zero value otherwise.
func (o *Cell) GetReferencedIds() []string {
	if o == nil || o.ReferencedIds == nil {
		var ret []string
		return ret
	}
	return o.ReferencedIds
}

// GetReferencedIdsOk returns a tuple with the ReferencedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cell) GetReferencedIdsOk() ([]string, bool) {
	if o == nil || o.ReferencedIds == nil {
		return nil, false
	}
	return o.ReferencedIds, true
}

// HasReferencedIds returns a boolean if a field has been set.
func (o *Cell) HasReferencedIds() bool {
	if o != nil && o.ReferencedIds != nil {
		return true
	}

	return false
}

// SetReferencedIds gets a reference to the given []string and assigns it to the ReferencedIds field.
func (o *Cell) SetReferencedIds(v []string) {
	o.ReferencedIds = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Cell) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cell) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Cell) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *Cell) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o Cell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.DependencyStatus != nil {
		toSerialize["dependencyStatus"] = o.DependencyStatus
	}
	if o.ReferencedIds != nil {
		toSerialize["referencedIds"] = o.ReferencedIds
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCell struct {
	value *Cell
	isSet bool
}

func (v NullableCell) Get() *Cell {
	return v.value
}

func (v *NullableCell) Set(val *Cell) {
	v.value = val
	v.isSet = true
}

func (v NullableCell) IsSet() bool {
	return v.isSet
}

func (v *NullableCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCell(val *Cell) *NullableCell {
	return &NullableCell{value: val, isSet: true}
}

func (v NullableCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


