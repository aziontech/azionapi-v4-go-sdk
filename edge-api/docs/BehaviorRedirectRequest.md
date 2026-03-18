# BehaviorRedirectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;redirect_to_301&#x60; - redirect_to_301 * &#x60;redirect_to_302&#x60; - redirect_to_302 | 
**Attributes** | [**BehaviorStringAttributesRequest**](BehaviorStringAttributesRequest.md) |  | 

## Methods

### NewBehaviorRedirectRequest

`func NewBehaviorRedirectRequest(type_ string, attributes BehaviorStringAttributesRequest, ) *BehaviorRedirectRequest`

NewBehaviorRedirectRequest instantiates a new BehaviorRedirectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRedirectRequestWithDefaults

`func NewBehaviorRedirectRequestWithDefaults() *BehaviorRedirectRequest`

NewBehaviorRedirectRequestWithDefaults instantiates a new BehaviorRedirectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorRedirectRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorRedirectRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorRedirectRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorRedirectRequest) GetAttributes() BehaviorStringAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorRedirectRequest) GetAttributesOk() (*BehaviorStringAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorRedirectRequest) SetAttributes(v BehaviorStringAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


