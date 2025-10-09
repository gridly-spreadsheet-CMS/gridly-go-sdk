# Automation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Completed** | Pointer to **bool** |  | [optional] 
**Triggers** | Pointer to [**[]AutomationTrigger**](AutomationTrigger.md) |  | [optional] 
**Actions** | Pointer to [**[]AutomationNode**](AutomationNode.md) |  | [optional] 

## Methods

### NewAutomation

`func NewAutomation() *Automation`

NewAutomation instantiates a new Automation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationWithDefaults

`func NewAutomationWithDefaults() *Automation`

NewAutomationWithDefaults instantiates a new Automation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Automation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Automation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Automation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Automation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Automation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Automation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Automation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Automation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Automation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Automation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Automation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Automation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *Automation) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Automation) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Automation) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Automation) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCompleted

`func (o *Automation) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *Automation) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *Automation) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *Automation) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetTriggers

`func (o *Automation) GetTriggers() []AutomationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *Automation) GetTriggersOk() (*[]AutomationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *Automation) SetTriggers(v []AutomationTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *Automation) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetActions

`func (o *Automation) GetActions() []AutomationNode`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Automation) GetActionsOk() (*[]AutomationNode, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Automation) SetActions(v []AutomationNode)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Automation) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


