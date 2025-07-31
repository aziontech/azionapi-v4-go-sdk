# ResponseAsyncLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**LoginResponse**](LoginResponse.md) |  | 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseAsyncLogin

`func NewResponseAsyncLogin(data LoginResponse, ) *ResponseAsyncLogin`

NewResponseAsyncLogin instantiates a new ResponseAsyncLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncLoginWithDefaults

`func NewResponseAsyncLoginWithDefaults() *ResponseAsyncLogin`

NewResponseAsyncLoginWithDefaults instantiates a new ResponseAsyncLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseAsyncLogin) GetData() LoginResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncLogin) GetDataOk() (*LoginResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncLogin) SetData(v LoginResponse)`

SetData sets Data field to given value.


### GetState

`func (o *ResponseAsyncLogin) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncLogin) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncLogin) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncLogin) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


