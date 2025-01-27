# ResponseListErrorResponses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**EdgeApplicationId** | **int32** |  | [readonly] 
**OriginId** | Pointer to **NullableInt32** |  | [optional] 
**ErrorResponses** | [**[]NestedErrorResponse**](NestedErrorResponse.md) |  | 

## Methods

### NewResponseListErrorResponses

`func NewResponseListErrorResponses(id int32, name string, edgeApplicationId int32, errorResponses []NestedErrorResponse, ) *ResponseListErrorResponses`

NewResponseListErrorResponses instantiates a new ResponseListErrorResponses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListErrorResponsesWithDefaults

`func NewResponseListErrorResponsesWithDefaults() *ResponseListErrorResponses`

NewResponseListErrorResponsesWithDefaults instantiates a new ResponseListErrorResponses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListErrorResponses) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListErrorResponses) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListErrorResponses) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListErrorResponses) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListErrorResponses) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListErrorResponses) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeApplicationId

`func (o *ResponseListErrorResponses) GetEdgeApplicationId() int32`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *ResponseListErrorResponses) GetEdgeApplicationIdOk() (*int32, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *ResponseListErrorResponses) SetEdgeApplicationId(v int32)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.


### GetOriginId

`func (o *ResponseListErrorResponses) GetOriginId() int32`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *ResponseListErrorResponses) GetOriginIdOk() (*int32, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *ResponseListErrorResponses) SetOriginId(v int32)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *ResponseListErrorResponses) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### SetOriginIdNil

`func (o *ResponseListErrorResponses) SetOriginIdNil(b bool)`

 SetOriginIdNil sets the value for OriginId to be an explicit nil

### UnsetOriginId
`func (o *ResponseListErrorResponses) UnsetOriginId()`

UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
### GetErrorResponses

`func (o *ResponseListErrorResponses) GetErrorResponses() []NestedErrorResponse`

GetErrorResponses returns the ErrorResponses field if non-nil, zero value otherwise.

### GetErrorResponsesOk

`func (o *ResponseListErrorResponses) GetErrorResponsesOk() (*[]NestedErrorResponse, bool)`

GetErrorResponsesOk returns a tuple with the ErrorResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorResponses

`func (o *ResponseListErrorResponses) SetErrorResponses(v []NestedErrorResponse)`

SetErrorResponses sets ErrorResponses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


