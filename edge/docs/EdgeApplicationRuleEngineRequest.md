# EdgeApplicationRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Phase** | **string** | * &#x60;default&#x60; - default * &#x60;request&#x60; - request * &#x60;response&#x60; - response | 
**Active** | Pointer to **bool** |  | [optional] 
**Behaviors** | [**[]EdgeApplicationBehaviorFieldRequest**](EdgeApplicationBehaviorFieldRequest.md) |  | 
**Criteria** | [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeApplicationRuleEngineRequest

`func NewEdgeApplicationRuleEngineRequest(name string, phase string, behaviors []EdgeApplicationBehaviorFieldRequest, criteria [][]EdgeApplicationCriterionFieldRequest, ) *EdgeApplicationRuleEngineRequest`

NewEdgeApplicationRuleEngineRequest instantiates a new EdgeApplicationRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineRequestWithDefaults

`func NewEdgeApplicationRuleEngineRequestWithDefaults() *EdgeApplicationRuleEngineRequest`

NewEdgeApplicationRuleEngineRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeApplicationRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPhase

`func (o *EdgeApplicationRuleEngineRequest) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *EdgeApplicationRuleEngineRequest) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *EdgeApplicationRuleEngineRequest) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetActive

`func (o *EdgeApplicationRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *EdgeApplicationRuleEngineRequest) GetBehaviors() []EdgeApplicationBehaviorFieldRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *EdgeApplicationRuleEngineRequest) GetBehaviorsOk() (*[]EdgeApplicationBehaviorFieldRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *EdgeApplicationRuleEngineRequest) SetBehaviors(v []EdgeApplicationBehaviorFieldRequest)`

SetBehaviors sets Behaviors field to given value.


### GetCriteria

`func (o *EdgeApplicationRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EdgeApplicationRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EdgeApplicationRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetDescription

`func (o *EdgeApplicationRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EdgeApplicationRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EdgeApplicationRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EdgeApplicationRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


