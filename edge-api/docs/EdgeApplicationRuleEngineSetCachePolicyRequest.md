# EdgeApplicationRuleEngineSetCachePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_cache_policy&#x60; - set_cache_policy | 
**Attributes** | [**EdgeApplicationRuleEngineSetCachePolicyAttributesRequest**](EdgeApplicationRuleEngineSetCachePolicyAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineSetCachePolicyRequest

`func NewEdgeApplicationRuleEngineSetCachePolicyRequest(type_ string, attributes EdgeApplicationRuleEngineSetCachePolicyAttributesRequest, ) *EdgeApplicationRuleEngineSetCachePolicyRequest`

NewEdgeApplicationRuleEngineSetCachePolicyRequest instantiates a new EdgeApplicationRuleEngineSetCachePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineSetCachePolicyRequestWithDefaults

`func NewEdgeApplicationRuleEngineSetCachePolicyRequestWithDefaults() *EdgeApplicationRuleEngineSetCachePolicyRequest`

NewEdgeApplicationRuleEngineSetCachePolicyRequestWithDefaults instantiates a new EdgeApplicationRuleEngineSetCachePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) GetAttributes() EdgeApplicationRuleEngineSetCachePolicyAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) GetAttributesOk() (*EdgeApplicationRuleEngineSetCachePolicyAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineSetCachePolicyRequest) SetAttributes(v EdgeApplicationRuleEngineSetCachePolicyAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


