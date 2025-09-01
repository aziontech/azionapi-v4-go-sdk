# ApplicationRuleEngineAddHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;add_request_header&#x60; - add_request_header | 
**Attributes** | [**ApplicationRuleEngineAddHeaderAttributesRequest**](ApplicationRuleEngineAddHeaderAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineAddHeaderRequest

`func NewApplicationRuleEngineAddHeaderRequest(type_ string, attributes ApplicationRuleEngineAddHeaderAttributesRequest, ) *ApplicationRuleEngineAddHeaderRequest`

NewApplicationRuleEngineAddHeaderRequest instantiates a new ApplicationRuleEngineAddHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineAddHeaderRequestWithDefaults

`func NewApplicationRuleEngineAddHeaderRequestWithDefaults() *ApplicationRuleEngineAddHeaderRequest`

NewApplicationRuleEngineAddHeaderRequestWithDefaults instantiates a new ApplicationRuleEngineAddHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineAddHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineAddHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineAddHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineAddHeaderRequest) GetAttributes() ApplicationRuleEngineAddHeaderAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineAddHeaderRequest) GetAttributesOk() (*ApplicationRuleEngineAddHeaderAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineAddHeaderRequest) SetAttributes(v ApplicationRuleEngineAddHeaderAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


