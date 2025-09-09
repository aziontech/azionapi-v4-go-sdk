# ApplicationRuleEngineRequestPhaseBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;add_request_cookie&#x60; - add_request_cookie | 
**Attributes** | [**ApplicationRuleEngineAddRequestCookieAttributesRequest**](ApplicationRuleEngineAddRequestCookieAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineRequestPhaseBehaviorsRequest

`func NewApplicationRuleEngineRequestPhaseBehaviorsRequest(type_ string, attributes ApplicationRuleEngineAddRequestCookieAttributesRequest, ) *ApplicationRuleEngineRequestPhaseBehaviorsRequest`

NewApplicationRuleEngineRequestPhaseBehaviorsRequest instantiates a new ApplicationRuleEngineRequestPhaseBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults

`func NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults() *ApplicationRuleEngineRequestPhaseBehaviorsRequest`

NewApplicationRuleEngineRequestPhaseBehaviorsRequestWithDefaults instantiates a new ApplicationRuleEngineRequestPhaseBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetAttributes() ApplicationRuleEngineAddRequestCookieAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) GetAttributesOk() (*ApplicationRuleEngineAddRequestCookieAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineRequestPhaseBehaviorsRequest) SetAttributes(v ApplicationRuleEngineAddRequestCookieAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


