# ResponseAsyncApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**Application**](Application.md) |  | 

## Methods

### NewResponseAsyncApplication

`func NewResponseAsyncApplication(data Application, ) *ResponseAsyncApplication`

NewResponseAsyncApplication instantiates a new ResponseAsyncApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncApplicationWithDefaults

`func NewResponseAsyncApplicationWithDefaults() *ResponseAsyncApplication`

NewResponseAsyncApplicationWithDefaults instantiates a new ResponseAsyncApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncApplication) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncApplication) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncApplication) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncApplication) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncApplication) GetData() Application`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncApplication) GetDataOk() (*Application, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncApplication) SetData(v Application)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


