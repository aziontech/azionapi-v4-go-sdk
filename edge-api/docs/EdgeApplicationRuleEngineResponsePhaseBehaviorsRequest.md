# EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;capture_match_groups&#x60; - capture_match_groups | 
**Attributes** | [**EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest**](EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequest

`func NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequest(type_ string, attributes EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest, ) *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest`

NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequest instantiates a new EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults

`func NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults() *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest`

NewEdgeApplicationRuleEngineResponsePhaseBehaviorsRequestWithDefaults instantiates a new EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) GetAttributes() EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) GetAttributesOk() (*EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest) SetAttributes(v EdgeApplicationRuleEngineCaptureMatchGroupsAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


