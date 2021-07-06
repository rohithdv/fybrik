# ActionOnColumns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**AllowableActionColumns**](AllowableActionColumns.md) |  | [optional] 
**Columns** | **[]string** | Represents the column names on which the actions have been taken | 

## Methods

### NewActionOnColumns

`func NewActionOnColumns(columns []string, ) *ActionOnColumns`

NewActionOnColumns instantiates a new ActionOnColumns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionOnColumnsWithDefaults

`func NewActionOnColumnsWithDefaults() *ActionOnColumns`

NewActionOnColumnsWithDefaults instantiates a new ActionOnColumns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionOnColumns) GetName() AllowableActionColumns`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionOnColumns) GetNameOk() (*AllowableActionColumns, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionOnColumns) SetName(v AllowableActionColumns)`

SetName sets Name field to given value.

### HasName

`func (o *ActionOnColumns) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColumns

`func (o *ActionOnColumns) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ActionOnColumns) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ActionOnColumns) SetColumns(v []string)`

SetColumns sets Columns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


