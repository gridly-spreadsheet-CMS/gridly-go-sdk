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

// CreateColumn struct for CreateColumn
type CreateColumn struct {
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	LanguageCode *string `json:"languageCode,omitempty"`
	Name *string `json:"name,omitempty"`
	NumberFormat *NumberFormat `json:"numberFormat,omitempty"`
	Reference *Reference `json:"reference,omitempty"`
	SelectionOptions *[]string `json:"selectionOptions,omitempty"`
	Type string `json:"type"`
}

// NewCreateColumn instantiates a new CreateColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateColumn(type_ string) *CreateColumn {
	this := CreateColumn{}
	this.Type = type_
	return &this
}

// NewCreateColumnWithDefaults instantiates a new CreateColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateColumnWithDefaults() *CreateColumn {
	this := CreateColumn{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateColumn) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateColumn) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateColumn) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateColumn) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateColumn) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateColumn) SetId(v string) {
	o.Id = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *CreateColumn) GetLanguageCode() string {
	if o == nil || o.LanguageCode == nil {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetLanguageCodeOk() (*string, bool) {
	if o == nil || o.LanguageCode == nil {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *CreateColumn) HasLanguageCode() bool {
	if o != nil && o.LanguageCode != nil {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *CreateColumn) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateColumn) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateColumn) SetName(v string) {
	o.Name = &v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *CreateColumn) GetNumberFormat() NumberFormat {
	if o == nil || o.NumberFormat == nil {
		var ret NumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetNumberFormatOk() (*NumberFormat, bool) {
	if o == nil || o.NumberFormat == nil {
		return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *CreateColumn) HasNumberFormat() bool {
	if o != nil && o.NumberFormat != nil {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given NumberFormat and assigns it to the NumberFormat field.
func (o *CreateColumn) SetNumberFormat(v NumberFormat) {
	o.NumberFormat = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *CreateColumn) GetReference() Reference {
	if o == nil || o.Reference == nil {
		var ret Reference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetReferenceOk() (*Reference, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *CreateColumn) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given Reference and assigns it to the Reference field.
func (o *CreateColumn) SetReference(v Reference) {
	o.Reference = &v
}

// GetSelectionOptions returns the SelectionOptions field value if set, zero value otherwise.
func (o *CreateColumn) GetSelectionOptions() []string {
	if o == nil || o.SelectionOptions == nil {
		var ret []string
		return ret
	}
	return *o.SelectionOptions
}

// GetSelectionOptionsOk returns a tuple with the SelectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetSelectionOptionsOk() (*[]string, bool) {
	if o == nil || o.SelectionOptions == nil {
		return nil, false
	}
	return o.SelectionOptions, true
}

// HasSelectionOptions returns a boolean if a field has been set.
func (o *CreateColumn) HasSelectionOptions() bool {
	if o != nil && o.SelectionOptions != nil {
		return true
	}

	return false
}

// SetSelectionOptions gets a reference to the given []string and assigns it to the SelectionOptions field.
func (o *CreateColumn) SetSelectionOptions(v []string) {
	o.SelectionOptions = &v
}

// GetType returns the Type field value
func (o *CreateColumn) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateColumn) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateColumn) SetType(v string) {
	o.Type = v
}

func (o CreateColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LanguageCode != nil {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NumberFormat != nil {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.SelectionOptions != nil {
		toSerialize["selectionOptions"] = o.SelectionOptions
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCreateColumn struct {
	value *CreateColumn
	isSet bool
}

func (v NullableCreateColumn) Get() *CreateColumn {
	return v.value
}

func (v *NullableCreateColumn) Set(val *CreateColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateColumn(val *CreateColumn) *NullableCreateColumn {
	return &NullableCreateColumn{value: val, isSet: true}
}

func (v NullableCreateColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


