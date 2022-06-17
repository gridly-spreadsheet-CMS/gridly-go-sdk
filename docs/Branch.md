# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]ViewColumn**](ViewColumn.md) |  | [optional] 
**DefaultAccessViewId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IsMaster** | Pointer to **bool** |  | [optional] 

## Methods

### NewBranch

`func NewBranch() *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Branch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Branch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetColumns

`func (o *Branch) GetColumns() []ViewColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Branch) GetColumnsOk() (*[]ViewColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Branch) SetColumns(v []ViewColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Branch) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDefaultAccessViewId

`func (o *Branch) GetDefaultAccessViewId() string`

GetDefaultAccessViewId returns the DefaultAccessViewId field if non-nil, zero value otherwise.

### GetDefaultAccessViewIdOk

`func (o *Branch) GetDefaultAccessViewIdOk() (*string, bool)`

GetDefaultAccessViewIdOk returns a tuple with the DefaultAccessViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessViewId

`func (o *Branch) SetDefaultAccessViewId(v string)`

SetDefaultAccessViewId sets DefaultAccessViewId field to given value.

### HasDefaultAccessViewId

`func (o *Branch) HasDefaultAccessViewId() bool`

HasDefaultAccessViewId returns a boolean if a field has been set.

### GetDescription

`func (o *Branch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Branch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Branch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Branch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *Branch) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Branch) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Branch) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Branch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Branch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Branch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Branch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Branch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Branch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsMaster

`func (o *Branch) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *Branch) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *Branch) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *Branch) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


