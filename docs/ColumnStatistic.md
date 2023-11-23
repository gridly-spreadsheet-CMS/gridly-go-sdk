# ColumnStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordCount** | Pointer to [**TranslationCount**](TranslationCount.md) |  | [optional] 
**WordCount** | Pointer to [**TranslationCount**](TranslationCount.md) |  | [optional] 

## Methods

### NewColumnStatistic

`func NewColumnStatistic() *ColumnStatistic`

NewColumnStatistic instantiates a new ColumnStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnStatisticWithDefaults

`func NewColumnStatisticWithDefaults() *ColumnStatistic`

NewColumnStatisticWithDefaults instantiates a new ColumnStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordCount

`func (o *ColumnStatistic) GetRecordCount() TranslationCount`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *ColumnStatistic) GetRecordCountOk() (*TranslationCount, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *ColumnStatistic) SetRecordCount(v TranslationCount)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *ColumnStatistic) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetWordCount

`func (o *ColumnStatistic) GetWordCount() TranslationCount`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *ColumnStatistic) GetWordCountOk() (*TranslationCount, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *ColumnStatistic) SetWordCount(v TranslationCount)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *ColumnStatistic) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


