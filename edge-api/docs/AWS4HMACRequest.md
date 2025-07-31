# AWS4HMACRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;aws4_hmac_sha256&#x60; - AWS for HMAC - SHA256 | 
**Attributes** | [**AWS4HMACAttributesRequest**](AWS4HMACAttributesRequest.md) |  | 

## Methods

### NewAWS4HMACRequest

`func NewAWS4HMACRequest(type_ string, attributes AWS4HMACAttributesRequest, ) *AWS4HMACRequest`

NewAWS4HMACRequest instantiates a new AWS4HMACRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWS4HMACRequestWithDefaults

`func NewAWS4HMACRequestWithDefaults() *AWS4HMACRequest`

NewAWS4HMACRequestWithDefaults instantiates a new AWS4HMACRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWS4HMACRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWS4HMACRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWS4HMACRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AWS4HMACRequest) GetAttributes() AWS4HMACAttributesRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AWS4HMACRequest) GetAttributesOk() (*AWS4HMACAttributesRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AWS4HMACRequest) SetAttributes(v AWS4HMACAttributesRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


