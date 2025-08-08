# EdgeApplicationRuleEngineStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**EdgeApplicationRuleEngineStringAttributesRequest**](EdgeApplicationRuleEngineStringAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineStringRequest

`func NewEdgeApplicationRuleEngineStringRequest(type_ string, attributes EdgeApplicationRuleEngineStringAttributesRequest, ) *EdgeApplicationRuleEngineStringRequest`

NewEdgeApplicationRuleEngineStringRequest instantiates a new EdgeApplicationRuleEngineStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineStringRequestWithDefaults

`func NewEdgeApplicationRuleEngineStringRequestWithDefaults() *EdgeApplicationRuleEngineStringRequest`

NewEdgeApplicationRuleEngineStringRequestWithDefaults instantiates a new EdgeApplicationRuleEngineStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineStringRequest) GetAttributes() EdgeApplicationRuleEngineStringAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineStringRequest) GetAttributesOk() (*EdgeApplicationRuleEngineStringAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineStringRequest) SetAttributes(v EdgeApplicationRuleEngineStringAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


