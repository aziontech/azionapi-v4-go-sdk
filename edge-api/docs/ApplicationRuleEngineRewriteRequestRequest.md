# ApplicationRuleEngineRewriteRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;rewrite_request&#x60; - rewrite_request | 
**Attributes** | [**ApplicationRuleEngineRewriteRequestAttributesRequest**](ApplicationRuleEngineRewriteRequestAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineRewriteRequestRequest

`func NewApplicationRuleEngineRewriteRequestRequest(type_ string, attributes ApplicationRuleEngineRewriteRequestAttributesRequest, ) *ApplicationRuleEngineRewriteRequestRequest`

NewApplicationRuleEngineRewriteRequestRequest instantiates a new ApplicationRuleEngineRewriteRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineRewriteRequestRequestWithDefaults

`func NewApplicationRuleEngineRewriteRequestRequestWithDefaults() *ApplicationRuleEngineRewriteRequestRequest`

NewApplicationRuleEngineRewriteRequestRequestWithDefaults instantiates a new ApplicationRuleEngineRewriteRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineRewriteRequestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineRewriteRequestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineRewriteRequestRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineRewriteRequestRequest) GetAttributes() ApplicationRuleEngineRewriteRequestAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineRewriteRequestRequest) GetAttributesOk() (*ApplicationRuleEngineRewriteRequestAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineRewriteRequestRequest) SetAttributes(v ApplicationRuleEngineRewriteRequestAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


