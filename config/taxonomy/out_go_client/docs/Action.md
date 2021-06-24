# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**ActionType**](ActionType.md) |  | 
**ProcessingLocation** | Pointer to [**GeographyName**](GeographyName.md) |  | [optional] 

## Methods

### NewAction

`func NewAction(actionType ActionType, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *Action) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *Action) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *Action) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.


### GetProcessingLocation

`func (o *Action) GetProcessingLocation() GeographyName`

GetProcessingLocation returns the ProcessingLocation field if non-nil, zero value otherwise.

### GetProcessingLocationOk

`func (o *Action) GetProcessingLocationOk() (*GeographyName, bool)`

GetProcessingLocationOk returns a tuple with the ProcessingLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingLocation

`func (o *Action) SetProcessingLocation(v GeographyName)`

SetProcessingLocation sets ProcessingLocation field to given value.

### HasProcessingLocation

`func (o *Action) HasProcessingLocation() bool`

HasProcessingLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


