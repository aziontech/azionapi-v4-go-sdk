# ResponseDeleteEdgeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableEdgeApplication**](EdgeApplication.md) |  | 

## Methods

### NewResponseDeleteEdgeApplication

`func NewResponseDeleteEdgeApplication(state string, data NullableEdgeApplication, ) *ResponseDeleteEdgeApplication`

NewResponseDeleteEdgeApplication instantiates a new ResponseDeleteEdgeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteEdgeApplicationWithDefaults

`func NewResponseDeleteEdgeApplicationWithDefaults() *ResponseDeleteEdgeApplication`

NewResponseDeleteEdgeApplicationWithDefaults instantiates a new ResponseDeleteEdgeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteEdgeApplication) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteEdgeApplication) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteEdgeApplication) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteEdgeApplication) GetData() EdgeApplication`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteEdgeApplication) GetDataOk() (*EdgeApplication, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteEdgeApplication) SetData(v EdgeApplication)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteEdgeApplication) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteEdgeApplication) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


