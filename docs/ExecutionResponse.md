# ExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AutomationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewExecutionResponse

`func NewExecutionResponse() *ExecutionResponse`

NewExecutionResponse instantiates a new ExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionResponseWithDefaults

`func NewExecutionResponseWithDefaults() *ExecutionResponse`

NewExecutionResponseWithDefaults instantiates a new ExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAutomationId

`func (o *ExecutionResponse) GetAutomationId() string`

GetAutomationId returns the AutomationId field if non-nil, zero value otherwise.

### GetAutomationIdOk

`func (o *ExecutionResponse) GetAutomationIdOk() (*string, bool)`

GetAutomationIdOk returns a tuple with the AutomationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationId

`func (o *ExecutionResponse) SetAutomationId(v string)`

SetAutomationId sets AutomationId field to given value.

### HasAutomationId

`func (o *ExecutionResponse) HasAutomationId() bool`

HasAutomationId returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedTime

`func (o *ExecutionResponse) GetStartedTime() time.Time`

GetStartedTime returns the StartedTime field if non-nil, zero value otherwise.

### GetStartedTimeOk

`func (o *ExecutionResponse) GetStartedTimeOk() (*time.Time, bool)`

GetStartedTimeOk returns a tuple with the StartedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedTime

`func (o *ExecutionResponse) SetStartedTime(v time.Time)`

SetStartedTime sets StartedTime field to given value.

### HasStartedTime

`func (o *ExecutionResponse) HasStartedTime() bool`

HasStartedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


