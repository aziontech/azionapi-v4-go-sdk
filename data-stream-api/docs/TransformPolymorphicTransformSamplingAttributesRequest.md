# TransformPolymorphicTransformSamplingAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;sampling&#x60; - Sampling | 
**Attributes** | [**TransformSamplingRequest**](TransformSamplingRequest.md) |  | 

## Methods

### NewTransformPolymorphicTransformSamplingAttributesRequest

`func NewTransformPolymorphicTransformSamplingAttributesRequest(type_ string, attributes TransformSamplingRequest, ) *TransformPolymorphicTransformSamplingAttributesRequest`

NewTransformPolymorphicTransformSamplingAttributesRequest instantiates a new TransformPolymorphicTransformSamplingAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformPolymorphicTransformSamplingAttributesRequestWithDefaults

`func NewTransformPolymorphicTransformSamplingAttributesRequestWithDefaults() *TransformPolymorphicTransformSamplingAttributesRequest`

NewTransformPolymorphicTransformSamplingAttributesRequestWithDefaults instantiates a new TransformPolymorphicTransformSamplingAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) GetAttributes() TransformSamplingRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) GetAttributesOk() (*TransformSamplingRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransformPolymorphicTransformSamplingAttributesRequest) SetAttributes(v TransformSamplingRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


