# NestedErrorResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**CodeEnum**](CodeEnum.md) |  | 
**Timeout** | **int32** |  | 
**Uri** | Pointer to **NullableString** |  | [optional] 
**CustomStatusCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNestedErrorResponseRequest

`func NewNestedErrorResponseRequest(code CodeEnum, timeout int32, ) *NestedErrorResponseRequest`

NewNestedErrorResponseRequest instantiates a new NestedErrorResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedErrorResponseRequestWithDefaults

`func NewNestedErrorResponseRequestWithDefaults() *NestedErrorResponseRequest`

NewNestedErrorResponseRequestWithDefaults instantiates a new NestedErrorResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NestedErrorResponseRequest) GetCode() CodeEnum`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NestedErrorResponseRequest) GetCodeOk() (*CodeEnum, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NestedErrorResponseRequest) SetCode(v CodeEnum)`

SetCode sets Code field to given value.


### GetTimeout

`func (o *NestedErrorResponseRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NestedErrorResponseRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NestedErrorResponseRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetUri

`func (o *NestedErrorResponseRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NestedErrorResponseRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NestedErrorResponseRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NestedErrorResponseRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *NestedErrorResponseRequest) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *NestedErrorResponseRequest) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetCustomStatusCode

`func (o *NestedErrorResponseRequest) GetCustomStatusCode() string`

GetCustomStatusCode returns the CustomStatusCode field if non-nil, zero value otherwise.

### GetCustomStatusCodeOk

`func (o *NestedErrorResponseRequest) GetCustomStatusCodeOk() (*string, bool)`

GetCustomStatusCodeOk returns a tuple with the CustomStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStatusCode

`func (o *NestedErrorResponseRequest) SetCustomStatusCode(v string)`

SetCustomStatusCode sets CustomStatusCode field to given value.

### HasCustomStatusCode

`func (o *NestedErrorResponseRequest) HasCustomStatusCode() bool`

HasCustomStatusCode returns a boolean if a field has been set.

### SetCustomStatusCodeNil

`func (o *NestedErrorResponseRequest) SetCustomStatusCodeNil(b bool)`

 SetCustomStatusCodeNil sets the value for CustomStatusCode to be an explicit nil

### UnsetCustomStatusCode
`func (o *NestedErrorResponseRequest) UnsetCustomStatusCode()`

UnsetCustomStatusCode ensures that no value is present for CustomStatusCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


