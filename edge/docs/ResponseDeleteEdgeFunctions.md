# ResponseDeleteEdgeFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableEdgeFunctions**](EdgeFunctions.md) |  | 

## Methods

### NewResponseDeleteEdgeFunctions

`func NewResponseDeleteEdgeFunctions(state string, data NullableEdgeFunctions, ) *ResponseDeleteEdgeFunctions`

NewResponseDeleteEdgeFunctions instantiates a new ResponseDeleteEdgeFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteEdgeFunctionsWithDefaults

`func NewResponseDeleteEdgeFunctionsWithDefaults() *ResponseDeleteEdgeFunctions`

NewResponseDeleteEdgeFunctionsWithDefaults instantiates a new ResponseDeleteEdgeFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteEdgeFunctions) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteEdgeFunctions) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteEdgeFunctions) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteEdgeFunctions) GetData() EdgeFunctions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteEdgeFunctions) GetDataOk() (*EdgeFunctions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteEdgeFunctions) SetData(v EdgeFunctions)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteEdgeFunctions) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteEdgeFunctions) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


