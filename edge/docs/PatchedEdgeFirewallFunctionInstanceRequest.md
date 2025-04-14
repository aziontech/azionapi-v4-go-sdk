# PatchedEdgeFirewallFunctionInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**JsonArgs** | Pointer to [**EdgeApplicationFunctionInstanceJsonArgs**](EdgeApplicationFunctionInstanceJsonArgs.md) |  | [optional] 
**EdgeFunction** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEdgeFirewallFunctionInstanceRequest

`func NewPatchedEdgeFirewallFunctionInstanceRequest() *PatchedEdgeFirewallFunctionInstanceRequest`

NewPatchedEdgeFirewallFunctionInstanceRequest instantiates a new PatchedEdgeFirewallFunctionInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeFirewallFunctionInstanceRequestWithDefaults

`func NewPatchedEdgeFirewallFunctionInstanceRequestWithDefaults() *PatchedEdgeFirewallFunctionInstanceRequest`

NewPatchedEdgeFirewallFunctionInstanceRequestWithDefaults instantiates a new PatchedEdgeFirewallFunctionInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetJsonArgs

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetJsonArgs() EdgeApplicationFunctionInstanceJsonArgs`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetJsonArgsOk() (*EdgeApplicationFunctionInstanceJsonArgs, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetJsonArgs(v EdgeApplicationFunctionInstanceJsonArgs)`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### GetEdgeFunction

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.

### HasEdgeFunction

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasEdgeFunction() bool`

HasEdgeFunction returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


