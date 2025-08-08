# EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;deny&#x60; - deny * &#x60;no_content&#x60; - no_content * &#x60;deliver&#x60; - deliver * &#x60;finish_request_phase&#x60; - finish_request_phase * &#x60;forward_cookies&#x60; - forward_cookies * &#x60;optimize_images&#x60; - optimize_images * &#x60;bypass_cache&#x60; - bypass_cache * &#x60;enable_gzip&#x60; - enable_gzip * &#x60;redirect_http_to_https&#x60; - redirect_http_to_https | 

## Methods

### NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequest

`func NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequest(type_ string, ) *EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest`

NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequest instantiates a new EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequestWithDefaults

`func NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequestWithDefaults() *EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest`

NewEdgeApplicationResponsePhaseBehaviorWithoutArgsRequestWithDefaults instantiates a new EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


