# ResponseEdgeFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | [**EdgeFirewall**](EdgeFirewall.md) |  | 

## Methods

### NewResponseEdgeFirewall

`func NewResponseEdgeFirewall(data EdgeFirewall, ) *ResponseEdgeFirewall`

NewResponseEdgeFirewall instantiates a new ResponseEdgeFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEdgeFirewallWithDefaults

`func NewResponseEdgeFirewallWithDefaults() *ResponseEdgeFirewall`

NewResponseEdgeFirewallWithDefaults instantiates a new ResponseEdgeFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseEdgeFirewall) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseEdgeFirewall) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseEdgeFirewall) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseEdgeFirewall) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseEdgeFirewall) GetData() EdgeFirewall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseEdgeFirewall) GetDataOk() (*EdgeFirewall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseEdgeFirewall) SetData(v EdgeFirewall)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


