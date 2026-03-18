# BehaviorFilterRequestHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_request_header&#x60; - filter_request_header | 
**Attributes** | [**BehaviorFilterHeaderAttributes**](BehaviorFilterHeaderAttributes.md) |  | 

## Methods

### NewBehaviorFilterRequestHeader

`func NewBehaviorFilterRequestHeader(type_ string, attributes BehaviorFilterHeaderAttributes, ) *BehaviorFilterRequestHeader`

NewBehaviorFilterRequestHeader instantiates a new BehaviorFilterRequestHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorFilterRequestHeaderWithDefaults

`func NewBehaviorFilterRequestHeaderWithDefaults() *BehaviorFilterRequestHeader`

NewBehaviorFilterRequestHeaderWithDefaults instantiates a new BehaviorFilterRequestHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorFilterRequestHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorFilterRequestHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorFilterRequestHeader) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorFilterRequestHeader) GetAttributes() BehaviorFilterHeaderAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorFilterRequestHeader) GetAttributesOk() (*BehaviorFilterHeaderAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorFilterRequestHeader) SetAttributes(v BehaviorFilterHeaderAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


