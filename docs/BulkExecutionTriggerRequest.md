# BulkExecutionTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomationIds** | Pointer to **[]string** |  | [optional] 
**Await** | Pointer to **bool** |  | [optional] 

## Methods

### NewBulkExecutionTriggerRequest

`func NewBulkExecutionTriggerRequest() *BulkExecutionTriggerRequest`

NewBulkExecutionTriggerRequest instantiates a new BulkExecutionTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkExecutionTriggerRequestWithDefaults

`func NewBulkExecutionTriggerRequestWithDefaults() *BulkExecutionTriggerRequest`

NewBulkExecutionTriggerRequestWithDefaults instantiates a new BulkExecutionTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomationIds

`func (o *BulkExecutionTriggerRequest) GetAutomationIds() []string`

GetAutomationIds returns the AutomationIds field if non-nil, zero value otherwise.

### GetAutomationIdsOk

`func (o *BulkExecutionTriggerRequest) GetAutomationIdsOk() (*[]string, bool)`

GetAutomationIdsOk returns a tuple with the AutomationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationIds

`func (o *BulkExecutionTriggerRequest) SetAutomationIds(v []string)`

SetAutomationIds sets AutomationIds field to given value.

### HasAutomationIds

`func (o *BulkExecutionTriggerRequest) HasAutomationIds() bool`

HasAutomationIds returns a boolean if a field has been set.

### GetAwait

`func (o *BulkExecutionTriggerRequest) GetAwait() bool`

GetAwait returns the Await field if non-nil, zero value otherwise.

### GetAwaitOk

`func (o *BulkExecutionTriggerRequest) GetAwaitOk() (*bool, bool)`

GetAwaitOk returns a tuple with the Await field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwait

`func (o *BulkExecutionTriggerRequest) SetAwait(v bool)`

SetAwait sets Await field to given value.

### HasAwait

`func (o *BulkExecutionTriggerRequest) HasAwait() bool`

HasAwait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


