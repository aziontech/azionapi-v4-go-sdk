# EdgeApplicationRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Phase** | **string** | * &#x60;default&#x60; - default * &#x60;request&#x60; - request * &#x60;response&#x60; - response | 
**Active** | Pointer to **bool** |  | [optional] 
**Behaviors** | [**[]EdgeApplicationBehaviorField**](EdgeApplicationBehaviorField.md) |  | 
**Criteria** | [**[][]EdgeApplicationCriterionField**]([]EdgeApplicationCriterionField.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | [readonly] 
**LastEditor** | **NullableString** |  | [readonly] 
**LastModified** | **NullableTime** |  | [readonly] 

## Methods

### NewEdgeApplicationRuleEngine

`func NewEdgeApplicationRuleEngine(id int64, name string, phase string, behaviors []EdgeApplicationBehaviorField, criteria [][]EdgeApplicationCriterionField, order int64, lastEditor NullableString, lastModified NullableTime, ) *EdgeApplicationRuleEngine`

NewEdgeApplicationRuleEngine instantiates a new EdgeApplicationRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineWithDefaults

`func NewEdgeApplicationRuleEngineWithDefaults() *EdgeApplicationRuleEngine`

NewEdgeApplicationRuleEngineWithDefaults instantiates a new EdgeApplicationRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeApplicationRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeApplicationRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeApplicationRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeApplicationRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeApplicationRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeApplicationRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetPhase

`func (o *EdgeApplicationRuleEngine) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *EdgeApplicationRuleEngine) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *EdgeApplicationRuleEngine) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetActive

`func (o *EdgeApplicationRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeApplicationRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeApplicationRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeApplicationRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *EdgeApplicationRuleEngine) GetBehaviors() []EdgeApplicationBehaviorField`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *EdgeApplicationRuleEngine) GetBehaviorsOk() (*[]EdgeApplicationBehaviorField, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *EdgeApplicationRuleEngine) SetBehaviors(v []EdgeApplicationBehaviorField)`

SetBehaviors sets Behaviors field to given value.


### GetCriteria

`func (o *EdgeApplicationRuleEngine) GetCriteria() [][]EdgeApplicationCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *EdgeApplicationRuleEngine) GetCriteriaOk() (*[][]EdgeApplicationCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *EdgeApplicationRuleEngine) SetCriteria(v [][]EdgeApplicationCriterionField)`

SetCriteria sets Criteria field to given value.


### GetDescription

`func (o *EdgeApplicationRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EdgeApplicationRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EdgeApplicationRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EdgeApplicationRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *EdgeApplicationRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EdgeApplicationRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EdgeApplicationRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetLastEditor

`func (o *EdgeApplicationRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeApplicationRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeApplicationRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### SetLastEditorNil

`func (o *EdgeApplicationRuleEngine) SetLastEditorNil(b bool)`

 SetLastEditorNil sets the value for LastEditor to be an explicit nil

### UnsetLastEditor
`func (o *EdgeApplicationRuleEngine) UnsetLastEditor()`

UnsetLastEditor ensures that no value is present for LastEditor, not even an explicit nil
### GetLastModified

`func (o *EdgeApplicationRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeApplicationRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeApplicationRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *EdgeApplicationRuleEngine) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *EdgeApplicationRuleEngine) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


