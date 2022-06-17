# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Privileges** | Pointer to [**[]Privilege**](Privilege.md) |  | [optional] 
**PrivilegeIds** | Pointer to **[]int64** |  | [optional] 
**IsSystemRole** | Pointer to **bool** |  | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *Role) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Role) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Role) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Role) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Role) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Role) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Role) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Role) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *Role) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Role) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Role) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Role) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetPrivileges

`func (o *Role) GetPrivileges() []Privilege`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *Role) GetPrivilegesOk() (*[]Privilege, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *Role) SetPrivileges(v []Privilege)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *Role) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetPrivilegeIds

`func (o *Role) GetPrivilegeIds() []int64`

GetPrivilegeIds returns the PrivilegeIds field if non-nil, zero value otherwise.

### GetPrivilegeIdsOk

`func (o *Role) GetPrivilegeIdsOk() (*[]int64, bool)`

GetPrivilegeIdsOk returns a tuple with the PrivilegeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeIds

`func (o *Role) SetPrivilegeIds(v []int64)`

SetPrivilegeIds sets PrivilegeIds field to given value.

### HasPrivilegeIds

`func (o *Role) HasPrivilegeIds() bool`

HasPrivilegeIds returns a boolean if a field has been set.

### GetIsSystemRole

`func (o *Role) GetIsSystemRole() bool`

GetIsSystemRole returns the IsSystemRole field if non-nil, zero value otherwise.

### GetIsSystemRoleOk

`func (o *Role) GetIsSystemRoleOk() (*bool, bool)`

GetIsSystemRoleOk returns a tuple with the IsSystemRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemRole

`func (o *Role) SetIsSystemRole(v bool)`

SetIsSystemRole sets IsSystemRole field to given value.

### HasIsSystemRole

`func (o *Role) HasIsSystemRole() bool`

HasIsSystemRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


