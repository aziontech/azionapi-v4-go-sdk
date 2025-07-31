# EdgeApplicationRuleEngineSetOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_origin&#x60; - set_origin | 
**Attributes** | [**EdgeApplicationRuleEngineSetOriginAttributesRequest**](EdgeApplicationRuleEngineSetOriginAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineSetOriginRequest

`func NewEdgeApplicationRuleEngineSetOriginRequest(type_ string, attributes EdgeApplicationRuleEngineSetOriginAttributesRequest, ) *EdgeApplicationRuleEngineSetOriginRequest`

NewEdgeApplicationRuleEngineSetOriginRequest instantiates a new EdgeApplicationRuleEngineSetOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineSetOriginRequestWithDefaults

`func NewEdgeApplicationRuleEngineSetOriginRequestWithDefaults() *EdgeApplicationRuleEngineSetOriginRequest`

NewEdgeApplicationRuleEngineSetOriginRequestWithDefaults instantiates a new EdgeApplicationRuleEngineSetOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineSetOriginRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineSetOriginRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineSetOriginRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineSetOriginRequest) GetAttributes() EdgeApplicationRuleEngineSetOriginAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineSetOriginRequest) GetAttributesOk() (*EdgeApplicationRuleEngineSetOriginAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineSetOriginRequest) SetAttributes(v EdgeApplicationRuleEngineSetOriginAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


