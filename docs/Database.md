# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Grids** | Pointer to [**[]Grid**](Grid.md) |  | [optional] 
**Groups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectId** | Pointer to **int64** |  | [optional] 
**Views** | Pointer to [**[]View**](View.md) |  | [optional] 

## Methods

### NewDatabase

`func NewDatabase() *Database`

NewDatabase instantiates a new Database object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Database) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Database) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Database) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Database) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGrids

`func (o *Database) GetGrids() []Grid`

GetGrids returns the Grids field if non-nil, zero value otherwise.

### GetGridsOk

`func (o *Database) GetGridsOk() (*[]Grid, bool)`

GetGridsOk returns a tuple with the Grids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrids

`func (o *Database) SetGrids(v []Grid)`

SetGrids sets Grids field to given value.

### HasGrids

`func (o *Database) HasGrids() bool`

HasGrids returns a boolean if a field has been set.

### GetGroups

`func (o *Database) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Database) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Database) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Database) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *Database) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Database) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Database) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Database) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Database) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Database) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *Database) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Database) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Database) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Database) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *Database) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Database) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Database) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Database) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetViews

`func (o *Database) GetViews() []View`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Database) GetViewsOk() (*[]View, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Database) SetViews(v []View)`

SetViews sets Views field to given value.

### HasViews

`func (o *Database) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


