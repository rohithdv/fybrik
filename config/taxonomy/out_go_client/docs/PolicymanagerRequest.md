# PolicymanagerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestContext** | Pointer to [**RequestContext**](RequestContext.md) |  | [optional] 
**Action** | [**Action**](Action.md) |  | 
**Resource** | [**Resource**](Resource.md) |  | 

## Methods

### NewPolicymanagerRequest

`func NewPolicymanagerRequest(action Action, resource Resource, ) *PolicymanagerRequest`

NewPolicymanagerRequest instantiates a new PolicymanagerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicymanagerRequestWithDefaults

`func NewPolicymanagerRequestWithDefaults() *PolicymanagerRequest`

NewPolicymanagerRequestWithDefaults instantiates a new PolicymanagerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestContext

`func (o *PolicymanagerRequest) GetRequestContext() RequestContext`

GetRequestContext returns the RequestContext field if non-nil, zero value otherwise.

### GetRequestContextOk

`func (o *PolicymanagerRequest) GetRequestContextOk() (*RequestContext, bool)`

GetRequestContextOk returns a tuple with the RequestContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContext

`func (o *PolicymanagerRequest) SetRequestContext(v RequestContext)`

SetRequestContext sets RequestContext field to given value.

### HasRequestContext

`func (o *PolicymanagerRequest) HasRequestContext() bool`

HasRequestContext returns a boolean if a field has been set.

### GetAction

`func (o *PolicymanagerRequest) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicymanagerRequest) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicymanagerRequest) SetAction(v Action)`

SetAction sets Action field to given value.


### GetResource

`func (o *PolicymanagerRequest) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicymanagerRequest) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicymanagerRequest) SetResource(v Resource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


