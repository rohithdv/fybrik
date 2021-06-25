# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the data set | 
**Creds** | Pointer to **string** | location of dataset credentials | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Columns** | Pointer to [**[]ResourceColumns**](ResourceColumns.md) | List of column names in the data set with their associated tags. They must be key value pairs. | [optional] 

## Methods

### NewResource

`func NewResource(name string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.


### GetCreds

`func (o *Resource) GetCreds() string`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *Resource) GetCredsOk() (*string, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *Resource) SetCreds(v string)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *Resource) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetTags

`func (o *Resource) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Resource) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Resource) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Resource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetColumns

`func (o *Resource) GetColumns() []ResourceColumns`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Resource) GetColumnsOk() (*[]ResourceColumns, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Resource) SetColumns(v []ResourceColumns)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Resource) HasColumns() bool`

HasColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


