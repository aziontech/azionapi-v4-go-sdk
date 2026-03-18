# ResponseBehaviorFilterHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;filter_response_header&#x60; - filter_response_header | 
**Attributes** | [**ResponseBehaviorFilterHeaderAttributes**](ResponseBehaviorFilterHeaderAttributes.md) |  | 

## Methods

### NewResponseBehaviorFilterHeader

`func NewResponseBehaviorFilterHeader(type_ string, attributes ResponseBehaviorFilterHeaderAttributes, ) *ResponseBehaviorFilterHeader`

NewResponseBehaviorFilterHeader instantiates a new ResponseBehaviorFilterHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBehaviorFilterHeaderWithDefaults

`func NewResponseBehaviorFilterHeaderWithDefaults() *ResponseBehaviorFilterHeader`

NewResponseBehaviorFilterHeaderWithDefaults instantiates a new ResponseBehaviorFilterHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponseBehaviorFilterHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseBehaviorFilterHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseBehaviorFilterHeader) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ResponseBehaviorFilterHeader) GetAttributes() ResponseBehaviorFilterHeaderAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResponseBehaviorFilterHeader) GetAttributesOk() (*ResponseBehaviorFilterHeaderAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResponseBehaviorFilterHeader) SetAttributes(v ResponseBehaviorFilterHeaderAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


