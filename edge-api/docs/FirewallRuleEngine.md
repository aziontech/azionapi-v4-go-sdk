# FirewallRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Criteria** | [**[][]EdgeFirewallCriterionField**]([]EdgeFirewallCriterionField.md) |  | 
**Behaviors** | [**[]FirewallBehaviors**](FirewallBehaviors.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Order** | **int64** |  | 

## Methods

### NewFirewallRuleEngine

`func NewFirewallRuleEngine(id int64, name string, lastEditor string, lastModified time.Time, criteria [][]EdgeFirewallCriterionField, behaviors []FirewallBehaviors, order int64, ) *FirewallRuleEngine`

NewFirewallRuleEngine instantiates a new FirewallRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleEngineWithDefaults

`func NewFirewallRuleEngineWithDefaults() *FirewallRuleEngine`

NewFirewallRuleEngineWithDefaults instantiates a new FirewallRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirewallRuleEngine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRuleEngine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRuleEngine) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *FirewallRuleEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRuleEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRuleEngine) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *FirewallRuleEngine) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *FirewallRuleEngine) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *FirewallRuleEngine) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *FirewallRuleEngine) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *FirewallRuleEngine) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *FirewallRuleEngine) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *FirewallRuleEngine) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FirewallRuleEngine) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FirewallRuleEngine) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FirewallRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCriteria

`func (o *FirewallRuleEngine) GetCriteria() [][]EdgeFirewallCriterionField`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *FirewallRuleEngine) GetCriteriaOk() (*[][]EdgeFirewallCriterionField, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *FirewallRuleEngine) SetCriteria(v [][]EdgeFirewallCriterionField)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *FirewallRuleEngine) GetBehaviors() []FirewallBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *FirewallRuleEngine) GetBehaviorsOk() (*[]FirewallBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *FirewallRuleEngine) SetBehaviors(v []FirewallBehaviors)`

SetBehaviors sets Behaviors field to given value.


### GetDescription

`func (o *FirewallRuleEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirewallRuleEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirewallRuleEngine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirewallRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *FirewallRuleEngine) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FirewallRuleEngine) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FirewallRuleEngine) SetOrder(v int64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


