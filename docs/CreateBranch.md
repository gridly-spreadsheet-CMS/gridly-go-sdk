# CreateBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**InheritGroupAccess** | Pointer to **bool** |  | [optional] 
**InheritAutomation** | Pointer to **bool** |  | [optional] 
**ViewId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateBranch

`func NewCreateBranch(name string, ) *CreateBranch`

NewCreateBranch instantiates a new CreateBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBranchWithDefaults

`func NewCreateBranchWithDefaults() *CreateBranch`

NewCreateBranchWithDefaults instantiates a new CreateBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBranch) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateBranch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBranch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBranch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateBranch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateBranch) GetCustomProperties() map[string]map[string]interface{}`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateBranch) GetCustomPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateBranch) SetCustomProperties(v map[string]map[string]interface{})`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateBranch) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetInheritGroupAccess

`func (o *CreateBranch) GetInheritGroupAccess() bool`

GetInheritGroupAccess returns the InheritGroupAccess field if non-nil, zero value otherwise.

### GetInheritGroupAccessOk

`func (o *CreateBranch) GetInheritGroupAccessOk() (*bool, bool)`

GetInheritGroupAccessOk returns a tuple with the InheritGroupAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritGroupAccess

`func (o *CreateBranch) SetInheritGroupAccess(v bool)`

SetInheritGroupAccess sets InheritGroupAccess field to given value.

### HasInheritGroupAccess

`func (o *CreateBranch) HasInheritGroupAccess() bool`

HasInheritGroupAccess returns a boolean if a field has been set.

### GetInheritAutomation

`func (o *CreateBranch) GetInheritAutomation() bool`

GetInheritAutomation returns the InheritAutomation field if non-nil, zero value otherwise.

### GetInheritAutomationOk

`func (o *CreateBranch) GetInheritAutomationOk() (*bool, bool)`

GetInheritAutomationOk returns a tuple with the InheritAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritAutomation

`func (o *CreateBranch) SetInheritAutomation(v bool)`

SetInheritAutomation sets InheritAutomation field to given value.

### HasInheritAutomation

`func (o *CreateBranch) HasInheritAutomation() bool`

HasInheritAutomation returns a boolean if a field has been set.

### GetViewId

`func (o *CreateBranch) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *CreateBranch) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *CreateBranch) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *CreateBranch) HasViewId() bool`

HasViewId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


