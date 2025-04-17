# BaseEdgeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [readonly] 
**Name** | **string** |  | 
**LastEditor** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Modules** | [**EdgeConnectorModules**](EdgeConnectorModules.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**ProductVersion** | **string** |  | [readonly] 
**Type** | **string** | * &#x60;http&#x60; - HTTP * &#x60;s3&#x60; - S3 * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**Tls** | Pointer to [**TLSEdgeConnector**](TLSEdgeConnector.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 

## Methods

### NewBaseEdgeConnector

`func NewBaseEdgeConnector(id int64, name string, lastEditor string, lastModified time.Time, modules EdgeConnectorModules, productVersion string, type_ string, ) *BaseEdgeConnector`

NewBaseEdgeConnector instantiates a new BaseEdgeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEdgeConnectorWithDefaults

`func NewBaseEdgeConnectorWithDefaults() *BaseEdgeConnector`

NewBaseEdgeConnectorWithDefaults instantiates a new BaseEdgeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseEdgeConnector) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEdgeConnector) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEdgeConnector) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *BaseEdgeConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseEdgeConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseEdgeConnector) SetName(v string)`

SetName sets Name field to given value.


### GetLastEditor

`func (o *BaseEdgeConnector) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *BaseEdgeConnector) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *BaseEdgeConnector) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *BaseEdgeConnector) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BaseEdgeConnector) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BaseEdgeConnector) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetModules

`func (o *BaseEdgeConnector) GetModules() EdgeConnectorModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *BaseEdgeConnector) GetModulesOk() (*EdgeConnectorModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *BaseEdgeConnector) SetModules(v EdgeConnectorModules)`

SetModules sets Modules field to given value.


### GetActive

`func (o *BaseEdgeConnector) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BaseEdgeConnector) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BaseEdgeConnector) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *BaseEdgeConnector) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetProductVersion

`func (o *BaseEdgeConnector) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *BaseEdgeConnector) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *BaseEdgeConnector) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetType

`func (o *BaseEdgeConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEdgeConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEdgeConnector) SetType(v string)`

SetType sets Type field to given value.


### GetAddresses

`func (o *BaseEdgeConnector) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BaseEdgeConnector) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BaseEdgeConnector) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BaseEdgeConnector) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *BaseEdgeConnector) GetTls() TLSEdgeConnector`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *BaseEdgeConnector) GetTlsOk() (*TLSEdgeConnector, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *BaseEdgeConnector) SetTls(v TLSEdgeConnector)`

SetTls sets Tls field to given value.

### HasTls

`func (o *BaseEdgeConnector) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *BaseEdgeConnector) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *BaseEdgeConnector) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *BaseEdgeConnector) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *BaseEdgeConnector) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *BaseEdgeConnector) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *BaseEdgeConnector) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *BaseEdgeConnector) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *BaseEdgeConnector) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *BaseEdgeConnector) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *BaseEdgeConnector) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *BaseEdgeConnector) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *BaseEdgeConnector) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *BaseEdgeConnector) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *BaseEdgeConnector) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *BaseEdgeConnector) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *BaseEdgeConnector) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *BaseEdgeConnector) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *BaseEdgeConnector) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *BaseEdgeConnector) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *BaseEdgeConnector) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


