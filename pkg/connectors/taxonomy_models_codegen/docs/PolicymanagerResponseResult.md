# PolicymanagerResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**Action1**](Action1.md) |  | 
**Policy** | **string** | The list of policies on which the decision was based. | 

## Methods

### NewPolicymanagerResponseResult

`func NewPolicymanagerResponseResult(action Action1, policy string, ) *PolicymanagerResponseResult`

NewPolicymanagerResponseResult instantiates a new PolicymanagerResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicymanagerResponseResultWithDefaults

`func NewPolicymanagerResponseResultWithDefaults() *PolicymanagerResponseResult`

NewPolicymanagerResponseResultWithDefaults instantiates a new PolicymanagerResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicymanagerResponseResult) GetAction() Action1`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicymanagerResponseResult) GetActionOk() (*Action1, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicymanagerResponseResult) SetAction(v Action1)`

SetAction sets Action field to given value.


### GetPolicy

`func (o *PolicymanagerResponseResult) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicymanagerResponseResult) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicymanagerResponseResult) SetPolicy(v string)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


