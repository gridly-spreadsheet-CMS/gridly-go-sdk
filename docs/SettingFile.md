# SettingFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**FileCategory**](FileCategory.md) |  | [optional] 
**Files** | Pointer to [**[]UploadedFile**](UploadedFile.md) |  | [optional] 

## Methods

### NewSettingFile

`func NewSettingFile() *SettingFile`

NewSettingFile instantiates a new SettingFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingFileWithDefaults

`func NewSettingFileWithDefaults() *SettingFile`

NewSettingFileWithDefaults instantiates a new SettingFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SettingFile) GetCategory() FileCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SettingFile) GetCategoryOk() (*FileCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SettingFile) SetCategory(v FileCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SettingFile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFiles

`func (o *SettingFile) GetFiles() []UploadedFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SettingFile) GetFilesOk() (*[]UploadedFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SettingFile) SetFiles(v []UploadedFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SettingFile) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


