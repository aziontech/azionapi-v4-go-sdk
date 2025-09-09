# ApplicationRuleEngineStringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**ApplicationRuleEngineStringAttributesRequest**](ApplicationRuleEngineStringAttributesRequest.md) |  | 

## Methods

### NewApplicationRuleEngineStringRequest

`func NewApplicationRuleEngineStringRequest(type_ string, attributes ApplicationRuleEngineStringAttributesRequest, ) *ApplicationRuleEngineStringRequest`

NewApplicationRuleEngineStringRequest instantiates a new ApplicationRuleEngineStringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRuleEngineStringRequestWithDefaults

`func NewApplicationRuleEngineStringRequestWithDefaults() *ApplicationRuleEngineStringRequest`

NewApplicationRuleEngineStringRequestWithDefaults instantiates a new ApplicationRuleEngineStringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRuleEngineStringRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRuleEngineStringRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRuleEngineStringRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ApplicationRuleEngineStringRequest) GetAttributes() ApplicationRuleEngineStringAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRuleEngineStringRequest) GetAttributesOk() (*ApplicationRuleEngineStringAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRuleEngineStringRequest) SetAttributes(v ApplicationRuleEngineStringAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


