# ResponseEdgeApplicationRuleEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**EdgeApplicationRuleEngine**](EdgeApplicationRuleEngine.md) |  | 

## Methods

### NewResponseEdgeApplicationRuleEngine

`func NewResponseEdgeApplicationRuleEngine(state string, data EdgeApplicationRuleEngine, ) *ResponseEdgeApplicationRuleEngine`

NewResponseEdgeApplicationRuleEngine instantiates a new ResponseEdgeApplicationRuleEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEdgeApplicationRuleEngineWithDefaults

`func NewResponseEdgeApplicationRuleEngineWithDefaults() *ResponseEdgeApplicationRuleEngine`

NewResponseEdgeApplicationRuleEngineWithDefaults instantiates a new ResponseEdgeApplicationRuleEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseEdgeApplicationRuleEngine) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseEdgeApplicationRuleEngine) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseEdgeApplicationRuleEngine) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseEdgeApplicationRuleEngine) GetData() EdgeApplicationRuleEngine`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseEdgeApplicationRuleEngine) GetDataOk() (*EdgeApplicationRuleEngine, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseEdgeApplicationRuleEngine) SetData(v EdgeApplicationRuleEngine)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


