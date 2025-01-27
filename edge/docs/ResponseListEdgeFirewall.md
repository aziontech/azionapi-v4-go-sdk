# ResponseListEdgeFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Modules** | Pointer to [**EdgeFirewallModules**](EdgeFirewallModules.md) |  | [optional] 
**DebugRules** | Pointer to **bool** |  | [optional] [default to false]
**Active** | Pointer to **bool** |  | [optional] [default to true]
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewResponseListEdgeFirewall

`func NewResponseListEdgeFirewall(id int32, name string, lastEditor string, lastModified time.Time, productVersion string, ) *ResponseListEdgeFirewall`

NewResponseListEdgeFirewall instantiates a new ResponseListEdgeFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListEdgeFirewallWithDefaults

`func NewResponseListEdgeFirewallWithDefaults() *ResponseListEdgeFirewall`

NewResponseListEdgeFirewallWithDefaults instantiates a new ResponseListEdgeFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponseListEdgeFirewall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponseListEdgeFirewall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponseListEdgeFirewall) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ResponseListEdgeFirewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseListEdgeFirewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseListEdgeFirewall) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *ResponseListEdgeFirewall) GetModules() EdgeFirewallModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ResponseListEdgeFirewall) GetModulesOk() (*EdgeFirewallModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ResponseListEdgeFirewall) SetModules(v EdgeFirewallModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ResponseListEdgeFirewall) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebugRules

`func (o *ResponseListEdgeFirewall) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *ResponseListEdgeFirewall) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *ResponseListEdgeFirewall) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.

### HasDebugRules

`func (o *ResponseListEdgeFirewall) HasDebugRules() bool`

HasDebugRules returns a boolean if a field has been set.

### GetActive

`func (o *ResponseListEdgeFirewall) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ResponseListEdgeFirewall) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ResponseListEdgeFirewall) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ResponseListEdgeFirewall) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *ResponseListEdgeFirewall) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ResponseListEdgeFirewall) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ResponseListEdgeFirewall) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ResponseListEdgeFirewall) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ResponseListEdgeFirewall) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ResponseListEdgeFirewall) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *ResponseListEdgeFirewall) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ResponseListEdgeFirewall) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ResponseListEdgeFirewall) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


