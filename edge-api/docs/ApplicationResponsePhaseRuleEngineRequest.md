# ApplicationResponsePhaseRuleEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionFieldRequest**]([]EdgeApplicationCriterionFieldRequest.md) |  | 
**Behaviors** | [**[]ApplicationRuleEngineResponsePhaseBehaviorsRequest**](ApplicationRuleEngineResponsePhaseBehaviorsRequest.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationResponsePhaseRuleEngineRequest

`func NewApplicationResponsePhaseRuleEngineRequest(name string, criteria [][]EdgeApplicationCriterionFieldRequest, behaviors []ApplicationRuleEngineResponsePhaseBehaviorsRequest, ) *ApplicationResponsePhaseRuleEngineRequest`

NewApplicationResponsePhaseRuleEngineRequest instantiates a new ApplicationResponsePhaseRuleEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponsePhaseRuleEngineRequestWithDefaults

`func NewApplicationResponsePhaseRuleEngineRequestWithDefaults() *ApplicationResponsePhaseRuleEngineRequest`

NewApplicationResponsePhaseRuleEngineRequestWithDefaults instantiates a new ApplicationResponsePhaseRuleEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponsePhaseRuleEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationResponsePhaseRuleEngineRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationResponsePhaseRuleEngineRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetCriteriaOk() (*[][]EdgeApplicationCriterionFieldRequest, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApplicationResponsePhaseRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetBehaviors() []ApplicationRuleEngineResponsePhaseBehaviorsRequest`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetBehaviorsOk() (*[]ApplicationRuleEngineResponsePhaseBehaviorsRequest, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ApplicationResponsePhaseRuleEngineRequest) SetBehaviors(v []ApplicationRuleEngineResponsePhaseBehaviorsRequest)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponsePhaseRuleEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponsePhaseRuleEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponsePhaseRuleEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


