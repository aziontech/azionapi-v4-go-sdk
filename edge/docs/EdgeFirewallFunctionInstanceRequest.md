# EdgeFirewallFunctionInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**EdgeFunction** | **int64** |  | 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeFirewallFunctionInstanceRequest

`func NewEdgeFirewallFunctionInstanceRequest(name string, edgeFunction int64, ) *EdgeFirewallFunctionInstanceRequest`

NewEdgeFirewallFunctionInstanceRequest instantiates a new EdgeFirewallFunctionInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallFunctionInstanceRequestWithDefaults

`func NewEdgeFirewallFunctionInstanceRequestWithDefaults() *EdgeFirewallFunctionInstanceRequest`

NewEdgeFirewallFunctionInstanceRequestWithDefaults instantiates a new EdgeFirewallFunctionInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFirewallFunctionInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallFunctionInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallFunctionInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetJsonArgs

`func (o *EdgeFirewallFunctionInstanceRequest) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *EdgeFirewallFunctionInstanceRequest) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *EdgeFirewallFunctionInstanceRequest) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *EdgeFirewallFunctionInstanceRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *EdgeFirewallFunctionInstanceRequest) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *EdgeFirewallFunctionInstanceRequest) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetEdgeFunction

`func (o *EdgeFirewallFunctionInstanceRequest) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *EdgeFirewallFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *EdgeFirewallFunctionInstanceRequest) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.


### GetActive

`func (o *EdgeFirewallFunctionInstanceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFirewallFunctionInstanceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFirewallFunctionInstanceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFirewallFunctionInstanceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


