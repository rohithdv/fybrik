# RequestContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intent** | [**Intent**](Intent.md) |  | 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 

## Methods

### NewRequestContext

`func NewRequestContext(intent Intent, ) *RequestContext`

NewRequestContext instantiates a new RequestContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestContextWithDefaults

`func NewRequestContextWithDefaults() *RequestContext`

NewRequestContextWithDefaults instantiates a new RequestContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntent

`func (o *RequestContext) GetIntent() Intent`

GetIntent returns the Intent field if non-nil, zero value otherwise.

### GetIntentOk

`func (o *RequestContext) GetIntentOk() (*Intent, bool)`

GetIntentOk returns a tuple with the Intent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntent

`func (o *RequestContext) SetIntent(v Intent)`

SetIntent sets Intent field to given value.


### GetRole

`func (o *RequestContext) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RequestContext) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RequestContext) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *RequestContext) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


