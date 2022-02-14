# SetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cells** | Pointer to [**[]SetCell**](SetCell.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewSetRecord

`func NewSetRecord() *SetRecord`

NewSetRecord instantiates a new SetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRecordWithDefaults

`func NewSetRecordWithDefaults() *SetRecord`

NewSetRecordWithDefaults instantiates a new SetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCells

`func (o *SetRecord) GetCells() []SetCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *SetRecord) GetCellsOk() (*[]SetCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *SetRecord) SetCells(v []SetCell)`

SetCells sets Cells field to given value.

### HasCells

`func (o *SetRecord) HasCells() bool`

HasCells returns a boolean if a field has been set.

### GetId

`func (o *SetRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SetRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *SetRecord) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SetRecord) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SetRecord) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SetRecord) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


