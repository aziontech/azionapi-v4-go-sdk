# ResponseAsyncEdgeFirewallFunctionInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**EdgeFirewallFunctionInstance**](EdgeFirewallFunctionInstance.md) |  | 

## Methods

### NewResponseAsyncEdgeFirewallFunctionInstance

`func NewResponseAsyncEdgeFirewallFunctionInstance(data EdgeFirewallFunctionInstance, ) *ResponseAsyncEdgeFirewallFunctionInstance`

NewResponseAsyncEdgeFirewallFunctionInstance instantiates a new ResponseAsyncEdgeFirewallFunctionInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAsyncEdgeFirewallFunctionInstanceWithDefaults

`func NewResponseAsyncEdgeFirewallFunctionInstanceWithDefaults() *ResponseAsyncEdgeFirewallFunctionInstance`

NewResponseAsyncEdgeFirewallFunctionInstanceWithDefaults instantiates a new ResponseAsyncEdgeFirewallFunctionInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) GetData() EdgeFirewallFunctionInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) GetDataOk() (*EdgeFirewallFunctionInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseAsyncEdgeFirewallFunctionInstance) SetData(v EdgeFirewallFunctionInstance)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


