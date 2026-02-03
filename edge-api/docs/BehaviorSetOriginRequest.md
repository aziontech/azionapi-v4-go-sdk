# BehaviorSetOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_origin&#x60; - set_origin | 
**Attributes** | [**BehaviorSetOriginAttributesRequest**](BehaviorSetOriginAttributesRequest.md) |  | 

## Methods

### NewBehaviorSetOriginRequest

`func NewBehaviorSetOriginRequest(type_ string, attributes BehaviorSetOriginAttributesRequest, ) *BehaviorSetOriginRequest`

NewBehaviorSetOriginRequest instantiates a new BehaviorSetOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorSetOriginRequestWithDefaults

`func NewBehaviorSetOriginRequestWithDefaults() *BehaviorSetOriginRequest`

NewBehaviorSetOriginRequestWithDefaults instantiates a new BehaviorSetOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorSetOriginRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorSetOriginRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorSetOriginRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorSetOriginRequest) GetAttributes() BehaviorSetOriginAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorSetOriginRequest) GetAttributesOk() (*BehaviorSetOriginAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorSetOriginRequest) SetAttributes(v BehaviorSetOriginAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


