# RecordHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]CellHistory**](CellHistory.md) |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRecordHistory

`func NewRecordHistory() *RecordHistory`

NewRecordHistory instantiates a new RecordHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHistoryWithDefaults

`func NewRecordHistoryWithDefaults() *RecordHistory`

NewRecordHistoryWithDefaults instantiates a new RecordHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RecordHistory) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RecordHistory) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RecordHistory) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RecordHistory) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCells

`func (o *RecordHistory) GetCells() []CellHistory`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *RecordHistory) GetCellsOk() (*[]CellHistory, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *RecordHistory) SetCells(v []CellHistory)`

SetCells sets Cells field to given value.

### HasCells

`func (o *RecordHistory) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *RecordHistory) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *RecordHistory) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *RecordHistory) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *RecordHistory) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *RecordHistory) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *RecordHistory) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *RecordHistory) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *RecordHistory) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


