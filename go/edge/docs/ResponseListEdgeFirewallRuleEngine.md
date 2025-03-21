# ResponseListEdgeFirewallRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**Behaviors** | [**[]EdgeFirewallBehaviorField**](EdgeFirewallBehaviorField.md) |  | 
**Criteria** | [**[][]EdgeFirewallCriterionField**]([]EdgeFirewallCriterionField.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | [readonly] 

## Methods

### NewResponseListEdgeFirewallRuleEngine

`func NewResponseListEdgeFirewallRuleEngine(id int64, name string, lastEditor string, lastModified time.Time, behaviors []EdgeFirewallBehaviorField, criteria [][]EdgeFirewallCriterionField, order int64, ) *ResponseListEdgeFirewallRuleEngine`

NewResponseListEdgeFirewallRuleEngine instantiates a new ResponseListEdgeFirewallRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListEdgeFirewallRuleEngineWithDefaults

`func NewResponseListEdgeFirewallRuleEngineWithDefaults() *ResponseListEdgeFirewallRuleEngine`

NewResponseListEdgeFirewallRuleEngineWithDefaults instantiates a new ResponseListEdgeFirewallRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListEdgeFirewallRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListEdgeFirewallRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListEdgeFirewallRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListEdgeFirewallRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *ResponseListEdgeFirewallRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListEdgeFirewallRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListEdgeFirewallRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListEdgeFirewallRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *ResponseListEdgeFirewallRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListEdgeFirewallRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListEdgeFirewallRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *ResponseListEdgeFirewallRuleEngine) GetBehaviors() []EdgeFirewallBehaviorField`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetBehaviorsOk() (*[]EdgeFirewallBehaviorField, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ResponseListEdgeFirewallRuleEngine) SetBehaviors(v []EdgeFirewallBehaviorField)`

SetBehaviors sets Behaviors field to given value.


### GetCriteria

`func (o *ResponseListEdgeFirewallRuleEngine) GetCriteria() [][]EdgeFirewallCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetCriteriaOk() (*[][]EdgeFirewallCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ResponseListEdgeFirewallRuleEngine) SetCriteria(v [][]EdgeFirewallCriterionField)`

SetCriteria sets Criteria field to given value.


### GetDescription

`func (o *ResponseListEdgeFirewallRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseListEdgeFirewallRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseListEdgeFirewallRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *ResponseListEdgeFirewallRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ResponseListEdgeFirewallRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ResponseListEdgeFirewallRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


