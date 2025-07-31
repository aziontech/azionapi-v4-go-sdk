# EdgeFirewallModulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeFunctions** | Pointer to [**EdgeFirewallModuleRequest**](EdgeFirewallModuleRequest.md) |  | [optional] 
**NetworkProtection** | Pointer to [**EdgeFirewallModuleRequest**](EdgeFirewallModuleRequest.md) |  | [optional] 
**Waf** | Pointer to [**EdgeFirewallModuleRequest**](EdgeFirewallModuleRequest.md) |  | [optional] 

## Methods

### NewEdgeFirewallModulesRequest

`func NewEdgeFirewallModulesRequest() *EdgeFirewallModulesRequest`

NewEdgeFirewallModulesRequest instantiates a new EdgeFirewallModulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallModulesRequestWithDefaults

`func NewEdgeFirewallModulesRequestWithDefaults() *EdgeFirewallModulesRequest`

NewEdgeFirewallModulesRequestWithDefaults instantiates a new EdgeFirewallModulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeFunctions

`func (o *EdgeFirewallModulesRequest) GetEdgeFunctions() EdgeFirewallModuleRequest`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *EdgeFirewallModulesRequest) GetEdgeFunctionsOk() (*EdgeFirewallModuleRequest, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *EdgeFirewallModulesRequest) SetEdgeFunctions(v EdgeFirewallModuleRequest)`

SetEdgeFunctions sets EdgeFunctions field to given value.

### HasEdgeFunctions

`func (o *EdgeFirewallModulesRequest) HasEdgeFunctions() bool`

HasEdgeFunctions returns a boolean if a field has been set.

### GetNetworkProtection

`func (o *EdgeFirewallModulesRequest) GetNetworkProtection() EdgeFirewallModuleRequest`

GetNetworkProtection returns the NetworkProtection field if non-nil, zero value otherwise.

### GetNetworkProtectionOk

`func (o *EdgeFirewallModulesRequest) GetNetworkProtectionOk() (*EdgeFirewallModuleRequest, bool)`

GetNetworkProtectionOk returns a tuple with the NetworkProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtection

`func (o *EdgeFirewallModulesRequest) SetNetworkProtection(v EdgeFirewallModuleRequest)`

SetNetworkProtection sets NetworkProtection field to given value.

### HasNetworkProtection

`func (o *EdgeFirewallModulesRequest) HasNetworkProtection() bool`

HasNetworkProtection returns a boolean if a field has been set.

### GetWaf

`func (o *EdgeFirewallModulesRequest) GetWaf() EdgeFirewallModuleRequest`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *EdgeFirewallModulesRequest) GetWafOk() (*EdgeFirewallModuleRequest, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *EdgeFirewallModulesRequest) SetWaf(v EdgeFirewallModuleRequest)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *EdgeFirewallModulesRequest) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


