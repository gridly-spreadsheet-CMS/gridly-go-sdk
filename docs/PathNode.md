# PathNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ParentPath** | Pointer to **string** |  | [optional] 

## Methods

### NewPathNode

`func NewPathNode() *PathNode`

NewPathNode instantiates a new PathNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathNodeWithDefaults

`func NewPathNodeWithDefaults() *PathNode`

NewPathNodeWithDefaults instantiates a new PathNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PathNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PathNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PathNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PathNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentPath

`func (o *PathNode) GetParentPath() string`

GetParentPath returns the ParentPath field if non-nil, zero value otherwise.

### GetParentPathOk

`func (o *PathNode) GetParentPathOk() (*string, bool)`

GetParentPathOk returns a tuple with the ParentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPath

`func (o *PathNode) SetParentPath(v string)`

SetParentPath sets ParentPath field to given value.

### HasParentPath

`func (o *PathNode) HasParentPath() bool`

HasParentPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


