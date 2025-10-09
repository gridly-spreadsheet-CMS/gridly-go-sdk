# LQAModelCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubCategories** | Pointer to [**[]LQAModelSubCategoryResponse**](LQAModelSubCategoryResponse.md) |  | [optional] 

## Methods

### NewLQAModelCategoryResponse

`func NewLQAModelCategoryResponse() *LQAModelCategoryResponse`

NewLQAModelCategoryResponse instantiates a new LQAModelCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLQAModelCategoryResponseWithDefaults

`func NewLQAModelCategoryResponseWithDefaults() *LQAModelCategoryResponse`

NewLQAModelCategoryResponseWithDefaults instantiates a new LQAModelCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LQAModelCategoryResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LQAModelCategoryResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LQAModelCategoryResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LQAModelCategoryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LQAModelCategoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LQAModelCategoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LQAModelCategoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LQAModelCategoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubCategories

`func (o *LQAModelCategoryResponse) GetSubCategories() []LQAModelSubCategoryResponse`

GetSubCategories returns the SubCategories field if non-nil, zero value otherwise.

### GetSubCategoriesOk

`func (o *LQAModelCategoryResponse) GetSubCategoriesOk() (*[]LQAModelSubCategoryResponse, bool)`

GetSubCategoriesOk returns a tuple with the SubCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategories

`func (o *LQAModelCategoryResponse) SetSubCategories(v []LQAModelSubCategoryResponse)`

SetSubCategories sets SubCategories field to given value.

### HasSubCategories

`func (o *LQAModelCategoryResponse) HasSubCategories() bool`

HasSubCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


