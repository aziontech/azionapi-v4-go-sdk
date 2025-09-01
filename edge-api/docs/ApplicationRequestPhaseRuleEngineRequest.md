# ApplicationRequestPhaseRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]ApplicationRuleEngineRequestPhaseBehaviorsRequest**](ApplicationRuleEngineRequestPhaseBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationRequestPhaseRuleEngineRequest

`func NewApplicationRequestPhaseRuleEngineRequest(name string, criteria [][]EdgeApplicationCriterionFieldRequest, behaviors []ApplicationRuleEngineRequestPhaseBehaviorsRequest, ) *ApplicationRequestPhaseRuleEngineRequest`

NewApplicationRequestPhaseRuleEngineRequest instantiates a new ApplicationRequestPhaseRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseRuleEngineRequestWithDefaults

`func NewApplicationRequestPhaseRuleEngineRequestWithDefaults() *ApplicationRequestPhaseRuleEngineRequest`

NewApplicationRequestPhaseRuleEngineRequestWithDefaults instantiates a new ApplicationRequestPhaseRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequestPhaseRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationRequestPhaseRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationRequestPhaseRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApplicationRequestPhaseRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetBehaviors() []ApplicationRuleEngineRequestPhaseBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetBehaviorsOk() (*[]ApplicationRuleEngineRequestPhaseBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ApplicationRequestPhaseRuleEngineRequest) SetBehaviors(v []ApplicationRuleEngineRequestPhaseBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRequestPhaseRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRequestPhaseRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRequestPhaseRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


