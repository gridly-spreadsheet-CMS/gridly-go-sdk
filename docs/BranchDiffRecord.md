# BranchDiffRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]BranchDiffCell**](BranchDiffCell.md) |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBranchDiffRecord

`func NewBranchDiffRecord() *BranchDiffRecord`

NewBranchDiffRecord instantiates a new BranchDiffRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchDiffRecordWithDefaults

`func NewBranchDiffRecordWithDefaults() *BranchDiffRecord`

NewBranchDiffRecordWithDefaults instantiates a new BranchDiffRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BranchDiffRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BranchDiffRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BranchDiffRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BranchDiffRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *BranchDiffRecord) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BranchDiffRecord) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BranchDiffRecord) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BranchDiffRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStatus

`func (o *BranchDiffRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BranchDiffRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BranchDiffRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BranchDiffRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCells

`func (o *BranchDiffRecord) GetCells() []BranchDiffCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *BranchDiffRecord) GetCellsOk() (*[]BranchDiffCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *BranchDiffRecord) SetCells(v []BranchDiffCell)`

SetCells sets Cells field to given value.

### HasCells

`func (o *BranchDiffRecord) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *BranchDiffRecord) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BranchDiffRecord) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BranchDiffRecord) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *BranchDiffRecord) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *BranchDiffRecord) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *BranchDiffRecord) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *BranchDiffRecord) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *BranchDiffRecord) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


