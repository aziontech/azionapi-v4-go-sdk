# TransformPolymorphicTransformRenderTemplateAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type identifier for this endpoint (render_template) | 
**Attributes** | [**TransformRenderTemplateRequest**](TransformRenderTemplateRequest.md) |  | 

## Methods

### NewTransformPolymorphicTransformRenderTemplateAttributesRequest

`func NewTransformPolymorphicTransformRenderTemplateAttributesRequest(type_ string, attributes TransformRenderTemplateRequest, ) *TransformPolymorphicTransformRenderTemplateAttributesRequest`

NewTransformPolymorphicTransformRenderTemplateAttributesRequest instantiates a new TransformPolymorphicTransformRenderTemplateAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformPolymorphicTransformRenderTemplateAttributesRequestWithDefaults

`func NewTransformPolymorphicTransformRenderTemplateAttributesRequestWithDefaults() *TransformPolymorphicTransformRenderTemplateAttributesRequest`

NewTransformPolymorphicTransformRenderTemplateAttributesRequestWithDefaults instantiates a new TransformPolymorphicTransformRenderTemplateAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) GetAttributes() TransformRenderTemplateRequest`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) GetAttributesOk() (*TransformRenderTemplateRequest, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TransformPolymorphicTransformRenderTemplateAttributesRequest) SetAttributes(v TransformRenderTemplateRequest)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


