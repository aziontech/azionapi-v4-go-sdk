# ResponseBehaviorSetCookieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;set_cookie&#x60; - set_cookie | 
**Attributes** | [**ResponseBehaviorSetCookieAttributesRequest**](ResponseBehaviorSetCookieAttributesRequest.md) |  | 

## Methods

### NewResponseBehaviorSetCookieRequest

`func NewResponseBehaviorSetCookieRequest(type_ string, attributes ResponseBehaviorSetCookieAttributesRequest, ) *ResponseBehaviorSetCookieRequest`

NewResponseBehaviorSetCookieRequest instantiates a new ResponseBehaviorSetCookieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBehaviorSetCookieRequestWithDefaults

`func NewResponseBehaviorSetCookieRequestWithDefaults() *ResponseBehaviorSetCookieRequest`

NewResponseBehaviorSetCookieRequestWithDefaults instantiates a new ResponseBehaviorSetCookieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseBehaviorSetCookieRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseBehaviorSetCookieRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseBehaviorSetCookieRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ResponseBehaviorSetCookieRequest) GetAttributes() ResponseBehaviorSetCookieAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResponseBehaviorSetCookieRequest) GetAttributesOk() (*ResponseBehaviorSetCookieAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResponseBehaviorSetCookieRequest) SetAttributes(v ResponseBehaviorSetCookieAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


