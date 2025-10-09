# ExportRecordHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnIds** | Pointer to **[]string** |  | [optional] 
**FromTime** | **time.Time** |  | 
**ToTime** | **time.Time** |  | 
**Format** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**LastModifiedByUsers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExportRecordHistory

`func NewExportRecordHistory(fromTime time.Time, toTime time.Time, ) *ExportRecordHistory`

NewExportRecordHistory instantiates a new ExportRecordHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportRecordHistoryWithDefaults

`func NewExportRecordHistoryWithDefaults() *ExportRecordHistory`

NewExportRecordHistoryWithDefaults instantiates a new ExportRecordHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnIds

`func (o *ExportRecordHistory) GetColumnIds() []string`

GetColumnIds returns the ColumnIds field if non-nil, zero value otherwise.

### GetColumnIdsOk

`func (o *ExportRecordHistory) GetColumnIdsOk() (*[]string, bool)`

GetColumnIdsOk returns a tuple with the ColumnIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIds

`func (o *ExportRecordHistory) SetColumnIds(v []string)`

SetColumnIds sets ColumnIds field to given value.

### HasColumnIds

`func (o *ExportRecordHistory) HasColumnIds() bool`

HasColumnIds returns a boolean if a field has been set.

### GetFromTime

`func (o *ExportRecordHistory) GetFromTime() time.Time`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *ExportRecordHistory) GetFromTimeOk() (*time.Time, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *ExportRecordHistory) SetFromTime(v time.Time)`

SetFromTime sets FromTime field to given value.


### GetToTime

`func (o *ExportRecordHistory) GetToTime() time.Time`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *ExportRecordHistory) GetToTimeOk() (*time.Time, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *ExportRecordHistory) SetToTime(v time.Time)`

SetToTime sets ToTime field to given value.


### GetFormat

`func (o *ExportRecordHistory) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportRecordHistory) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportRecordHistory) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ExportRecordHistory) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMode

`func (o *ExportRecordHistory) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ExportRecordHistory) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ExportRecordHistory) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ExportRecordHistory) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetLastModifiedByUsers

`func (o *ExportRecordHistory) GetLastModifiedByUsers() []string`

GetLastModifiedByUsers returns the LastModifiedByUsers field if non-nil, zero value otherwise.

### GetLastModifiedByUsersOk

`func (o *ExportRecordHistory) GetLastModifiedByUsersOk() (*[]string, bool)`

GetLastModifiedByUsersOk returns a tuple with the LastModifiedByUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUsers

`func (o *ExportRecordHistory) SetLastModifiedByUsers(v []string)`

SetLastModifiedByUsers sets LastModifiedByUsers field to given value.

### HasLastModifiedByUsers

`func (o *ExportRecordHistory) HasLastModifiedByUsers() bool`

HasLastModifiedByUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


