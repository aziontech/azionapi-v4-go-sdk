# PatchedApplicationRequestPhaseRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]ApplicationRuleEngineRequestPhaseBehaviorsRequest**](ApplicationRuleEngineRequestPhaseBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedApplicationRequestPhaseRuleEngineRequest

`func NewPatchedApplicationRequestPhaseRuleEngineRequest() *PatchedApplicationRequestPhaseRuleEngineRequest`

NewPatchedApplicationRequestPhaseRuleEngineRequest instantiates a new PatchedApplicationRequestPhaseRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApplicationRequestPhaseRuleEngineRequestWithDefaults

`func NewPatchedApplicationRequestPhaseRuleEngineRequestWithDefaults() *PatchedApplicationRequestPhaseRuleEngineRequest`

NewPatchedApplicationRequestPhaseRuleEngineRequestWithDefaults instantiates a new PatchedApplicationRequestPhaseRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetBehaviors() []ApplicationRuleEngineRequestPhaseBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetBehaviorsOk() (*[]ApplicationRuleEngineRequestPhaseBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) SetBehaviors(v []ApplicationRuleEngineRequestPhaseBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedApplicationRequestPhaseRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


