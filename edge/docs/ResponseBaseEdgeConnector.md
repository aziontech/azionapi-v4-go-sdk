# ResponseBaseEdgeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**BaseEdgeConnector**](BaseEdgeConnector.md) |  | 

## Methods

### NewResponseBaseEdgeConnector

`func NewResponseBaseEdgeConnector(state string, data BaseEdgeConnector, ) *ResponseBaseEdgeConnector`

NewResponseBaseEdgeConnector instantiates a new ResponseBaseEdgeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBaseEdgeConnectorWithDefaults

`func NewResponseBaseEdgeConnectorWithDefaults() *ResponseBaseEdgeConnector`

NewResponseBaseEdgeConnectorWithDefaults instantiates a new ResponseBaseEdgeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseBaseEdgeConnector) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseBaseEdgeConnector) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseBaseEdgeConnector) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseBaseEdgeConnector) GetData() BaseEdgeConnector`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseBaseEdgeConnector) GetDataOk() (*BaseEdgeConnector, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseBaseEdgeConnector) SetData(v BaseEdgeConnector)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


