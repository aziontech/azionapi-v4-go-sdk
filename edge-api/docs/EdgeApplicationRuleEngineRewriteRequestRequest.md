# EdgeApplicationRuleEngineRewriteRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;rewrite_request&#x60; - rewrite_request | 
**Attributes** | [**EdgeApllicationRuleEngineRewriteRequestAttributesRequest**](EdgeApllicationRuleEngineRewriteRequestAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineRewriteRequestRequest

`func NewEdgeApplicationRuleEngineRewriteRequestRequest(type_ string, attributes EdgeApllicationRuleEngineRewriteRequestAttributesRequest, ) *EdgeApplicationRuleEngineRewriteRequestRequest`

NewEdgeApplicationRuleEngineRewriteRequestRequest instantiates a new EdgeApplicationRuleEngineRewriteRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineRewriteRequestRequestWithDefaults

`func NewEdgeApplicationRuleEngineRewriteRequestRequestWithDefaults() *EdgeApplicationRuleEngineRewriteRequestRequest`

NewEdgeApplicationRuleEngineRewriteRequestRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRewriteRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) GetAttributes() EdgeApllicationRuleEngineRewriteRequestAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) GetAttributesOk() (*EdgeApllicationRuleEngineRewriteRequestAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineRewriteRequestRequest) SetAttributes(v EdgeApllicationRuleEngineRewriteRequestAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


