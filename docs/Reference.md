# Reference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridId** | **string** |  | 
**BranchId** | Pointer to **string** |  | [optional] 
**ColumnId** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**SelectionType** | Pointer to **string** |  | [optional] 

## Methods

### NewReference

`func NewReference(gridId string, columnId string, ) *Reference`

NewReference instantiates a new Reference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceWithDefaults

`func NewReferenceWithDefaults() *Reference`

NewReferenceWithDefaults instantiates a new Reference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGridId

`func (o *Reference) GetGridId() string`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *Reference) GetGridIdOk() (*string, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *Reference) SetGridId(v string)`

SetGridId sets GridId field to given value.


### GetBranchId

`func (o *Reference) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *Reference) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *Reference) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *Reference) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetColumnId

`func (o *Reference) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *Reference) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *Reference) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetType

`func (o *Reference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Reference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSelectionType

`func (o *Reference) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *Reference) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *Reference) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *Reference) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


