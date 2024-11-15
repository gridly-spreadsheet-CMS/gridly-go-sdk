# CDN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**GridId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LastGeneratedTime** | Pointer to **time.Time** |  | [optional] 
**LastPublishedTime** | Pointer to **time.Time** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**WhiteListIP** | Pointer to **string** |  | [optional] 

## Methods

### NewCDN

`func NewCDN() *CDN`

NewCDN instantiates a new CDN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDNWithDefaults

`func NewCDNWithDefaults() *CDN`

NewCDNWithDefaults instantiates a new CDN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CDN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CDN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CDN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CDN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CDN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CDN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CDN) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CDN) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CDN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CDN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CDN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CDN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *CDN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CDN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CDN) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CDN) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetGridId

`func (o *CDN) GetGridId() string`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *CDN) GetGridIdOk() (*string, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *CDN) SetGridId(v string)`

SetGridId sets GridId field to given value.

### HasGridId

`func (o *CDN) HasGridId() bool`

HasGridId returns a boolean if a field has been set.

### GetStatus

`func (o *CDN) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CDN) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CDN) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CDN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastGeneratedTime

`func (o *CDN) GetLastGeneratedTime() time.Time`

GetLastGeneratedTime returns the LastGeneratedTime field if non-nil, zero value otherwise.

### GetLastGeneratedTimeOk

`func (o *CDN) GetLastGeneratedTimeOk() (*time.Time, bool)`

GetLastGeneratedTimeOk returns a tuple with the LastGeneratedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastGeneratedTime

`func (o *CDN) SetLastGeneratedTime(v time.Time)`

SetLastGeneratedTime sets LastGeneratedTime field to given value.

### HasLastGeneratedTime

`func (o *CDN) HasLastGeneratedTime() bool`

HasLastGeneratedTime returns a boolean if a field has been set.

### GetLastPublishedTime

`func (o *CDN) GetLastPublishedTime() time.Time`

GetLastPublishedTime returns the LastPublishedTime field if non-nil, zero value otherwise.

### GetLastPublishedTimeOk

`func (o *CDN) GetLastPublishedTimeOk() (*time.Time, bool)`

GetLastPublishedTimeOk returns a tuple with the LastPublishedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublishedTime

`func (o *CDN) SetLastPublishedTime(v time.Time)`

SetLastPublishedTime sets LastPublishedTime field to given value.

### HasLastPublishedTime

`func (o *CDN) HasLastPublishedTime() bool`

HasLastPublishedTime returns a boolean if a field has been set.

### GetPermission

`func (o *CDN) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CDN) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CDN) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CDN) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetType

`func (o *CDN) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CDN) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CDN) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CDN) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CDN) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CDN) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CDN) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CDN) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *CDN) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CDN) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CDN) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *CDN) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CDN) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CDN) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CDN) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CDN) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *CDN) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *CDN) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *CDN) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *CDN) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetStartTime

`func (o *CDN) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CDN) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CDN) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CDN) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CDN) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CDN) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CDN) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CDN) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetWhiteListIP

`func (o *CDN) GetWhiteListIP() string`

GetWhiteListIP returns the WhiteListIP field if non-nil, zero value otherwise.

### GetWhiteListIPOk

`func (o *CDN) GetWhiteListIPOk() (*string, bool)`

GetWhiteListIPOk returns a tuple with the WhiteListIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListIP

`func (o *CDN) SetWhiteListIP(v string)`

SetWhiteListIP sets WhiteListIP field to given value.

### HasWhiteListIP

`func (o *CDN) HasWhiteListIP() bool`

HasWhiteListIP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


