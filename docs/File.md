# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absolute** | Pointer to **bool** |  | [optional] 
**AbsoluteFile** | Pointer to [**File**](File.md) |  | [optional] 
**AbsolutePath** | Pointer to **string** |  | [optional] 
**CanonicalFile** | Pointer to [**File**](File.md) |  | [optional] 
**CanonicalPath** | Pointer to **string** |  | [optional] 
**Directory** | Pointer to **bool** |  | [optional] 
**File** | Pointer to **bool** |  | [optional] 
**FreeSpace** | Pointer to **int64** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**ParentFile** | Pointer to [**File**](File.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**TotalSpace** | Pointer to **int64** |  | [optional] 
**UsableSpace** | Pointer to **int64** |  | [optional] 

## Methods

### NewFile

`func NewFile() *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolute

`func (o *File) GetAbsolute() bool`

GetAbsolute returns the Absolute field if non-nil, zero value otherwise.

### GetAbsoluteOk

`func (o *File) GetAbsoluteOk() (*bool, bool)`

GetAbsoluteOk returns a tuple with the Absolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolute

`func (o *File) SetAbsolute(v bool)`

SetAbsolute sets Absolute field to given value.

### HasAbsolute

`func (o *File) HasAbsolute() bool`

HasAbsolute returns a boolean if a field has been set.

### GetAbsoluteFile

`func (o *File) GetAbsoluteFile() File`

GetAbsoluteFile returns the AbsoluteFile field if non-nil, zero value otherwise.

### GetAbsoluteFileOk

`func (o *File) GetAbsoluteFileOk() (*File, bool)`

GetAbsoluteFileOk returns a tuple with the AbsoluteFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteFile

`func (o *File) SetAbsoluteFile(v File)`

SetAbsoluteFile sets AbsoluteFile field to given value.

### HasAbsoluteFile

`func (o *File) HasAbsoluteFile() bool`

HasAbsoluteFile returns a boolean if a field has been set.

### GetAbsolutePath

`func (o *File) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *File) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *File) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.

### HasAbsolutePath

`func (o *File) HasAbsolutePath() bool`

HasAbsolutePath returns a boolean if a field has been set.

### GetCanonicalFile

`func (o *File) GetCanonicalFile() File`

GetCanonicalFile returns the CanonicalFile field if non-nil, zero value otherwise.

### GetCanonicalFileOk

`func (o *File) GetCanonicalFileOk() (*File, bool)`

GetCanonicalFileOk returns a tuple with the CanonicalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalFile

`func (o *File) SetCanonicalFile(v File)`

SetCanonicalFile sets CanonicalFile field to given value.

### HasCanonicalFile

`func (o *File) HasCanonicalFile() bool`

HasCanonicalFile returns a boolean if a field has been set.

### GetCanonicalPath

`func (o *File) GetCanonicalPath() string`

GetCanonicalPath returns the CanonicalPath field if non-nil, zero value otherwise.

### GetCanonicalPathOk

`func (o *File) GetCanonicalPathOk() (*string, bool)`

GetCanonicalPathOk returns a tuple with the CanonicalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalPath

`func (o *File) SetCanonicalPath(v string)`

SetCanonicalPath sets CanonicalPath field to given value.

### HasCanonicalPath

`func (o *File) HasCanonicalPath() bool`

HasCanonicalPath returns a boolean if a field has been set.

### GetDirectory

`func (o *File) GetDirectory() bool`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *File) GetDirectoryOk() (*bool, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *File) SetDirectory(v bool)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *File) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetFile

`func (o *File) GetFile() bool`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *File) GetFileOk() (*bool, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *File) SetFile(v bool)`

SetFile sets File field to given value.

### HasFile

`func (o *File) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetFreeSpace

`func (o *File) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *File) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *File) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *File) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetHidden

`func (o *File) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *File) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *File) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *File) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *File) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *File) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *File) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *File) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *File) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParentFile

`func (o *File) GetParentFile() File`

GetParentFile returns the ParentFile field if non-nil, zero value otherwise.

### GetParentFileOk

`func (o *File) GetParentFileOk() (*File, bool)`

GetParentFileOk returns a tuple with the ParentFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFile

`func (o *File) SetParentFile(v File)`

SetParentFile sets ParentFile field to given value.

### HasParentFile

`func (o *File) HasParentFile() bool`

HasParentFile returns a boolean if a field has been set.

### GetPath

`func (o *File) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *File) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *File) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *File) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTotalSpace

`func (o *File) GetTotalSpace() int64`

GetTotalSpace returns the TotalSpace field if non-nil, zero value otherwise.

### GetTotalSpaceOk

`func (o *File) GetTotalSpaceOk() (*int64, bool)`

GetTotalSpaceOk returns a tuple with the TotalSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpace

`func (o *File) SetTotalSpace(v int64)`

SetTotalSpace sets TotalSpace field to given value.

### HasTotalSpace

`func (o *File) HasTotalSpace() bool`

HasTotalSpace returns a boolean if a field has been set.

### GetUsableSpace

`func (o *File) GetUsableSpace() int64`

GetUsableSpace returns the UsableSpace field if non-nil, zero value otherwise.

### GetUsableSpaceOk

`func (o *File) GetUsableSpaceOk() (*int64, bool)`

GetUsableSpaceOk returns a tuple with the UsableSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsableSpace

`func (o *File) SetUsableSpace(v int64)`

SetUsableSpace sets UsableSpace field to given value.

### HasUsableSpace

`func (o *File) HasUsableSpace() bool`

HasUsableSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


