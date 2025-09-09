# ApplicationResponsePhaseRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeApplicationCriterionField**]([]EdgeApplicationCriterionField.md) |  | 
**Behaviors** | [**[]ApplicationRuleEngineResponsePhaseBehaviors**](ApplicationRuleEngineResponsePhaseBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 
**LastEditor** | **NullableString** |  | 
**LastModified** | **NullableTime** |  | 

## Methods

### NewApplicationResponsePhaseRuleEngine

`func NewApplicationResponsePhaseRuleEngine(id int64, name string, criteria [][]EdgeApplicationCriterionField, behaviors []ApplicationRuleEngineResponsePhaseBehaviors, order int64, lastEditor NullableString, lastModified NullableTime, ) *ApplicationResponsePhaseRuleEngine`

NewApplicationResponsePhaseRuleEngine instantiates a new ApplicationResponsePhaseRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponsePhaseRuleEngineWithDefaults

`func NewApplicationResponsePhaseRuleEngineWithDefaults() *ApplicationResponsePhaseRuleEngine`

NewApplicationResponsePhaseRuleEngineWithDefaults instantiates a new ApplicationResponsePhaseRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResponsePhaseRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponsePhaseRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponsePhaseRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationResponsePhaseRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponsePhaseRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponsePhaseRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationResponsePhaseRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationResponsePhaseRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationResponsePhaseRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationResponsePhaseRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *ApplicationResponsePhaseRuleEngine) GetCriteria() [][]EdgeApplicationCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApplicationResponsePhaseRuleEngine) GetCriteriaOk() (*[][]EdgeApplicationCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApplicationResponsePhaseRuleEngine) SetCriteria(v [][]EdgeApplicationCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *ApplicationResponsePhaseRuleEngine) GetBehaviors() []ApplicationRuleEngineResponsePhaseBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ApplicationResponsePhaseRuleEngine) GetBehaviorsOk() (*[]ApplicationRuleEngineResponsePhaseBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ApplicationResponsePhaseRuleEngine) SetBehaviors(v []ApplicationRuleEngineResponsePhaseBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *ApplicationResponsePhaseRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponsePhaseRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponsePhaseRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponsePhaseRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *ApplicationResponsePhaseRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApplicationResponsePhaseRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApplicationResponsePhaseRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetLastEditor

`func (o *ApplicationResponsePhaseRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ApplicationResponsePhaseRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ApplicationResponsePhaseRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### SetLastEditorNil

`func (o *ApplicationResponsePhaseRuleEngine) SetLastEditorNil(b bool)`

 SetLastEditorNil sets the value for LastEditor to be an explicit nil

### UnsetLastEditor
`func (o *ApplicationResponsePhaseRuleEngine) UnsetLastEditor()`

UnsetLastEditor ensures that no value is present for LastEditor, not even an explicit nil
### GetLastModified

`func (o *ApplicationResponsePhaseRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApplicationResponsePhaseRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApplicationResponsePhaseRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *ApplicationResponsePhaseRuleEngine) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ApplicationResponsePhaseRuleEngine) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


