# ApplicationRequestPhaseRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionField**]([]EdgeApplicationCriterionField.md) |  | 
**Behaviors** | [**[]ApplicationRuleEngineRequestPhaseBehaviors**](ApplicationRuleEngineRequestPhaseBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 
**LastEditor** | **NullableString** |  | 
**LastModified** | **NullableTime** |  | 

## Methods

### NewApplicationRequestPhaseRuleEngine

`func NewApplicationRequestPhaseRuleEngine(id int64, name string, criteria [][]EdgeApplicationCriterionField, behaviors []ApplicationRuleEngineRequestPhaseBehaviors, order int64, lastEditor NullableString, lastModified NullableTime, ) *ApplicationRequestPhaseRuleEngine`

NewApplicationRequestPhaseRuleEngine instantiates a new ApplicationRequestPhaseRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestPhaseRuleEngineWithDefaults

`func NewApplicationRequestPhaseRuleEngineWithDefaults() *ApplicationRequestPhaseRuleEngine`

NewApplicationRequestPhaseRuleEngineWithDefaults instantiates a new ApplicationRequestPhaseRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRequestPhaseRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRequestPhaseRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRequestPhaseRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationRequestPhaseRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequestPhaseRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequestPhaseRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationRequestPhaseRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationRequestPhaseRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationRequestPhaseRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationRequestPhaseRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *ApplicationRequestPhaseRuleEngine) GetCriteria() [][]EdgeApplicationCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApplicationRequestPhaseRuleEngine) GetCriteriaOk() (*[][]EdgeApplicationCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApplicationRequestPhaseRuleEngine) SetCriteria(v [][]EdgeApplicationCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *ApplicationRequestPhaseRuleEngine) GetBehaviors() []ApplicationRuleEngineRequestPhaseBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ApplicationRequestPhaseRuleEngine) GetBehaviorsOk() (*[]ApplicationRuleEngineRequestPhaseBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ApplicationRequestPhaseRuleEngine) SetBehaviors(v []ApplicationRuleEngineRequestPhaseBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *ApplicationRequestPhaseRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRequestPhaseRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRequestPhaseRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRequestPhaseRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *ApplicationRequestPhaseRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApplicationRequestPhaseRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApplicationRequestPhaseRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetLastEditor

`func (o *ApplicationRequestPhaseRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ApplicationRequestPhaseRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ApplicationRequestPhaseRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### SetLastEditorNil

`func (o *ApplicationRequestPhaseRuleEngine) SetLastEditorNil(b bool)`

 SetLastEditorNil sets the value for LastEditor to be an explicit nil

### UnsetLastEditor
`func (o *ApplicationRequestPhaseRuleEngine) UnsetLastEditor()`

UnsetLastEditor ensures that no value is present for LastEditor, not even an explicit nil
### GetLastModified

`func (o *ApplicationRequestPhaseRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApplicationRequestPhaseRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApplicationRequestPhaseRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *ApplicationRequestPhaseRuleEngine) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ApplicationRequestPhaseRuleEngine) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


