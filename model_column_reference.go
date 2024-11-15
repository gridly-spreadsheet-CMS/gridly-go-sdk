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

// ColumnReference struct for ColumnReference
type ColumnReference struct {
	Grid *ReferencedGrid `json:"grid,omitempty"`
	Column *ReferencedColumn `json:"column,omitempty"`
	Branch *ReferencedGrid `json:"branch,omitempty"`
	Type *string `json:"type,omitempty"`
	SelectionType *string `json:"selectionType,omitempty"`
}

// NewColumnReference instantiates a new ColumnReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnReference() *ColumnReference {
	this := ColumnReference{}
	return &this
}

// NewColumnReferenceWithDefaults instantiates a new ColumnReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnReferenceWithDefaults() *ColumnReference {
	this := ColumnReference{}
	return &this
}

// GetGrid returns the Grid field value if set, zero value otherwise.
func (o *ColumnReference) GetGrid() ReferencedGrid {
	if o == nil || isNil(o.Grid) {
		var ret ReferencedGrid
		return ret
	}
	return *o.Grid
}

// GetGridOk returns a tuple with the Grid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnReference) GetGridOk() (*ReferencedGrid, bool) {
	if o == nil || isNil(o.Grid) {
    return nil, false
	}
	return o.Grid, true
}

// HasGrid returns a boolean if a field has been set.
func (o *ColumnReference) HasGrid() bool {
	if o != nil && !isNil(o.Grid) {
		return true
	}

	return false
}

// SetGrid gets a reference to the given ReferencedGrid and assigns it to the Grid field.
func (o *ColumnReference) SetGrid(v ReferencedGrid) {
	o.Grid = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *ColumnReference) GetColumn() ReferencedColumn {
	if o == nil || isNil(o.Column) {
		var ret ReferencedColumn
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnReference) GetColumnOk() (*ReferencedColumn, bool) {
	if o == nil || isNil(o.Column) {
    return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *ColumnReference) HasColumn() bool {
	if o != nil && !isNil(o.Column) {
		return true
	}

	return false
}

// SetColumn gets a reference to the given ReferencedColumn and assigns it to the Column field.
func (o *ColumnReference) SetColumn(v ReferencedColumn) {
	o.Column = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ColumnReference) GetBranch() ReferencedGrid {
	if o == nil || isNil(o.Branch) {
		var ret ReferencedGrid
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnReference) GetBranchOk() (*ReferencedGrid, bool) {
	if o == nil || isNil(o.Branch) {
    return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ColumnReference) HasBranch() bool {
	if o != nil && !isNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given ReferencedGrid and assigns it to the Branch field.
func (o *ColumnReference) SetBranch(v ReferencedGrid) {
	o.Branch = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ColumnReference) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnReference) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ColumnReference) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ColumnReference) SetType(v string) {
	o.Type = &v
}

// GetSelectionType returns the SelectionType field value if set, zero value otherwise.
func (o *ColumnReference) GetSelectionType() string {
	if o == nil || isNil(o.SelectionType) {
		var ret string
		return ret
	}
	return *o.SelectionType
}

// GetSelectionTypeOk returns a tuple with the SelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnReference) GetSelectionTypeOk() (*string, bool) {
	if o == nil || isNil(o.SelectionType) {
    return nil, false
	}
	return o.SelectionType, true
}

// HasSelectionType returns a boolean if a field has been set.
func (o *ColumnReference) HasSelectionType() bool {
	if o != nil && !isNil(o.SelectionType) {
		return true
	}

	return false
}

// SetSelectionType gets a reference to the given string and assigns it to the SelectionType field.
func (o *ColumnReference) SetSelectionType(v string) {
	o.SelectionType = &v
}

func (o ColumnReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Grid) {
		toSerialize["grid"] = o.Grid
	}
	if !isNil(o.Column) {
		toSerialize["column"] = o.Column
	}
	if !isNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.SelectionType) {
		toSerialize["selectionType"] = o.SelectionType
	}
	return json.Marshal(toSerialize)
}

type NullableColumnReference struct {
	value *ColumnReference
	isSet bool
}

func (v NullableColumnReference) Get() *ColumnReference {
	return v.value
}

func (v *NullableColumnReference) Set(val *ColumnReference) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnReference) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnReference(val *ColumnReference) *NullableColumnReference {
	return &NullableColumnReference{value: val, isSet: true}
}

func (v NullableColumnReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


