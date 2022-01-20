# URI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absolute** | Pointer to **bool** |  | [optional] 
**Authority** | Pointer to **string** |  | [optional] 
**Fragment** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Opaque** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**RawAuthority** | Pointer to **string** |  | [optional] 
**RawFragment** | Pointer to **string** |  | [optional] 
**RawPath** | Pointer to **string** |  | [optional] 
**RawQuery** | Pointer to **string** |  | [optional] 
**RawSchemeSpecificPart** | Pointer to **string** |  | [optional] 
**RawUserInfo** | Pointer to **string** |  | [optional] 
**Scheme** | Pointer to **string** |  | [optional] 
**SchemeSpecificPart** | Pointer to **string** |  | [optional] 
**UserInfo** | Pointer to **string** |  | [optional] 

## Methods

### NewURI

`func NewURI() *URI`

NewURI instantiates a new URI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewURIWithDefaults

`func NewURIWithDefaults() *URI`

NewURIWithDefaults instantiates a new URI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolute

`func (o *URI) GetAbsolute() bool`

GetAbsolute returns the Absolute field if non-nil, zero value otherwise.

### GetAbsoluteOk

`func (o *URI) GetAbsoluteOk() (*bool, bool)`

GetAbsoluteOk returns a tuple with the Absolute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolute

`func (o *URI) SetAbsolute(v bool)`

SetAbsolute sets Absolute field to given value.

### HasAbsolute

`func (o *URI) HasAbsolute() bool`

HasAbsolute returns a boolean if a field has been set.

### GetAuthority

`func (o *URI) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *URI) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *URI) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *URI) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetFragment

`func (o *URI) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *URI) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *URI) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *URI) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### GetHost

`func (o *URI) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *URI) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *URI) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *URI) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOpaque

`func (o *URI) GetOpaque() bool`

GetOpaque returns the Opaque field if non-nil, zero value otherwise.

### GetOpaqueOk

`func (o *URI) GetOpaqueOk() (*bool, bool)`

GetOpaqueOk returns a tuple with the Opaque field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaque

`func (o *URI) SetOpaque(v bool)`

SetOpaque sets Opaque field to given value.

### HasOpaque

`func (o *URI) HasOpaque() bool`

HasOpaque returns a boolean if a field has been set.

### GetPath

`func (o *URI) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *URI) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *URI) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *URI) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *URI) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *URI) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *URI) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *URI) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetQuery

`func (o *URI) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *URI) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *URI) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *URI) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRawAuthority

`func (o *URI) GetRawAuthority() string`

GetRawAuthority returns the RawAuthority field if non-nil, zero value otherwise.

### GetRawAuthorityOk

`func (o *URI) GetRawAuthorityOk() (*string, bool)`

GetRawAuthorityOk returns a tuple with the RawAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAuthority

`func (o *URI) SetRawAuthority(v string)`

SetRawAuthority sets RawAuthority field to given value.

### HasRawAuthority

`func (o *URI) HasRawAuthority() bool`

HasRawAuthority returns a boolean if a field has been set.

### GetRawFragment

`func (o *URI) GetRawFragment() string`

GetRawFragment returns the RawFragment field if non-nil, zero value otherwise.

### GetRawFragmentOk

`func (o *URI) GetRawFragmentOk() (*string, bool)`

GetRawFragmentOk returns a tuple with the RawFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawFragment

`func (o *URI) SetRawFragment(v string)`

SetRawFragment sets RawFragment field to given value.

### HasRawFragment

`func (o *URI) HasRawFragment() bool`

HasRawFragment returns a boolean if a field has been set.

### GetRawPath

`func (o *URI) GetRawPath() string`

GetRawPath returns the RawPath field if non-nil, zero value otherwise.

### GetRawPathOk

`func (o *URI) GetRawPathOk() (*string, bool)`

GetRawPathOk returns a tuple with the RawPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPath

`func (o *URI) SetRawPath(v string)`

SetRawPath sets RawPath field to given value.

### HasRawPath

`func (o *URI) HasRawPath() bool`

HasRawPath returns a boolean if a field has been set.

### GetRawQuery

`func (o *URI) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *URI) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *URI) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *URI) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### GetRawSchemeSpecificPart

`func (o *URI) GetRawSchemeSpecificPart() string`

GetRawSchemeSpecificPart returns the RawSchemeSpecificPart field if non-nil, zero value otherwise.

### GetRawSchemeSpecificPartOk

`func (o *URI) GetRawSchemeSpecificPartOk() (*string, bool)`

GetRawSchemeSpecificPartOk returns a tuple with the RawSchemeSpecificPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSchemeSpecificPart

`func (o *URI) SetRawSchemeSpecificPart(v string)`

SetRawSchemeSpecificPart sets RawSchemeSpecificPart field to given value.

### HasRawSchemeSpecificPart

`func (o *URI) HasRawSchemeSpecificPart() bool`

HasRawSchemeSpecificPart returns a boolean if a field has been set.

### GetRawUserInfo

`func (o *URI) GetRawUserInfo() string`

GetRawUserInfo returns the RawUserInfo field if non-nil, zero value otherwise.

### GetRawUserInfoOk

`func (o *URI) GetRawUserInfoOk() (*string, bool)`

GetRawUserInfoOk returns a tuple with the RawUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawUserInfo

`func (o *URI) SetRawUserInfo(v string)`

SetRawUserInfo sets RawUserInfo field to given value.

### HasRawUserInfo

`func (o *URI) HasRawUserInfo() bool`

HasRawUserInfo returns a boolean if a field has been set.

### GetScheme

`func (o *URI) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *URI) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *URI) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *URI) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSchemeSpecificPart

`func (o *URI) GetSchemeSpecificPart() string`

GetSchemeSpecificPart returns the SchemeSpecificPart field if non-nil, zero value otherwise.

### GetSchemeSpecificPartOk

`func (o *URI) GetSchemeSpecificPartOk() (*string, bool)`

GetSchemeSpecificPartOk returns a tuple with the SchemeSpecificPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeSpecificPart

`func (o *URI) SetSchemeSpecificPart(v string)`

SetSchemeSpecificPart sets SchemeSpecificPart field to given value.

### HasSchemeSpecificPart

`func (o *URI) HasSchemeSpecificPart() bool`

HasSchemeSpecificPart returns a boolean if a field has been set.

### GetUserInfo

`func (o *URI) GetUserInfo() string`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *URI) GetUserInfoOk() (*string, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *URI) SetUserInfo(v string)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *URI) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


