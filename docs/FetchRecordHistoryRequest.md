# FetchRecordHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnIds** | Pointer to **[]string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to **string** |  | [optional] 
**GroupBy** | Pointer to **[]string** |  | [optional] 
**Page** | Pointer to **string** |  | [optional] 
**FetchOptions** | Pointer to **string** |  | [optional] 
**IncludeSystemUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewFetchRecordHistoryRequest

`func NewFetchRecordHistoryRequest() *FetchRecordHistoryRequest`

NewFetchRecordHistoryRequest instantiates a new FetchRecordHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchRecordHistoryRequestWithDefaults

`func NewFetchRecordHistoryRequestWithDefaults() *FetchRecordHistoryRequest`

NewFetchRecordHistoryRequestWithDefaults instantiates a new FetchRecordHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnIds

`func (o *FetchRecordHistoryRequest) GetColumnIds() []string`

GetColumnIds returns the ColumnIds field if non-nil, zero value otherwise.

### GetColumnIdsOk

`func (o *FetchRecordHistoryRequest) GetColumnIdsOk() (*[]string, bool)`

GetColumnIdsOk returns a tuple with the ColumnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIds

`func (o *FetchRecordHistoryRequest) SetColumnIds(v []string)`

SetColumnIds sets ColumnIds field to given value.

### HasColumnIds

`func (o *FetchRecordHistoryRequest) HasColumnIds() bool`

HasColumnIds returns a boolean if a field has been set.

### GetQuery

`func (o *FetchRecordHistoryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *FetchRecordHistoryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *FetchRecordHistoryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *FetchRecordHistoryRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSort

`func (o *FetchRecordHistoryRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *FetchRecordHistoryRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *FetchRecordHistoryRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *FetchRecordHistoryRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroupBy

`func (o *FetchRecordHistoryRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *FetchRecordHistoryRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *FetchRecordHistoryRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *FetchRecordHistoryRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetPage

`func (o *FetchRecordHistoryRequest) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FetchRecordHistoryRequest) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FetchRecordHistoryRequest) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *FetchRecordHistoryRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetFetchOptions

`func (o *FetchRecordHistoryRequest) GetFetchOptions() string`

GetFetchOptions returns the FetchOptions field if non-nil, zero value otherwise.

### GetFetchOptionsOk

`func (o *FetchRecordHistoryRequest) GetFetchOptionsOk() (*string, bool)`

GetFetchOptionsOk returns a tuple with the FetchOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchOptions

`func (o *FetchRecordHistoryRequest) SetFetchOptions(v string)`

SetFetchOptions sets FetchOptions field to given value.

### HasFetchOptions

`func (o *FetchRecordHistoryRequest) HasFetchOptions() bool`

HasFetchOptions returns a boolean if a field has been set.

### GetIncludeSystemUser

`func (o *FetchRecordHistoryRequest) GetIncludeSystemUser() bool`

GetIncludeSystemUser returns the IncludeSystemUser field if non-nil, zero value otherwise.

### GetIncludeSystemUserOk

`func (o *FetchRecordHistoryRequest) GetIncludeSystemUserOk() (*bool, bool)`

GetIncludeSystemUserOk returns a tuple with the IncludeSystemUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSystemUser

`func (o *FetchRecordHistoryRequest) SetIncludeSystemUser(v bool)`

SetIncludeSystemUser sets IncludeSystemUser field to given value.

### HasIncludeSystemUser

`func (o *FetchRecordHistoryRequest) HasIncludeSystemUser() bool`

HasIncludeSystemUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


