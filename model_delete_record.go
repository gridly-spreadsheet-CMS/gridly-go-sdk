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

// DeleteRecord struct for DeleteRecord
type DeleteRecord struct {
	Identifiers *[]RecordIdentifierWrapper `json:"identifiers,omitempty"`
	Ids *[]string `json:"ids,omitempty"`
}

// NewDeleteRecord instantiates a new DeleteRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRecord() *DeleteRecord {
	this := DeleteRecord{}
	return &this
}

// NewDeleteRecordWithDefaults instantiates a new DeleteRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRecordWithDefaults() *DeleteRecord {
	this := DeleteRecord{}
	return &this
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *DeleteRecord) GetIdentifiers() []RecordIdentifierWrapper {
	if o == nil || o.Identifiers == nil {
		var ret []RecordIdentifierWrapper
		return ret
	}
	return *o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRecord) GetIdentifiersOk() (*[]RecordIdentifierWrapper, bool) {
	if o == nil || o.Identifiers == nil {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *DeleteRecord) HasIdentifiers() bool {
	if o != nil && o.Identifiers != nil {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []RecordIdentifierWrapper and assigns it to the Identifiers field.
func (o *DeleteRecord) SetIdentifiers(v []RecordIdentifierWrapper) {
	o.Identifiers = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteRecord) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRecord) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteRecord) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteRecord) SetIds(v []string) {
	o.Ids = &v
}

func (o DeleteRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifiers != nil {
		toSerialize["identifiers"] = o.Identifiers
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRecord struct {
	value *DeleteRecord
	isSet bool
}

func (v NullableDeleteRecord) Get() *DeleteRecord {
	return v.value
}

func (v *NullableDeleteRecord) Set(val *DeleteRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRecord(val *DeleteRecord) *NullableDeleteRecord {
	return &NullableDeleteRecord{value: val, isSet: true}
}

func (v NullableDeleteRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


