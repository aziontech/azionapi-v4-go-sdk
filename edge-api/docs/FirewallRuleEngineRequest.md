# FirewallRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeFirewallCriterionFieldRequest**]([]EdgeFirewallCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]FirewallBehaviorsRequest**](FirewallBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallRuleEngineRequest

`func NewFirewallRuleEngineRequest(name string, criteria [][]EdgeFirewallCriterionFieldRequest, behaviors []FirewallBehaviorsRequest, ) *FirewallRuleEngineRequest`

NewFirewallRuleEngineRequest instantiates a new FirewallRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleEngineRequestWithDefaults

`func NewFirewallRuleEngineRequestWithDefaults() *FirewallRuleEngineRequest`

NewFirewallRuleEngineRequestWithDefaults instantiates a new FirewallRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FirewallRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *FirewallRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FirewallRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FirewallRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FirewallRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *FirewallRuleEngineRequest) GetCriteria() [][]EdgeFirewallCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *FirewallRuleEngineRequest) GetCriteriaOk() (*[][]EdgeFirewallCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *FirewallRuleEngineRequest) SetCriteria(v [][]EdgeFirewallCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *FirewallRuleEngineRequest) GetBehaviors() []FirewallBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *FirewallRuleEngineRequest) GetBehaviorsOk() (*[]FirewallBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *FirewallRuleEngineRequest) SetBehaviors(v []FirewallBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *FirewallRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


