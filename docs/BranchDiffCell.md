# BranchDiffCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDependencyStatus** | Pointer to **string** |  | [optional] 
**DestinationDependencyStatus** | Pointer to **string** |  | [optional] 
**SourceSourceStatus** | Pointer to **string** |  | [optional] 
**DestinationSourceStatus** | Pointer to **string** |  | [optional] 
**SourceLengthSetting** | Pointer to **int32** |  | [optional] 
**DestinationLengthSetting** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**SourceValue** | Pointer to **map[string]interface{}** |  | [optional] 
**DestinationValue** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBranchDiffCell

`func NewBranchDiffCell() *BranchDiffCell`

NewBranchDiffCell instantiates a new BranchDiffCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchDiffCellWithDefaults

`func NewBranchDiffCellWithDefaults() *BranchDiffCell`

NewBranchDiffCellWithDefaults instantiates a new BranchDiffCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDependencyStatus

`func (o *BranchDiffCell) GetSourceDependencyStatus() string`

GetSourceDependencyStatus returns the SourceDependencyStatus field if non-nil, zero value otherwise.

### GetSourceDependencyStatusOk

`func (o *BranchDiffCell) GetSourceDependencyStatusOk() (*string, bool)`

GetSourceDependencyStatusOk returns a tuple with the SourceDependencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDependencyStatus

`func (o *BranchDiffCell) SetSourceDependencyStatus(v string)`

SetSourceDependencyStatus sets SourceDependencyStatus field to given value.

### HasSourceDependencyStatus

`func (o *BranchDiffCell) HasSourceDependencyStatus() bool`

HasSourceDependencyStatus returns a boolean if a field has been set.

### GetDestinationDependencyStatus

`func (o *BranchDiffCell) GetDestinationDependencyStatus() string`

GetDestinationDependencyStatus returns the DestinationDependencyStatus field if non-nil, zero value otherwise.

### GetDestinationDependencyStatusOk

`func (o *BranchDiffCell) GetDestinationDependencyStatusOk() (*string, bool)`

GetDestinationDependencyStatusOk returns a tuple with the DestinationDependencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDependencyStatus

`func (o *BranchDiffCell) SetDestinationDependencyStatus(v string)`

SetDestinationDependencyStatus sets DestinationDependencyStatus field to given value.

### HasDestinationDependencyStatus

`func (o *BranchDiffCell) HasDestinationDependencyStatus() bool`

HasDestinationDependencyStatus returns a boolean if a field has been set.

### GetSourceSourceStatus

`func (o *BranchDiffCell) GetSourceSourceStatus() string`

GetSourceSourceStatus returns the SourceSourceStatus field if non-nil, zero value otherwise.

### GetSourceSourceStatusOk

`func (o *BranchDiffCell) GetSourceSourceStatusOk() (*string, bool)`

GetSourceSourceStatusOk returns a tuple with the SourceSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSourceStatus

`func (o *BranchDiffCell) SetSourceSourceStatus(v string)`

SetSourceSourceStatus sets SourceSourceStatus field to given value.

### HasSourceSourceStatus

`func (o *BranchDiffCell) HasSourceSourceStatus() bool`

HasSourceSourceStatus returns a boolean if a field has been set.

### GetDestinationSourceStatus

`func (o *BranchDiffCell) GetDestinationSourceStatus() string`

GetDestinationSourceStatus returns the DestinationSourceStatus field if non-nil, zero value otherwise.

### GetDestinationSourceStatusOk

`func (o *BranchDiffCell) GetDestinationSourceStatusOk() (*string, bool)`

GetDestinationSourceStatusOk returns a tuple with the DestinationSourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSourceStatus

`func (o *BranchDiffCell) SetDestinationSourceStatus(v string)`

SetDestinationSourceStatus sets DestinationSourceStatus field to given value.

### HasDestinationSourceStatus

`func (o *BranchDiffCell) HasDestinationSourceStatus() bool`

HasDestinationSourceStatus returns a boolean if a field has been set.

### GetSourceLengthSetting

`func (o *BranchDiffCell) GetSourceLengthSetting() int32`

GetSourceLengthSetting returns the SourceLengthSetting field if non-nil, zero value otherwise.

### GetSourceLengthSettingOk

`func (o *BranchDiffCell) GetSourceLengthSettingOk() (*int32, bool)`

GetSourceLengthSettingOk returns a tuple with the SourceLengthSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLengthSetting

`func (o *BranchDiffCell) SetSourceLengthSetting(v int32)`

SetSourceLengthSetting sets SourceLengthSetting field to given value.

### HasSourceLengthSetting

`func (o *BranchDiffCell) HasSourceLengthSetting() bool`

HasSourceLengthSetting returns a boolean if a field has been set.

### GetDestinationLengthSetting

`func (o *BranchDiffCell) GetDestinationLengthSetting() int32`

GetDestinationLengthSetting returns the DestinationLengthSetting field if non-nil, zero value otherwise.

### GetDestinationLengthSettingOk

`func (o *BranchDiffCell) GetDestinationLengthSettingOk() (*int32, bool)`

GetDestinationLengthSettingOk returns a tuple with the DestinationLengthSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLengthSetting

`func (o *BranchDiffCell) SetDestinationLengthSetting(v int32)`

SetDestinationLengthSetting sets DestinationLengthSetting field to given value.

### HasDestinationLengthSetting

`func (o *BranchDiffCell) HasDestinationLengthSetting() bool`

HasDestinationLengthSetting returns a boolean if a field has been set.

### GetStatus

`func (o *BranchDiffCell) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BranchDiffCell) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BranchDiffCell) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BranchDiffCell) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetColumnId

`func (o *BranchDiffCell) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *BranchDiffCell) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *BranchDiffCell) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *BranchDiffCell) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetSourceValue

`func (o *BranchDiffCell) GetSourceValue() map[string]interface{}`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *BranchDiffCell) GetSourceValueOk() (*map[string]interface{}, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *BranchDiffCell) SetSourceValue(v map[string]interface{})`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *BranchDiffCell) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### GetDestinationValue

`func (o *BranchDiffCell) GetDestinationValue() map[string]interface{}`

GetDestinationValue returns the DestinationValue field if non-nil, zero value otherwise.

### GetDestinationValueOk

`func (o *BranchDiffCell) GetDestinationValueOk() (*map[string]interface{}, bool)`

GetDestinationValueOk returns a tuple with the DestinationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationValue

`func (o *BranchDiffCell) SetDestinationValue(v map[string]interface{})`

SetDestinationValue sets DestinationValue field to given value.

### HasDestinationValue

`func (o *BranchDiffCell) HasDestinationValue() bool`

HasDestinationValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


