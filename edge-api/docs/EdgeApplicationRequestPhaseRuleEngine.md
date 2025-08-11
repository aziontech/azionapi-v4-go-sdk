# EdgeApplicationRequestPhaseRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionField**]([]EdgeApplicationCriterionField.md) |  | 
**Behaviors** | [**[]EdgeApplicationRuleEngineRequestPhaseBehaviors**](EdgeApplicationRuleEngineRequestPhaseBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 
**LastEditor** | **NullableString** |  | 
**LastModified** | **NullableTime** |  | 

## Methods

### NewEdgeApplicationRequestPhaseRuleEngine

`func NewEdgeApplicationRequestPhaseRuleEngine(id int64, name string, criteria [][]EdgeApplicationCriterionField, behaviors []EdgeApplicationRuleEngineRequestPhaseBehaviors, order int64, lastEditor NullableString, lastModified NullableTime, ) *EdgeApplicationRequestPhaseRuleEngine`

NewEdgeApplicationRequestPhaseRuleEngine instantiates a new EdgeApplicationRequestPhaseRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRequestPhaseRuleEngineWithDefaults

`func NewEdgeApplicationRequestPhaseRuleEngineWithDefaults() *EdgeApplicationRequestPhaseRuleEngine`

NewEdgeApplicationRequestPhaseRuleEngineWithDefaults instantiates a new EdgeApplicationRequestPhaseRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationRequestPhaseRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetCriteria() [][]EdgeApplicationCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetCriteriaOk() (*[][]EdgeApplicationCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetCriteria(v [][]EdgeApplicationCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetBehaviors() []EdgeApplicationRuleEngineRequestPhaseBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetBehaviorsOk() (*[]EdgeApplicationRuleEngineRequestPhaseBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetBehaviors(v []EdgeApplicationRuleEngineRequestPhaseBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EdgeApplicationRequestPhaseRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetLastEditor

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### SetLastEditorNil

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetLastEditorNil(b bool)`

 SetLastEditorNil sets the value for LastEditor to be an explicit nil

### UnsetLastEditor
`func (o *EdgeApplicationRequestPhaseRuleEngine) UnsetLastEditor()`

UnsetLastEditor ensures that no value is present for LastEditor, not even an explicit nil
### GetLastModified

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeApplicationRequestPhaseRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *EdgeApplicationRequestPhaseRuleEngine) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *EdgeApplicationRequestPhaseRuleEngine) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


