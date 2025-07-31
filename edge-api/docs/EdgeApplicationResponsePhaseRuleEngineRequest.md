# EdgeApplicationResponsePhaseRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest**](EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeApplicationResponsePhaseRuleEngineRequest

`func NewEdgeApplicationResponsePhaseRuleEngineRequest(name string, criteria [][]EdgeApplicationCriterionFieldRequest, behaviors []EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest, ) *EdgeApplicationResponsePhaseRuleEngineRequest`

NewEdgeApplicationResponsePhaseRuleEngineRequest instantiates a new EdgeApplicationResponsePhaseRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationResponsePhaseRuleEngineRequestWithDefaults

`func NewEdgeApplicationResponsePhaseRuleEngineRequestWithDefaults() *EdgeApplicationResponsePhaseRuleEngineRequest`

NewEdgeApplicationResponsePhaseRuleEngineRequestWithDefaults instantiates a new EdgeApplicationResponsePhaseRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetBehaviors() []EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetBehaviorsOk() (*[]EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) SetBehaviors(v []EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EdgeApplicationResponsePhaseRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


