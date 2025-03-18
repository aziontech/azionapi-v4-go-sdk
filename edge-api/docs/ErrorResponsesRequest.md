# ErrorResponsesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | Pointer to **NullableInt32** |  | [optional] 
**ErrorResponses** | [**[]NestedErrorResponseRequest**](NestedErrorResponseRequest.md) |  | 

## Methods

### NewErrorResponsesRequest

`func NewErrorResponsesRequest(errorResponses []NestedErrorResponseRequest, ) *ErrorResponsesRequest`

NewErrorResponsesRequest instantiates a new ErrorResponsesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponsesRequestWithDefaults

`func NewErrorResponsesRequestWithDefaults() *ErrorResponsesRequest`

NewErrorResponsesRequestWithDefaults instantiates a new ErrorResponsesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *ErrorResponsesRequest) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *ErrorResponsesRequest) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *ErrorResponsesRequest) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *ErrorResponsesRequest) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### SetOriginIdNil

`func (o *ErrorResponsesRequest) SetOriginIdNil(b bool)`

 SetOriginIdNil sets the value for OriginId to be an explicit nil

### UnsetOriginId
`func (o *ErrorResponsesRequest) UnsetOriginId()`

UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
### GetErrorResponses

`func (o *ErrorResponsesRequest) GetErrorResponses() []NestedErrorResponseRequest`

GetErrorResponses returns the ErrorResponses field if non-nil, zero value otherwise.

### GetErrorResponsesOk

`func (o *ErrorResponsesRequest) GetErrorResponsesOk() (*[]NestedErrorResponseRequest, bool)`

GetErrorResponsesOk returns a tuple with the ErrorResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponses

`func (o *ErrorResponsesRequest) SetErrorResponses(v []NestedErrorResponseRequest)`

SetErrorResponses sets ErrorResponses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


