/*
Gridly API

Gridly API documentation

API version: 4.29.1
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// ColumnStatistic struct for ColumnStatistic
type ColumnStatistic struct {
	RecordCount *TranslationCount `json:"recordCount,omitempty"`
	WordCount *TranslationCount `json:"wordCount,omitempty"`
}

// NewColumnStatistic instantiates a new ColumnStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumnStatistic() *ColumnStatistic {
	this := ColumnStatistic{}
	return &this
}

// NewColumnStatisticWithDefaults instantiates a new ColumnStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnStatisticWithDefaults() *ColumnStatistic {
	this := ColumnStatistic{}
	return &this
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *ColumnStatistic) GetRecordCount() TranslationCount {
	if o == nil || o.RecordCount == nil {
		var ret TranslationCount
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnStatistic) GetRecordCountOk() (*TranslationCount, bool) {
	if o == nil || o.RecordCount == nil {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *ColumnStatistic) HasRecordCount() bool {
	if o != nil && o.RecordCount != nil {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given TranslationCount and assigns it to the RecordCount field.
func (o *ColumnStatistic) SetRecordCount(v TranslationCount) {
	o.RecordCount = &v
}

// GetWordCount returns the WordCount field value if set, zero value otherwise.
func (o *ColumnStatistic) GetWordCount() TranslationCount {
	if o == nil || o.WordCount == nil {
		var ret TranslationCount
		return ret
	}
	return *o.WordCount
}

// GetWordCountOk returns a tuple with the WordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColumnStatistic) GetWordCountOk() (*TranslationCount, bool) {
	if o == nil || o.WordCount == nil {
		return nil, false
	}
	return o.WordCount, true
}

// HasWordCount returns a boolean if a field has been set.
func (o *ColumnStatistic) HasWordCount() bool {
	if o != nil && o.WordCount != nil {
		return true
	}

	return false
}

// SetWordCount gets a reference to the given TranslationCount and assigns it to the WordCount field.
func (o *ColumnStatistic) SetWordCount(v TranslationCount) {
	o.WordCount = &v
}

func (o ColumnStatistic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecordCount != nil {
		toSerialize["recordCount"] = o.RecordCount
	}
	if o.WordCount != nil {
		toSerialize["wordCount"] = o.WordCount
	}
	return json.Marshal(toSerialize)
}

type NullableColumnStatistic struct {
	value *ColumnStatistic
	isSet bool
}

func (v NullableColumnStatistic) Get() *ColumnStatistic {
	return v.value
}

func (v *NullableColumnStatistic) Set(val *ColumnStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullableColumnStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullableColumnStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumnStatistic(val *ColumnStatistic) *NullableColumnStatistic {
	return &NullableColumnStatistic{value: val, isSet: true}
}

func (v NullableColumnStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumnStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

