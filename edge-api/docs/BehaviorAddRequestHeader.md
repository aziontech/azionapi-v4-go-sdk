# BehaviorAddRequestHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;add_request_header&#x60; - add_request_header | 
**Attributes** | [**BehaviorAddHeaderAttributes**](BehaviorAddHeaderAttributes.md) |  | 

## Methods

### NewBehaviorAddRequestHeader

`func NewBehaviorAddRequestHeader(type_ string, attributes BehaviorAddHeaderAttributes, ) *BehaviorAddRequestHeader`

NewBehaviorAddRequestHeader instantiates a new BehaviorAddRequestHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorAddRequestHeaderWithDefaults

`func NewBehaviorAddRequestHeaderWithDefaults() *BehaviorAddRequestHeader`

NewBehaviorAddRequestHeaderWithDefaults instantiates a new BehaviorAddRequestHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BehaviorAddRequestHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BehaviorAddRequestHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BehaviorAddRequestHeader) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BehaviorAddRequestHeader) GetAttributes() BehaviorAddHeaderAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BehaviorAddRequestHeader) GetAttributesOk() (*BehaviorAddHeaderAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BehaviorAddRequestHeader) SetAttributes(v BehaviorAddHeaderAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


