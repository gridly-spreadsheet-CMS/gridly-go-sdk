# MergeBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeRecordOptions** | Pointer to **[]string** |  | [optional] 
**MergeRecordConflicts** | Pointer to [**[]MergeRecordConflict**](MergeRecordConflict.md) |  | [optional] 
**UseLastMergeResolve** | Pointer to **bool** |  | [optional] 
**CheckMismatchedColumnType** | Pointer to **bool** |  | [optional] 
**Query** | Pointer to [**[]FilterField**](FilterField.md) |  | [optional] 

## Methods

### NewMergeBranchRequest

`func NewMergeBranchRequest() *MergeBranchRequest`

NewMergeBranchRequest instantiates a new MergeBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeBranchRequestWithDefaults

`func NewMergeBranchRequestWithDefaults() *MergeBranchRequest`

NewMergeBranchRequestWithDefaults instantiates a new MergeBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeRecordOptions

`func (o *MergeBranchRequest) GetMergeRecordOptions() []string`

GetMergeRecordOptions returns the MergeRecordOptions field if non-nil, zero value otherwise.

### GetMergeRecordOptionsOk

`func (o *MergeBranchRequest) GetMergeRecordOptionsOk() (*[]string, bool)`

GetMergeRecordOptionsOk returns a tuple with the MergeRecordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRecordOptions

`func (o *MergeBranchRequest) SetMergeRecordOptions(v []string)`

SetMergeRecordOptions sets MergeRecordOptions field to given value.

### HasMergeRecordOptions

`func (o *MergeBranchRequest) HasMergeRecordOptions() bool`

HasMergeRecordOptions returns a boolean if a field has been set.

### GetMergeRecordConflicts

`func (o *MergeBranchRequest) GetMergeRecordConflicts() []MergeRecordConflict`

GetMergeRecordConflicts returns the MergeRecordConflicts field if non-nil, zero value otherwise.

### GetMergeRecordConflictsOk

`func (o *MergeBranchRequest) GetMergeRecordConflictsOk() (*[]MergeRecordConflict, bool)`

GetMergeRecordConflictsOk returns a tuple with the MergeRecordConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRecordConflicts

`func (o *MergeBranchRequest) SetMergeRecordConflicts(v []MergeRecordConflict)`

SetMergeRecordConflicts sets MergeRecordConflicts field to given value.

### HasMergeRecordConflicts

`func (o *MergeBranchRequest) HasMergeRecordConflicts() bool`

HasMergeRecordConflicts returns a boolean if a field has been set.

### GetUseLastMergeResolve

`func (o *MergeBranchRequest) GetUseLastMergeResolve() bool`

GetUseLastMergeResolve returns the UseLastMergeResolve field if non-nil, zero value otherwise.

### GetUseLastMergeResolveOk

`func (o *MergeBranchRequest) GetUseLastMergeResolveOk() (*bool, bool)`

GetUseLastMergeResolveOk returns a tuple with the UseLastMergeResolve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLastMergeResolve

`func (o *MergeBranchRequest) SetUseLastMergeResolve(v bool)`

SetUseLastMergeResolve sets UseLastMergeResolve field to given value.

### HasUseLastMergeResolve

`func (o *MergeBranchRequest) HasUseLastMergeResolve() bool`

HasUseLastMergeResolve returns a boolean if a field has been set.

### GetCheckMismatchedColumnType

`func (o *MergeBranchRequest) GetCheckMismatchedColumnType() bool`

GetCheckMismatchedColumnType returns the CheckMismatchedColumnType field if non-nil, zero value otherwise.

### GetCheckMismatchedColumnTypeOk

`func (o *MergeBranchRequest) GetCheckMismatchedColumnTypeOk() (*bool, bool)`

GetCheckMismatchedColumnTypeOk returns a tuple with the CheckMismatchedColumnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMismatchedColumnType

`func (o *MergeBranchRequest) SetCheckMismatchedColumnType(v bool)`

SetCheckMismatchedColumnType sets CheckMismatchedColumnType field to given value.

### HasCheckMismatchedColumnType

`func (o *MergeBranchRequest) HasCheckMismatchedColumnType() bool`

HasCheckMismatchedColumnType returns a boolean if a field has been set.

### GetQuery

`func (o *MergeBranchRequest) GetQuery() []FilterField`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MergeBranchRequest) GetQueryOk() (*[]FilterField, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MergeBranchRequest) SetQuery(v []FilterField)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MergeBranchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


