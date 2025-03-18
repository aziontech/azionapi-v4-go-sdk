# PatchedEdgeApplicationRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** | * &#x60;default&#x60; - default * &#x60;request&#x60; - request * &#x60;response&#x60; - response | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Behaviors** | Pointer to [**[]EdgeApplicationBehaviorFieldRequest**](EdgeApplicationBehaviorFieldRequest.md) |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedEdgeApplicationRuleEngineRequest

`func NewPatchedEdgeApplicationRuleEngineRequest() *PatchedEdgeApplicationRuleEngineRequest`

NewPatchedEdgeApplicationRuleEngineRequest instantiates a new PatchedEdgeApplicationRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeApplicationRuleEngineRequestWithDefaults

`func NewPatchedEdgeApplicationRuleEngineRequestWithDefaults() *PatchedEdgeApplicationRuleEngineRequest`

NewPatchedEdgeApplicationRuleEngineRequestWithDefaults instantiates a new PatchedEdgeApplicationRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhase

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetBehaviors() []EdgeApplicationBehaviorFieldRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetBehaviorsOk() (*[]EdgeApplicationBehaviorFieldRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetBehaviors(v []EdgeApplicationBehaviorFieldRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedEdgeApplicationRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedEdgeApplicationRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedEdgeApplicationRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


