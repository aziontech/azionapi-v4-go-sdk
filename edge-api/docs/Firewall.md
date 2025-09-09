# Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Modules** | Pointer to [**FirewallModules**](FirewallModules.md) |  | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | **string** |  | 
**LastModified** | **time.Time** |  | 
**ProductVersion** | **string** |  | 

## Methods

### NewFirewall

`func NewFirewall(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, ) *Firewall`

NewFirewall instantiates a new Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallWithDefaults

`func NewFirewallWithDefaults() *Firewall`

NewFirewallWithDefaults instantiates a new Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Firewall) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Firewall) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Firewall) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Firewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Firewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Firewall) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *Firewall) GetModules() FirewallModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *Firewall) GetModulesOk() (*FirewallModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *Firewall) SetModules(v FirewallModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *Firewall) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetDebug

`func (o *Firewall) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Firewall) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Firewall) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *Firewall) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetActive

`func (o *Firewall) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Firewall) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Firewall) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Firewall) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *Firewall) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *Firewall) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *Firewall) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *Firewall) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Firewall) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Firewall) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetProductVersion

`func (o *Firewall) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *Firewall) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *Firewall) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


