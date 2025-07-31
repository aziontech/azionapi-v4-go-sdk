# EdgeFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**Modules** | Pointer to [**EdgeFirewallModules**](EdgeFirewallModules.md) |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**ProductVersion** | **string** |  | [readonly] 

## Methods

### NewEdgeFirewall

`func NewEdgeFirewall(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, ) *EdgeFirewall`

NewEdgeFirewall instantiates a new EdgeFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallWithDefaults

`func NewEdgeFirewallWithDefaults() *EdgeFirewall`

NewEdgeFirewallWithDefaults instantiates a new EdgeFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeFirewall) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFirewall) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFirewall) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *EdgeFirewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewall) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *EdgeFirewall) GetModules() EdgeFirewallModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeFirewall) GetModulesOk() (*EdgeFirewallModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeFirewall) SetModules(v EdgeFirewallModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *EdgeFirewall) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebug

`func (o *EdgeFirewall) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *EdgeFirewall) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *EdgeFirewall) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *EdgeFirewall) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetActive

`func (o *EdgeFirewall) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeFirewall) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeFirewall) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeFirewall) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *EdgeFirewall) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeFirewall) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeFirewall) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *EdgeFirewall) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeFirewall) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeFirewall) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *EdgeFirewall) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *EdgeFirewall) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *EdgeFirewall) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


