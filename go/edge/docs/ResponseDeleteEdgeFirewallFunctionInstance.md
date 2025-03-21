# ResponseDeleteEdgeFirewallFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | * &#x60;pending&#x60; - pending * &#x60;executed&#x60; - executed | 
**Data** | [**NullableEdgeFirewallFunctionInstance**](EdgeFirewallFunctionInstance.md) |  | 

## Methods

### NewResponseDeleteEdgeFirewallFunctionInstance

`func NewResponseDeleteEdgeFirewallFunctionInstance(state string, data NullableEdgeFirewallFunctionInstance, ) *ResponseDeleteEdgeFirewallFunctionInstance`

NewResponseDeleteEdgeFirewallFunctionInstance instantiates a new ResponseDeleteEdgeFirewallFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteEdgeFirewallFunctionInstanceWithDefaults

`func NewResponseDeleteEdgeFirewallFunctionInstanceWithDefaults() *ResponseDeleteEdgeFirewallFunctionInstance`

NewResponseDeleteEdgeFirewallFunctionInstanceWithDefaults instantiates a new ResponseDeleteEdgeFirewallFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) SetState(v string)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) GetData() EdgeFirewallFunctionInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) GetDataOk() (*EdgeFirewallFunctionInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) SetData(v EdgeFirewallFunctionInstance)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteEdgeFirewallFunctionInstance) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteEdgeFirewallFunctionInstance) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


