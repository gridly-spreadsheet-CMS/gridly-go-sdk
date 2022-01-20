# DeleteRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifiers** | Pointer to [**[]RecordIdentifierWrapper**](RecordIdentifierWrapper.md) |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDeleteRecord

`func NewDeleteRecord() *DeleteRecord`

NewDeleteRecord instantiates a new DeleteRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRecordWithDefaults

`func NewDeleteRecordWithDefaults() *DeleteRecord`

NewDeleteRecordWithDefaults instantiates a new DeleteRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifiers

`func (o *DeleteRecord) GetIdentifiers() []RecordIdentifierWrapper`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *DeleteRecord) GetIdentifiersOk() (*[]RecordIdentifierWrapper, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *DeleteRecord) SetIdentifiers(v []RecordIdentifierWrapper)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *DeleteRecord) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetIds

`func (o *DeleteRecord) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *DeleteRecord) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *DeleteRecord) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *DeleteRecord) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


