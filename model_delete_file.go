/*
Gridly API

Gridly API documentation

API version: 3.27.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// DeleteFile struct for DeleteFile
type DeleteFile struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewDeleteFile instantiates a new DeleteFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFile() *DeleteFile {
	this := DeleteFile{}
	return &this
}

// NewDeleteFileWithDefaults instantiates a new DeleteFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFileWithDefaults() *DeleteFile {
	this := DeleteFile{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteFile) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFile) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteFile) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteFile) SetIds(v []string) {
	o.Ids = &v
}

func (o DeleteFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteFile struct {
	value *DeleteFile
	isSet bool
}

func (v NullableDeleteFile) Get() *DeleteFile {
	return v.value
}

func (v *NullableDeleteFile) Set(val *DeleteFile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFile(val *DeleteFile) *NullableDeleteFile {
	return &NullableDeleteFile{value: val, isSet: true}
}

func (v NullableDeleteFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


