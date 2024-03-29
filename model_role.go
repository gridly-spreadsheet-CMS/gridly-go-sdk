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

// Role struct for Role
type Role struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	Level *string `json:"level,omitempty"`
	Privileges []Privilege `json:"privileges,omitempty"`
	PrivilegeIds []int64 `json:"privilegeIds,omitempty"`
	IsSystemRole *bool `json:"isSystemRole,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Role) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Role) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Role) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Role) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Role) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Role) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Role) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Role) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Role) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Role) SetType(v string) {
	o.Type = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Role) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Role) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *Role) SetLevel(v string) {
	o.Level = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *Role) GetPrivileges() []Privilege {
	if o == nil || o.Privileges == nil {
		var ret []Privilege
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetPrivilegesOk() ([]Privilege, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *Role) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []Privilege and assigns it to the Privileges field.
func (o *Role) SetPrivileges(v []Privilege) {
	o.Privileges = v
}

// GetPrivilegeIds returns the PrivilegeIds field value if set, zero value otherwise.
func (o *Role) GetPrivilegeIds() []int64 {
	if o == nil || o.PrivilegeIds == nil {
		var ret []int64
		return ret
	}
	return o.PrivilegeIds
}

// GetPrivilegeIdsOk returns a tuple with the PrivilegeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetPrivilegeIdsOk() ([]int64, bool) {
	if o == nil || o.PrivilegeIds == nil {
		return nil, false
	}
	return o.PrivilegeIds, true
}

// HasPrivilegeIds returns a boolean if a field has been set.
func (o *Role) HasPrivilegeIds() bool {
	if o != nil && o.PrivilegeIds != nil {
		return true
	}

	return false
}

// SetPrivilegeIds gets a reference to the given []int64 and assigns it to the PrivilegeIds field.
func (o *Role) SetPrivilegeIds(v []int64) {
	o.PrivilegeIds = v
}

// GetIsSystemRole returns the IsSystemRole field value if set, zero value otherwise.
func (o *Role) GetIsSystemRole() bool {
	if o == nil || o.IsSystemRole == nil {
		var ret bool
		return ret
	}
	return *o.IsSystemRole
}

// GetIsSystemRoleOk returns a tuple with the IsSystemRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIsSystemRoleOk() (*bool, bool) {
	if o == nil || o.IsSystemRole == nil {
		return nil, false
	}
	return o.IsSystemRole, true
}

// HasIsSystemRole returns a boolean if a field has been set.
func (o *Role) HasIsSystemRole() bool {
	if o != nil && o.IsSystemRole != nil {
		return true
	}

	return false
}

// SetIsSystemRole gets a reference to the given bool and assigns it to the IsSystemRole field.
func (o *Role) SetIsSystemRole(v bool) {
	o.IsSystemRole = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.PrivilegeIds != nil {
		toSerialize["privilegeIds"] = o.PrivilegeIds
	}
	if o.IsSystemRole != nil {
		toSerialize["isSystemRole"] = o.IsSystemRole
	}
	return json.Marshal(toSerialize)
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


