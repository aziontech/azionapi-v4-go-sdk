# PatchedApplicationResponsePhaseRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]ApplicationRuleEngineResponsePhaseBehaviorsRequest**](ApplicationRuleEngineResponsePhaseBehaviorsRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedApplicationResponsePhaseRuleEngineRequest

`func NewPatchedApplicationResponsePhaseRuleEngineRequest() *PatchedApplicationResponsePhaseRuleEngineRequest`

NewPatchedApplicationResponsePhaseRuleEngineRequest instantiates a new PatchedApplicationResponsePhaseRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApplicationResponsePhaseRuleEngineRequestWithDefaults

`func NewPatchedApplicationResponsePhaseRuleEngineRequestWithDefaults() *PatchedApplicationResponsePhaseRuleEngineRequest`

NewPatchedApplicationResponsePhaseRuleEngineRequestWithDefaults instantiates a new PatchedApplicationResponsePhaseRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetBehaviors() []ApplicationRuleEngineResponsePhaseBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetBehaviorsOk() (*[]ApplicationRuleEngineResponsePhaseBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) SetBehaviors(v []ApplicationRuleEngineResponsePhaseBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedApplicationResponsePhaseRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


