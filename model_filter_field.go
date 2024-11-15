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

// FilterField struct for FilterField
type FilterField struct {
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	ColumnId *string `json:"columnId,omitempty"`
	DynamicColumn *string `json:"dynamicColumn,omitempty"`
	Operator string `json:"operator"`
	QueryPathTagViaId *bool `json:"queryPathTagViaId,omitempty"`
	SubField *string `json:"subField,omitempty"`
	Values []map[string]interface{} `json:"values,omitempty"`
}

// NewFilterField instantiates a new FilterField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterField(operator string) *FilterField {
	this := FilterField{}
	this.Operator = operator
	return &this
}

// NewFilterFieldWithDefaults instantiates a new FilterField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterFieldWithDefaults() *FilterField {
	this := FilterField{}
	return &this
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *FilterField) GetCaseSensitive() bool {
	if o == nil || isNil(o.CaseSensitive) {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || isNil(o.CaseSensitive) {
    return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *FilterField) HasCaseSensitive() bool {
	if o != nil && !isNil(o.CaseSensitive) {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *FilterField) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *FilterField) GetColumnId() string {
	if o == nil || isNil(o.ColumnId) {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetColumnIdOk() (*string, bool) {
	if o == nil || isNil(o.ColumnId) {
    return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *FilterField) HasColumnId() bool {
	if o != nil && !isNil(o.ColumnId) {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *FilterField) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetDynamicColumn returns the DynamicColumn field value if set, zero value otherwise.
func (o *FilterField) GetDynamicColumn() string {
	if o == nil || isNil(o.DynamicColumn) {
		var ret string
		return ret
	}
	return *o.DynamicColumn
}

// GetDynamicColumnOk returns a tuple with the DynamicColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetDynamicColumnOk() (*string, bool) {
	if o == nil || isNil(o.DynamicColumn) {
    return nil, false
	}
	return o.DynamicColumn, true
}

// HasDynamicColumn returns a boolean if a field has been set.
func (o *FilterField) HasDynamicColumn() bool {
	if o != nil && !isNil(o.DynamicColumn) {
		return true
	}

	return false
}

// SetDynamicColumn gets a reference to the given string and assigns it to the DynamicColumn field.
func (o *FilterField) SetDynamicColumn(v string) {
	o.DynamicColumn = &v
}

// GetOperator returns the Operator field value
func (o *FilterField) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *FilterField) GetOperatorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *FilterField) SetOperator(v string) {
	o.Operator = v
}

// GetQueryPathTagViaId returns the QueryPathTagViaId field value if set, zero value otherwise.
func (o *FilterField) GetQueryPathTagViaId() bool {
	if o == nil || isNil(o.QueryPathTagViaId) {
		var ret bool
		return ret
	}
	return *o.QueryPathTagViaId
}

// GetQueryPathTagViaIdOk returns a tuple with the QueryPathTagViaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetQueryPathTagViaIdOk() (*bool, bool) {
	if o == nil || isNil(o.QueryPathTagViaId) {
    return nil, false
	}
	return o.QueryPathTagViaId, true
}

// HasQueryPathTagViaId returns a boolean if a field has been set.
func (o *FilterField) HasQueryPathTagViaId() bool {
	if o != nil && !isNil(o.QueryPathTagViaId) {
		return true
	}

	return false
}

// SetQueryPathTagViaId gets a reference to the given bool and assigns it to the QueryPathTagViaId field.
func (o *FilterField) SetQueryPathTagViaId(v bool) {
	o.QueryPathTagViaId = &v
}

// GetSubField returns the SubField field value if set, zero value otherwise.
func (o *FilterField) GetSubField() string {
	if o == nil || isNil(o.SubField) {
		var ret string
		return ret
	}
	return *o.SubField
}

// GetSubFieldOk returns a tuple with the SubField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetSubFieldOk() (*string, bool) {
	if o == nil || isNil(o.SubField) {
    return nil, false
	}
	return o.SubField, true
}

// HasSubField returns a boolean if a field has been set.
func (o *FilterField) HasSubField() bool {
	if o != nil && !isNil(o.SubField) {
		return true
	}

	return false
}

// SetSubField gets a reference to the given string and assigns it to the SubField field.
func (o *FilterField) SetSubField(v string) {
	o.SubField = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FilterField) GetValues() []map[string]interface{} {
	if o == nil || isNil(o.Values) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterField) GetValuesOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.Values) {
    return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FilterField) HasValues() bool {
	if o != nil && !isNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []map[string]interface{} and assigns it to the Values field.
func (o *FilterField) SetValues(v []map[string]interface{}) {
	o.Values = v
}

func (o FilterField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaseSensitive) {
		toSerialize["caseSensitive"] = o.CaseSensitive
	}
	if !isNil(o.ColumnId) {
		toSerialize["columnId"] = o.ColumnId
	}
	if !isNil(o.DynamicColumn) {
		toSerialize["dynamicColumn"] = o.DynamicColumn
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if !isNil(o.QueryPathTagViaId) {
		toSerialize["queryPathTagViaId"] = o.QueryPathTagViaId
	}
	if !isNil(o.SubField) {
		toSerialize["subField"] = o.SubField
	}
	if !isNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableFilterField struct {
	value *FilterField
	isSet bool
}

func (v NullableFilterField) Get() *FilterField {
	return v.value
}

func (v *NullableFilterField) Set(val *FilterField) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterField) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterField(val *FilterField) *NullableFilterField {
	return &NullableFilterField{value: val, isSet: true}
}

func (v NullableFilterField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


