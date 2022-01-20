# ColumnReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to [**ReferencedColumn**](ReferencedColumn.md) |  | [optional] 
**Grid** | Pointer to [**ReferencedGrid**](ReferencedGrid.md) |  | [optional] 
**SelectionType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewColumnReference

`func NewColumnReference() *ColumnReference`

NewColumnReference instantiates a new ColumnReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnReferenceWithDefaults

`func NewColumnReferenceWithDefaults() *ColumnReference`

NewColumnReferenceWithDefaults instantiates a new ColumnReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *ColumnReference) GetColumn() ReferencedColumn`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *ColumnReference) GetColumnOk() (*ReferencedColumn, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *ColumnReference) SetColumn(v ReferencedColumn)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *ColumnReference) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetGrid

`func (o *ColumnReference) GetGrid() ReferencedGrid`

GetGrid returns the Grid field if non-nil, zero value otherwise.

### GetGridOk

`func (o *ColumnReference) GetGridOk() (*ReferencedGrid, bool)`

GetGridOk returns a tuple with the Grid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrid

`func (o *ColumnReference) SetGrid(v ReferencedGrid)`

SetGrid sets Grid field to given value.

### HasGrid

`func (o *ColumnReference) HasGrid() bool`

HasGrid returns a boolean if a field has been set.

### GetSelectionType

`func (o *ColumnReference) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *ColumnReference) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *ColumnReference) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *ColumnReference) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetType

`func (o *ColumnReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ColumnReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ColumnReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ColumnReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


