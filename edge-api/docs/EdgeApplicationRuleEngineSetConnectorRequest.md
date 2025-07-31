# EdgeApplicationRuleEngineSetConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_edge_connector&#x60; - set_edge_connector | 
**Attributes** | [**EdgeApplicationRuleEngineSetConnectorAttributesRequest**](EdgeApplicationRuleEngineSetConnectorAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineSetConnectorRequest

`func NewEdgeApplicationRuleEngineSetConnectorRequest(type_ string, attributes EdgeApplicationRuleEngineSetConnectorAttributesRequest, ) *EdgeApplicationRuleEngineSetConnectorRequest`

NewEdgeApplicationRuleEngineSetConnectorRequest instantiates a new EdgeApplicationRuleEngineSetConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineSetConnectorRequestWithDefaults

`func NewEdgeApplicationRuleEngineSetConnectorRequestWithDefaults() *EdgeApplicationRuleEngineSetConnectorRequest`

NewEdgeApplicationRuleEngineSetConnectorRequestWithDefaults instantiates a new EdgeApplicationRuleEngineSetConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) GetAttributes() EdgeApplicationRuleEngineSetConnectorAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) GetAttributesOk() (*EdgeApplicationRuleEngineSetConnectorAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineSetConnectorRequest) SetAttributes(v EdgeApplicationRuleEngineSetConnectorAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


