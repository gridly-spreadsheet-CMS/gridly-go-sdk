# LQAModelSubCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Severities** | Pointer to [**[]SubCategorySeverityResponse**](SubCategorySeverityResponse.md) |  | [optional] 

## Methods

### NewLQAModelSubCategoryResponse

`func NewLQAModelSubCategoryResponse() *LQAModelSubCategoryResponse`

NewLQAModelSubCategoryResponse instantiates a new LQAModelSubCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLQAModelSubCategoryResponseWithDefaults

`func NewLQAModelSubCategoryResponseWithDefaults() *LQAModelSubCategoryResponse`

NewLQAModelSubCategoryResponseWithDefaults instantiates a new LQAModelSubCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LQAModelSubCategoryResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LQAModelSubCategoryResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LQAModelSubCategoryResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LQAModelSubCategoryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LQAModelSubCategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LQAModelSubCategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LQAModelSubCategoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LQAModelSubCategoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeverities

`func (o *LQAModelSubCategoryResponse) GetSeverities() []SubCategorySeverityResponse`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *LQAModelSubCategoryResponse) GetSeveritiesOk() (*[]SubCategorySeverityResponse, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *LQAModelSubCategoryResponse) SetSeverities(v []SubCategorySeverityResponse)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *LQAModelSubCategoryResponse) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


