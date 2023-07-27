# CellHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyStatus** | Pointer to **string** |  | [optional] 
**SourceStatus** | Pointer to **string** |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCellHistory

`func NewCellHistory() *CellHistory`

NewCellHistory instantiates a new CellHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCellHistoryWithDefaults

`func NewCellHistoryWithDefaults() *CellHistory`

NewCellHistoryWithDefaults instantiates a new CellHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyStatus

`func (o *CellHistory) GetDependencyStatus() string`

GetDependencyStatus returns the DependencyStatus field if non-nil, zero value otherwise.

### GetDependencyStatusOk

`func (o *CellHistory) GetDependencyStatusOk() (*string, bool)`

GetDependencyStatusOk returns a tuple with the DependencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyStatus

`func (o *CellHistory) SetDependencyStatus(v string)`

SetDependencyStatus sets DependencyStatus field to given value.

### HasDependencyStatus

`func (o *CellHistory) HasDependencyStatus() bool`

HasDependencyStatus returns a boolean if a field has been set.

### GetSourceStatus

`func (o *CellHistory) GetSourceStatus() string`

GetSourceStatus returns the SourceStatus field if non-nil, zero value otherwise.

### GetSourceStatusOk

`func (o *CellHistory) GetSourceStatusOk() (*string, bool)`

GetSourceStatusOk returns a tuple with the SourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatus

`func (o *CellHistory) SetSourceStatus(v string)`

SetSourceStatus sets SourceStatus field to given value.

### HasSourceStatus

`func (o *CellHistory) HasSourceStatus() bool`

HasSourceStatus returns a boolean if a field has been set.

### GetColumnId

`func (o *CellHistory) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *CellHistory) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *CellHistory) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *CellHistory) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetValue

`func (o *CellHistory) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CellHistory) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CellHistory) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CellHistory) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


