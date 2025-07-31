# EdgeApplicationRuleEngineRunFunctionResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;run_function&#x60; - run_function | 
**Attributes** | [**EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest**](EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest.md) |  | 

## Methods

### NewEdgeApplicationRuleEngineRunFunctionResponseRequest

`func NewEdgeApplicationRuleEngineRunFunctionResponseRequest(type_ string, attributes EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest, ) *EdgeApplicationRuleEngineRunFunctionResponseRequest`

NewEdgeApplicationRuleEngineRunFunctionResponseRequest instantiates a new EdgeApplicationRuleEngineRunFunctionResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationRuleEngineRunFunctionResponseRequestWithDefaults

`func NewEdgeApplicationRuleEngineRunFunctionResponseRequestWithDefaults() *EdgeApplicationRuleEngineRunFunctionResponseRequest`

NewEdgeApplicationRuleEngineRunFunctionResponseRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRunFunctionResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) GetAttributes() EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) GetAttributesOk() (*EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EdgeApplicationRuleEngineRunFunctionResponseRequest) SetAttributes(v EdgeApplicationRuleEngineRunFunctionResponseAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


