# PatchedEdgeFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to [**EdgeFirewallModulesRequest**](EdgeFirewallModulesRequest.md) |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedEdgeFirewallRequest

`func NewPatchedEdgeFirewallRequest() *PatchedEdgeFirewallRequest`

NewPatchedEdgeFirewallRequest instantiates a new PatchedEdgeFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeFirewallRequestWithDefaults

`func NewPatchedEdgeFirewallRequestWithDefaults() *PatchedEdgeFirewallRequest`

NewPatchedEdgeFirewallRequestWithDefaults instantiates a new PatchedEdgeFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeFirewallRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeFirewallRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchedEdgeFirewallRequest) GetModules() EdgeFirewallModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchedEdgeFirewallRequest) GetModulesOk() (*EdgeFirewallModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchedEdgeFirewallRequest) SetModules(v EdgeFirewallModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchedEdgeFirewallRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebug

`func (o *PatchedEdgeFirewallRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *PatchedEdgeFirewallRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *PatchedEdgeFirewallRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *PatchedEdgeFirewallRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeFirewallRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeFirewallRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeFirewallRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeFirewallRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


