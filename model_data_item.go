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

// DataItem struct for DataItem
type DataItem struct {
}

// NewDataItem instantiates a new DataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataItem() *DataItem {
	this := DataItem{}
	return &this
}

// NewDataItemWithDefaults instantiates a new DataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataItemWithDefaults() *DataItem {
	this := DataItem{}
	return &this
}

func (o DataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableDataItem struct {
	value *DataItem
	isSet bool
}

func (v NullableDataItem) Get() *DataItem {
	return v.value
}

func (v *NullableDataItem) Set(val *DataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataItem(val *DataItem) *NullableDataItem {
	return &NullableDataItem{value: val, isSet: true}
}

func (v NullableDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


