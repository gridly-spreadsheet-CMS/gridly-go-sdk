# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **int64** |  | [optional] 
**CompanyId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ShareType** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] 
**LastModifiedDate** | Pointer to **time.Time** |  | [optional] 
**SystemGroup** | Pointer to **bool** |  | [optional] 
**IsSystemGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *Group) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Group) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Group) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Group) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Group) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Group) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Group) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Group) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Group) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Group) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Group) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Group) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetType

`func (o *Group) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Group) HasType() bool`

HasType returns a boolean if a field has been set.

### GetShareType

`func (o *Group) GetShareType() string`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *Group) GetShareTypeOk() (*string, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *Group) SetShareType(v string)`

SetShareType sets ShareType field to given value.

### HasShareType

`func (o *Group) HasShareType() bool`

HasShareType returns a boolean if a field has been set.

### GetCreatedDate

`func (o *Group) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Group) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Group) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Group) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *Group) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *Group) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *Group) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *Group) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetSystemGroup

`func (o *Group) GetSystemGroup() bool`

GetSystemGroup returns the SystemGroup field if non-nil, zero value otherwise.

### GetSystemGroupOk

`func (o *Group) GetSystemGroupOk() (*bool, bool)`

GetSystemGroupOk returns a tuple with the SystemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemGroup

`func (o *Group) SetSystemGroup(v bool)`

SetSystemGroup sets SystemGroup field to given value.

### HasSystemGroup

`func (o *Group) HasSystemGroup() bool`

HasSystemGroup returns a boolean if a field has been set.

### GetIsSystemGroup

`func (o *Group) GetIsSystemGroup() bool`

GetIsSystemGroup returns the IsSystemGroup field if non-nil, zero value otherwise.

### GetIsSystemGroupOk

`func (o *Group) GetIsSystemGroupOk() (*bool, bool)`

GetIsSystemGroupOk returns a tuple with the IsSystemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemGroup

`func (o *Group) SetIsSystemGroup(v bool)`

SetIsSystemGroup sets IsSystemGroup field to given value.

### HasIsSystemGroup

`func (o *Group) HasIsSystemGroup() bool`

HasIsSystemGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


