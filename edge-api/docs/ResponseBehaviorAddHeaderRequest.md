# ResponseBehaviorAddHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;add_response_header&#x60; - add_response_header | 
**Attributes** | [**ResponseBehaviorAddHeaderAttributesRequest**](ResponseBehaviorAddHeaderAttributesRequest.md) |  | 

## Methods

### NewResponseBehaviorAddHeaderRequest

`func NewResponseBehaviorAddHeaderRequest(type_ string, attributes ResponseBehaviorAddHeaderAttributesRequest, ) *ResponseBehaviorAddHeaderRequest`

NewResponseBehaviorAddHeaderRequest instantiates a new ResponseBehaviorAddHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBehaviorAddHeaderRequestWithDefaults

`func NewResponseBehaviorAddHeaderRequestWithDefaults() *ResponseBehaviorAddHeaderRequest`

NewResponseBehaviorAddHeaderRequestWithDefaults instantiates a new ResponseBehaviorAddHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseBehaviorAddHeaderRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseBehaviorAddHeaderRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseBehaviorAddHeaderRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ResponseBehaviorAddHeaderRequest) GetAttributes() ResponseBehaviorAddHeaderAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResponseBehaviorAddHeaderRequest) GetAttributesOk() (*ResponseBehaviorAddHeaderAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResponseBehaviorAddHeaderRequest) SetAttributes(v ResponseBehaviorAddHeaderAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


