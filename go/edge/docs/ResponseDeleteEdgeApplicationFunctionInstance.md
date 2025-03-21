# ResponseDeleteEdgeApplicationFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableEdgeApplicationFunctionInstance**](EdgeApplicationFunctionInstance.md) |  | 

## Methods

### NewResponseDeleteEdgeApplicationFunctionInstance

`func NewResponseDeleteEdgeApplicationFunctionInstance(state string, data NullableEdgeApplicationFunctionInstance, ) *ResponseDeleteEdgeApplicationFunctionInstance`

NewResponseDeleteEdgeApplicationFunctionInstance instantiates a new ResponseDeleteEdgeApplicationFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteEdgeApplicationFunctionInstanceWithDefaults

`func NewResponseDeleteEdgeApplicationFunctionInstanceWithDefaults() *ResponseDeleteEdgeApplicationFunctionInstance`

NewResponseDeleteEdgeApplicationFunctionInstanceWithDefaults instantiates a new ResponseDeleteEdgeApplicationFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) GetData() EdgeApplicationFunctionInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) GetDataOk() (*EdgeApplicationFunctionInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) SetData(v EdgeApplicationFunctionInstance)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteEdgeApplicationFunctionInstance) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteEdgeApplicationFunctionInstance) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


