# UploadSettingFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUploadSettingFileRequest

`func NewUploadSettingFileRequest() *UploadSettingFileRequest`

NewUploadSettingFileRequest instantiates a new UploadSettingFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSettingFileRequestWithDefaults

`func NewUploadSettingFileRequestWithDefaults() *UploadSettingFileRequest`

NewUploadSettingFileRequestWithDefaults instantiates a new UploadSettingFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *UploadSettingFileRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadSettingFileRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadSettingFileRequest) SetFile(v *os.File)`

SetFile sets File field to given value.

### HasFile

`func (o *UploadSettingFileRequest) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


