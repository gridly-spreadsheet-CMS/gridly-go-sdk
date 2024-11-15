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

// Database struct for Database
type Database struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ProjectId *int64 `json:"projectId,omitempty"`
	Grids []Grid `json:"grids,omitempty"`
	Views []View `json:"views,omitempty"`
	Project *Project `json:"project,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

// NewDatabase instantiates a new Database object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabase() *Database {
	this := Database{}
	return &this
}

// NewDatabaseWithDefaults instantiates a new Database object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWithDefaults() *Database {
	this := Database{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Database) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Database) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Database) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Database) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Database) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Database) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Database) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Database) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Database) SetDescription(v string) {
	o.Description = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Database) GetProjectId() int64 {
	if o == nil || isNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetProjectIdOk() (*int64, bool) {
	if o == nil || isNil(o.ProjectId) {
    return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Database) HasProjectId() bool {
	if o != nil && !isNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *Database) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetGrids returns the Grids field value if set, zero value otherwise.
func (o *Database) GetGrids() []Grid {
	if o == nil || isNil(o.Grids) {
		var ret []Grid
		return ret
	}
	return o.Grids
}

// GetGridsOk returns a tuple with the Grids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetGridsOk() ([]Grid, bool) {
	if o == nil || isNil(o.Grids) {
    return nil, false
	}
	return o.Grids, true
}

// HasGrids returns a boolean if a field has been set.
func (o *Database) HasGrids() bool {
	if o != nil && !isNil(o.Grids) {
		return true
	}

	return false
}

// SetGrids gets a reference to the given []Grid and assigns it to the Grids field.
func (o *Database) SetGrids(v []Grid) {
	o.Grids = v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *Database) GetViews() []View {
	if o == nil || isNil(o.Views) {
		var ret []View
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetViewsOk() ([]View, bool) {
	if o == nil || isNil(o.Views) {
    return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *Database) HasViews() bool {
	if o != nil && !isNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []View and assigns it to the Views field.
func (o *Database) SetViews(v []View) {
	o.Views = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Database) GetProject() Project {
	if o == nil || isNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetProjectOk() (*Project, bool) {
	if o == nil || isNil(o.Project) {
    return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Database) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *Database) SetProject(v Project) {
	o.Project = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Database) GetGroups() []Group {
	if o == nil || isNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Database) GetGroupsOk() ([]Group, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Database) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *Database) SetGroups(v []Group) {
	o.Groups = v
}

func (o Database) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !isNil(o.Grids) {
		toSerialize["grids"] = o.Grids
	}
	if !isNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableDatabase struct {
	value *Database
	isSet bool
}

func (v NullableDatabase) Get() *Database {
	return v.value
}

func (v *NullableDatabase) Set(val *Database) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabase(val *Database) *NullableDatabase {
	return &NullableDatabase{value: val, isSet: true}
}

func (v NullableDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


