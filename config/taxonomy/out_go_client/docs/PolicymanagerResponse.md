# PolicymanagerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionId** | Pointer to **string** | This is the id returned by the governance engine | [optional] 
**Result** | [**[]PolicymanagerResponseResult**](PolicymanagerResponseResult.md) | While showing the result, action contains the action type and the associated entity on which action has been taken. | 

## Methods

### NewPolicymanagerResponse

`func NewPolicymanagerResponse(result []PolicymanagerResponseResult, ) *PolicymanagerResponse`

NewPolicymanagerResponse instantiates a new PolicymanagerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicymanagerResponseWithDefaults

`func NewPolicymanagerResponseWithDefaults() *PolicymanagerResponse`

NewPolicymanagerResponseWithDefaults instantiates a new PolicymanagerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisionId

`func (o *PolicymanagerResponse) GetDecisionId() string`

GetDecisionId returns the DecisionId field if non-nil, zero value otherwise.

### GetDecisionIdOk

`func (o *PolicymanagerResponse) GetDecisionIdOk() (*string, bool)`

GetDecisionIdOk returns a tuple with the DecisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionId

`func (o *PolicymanagerResponse) SetDecisionId(v string)`

SetDecisionId sets DecisionId field to given value.

### HasDecisionId

`func (o *PolicymanagerResponse) HasDecisionId() bool`

HasDecisionId returns a boolean if a field has been set.

### GetResult

`func (o *PolicymanagerResponse) GetResult() []PolicymanagerResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PolicymanagerResponse) GetResultOk() (*[]PolicymanagerResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PolicymanagerResponse) SetResult(v []PolicymanagerResponseResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


