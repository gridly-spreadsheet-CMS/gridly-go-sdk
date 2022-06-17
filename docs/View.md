# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Columns** | Pointer to [**[]ViewColumn**](ViewColumn.md) |  | [optional] 
**GridId** | Pointer to **string** |  | [optional] 
**GridStatus** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Records** | Pointer to [**[]Record**](Record.md) |  | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *View) HasId() bool`

HasId returns a boolean if a field has been set.

### GetColumns

`func (o *View) GetColumns() []ViewColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *View) GetColumnsOk() (*[]ViewColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *View) SetColumns(v []ViewColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *View) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetGridId

`func (o *View) GetGridId() string`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *View) GetGridIdOk() (*string, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *View) SetGridId(v string)`

SetGridId sets GridId field to given value.

### HasGridId

`func (o *View) HasGridId() bool`

HasGridId returns a boolean if a field has been set.

### GetGridStatus

`func (o *View) GetGridStatus() string`

GetGridStatus returns the GridStatus field if non-nil, zero value otherwise.

### GetGridStatusOk

`func (o *View) GetGridStatusOk() (*string, bool)`

GetGridStatusOk returns a tuple with the GridStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridStatus

`func (o *View) SetGridStatus(v string)`

SetGridStatus sets GridStatus field to given value.

### HasGridStatus

`func (o *View) HasGridStatus() bool`

HasGridStatus returns a boolean if a field has been set.

### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *View) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecords

`func (o *View) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *View) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *View) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *View) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


