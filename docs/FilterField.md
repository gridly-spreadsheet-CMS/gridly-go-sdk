# FilterField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseSensitive** | Pointer to **bool** |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**DynamicColumn** | Pointer to **string** |  | [optional] [readonly] 
**Operator** | **string** |  | 
**QueryPathTagViaId** | Pointer to **bool** |  | [optional] 
**SubField** | Pointer to **string** |  | [optional] 
**Values** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewFilterField

`func NewFilterField(operator string, ) *FilterField`

NewFilterField instantiates a new FilterField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterFieldWithDefaults

`func NewFilterFieldWithDefaults() *FilterField`

NewFilterFieldWithDefaults instantiates a new FilterField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseSensitive

`func (o *FilterField) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *FilterField) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *FilterField) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *FilterField) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetColumnId

`func (o *FilterField) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *FilterField) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *FilterField) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *FilterField) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetDynamicColumn

`func (o *FilterField) GetDynamicColumn() string`

GetDynamicColumn returns the DynamicColumn field if non-nil, zero value otherwise.

### GetDynamicColumnOk

`func (o *FilterField) GetDynamicColumnOk() (*string, bool)`

GetDynamicColumnOk returns a tuple with the DynamicColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicColumn

`func (o *FilterField) SetDynamicColumn(v string)`

SetDynamicColumn sets DynamicColumn field to given value.

### HasDynamicColumn

`func (o *FilterField) HasDynamicColumn() bool`

HasDynamicColumn returns a boolean if a field has been set.

### GetOperator

`func (o *FilterField) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterField) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterField) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetQueryPathTagViaId

`func (o *FilterField) GetQueryPathTagViaId() bool`

GetQueryPathTagViaId returns the QueryPathTagViaId field if non-nil, zero value otherwise.

### GetQueryPathTagViaIdOk

`func (o *FilterField) GetQueryPathTagViaIdOk() (*bool, bool)`

GetQueryPathTagViaIdOk returns a tuple with the QueryPathTagViaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPathTagViaId

`func (o *FilterField) SetQueryPathTagViaId(v bool)`

SetQueryPathTagViaId sets QueryPathTagViaId field to given value.

### HasQueryPathTagViaId

`func (o *FilterField) HasQueryPathTagViaId() bool`

HasQueryPathTagViaId returns a boolean if a field has been set.

### GetSubField

`func (o *FilterField) GetSubField() string`

GetSubField returns the SubField field if non-nil, zero value otherwise.

### GetSubFieldOk

`func (o *FilterField) GetSubFieldOk() (*string, bool)`

GetSubFieldOk returns a tuple with the SubField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubField

`func (o *FilterField) SetSubField(v string)`

SetSubField sets SubField field to given value.

### HasSubField

`func (o *FilterField) HasSubField() bool`

HasSubField returns a boolean if a field has been set.

### GetValues

`func (o *FilterField) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FilterField) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FilterField) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *FilterField) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


