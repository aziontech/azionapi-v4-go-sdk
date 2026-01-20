# BehaviorFilterRequestHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_request_header&#x60; - filter_request_header | 
**Attributes** | [**BehaviorFilterHeaderAttributesRequest**](BehaviorFilterHeaderAttributesRequest.md) |  | 

## Methods

### NewBehaviorFilterRequestHeaderRequest

`func NewBehaviorFilterRequestHeaderRequest(type_ string, attributes BehaviorFilterHeaderAttributesRequest, ) *BehaviorFilterRequestHeaderRequest`

NewBehaviorFilterRequestHeaderRequest instantiates a new BehaviorFilterRequestHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFilterRequestHeaderRequestWithDefaults

`func NewBehaviorFilterRequestHeaderRequestWithDefaults() *BehaviorFilterRequestHeaderRequest`

NewBehaviorFilterRequestHeaderRequestWithDefaults instantiates a new BehaviorFilterRequestHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFilterRequestHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFilterRequestHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFilterRequestHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFilterRequestHeaderRequest) GetAttributes() BehaviorFilterHeaderAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFilterRequestHeaderRequest) GetAttributesOk() (*BehaviorFilterHeaderAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFilterRequestHeaderRequest) SetAttributes(v BehaviorFilterHeaderAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


