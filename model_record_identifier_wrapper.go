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

// RecordIdentifierWrapper struct for RecordIdentifierWrapper
type RecordIdentifierWrapper struct {
	Id string `json:"id"`
	Path string `json:"path"`
}

// NewRecordIdentifierWrapper instantiates a new RecordIdentifierWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordIdentifierWrapper(id string, path string) *RecordIdentifierWrapper {
	this := RecordIdentifierWrapper{}
	this.Id = id
	this.Path = path
	return &this
}

// NewRecordIdentifierWrapperWithDefaults instantiates a new RecordIdentifierWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordIdentifierWrapperWithDefaults() *RecordIdentifierWrapper {
	this := RecordIdentifierWrapper{}
	return &this
}

// GetId returns the Id field value
func (o *RecordIdentifierWrapper) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecordIdentifierWrapper) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RecordIdentifierWrapper) SetId(v string) {
	o.Id = v
}

// GetPath returns the Path field value
func (o *RecordIdentifierWrapper) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RecordIdentifierWrapper) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RecordIdentifierWrapper) SetPath(v string) {
	o.Path = v
}

func (o RecordIdentifierWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableRecordIdentifierWrapper struct {
	value *RecordIdentifierWrapper
	isSet bool
}

func (v NullableRecordIdentifierWrapper) Get() *RecordIdentifierWrapper {
	return v.value
}

func (v *NullableRecordIdentifierWrapper) Set(val *RecordIdentifierWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordIdentifierWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordIdentifierWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordIdentifierWrapper(val *RecordIdentifierWrapper) *NullableRecordIdentifierWrapper {
	return &NullableRecordIdentifierWrapper{value: val, isSet: true}
}

func (v NullableRecordIdentifierWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordIdentifierWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


