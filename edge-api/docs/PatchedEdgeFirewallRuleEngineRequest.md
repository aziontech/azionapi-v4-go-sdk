# PatchedEdgeFirewallRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeFirewallCriterionFieldRequest**]([]EdgeFirewallCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]EdgeFirewallBehaviorsRequest**](EdgeFirewallBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedEdgeFirewallRuleEngineRequest

`func NewPatchedEdgeFirewallRuleEngineRequest() *PatchedEdgeFirewallRuleEngineRequest`

NewPatchedEdgeFirewallRuleEngineRequest instantiates a new PatchedEdgeFirewallRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeFirewallRuleEngineRequestWithDefaults

`func NewPatchedEdgeFirewallRuleEngineRequestWithDefaults() *PatchedEdgeFirewallRuleEngineRequest`

NewPatchedEdgeFirewallRuleEngineRequestWithDefaults instantiates a new PatchedEdgeFirewallRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeFirewallRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeFirewallRuleEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeFirewallRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeFirewallRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetCriteria() [][]EdgeFirewallCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetCriteriaOk() (*[][]EdgeFirewallCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedEdgeFirewallRuleEngineRequest) SetCriteria(v [][]EdgeFirewallCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedEdgeFirewallRuleEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetBehaviors() []EdgeFirewallBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetBehaviorsOk() (*[]EdgeFirewallBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedEdgeFirewallRuleEngineRequest) SetBehaviors(v []EdgeFirewallBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedEdgeFirewallRuleEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEdgeFirewallRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEdgeFirewallRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEdgeFirewallRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


