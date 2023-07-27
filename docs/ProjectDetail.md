# ProjectDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CompanyId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**Databases** | Pointer to [**[]Database**](Database.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 

## Methods

### NewProjectDetail

`func NewProjectDetail() *ProjectDetail`

NewProjectDetail instantiates a new ProjectDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailWithDefaults

`func NewProjectDetailWithDefaults() *ProjectDetail`

NewProjectDetailWithDefaults instantiates a new ProjectDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompanyId

`func (o *ProjectDetail) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ProjectDetail) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ProjectDetail) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ProjectDetail) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetName

`func (o *ProjectDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRole

`func (o *ProjectDetail) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectDetail) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectDetail) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProjectDetail) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetDatabases

`func (o *ProjectDetail) GetDatabases() []Database`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *ProjectDetail) GetDatabasesOk() (*[]Database, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *ProjectDetail) SetDatabases(v []Database)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *ProjectDetail) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetGroups

`func (o *ProjectDetail) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectDetail) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectDetail) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ProjectDetail) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


