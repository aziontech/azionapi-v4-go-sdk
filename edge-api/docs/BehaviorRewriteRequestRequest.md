# BehaviorRewriteRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;rewrite_request&#x60; - rewrite_request | 
**Attributes** | [**BehaviorRewriteRequestAttributesRequest**](BehaviorRewriteRequestAttributesRequest.md) |  | 

## Methods

### NewBehaviorRewriteRequestRequest

`func NewBehaviorRewriteRequestRequest(type_ string, attributes BehaviorRewriteRequestAttributesRequest, ) *BehaviorRewriteRequestRequest`

NewBehaviorRewriteRequestRequest instantiates a new BehaviorRewriteRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRewriteRequestRequestWithDefaults

`func NewBehaviorRewriteRequestRequestWithDefaults() *BehaviorRewriteRequestRequest`

NewBehaviorRewriteRequestRequestWithDefaults instantiates a new BehaviorRewriteRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorRewriteRequestRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRewriteRequestRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRewriteRequestRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorRewriteRequestRequest) GetAttributes() BehaviorRewriteRequestAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorRewriteRequestRequest) GetAttributesOk() (*BehaviorRewriteRequestAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorRewriteRequestRequest) SetAttributes(v BehaviorRewriteRequestAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


