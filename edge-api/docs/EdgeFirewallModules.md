# EdgeFirewallModules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdosProtection** | [**EdgeFirewallModule**](EdgeFirewallModule.md) |  | 
**EdgeFunctions** | Pointer to [**EdgeFirewallModule**](EdgeFirewallModule.md) |  | [optional] 
**NetworkProtection** | Pointer to [**EdgeFirewallModule**](EdgeFirewallModule.md) |  | [optional] 
**Waf** | Pointer to [**EdgeFirewallModule**](EdgeFirewallModule.md) |  | [optional] 

## Methods

### NewEdgeFirewallModules

`func NewEdgeFirewallModules(ddosProtection EdgeFirewallModule, ) *EdgeFirewallModules`

NewEdgeFirewallModules instantiates a new EdgeFirewallModules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallModulesWithDefaults

`func NewEdgeFirewallModulesWithDefaults() *EdgeFirewallModules`

NewEdgeFirewallModulesWithDefaults instantiates a new EdgeFirewallModules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdosProtection

`func (o *EdgeFirewallModules) GetDdosProtection() EdgeFirewallModule`

GetDdosProtection returns the DdosProtection field if non-nil, zero value otherwise.

### GetDdosProtectionOk

`func (o *EdgeFirewallModules) GetDdosProtectionOk() (*EdgeFirewallModule, bool)`

GetDdosProtectionOk returns a tuple with the DdosProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosProtection

`func (o *EdgeFirewallModules) SetDdosProtection(v EdgeFirewallModule)`

SetDdosProtection sets DdosProtection field to given value.


### GetEdgeFunctions

`func (o *EdgeFirewallModules) GetEdgeFunctions() EdgeFirewallModule`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *EdgeFirewallModules) GetEdgeFunctionsOk() (*EdgeFirewallModule, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *EdgeFirewallModules) SetEdgeFunctions(v EdgeFirewallModule)`

SetEdgeFunctions sets EdgeFunctions field to given value.

### HasEdgeFunctions

`func (o *EdgeFirewallModules) HasEdgeFunctions() bool`

HasEdgeFunctions returns a boolean if a field has been set.

### GetNetworkProtection

`func (o *EdgeFirewallModules) GetNetworkProtection() EdgeFirewallModule`

GetNetworkProtection returns the NetworkProtection field if non-nil, zero value otherwise.

### GetNetworkProtectionOk

`func (o *EdgeFirewallModules) GetNetworkProtectionOk() (*EdgeFirewallModule, bool)`

GetNetworkProtectionOk returns a tuple with the NetworkProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtection

`func (o *EdgeFirewallModules) SetNetworkProtection(v EdgeFirewallModule)`

SetNetworkProtection sets NetworkProtection field to given value.

### HasNetworkProtection

`func (o *EdgeFirewallModules) HasNetworkProtection() bool`

HasNetworkProtection returns a boolean if a field has been set.

### GetWaf

`func (o *EdgeFirewallModules) GetWaf() EdgeFirewallModule`

GetWaf returns the Waf field if non-nil, zero value otherwise.

### GetWafOk

`func (o *EdgeFirewallModules) GetWafOk() (*EdgeFirewallModule, bool)`

GetWafOk returns a tuple with the Waf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaf

`func (o *EdgeFirewallModules) SetWaf(v EdgeFirewallModule)`

SetWaf sets Waf field to given value.

### HasWaf

`func (o *EdgeFirewallModules) HasWaf() bool`

HasWaf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


