# CreateView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**GridId** | **string** |  | 
**Columns** | Pointer to [**[]AddViewColumn**](AddViewColumn.md) |  | [optional] 

## Methods

### NewCreateView

`func NewCreateView(name string, gridId string, ) *CreateView`

NewCreateView instantiates a new CreateView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateViewWithDefaults

`func NewCreateViewWithDefaults() *CreateView`

NewCreateViewWithDefaults instantiates a new CreateView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateView) SetName(v string)`

SetName sets Name field to given value.


### GetGridId

`func (o *CreateView) GetGridId() string`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *CreateView) GetGridIdOk() (*string, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *CreateView) SetGridId(v string)`

SetGridId sets GridId field to given value.


### GetColumns

`func (o *CreateView) GetColumns() []AddViewColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CreateView) GetColumnsOk() (*[]AddViewColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CreateView) SetColumns(v []AddViewColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *CreateView) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


