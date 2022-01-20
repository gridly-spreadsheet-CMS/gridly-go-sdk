# URL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **map[string]interface{}** |  | [optional] 
**DefaultPort** | Pointer to **int32** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**UserInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewURL

`func NewURL() *URL`

NewURL instantiates a new URL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURLWithDefaults

`func NewURLWithDefaults() *URL`

NewURLWithDefaults instantiates a new URL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *URL) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *URL) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *URL) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *URL) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetContent

`func (o *URL) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *URL) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *URL) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *URL) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDefaultPort

`func (o *URL) GetDefaultPort() int32`

GetDefaultPort returns the DefaultPort field if non-nil, zero value otherwise.

### GetDefaultPortOk

`func (o *URL) GetDefaultPortOk() (*int32, bool)`

GetDefaultPortOk returns a tuple with the DefaultPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPort

`func (o *URL) SetDefaultPort(v int32)`

SetDefaultPort sets DefaultPort field to given value.

### HasDefaultPort

`func (o *URL) HasDefaultPort() bool`

HasDefaultPort returns a boolean if a field has been set.

### GetFile

`func (o *URL) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *URL) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *URL) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *URL) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetHost

`func (o *URL) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *URL) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *URL) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *URL) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPath

`func (o *URL) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *URL) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *URL) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *URL) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *URL) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *URL) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *URL) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *URL) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *URL) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *URL) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *URL) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *URL) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetQuery

`func (o *URL) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *URL) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *URL) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *URL) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRef

`func (o *URL) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *URL) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *URL) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *URL) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUserInfo

`func (o *URL) GetUserInfo() string`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *URL) GetUserInfoOk() (*string, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *URL) SetUserInfo(v string)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *URL) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


