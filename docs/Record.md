# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]Cell**](Cell.md) |  | [optional] 
**LastModifiedBy** | Pointer to **string** |  | [optional] 
**LastModifiedTime** | Pointer to **time.Time** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewRecord

`func NewRecord() *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Record) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Record) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Record) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Record) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCells

`func (o *Record) GetCells() []Cell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *Record) GetCellsOk() (*[]Cell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *Record) SetCells(v []Cell)`

SetCells sets Cells field to given value.

### HasCells

`func (o *Record) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *Record) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *Record) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *Record) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *Record) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *Record) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *Record) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *Record) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *Record) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetPath

`func (o *Record) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Record) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Record) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Record) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


