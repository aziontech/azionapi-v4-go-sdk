# ResponseDeleteEdgeFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Data** | [**NullableEdgeFirewall**](EdgeFirewall.md) |  | 

## Methods

### NewResponseDeleteEdgeFirewall

`func NewResponseDeleteEdgeFirewall(state StateEnum, data NullableEdgeFirewall, ) *ResponseDeleteEdgeFirewall`

NewResponseDeleteEdgeFirewall instantiates a new ResponseDeleteEdgeFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteEdgeFirewallWithDefaults

`func NewResponseDeleteEdgeFirewallWithDefaults() *ResponseDeleteEdgeFirewall`

NewResponseDeleteEdgeFirewallWithDefaults instantiates a new ResponseDeleteEdgeFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteEdgeFirewall) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteEdgeFirewall) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteEdgeFirewall) SetState(v StateEnum)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteEdgeFirewall) GetData() EdgeFirewall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteEdgeFirewall) GetDataOk() (*EdgeFirewall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteEdgeFirewall) SetData(v EdgeFirewall)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteEdgeFirewall) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteEdgeFirewall) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


