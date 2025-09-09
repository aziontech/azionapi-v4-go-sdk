# ResponseApplicationFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**ApplicationFunctionInstance**](ApplicationFunctionInstance.md) |  | 

## Methods

### NewResponseApplicationFunctionInstance

`func NewResponseApplicationFunctionInstance(data ApplicationFunctionInstance, ) *ResponseApplicationFunctionInstance`

NewResponseApplicationFunctionInstance instantiates a new ResponseApplicationFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseApplicationFunctionInstanceWithDefaults

`func NewResponseApplicationFunctionInstanceWithDefaults() *ResponseApplicationFunctionInstance`

NewResponseApplicationFunctionInstanceWithDefaults instantiates a new ResponseApplicationFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseApplicationFunctionInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseApplicationFunctionInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseApplicationFunctionInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseApplicationFunctionInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseApplicationFunctionInstance) GetData() ApplicationFunctionInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseApplicationFunctionInstance) GetDataOk() (*ApplicationFunctionInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseApplicationFunctionInstance) SetData(v ApplicationFunctionInstance)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


