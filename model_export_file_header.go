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
	"fmt"
)

// ExportFileHeader the model 'ExportFileHeader'
type ExportFileHeader string

// List of ExportFileHeader
const (
	none ExportFileHeader = "none"
	columnName ExportFileHeader = "columnName"
)

// All allowed values of ExportFileHeader enum
var AllowedExportFileHeaderEnumValues = []ExportFileHeader{
	"none",
	"columnName",
}

func (v *ExportFileHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportFileHeader(value)
	for _, existing := range AllowedExportFileHeaderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportFileHeader", value)
}

// NewExportFileHeaderFromValue returns a pointer to a valid ExportFileHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportFileHeaderFromValue(v string) (*ExportFileHeader, error) {
	ev := ExportFileHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportFileHeader: valid values are %v", v, AllowedExportFileHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportFileHeader) IsValid() bool {
	for _, existing := range AllowedExportFileHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportFileHeader value
func (v ExportFileHeader) Ptr() *ExportFileHeader {
	return &v
}

type NullableExportFileHeader struct {
	value *ExportFileHeader
	isSet bool
}

func (v NullableExportFileHeader) Get() *ExportFileHeader {
	return v.value
}

func (v *NullableExportFileHeader) Set(val *ExportFileHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableExportFileHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableExportFileHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportFileHeader(val *ExportFileHeader) *NullableExportFileHeader {
	return &NullableExportFileHeader{value: val, isSet: true}
}

func (v NullableExportFileHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportFileHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

