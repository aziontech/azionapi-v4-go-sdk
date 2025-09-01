# ApplicationRuleEngineRewriteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;rewrite_request&#x60; - rewrite_request | 
**Attributes** | [**ApplicationRuleEngineRewriteRequestAttributes**](ApplicationRuleEngineRewriteRequestAttributes.md) |  | 

## Methods

### NewApplicationRuleEngineRewriteRequest

`func NewApplicationRuleEngineRewriteRequest(type_ string, attributes ApplicationRuleEngineRewriteRequestAttributes, ) *ApplicationRuleEngineRewriteRequest`

NewApplicationRuleEngineRewriteRequest instantiates a new ApplicationRuleEngineRewriteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineRewriteRequestWithDefaults

`func NewApplicationRuleEngineRewriteRequestWithDefaults() *ApplicationRuleEngineRewriteRequest`

NewApplicationRuleEngineRewriteRequestWithDefaults instantiates a new ApplicationRuleEngineRewriteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineRewriteRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineRewriteRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineRewriteRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineRewriteRequest) GetAttributes() ApplicationRuleEngineRewriteRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineRewriteRequest) GetAttributesOk() (*ApplicationRuleEngineRewriteRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineRewriteRequest) SetAttributes(v ApplicationRuleEngineRewriteRequestAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


