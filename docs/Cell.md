# Cell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | Pointer to **string** |  | [optional] 
**DependencyStatus** | Pointer to **string** |  | [optional] 
**LengthLimit** | Pointer to **int32** |  | [optional] 
**ReferencedIds** | Pointer to **[]string** |  | [optional] 
**SourceStatus** | Pointer to **string** |  | [optional] 
**Tm** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCell

`func NewCell() *Cell`

NewCell instantiates a new Cell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellWithDefaults

`func NewCellWithDefaults() *Cell`

NewCellWithDefaults instantiates a new Cell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *Cell) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *Cell) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *Cell) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *Cell) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetDependencyStatus

`func (o *Cell) GetDependencyStatus() string`

GetDependencyStatus returns the DependencyStatus field if non-nil, zero value otherwise.

### GetDependencyStatusOk

`func (o *Cell) GetDependencyStatusOk() (*string, bool)`

GetDependencyStatusOk returns a tuple with the DependencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyStatus

`func (o *Cell) SetDependencyStatus(v string)`

SetDependencyStatus sets DependencyStatus field to given value.

### HasDependencyStatus

`func (o *Cell) HasDependencyStatus() bool`

HasDependencyStatus returns a boolean if a field has been set.

### GetLengthLimit

`func (o *Cell) GetLengthLimit() int32`

GetLengthLimit returns the LengthLimit field if non-nil, zero value otherwise.

### GetLengthLimitOk

`func (o *Cell) GetLengthLimitOk() (*int32, bool)`

GetLengthLimitOk returns a tuple with the LengthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthLimit

`func (o *Cell) SetLengthLimit(v int32)`

SetLengthLimit sets LengthLimit field to given value.

### HasLengthLimit

`func (o *Cell) HasLengthLimit() bool`

HasLengthLimit returns a boolean if a field has been set.

### GetReferencedIds

`func (o *Cell) GetReferencedIds() []string`

GetReferencedIds returns the ReferencedIds field if non-nil, zero value otherwise.

### GetReferencedIdsOk

`func (o *Cell) GetReferencedIdsOk() (*[]string, bool)`

GetReferencedIdsOk returns a tuple with the ReferencedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedIds

`func (o *Cell) SetReferencedIds(v []string)`

SetReferencedIds sets ReferencedIds field to given value.

### HasReferencedIds

`func (o *Cell) HasReferencedIds() bool`

HasReferencedIds returns a boolean if a field has been set.

### GetSourceStatus

`func (o *Cell) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *Cell) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *Cell) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *Cell) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.

### GetTm

`func (o *Cell) GetTm() bool`

GetTm returns the Tm field if non-nil, zero value otherwise.

### GetTmOk

`func (o *Cell) GetTmOk() (*bool, bool)`

GetTmOk returns a tuple with the Tm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTm

`func (o *Cell) SetTm(v bool)`

SetTm sets Tm field to given value.

### HasTm

`func (o *Cell) HasTm() bool`

HasTm returns a boolean if a field has been set.

### GetValue

`func (o *Cell) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Cell) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Cell) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *Cell) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


