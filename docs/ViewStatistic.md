# ViewStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordCount** | Pointer to **int64** |  | [optional] 
**Translations** | Pointer to [**map[string]ColumnStatistic**](ColumnStatistic.md) |  | [optional] 

## Methods

### NewViewStatistic

`func NewViewStatistic() *ViewStatistic`

NewViewStatistic instantiates a new ViewStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewStatisticWithDefaults

`func NewViewStatisticWithDefaults() *ViewStatistic`

NewViewStatisticWithDefaults instantiates a new ViewStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordCount

`func (o *ViewStatistic) GetRecordCount() int64`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *ViewStatistic) GetRecordCountOk() (*int64, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *ViewStatistic) SetRecordCount(v int64)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *ViewStatistic) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetTranslations

`func (o *ViewStatistic) GetTranslations() map[string]ColumnStatistic`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *ViewStatistic) GetTranslationsOk() (*map[string]ColumnStatistic, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *ViewStatistic) SetTranslations(v map[string]ColumnStatistic)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *ViewStatistic) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


