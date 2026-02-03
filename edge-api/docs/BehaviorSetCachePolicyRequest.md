# BehaviorSetCachePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_cache_policy&#x60; - set_cache_policy | 
**Attributes** | [**BehaviorSetCachePolicyAttributesRequest**](BehaviorSetCachePolicyAttributesRequest.md) |  | 

## Methods

### NewBehaviorSetCachePolicyRequest

`func NewBehaviorSetCachePolicyRequest(type_ string, attributes BehaviorSetCachePolicyAttributesRequest, ) *BehaviorSetCachePolicyRequest`

NewBehaviorSetCachePolicyRequest instantiates a new BehaviorSetCachePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorSetCachePolicyRequestWithDefaults

`func NewBehaviorSetCachePolicyRequestWithDefaults() *BehaviorSetCachePolicyRequest`

NewBehaviorSetCachePolicyRequestWithDefaults instantiates a new BehaviorSetCachePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorSetCachePolicyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorSetCachePolicyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorSetCachePolicyRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorSetCachePolicyRequest) GetAttributes() BehaviorSetCachePolicyAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorSetCachePolicyRequest) GetAttributesOk() (*BehaviorSetCachePolicyAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorSetCachePolicyRequest) SetAttributes(v BehaviorSetCachePolicyAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


