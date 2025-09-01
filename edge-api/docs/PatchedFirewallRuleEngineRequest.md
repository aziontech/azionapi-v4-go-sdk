# PatchedFirewallRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeFirewallCriterionFieldRequest**]([]EdgeFirewallCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]FirewallBehaviorsRequest**](FirewallBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedFirewallRuleEngineRequest

`func NewPatchedFirewallRuleEngineRequest() *PatchedFirewallRuleEngineRequest`

NewPatchedFirewallRuleEngineRequest instantiates a new PatchedFirewallRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFirewallRuleEngineRequestWithDefaults

`func NewPatchedFirewallRuleEngineRequestWithDefaults() *PatchedFirewallRuleEngineRequest`

NewPatchedFirewallRuleEngineRequestWithDefaults instantiates a new PatchedFirewallRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedFirewallRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedFirewallRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedFirewallRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedFirewallRuleEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedFirewallRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedFirewallRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedFirewallRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedFirewallRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedFirewallRuleEngineRequest) GetCriteria() [][]EdgeFirewallCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedFirewallRuleEngineRequest) GetCriteriaOk() (*[][]EdgeFirewallCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedFirewallRuleEngineRequest) SetCriteria(v [][]EdgeFirewallCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedFirewallRuleEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedFirewallRuleEngineRequest) GetBehaviors() []FirewallBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedFirewallRuleEngineRequest) GetBehaviorsOk() (*[]FirewallBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedFirewallRuleEngineRequest) SetBehaviors(v []FirewallBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedFirewallRuleEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedFirewallRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedFirewallRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedFirewallRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedFirewallRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


