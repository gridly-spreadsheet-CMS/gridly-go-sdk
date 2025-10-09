# AutomationNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 

## Methods

### NewAutomationNode

`func NewAutomationNode() *AutomationNode`

NewAutomationNode instantiates a new AutomationNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationNodeWithDefaults

`func NewAutomationNodeWithDefaults() *AutomationNode`

NewAutomationNodeWithDefaults instantiates a new AutomationNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutomationNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomationNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomationNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutomationNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AutomationNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutomationNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutomationNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AutomationNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AutomationNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomationNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomationNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AutomationNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AutomationNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AutomationNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AutomationNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AutomationNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParams

`func (o *AutomationNode) GetParams() map[string]map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *AutomationNode) GetParamsOk() (*map[string]map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *AutomationNode) SetParams(v map[string]map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *AutomationNode) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetEnabled

`func (o *AutomationNode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutomationNode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutomationNode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutomationNode) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCompleted

`func (o *AutomationNode) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *AutomationNode) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *AutomationNode) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *AutomationNode) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


