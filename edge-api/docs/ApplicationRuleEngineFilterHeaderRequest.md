# ApplicationRuleEngineFilterHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_request_header&#x60; - filter_request_header | 
**Attributes** | [**ApplicationRuleEngineFilterHeaderAttributesRequest**](ApplicationRuleEngineFilterHeaderAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineFilterHeaderRequest

`func NewApplicationRuleEngineFilterHeaderRequest(type_ string, attributes ApplicationRuleEngineFilterHeaderAttributesRequest, ) *ApplicationRuleEngineFilterHeaderRequest`

NewApplicationRuleEngineFilterHeaderRequest instantiates a new ApplicationRuleEngineFilterHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineFilterHeaderRequestWithDefaults

`func NewApplicationRuleEngineFilterHeaderRequestWithDefaults() *ApplicationRuleEngineFilterHeaderRequest`

NewApplicationRuleEngineFilterHeaderRequestWithDefaults instantiates a new ApplicationRuleEngineFilterHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineFilterHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineFilterHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineFilterHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineFilterHeaderRequest) GetAttributes() ApplicationRuleEngineFilterHeaderAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineFilterHeaderRequest) GetAttributesOk() (*ApplicationRuleEngineFilterHeaderAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineFilterHeaderRequest) SetAttributes(v ApplicationRuleEngineFilterHeaderAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


