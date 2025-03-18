# ResponseBadRequestEdgeFirewallRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **[]string** |  | [optional] 
**LastEditor** | Pointer to **[]string** |  | [optional] 
**LastModified** | Pointer to **[]string** |  | [optional] 
**Active** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **[]string** |  | [optional] 
**Order** | Pointer to **[]string** |  | [optional] 
**Behaviors** | Pointer to [**ResponseBadRequestEdgeApplicationRuleEngineBehaviors**](ResponseBadRequestEdgeApplicationRuleEngineBehaviors.md) |  | [optional] 
**Criteria** | Pointer to [**ResponseBadRequestEdgeApplicationRuleEngineBehaviors**](ResponseBadRequestEdgeApplicationRuleEngineBehaviors.md) |  | [optional] 
**Detail** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseBadRequestEdgeFirewallRuleEngine

`func NewResponseBadRequestEdgeFirewallRuleEngine() *ResponseBadRequestEdgeFirewallRuleEngine`

NewResponseBadRequestEdgeFirewallRuleEngine instantiates a new ResponseBadRequestEdgeFirewallRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBadRequestEdgeFirewallRuleEngineWithDefaults

`func NewResponseBadRequestEdgeFirewallRuleEngineWithDefaults() *ResponseBadRequestEdgeFirewallRuleEngine`

NewResponseBadRequestEdgeFirewallRuleEngineWithDefaults instantiates a new ResponseBadRequestEdgeFirewallRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetLastEditor() []string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetLastEditorOk() (*[]string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetLastEditor(v []string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetLastModified() []string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetLastModifiedOk() (*[]string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetLastModified(v []string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetActive

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetActive() []string`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetActiveOk() (*[]string, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetActive(v []string)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrder

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetOrder() []string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetOrderOk() (*[]string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetOrder(v []string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetBehaviors

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetBehaviors() ResponseBadRequestEdgeApplicationRuleEngineBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetBehaviorsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetBehaviors(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetCriteria() ResponseBadRequestEdgeApplicationRuleEngineBehaviors`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetCriteriaOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetCriteria(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDetail

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponseBadRequestEdgeFirewallRuleEngine) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


