# ApplicationRuleEngineResponsePhaseBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;capture_match_groups&#x60; - capture_match_groups | 
**Attributes** | [**ApplicationRuleEngineCaptureMatchGroupsAttributesRequest**](ApplicationRuleEngineCaptureMatchGroupsAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineResponsePhaseBehaviorsRequest

`func NewApplicationRuleEngineResponsePhaseBehaviorsRequest(type_ string, attributes ApplicationRuleEngineCaptureMatchGroupsAttributesRequest, ) *ApplicationRuleEngineResponsePhaseBehaviorsRequest`

NewApplicationRuleEngineResponsePhaseBehaviorsRequest instantiates a new ApplicationRuleEngineResponsePhaseBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults

`func NewApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults() *ApplicationRuleEngineResponsePhaseBehaviorsRequest`

NewApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults instantiates a new ApplicationRuleEngineResponsePhaseBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) GetAttributes() ApplicationRuleEngineCaptureMatchGroupsAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) GetAttributesOk() (*ApplicationRuleEngineCaptureMatchGroupsAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineResponsePhaseBehaviorsRequest) SetAttributes(v ApplicationRuleEngineCaptureMatchGroupsAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


