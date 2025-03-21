# PatchedErrorResponsesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | Pointer to **NullableInt64** |  | [optional] 
**ErrorResponses** | Pointer to [**[]NestedErrorResponseRequest**](NestedErrorResponseRequest.md) |  | [optional] 

## Methods

### NewPatchedErrorResponsesRequest

`func NewPatchedErrorResponsesRequest() *PatchedErrorResponsesRequest`

NewPatchedErrorResponsesRequest instantiates a new PatchedErrorResponsesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedErrorResponsesRequestWithDefaults

`func NewPatchedErrorResponsesRequestWithDefaults() *PatchedErrorResponsesRequest`

NewPatchedErrorResponsesRequestWithDefaults instantiates a new PatchedErrorResponsesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *PatchedErrorResponsesRequest) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *PatchedErrorResponsesRequest) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *PatchedErrorResponsesRequest) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *PatchedErrorResponsesRequest) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### SetOriginIdNil

`func (o *PatchedErrorResponsesRequest) SetOriginIdNil(b bool)`

 SetOriginIdNil sets the value for OriginId to be an explicit nil

### UnsetOriginId
`func (o *PatchedErrorResponsesRequest) UnsetOriginId()`

UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
### GetErrorResponses

`func (o *PatchedErrorResponsesRequest) GetErrorResponses() []NestedErrorResponseRequest`

GetErrorResponses returns the ErrorResponses field if non-nil, zero value otherwise.

### GetErrorResponsesOk

`func (o *PatchedErrorResponsesRequest) GetErrorResponsesOk() (*[]NestedErrorResponseRequest, bool)`

GetErrorResponsesOk returns a tuple with the ErrorResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponses

`func (o *PatchedErrorResponsesRequest) SetErrorResponses(v []NestedErrorResponseRequest)`

SetErrorResponses sets ErrorResponses field to given value.

### HasErrorResponses

`func (o *PatchedErrorResponsesRequest) HasErrorResponses() bool`

HasErrorResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


