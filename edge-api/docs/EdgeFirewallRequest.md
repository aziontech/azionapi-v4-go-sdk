# EdgeFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Modules** | Pointer to [**EdgeFirewallModulesRequest**](EdgeFirewallModulesRequest.md) |  | [optional] 
**DebugRules** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewEdgeFirewallRequest

`func NewEdgeFirewallRequest(name string, ) *EdgeFirewallRequest`

NewEdgeFirewallRequest instantiates a new EdgeFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallRequestWithDefaults

`func NewEdgeFirewallRequestWithDefaults() *EdgeFirewallRequest`

NewEdgeFirewallRequestWithDefaults instantiates a new EdgeFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewallRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *EdgeFirewallRequest) GetModules() EdgeFirewallModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeFirewallRequest) GetModulesOk() (*EdgeFirewallModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeFirewallRequest) SetModules(v EdgeFirewallModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *EdgeFirewallRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebugRules

`func (o *EdgeFirewallRequest) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *EdgeFirewallRequest) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *EdgeFirewallRequest) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.

### HasDebugRules

`func (o *EdgeFirewallRequest) HasDebugRules() bool`

HasDebugRules returns a boolean if a field has been set.

### GetActive

`func (o *EdgeFirewallRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFirewallRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFirewallRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFirewallRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


