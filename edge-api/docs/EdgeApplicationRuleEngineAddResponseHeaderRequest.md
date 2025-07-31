# EdgeApplicationRuleEngineAddResponseHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;add_response_header&#x60; - add_response_header | 
**Attributes** | [**EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest**](EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineAddResponseHeaderRequest

`func NewEdgeApplicationRuleEngineAddResponseHeaderRequest(type_ string, attributes EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest, ) *EdgeApplicationRuleEngineAddResponseHeaderRequest`

NewEdgeApplicationRuleEngineAddResponseHeaderRequest instantiates a new EdgeApplicationRuleEngineAddResponseHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineAddResponseHeaderRequestWithDefaults

`func NewEdgeApplicationRuleEngineAddResponseHeaderRequestWithDefaults() *EdgeApplicationRuleEngineAddResponseHeaderRequest`

NewEdgeApplicationRuleEngineAddResponseHeaderRequestWithDefaults instantiates a new EdgeApplicationRuleEngineAddResponseHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) GetAttributes() EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) GetAttributesOk() (*EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineAddResponseHeaderRequest) SetAttributes(v EdgeApplicationRuleEngineAddResponseHeaderAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


