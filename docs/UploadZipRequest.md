# UploadZipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColumnId** | **string** |  | 
**FileMappings** | **string** |  | 
**File** | ***os.File** |  | 

## Methods

### NewUploadZipRequest

`func NewUploadZipRequest(columnId string, fileMappings string, file *os.File, ) *UploadZipRequest`

NewUploadZipRequest instantiates a new UploadZipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadZipRequestWithDefaults

`func NewUploadZipRequestWithDefaults() *UploadZipRequest`

NewUploadZipRequestWithDefaults instantiates a new UploadZipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumnId

`func (o *UploadZipRequest) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *UploadZipRequest) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *UploadZipRequest) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.


### GetFileMappings

`func (o *UploadZipRequest) GetFileMappings() string`

GetFileMappings returns the FileMappings field if non-nil, zero value otherwise.

### GetFileMappingsOk

`func (o *UploadZipRequest) GetFileMappingsOk() (*string, bool)`

GetFileMappingsOk returns a tuple with the FileMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMappings

`func (o *UploadZipRequest) SetFileMappings(v string)`

SetFileMappings sets FileMappings field to given value.


### GetFile

`func (o *UploadZipRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadZipRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadZipRequest) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


