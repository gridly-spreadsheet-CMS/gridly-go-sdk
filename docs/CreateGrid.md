# CreateGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TemplateGridId** | Pointer to **string** |  | [optional] 
**RecordIdentifierType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCreateGrid

`func NewCreateGrid(name string, ) *CreateGrid`

NewCreateGrid instantiates a new CreateGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGridWithDefaults

`func NewCreateGridWithDefaults() *CreateGrid`

NewCreateGridWithDefaults instantiates a new CreateGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGrid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGrid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGrid) SetName(v string)`

SetName sets Name field to given value.


### GetTemplateGridId

`func (o *CreateGrid) GetTemplateGridId() string`

GetTemplateGridId returns the TemplateGridId field if non-nil, zero value otherwise.

### GetTemplateGridIdOk

`func (o *CreateGrid) GetTemplateGridIdOk() (*string, bool)`

GetTemplateGridIdOk returns a tuple with the TemplateGridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateGridId

`func (o *CreateGrid) SetTemplateGridId(v string)`

SetTemplateGridId sets TemplateGridId field to given value.

### HasTemplateGridId

`func (o *CreateGrid) HasTemplateGridId() bool`

HasTemplateGridId returns a boolean if a field has been set.

### GetRecordIdentifierType

`func (o *CreateGrid) GetRecordIdentifierType() string`

GetRecordIdentifierType returns the RecordIdentifierType field if non-nil, zero value otherwise.

### GetRecordIdentifierTypeOk

`func (o *CreateGrid) GetRecordIdentifierTypeOk() (*string, bool)`

GetRecordIdentifierTypeOk returns a tuple with the RecordIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdentifierType

`func (o *CreateGrid) SetRecordIdentifierType(v string)`

SetRecordIdentifierType sets RecordIdentifierType field to given value.

### HasRecordIdentifierType

`func (o *CreateGrid) HasRecordIdentifierType() bool`

HasRecordIdentifierType returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateGrid) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateGrid) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateGrid) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateGrid) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


