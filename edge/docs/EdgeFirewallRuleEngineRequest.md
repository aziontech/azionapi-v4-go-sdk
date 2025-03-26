# EdgeFirewallRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Behaviors** | [**[]EdgeFirewallBehaviorFieldRequest**](EdgeFirewallBehaviorFieldRequest.md) |  | 
**Criteria** | [**[][]EdgeFirewallCriterionFieldRequest**]([]EdgeFirewallCriterionFieldRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeFirewallRuleEngineRequest

`func NewEdgeFirewallRuleEngineRequest(name string, behaviors []EdgeFirewallBehaviorFieldRequest, criteria [][]EdgeFirewallCriterionFieldRequest, ) *EdgeFirewallRuleEngineRequest`

NewEdgeFirewallRuleEngineRequest instantiates a new EdgeFirewallRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallRuleEngineRequestWithDefaults

`func NewEdgeFirewallRuleEngineRequestWithDefaults() *EdgeFirewallRuleEngineRequest`

NewEdgeFirewallRuleEngineRequestWithDefaults instantiates a new EdgeFirewallRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFirewallRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *EdgeFirewallRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFirewallRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFirewallRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFirewallRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *EdgeFirewallRuleEngineRequest) GetBehaviors() []EdgeFirewallBehaviorFieldRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *EdgeFirewallRuleEngineRequest) GetBehaviorsOk() (*[]EdgeFirewallBehaviorFieldRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *EdgeFirewallRuleEngineRequest) SetBehaviors(v []EdgeFirewallBehaviorFieldRequest)`

SetBehaviors sets Behaviors field to given value.


### GetCriteria

`func (o *EdgeFirewallRuleEngineRequest) GetCriteria() [][]EdgeFirewallCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EdgeFirewallRuleEngineRequest) GetCriteriaOk() (*[][]EdgeFirewallCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EdgeFirewallRuleEngineRequest) SetCriteria(v [][]EdgeFirewallCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetDescription

`func (o *EdgeFirewallRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EdgeFirewallRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EdgeFirewallRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EdgeFirewallRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


