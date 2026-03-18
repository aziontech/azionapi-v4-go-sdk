# PatchedRequestPhaseRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | Pointer to [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | [optional] 
**Behaviors** | Pointer to [**[]RequestPhaseBehaviorRequest**](RequestPhaseBehaviorRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedRequestPhaseRuleRequest

`func NewPatchedRequestPhaseRuleRequest() *PatchedRequestPhaseRuleRequest`

NewPatchedRequestPhaseRuleRequest instantiates a new PatchedRequestPhaseRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRequestPhaseRuleRequestWithDefaults

`func NewPatchedRequestPhaseRuleRequestWithDefaults() *PatchedRequestPhaseRuleRequest`

NewPatchedRequestPhaseRuleRequestWithDefaults instantiates a new PatchedRequestPhaseRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRequestPhaseRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRequestPhaseRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRequestPhaseRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRequestPhaseRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *PatchedRequestPhaseRuleRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedRequestPhaseRuleRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedRequestPhaseRuleRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedRequestPhaseRuleRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchedRequestPhaseRuleRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchedRequestPhaseRuleRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchedRequestPhaseRuleRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchedRequestPhaseRuleRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchedRequestPhaseRuleRequest) GetBehaviors() []RequestPhaseBehaviorRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchedRequestPhaseRuleRequest) GetBehaviorsOk() (*[]RequestPhaseBehaviorRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchedRequestPhaseRuleRequest) SetBehaviors(v []RequestPhaseBehaviorRequest)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchedRequestPhaseRuleRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRequestPhaseRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRequestPhaseRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRequestPhaseRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRequestPhaseRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


