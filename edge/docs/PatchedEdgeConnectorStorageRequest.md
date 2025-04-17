# PatchedEdgeConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to [**EdgeConnectorModulesRequest**](EdgeConnectorModulesRequest.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** | * &#x60;http&#x60; - HTTP * &#x60;s3&#x60; - S3 * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | [optional] 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Tls** | Pointer to [**TLSEdgeConnectorRequest**](TLSEdgeConnectorRequest.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | Pointer to [**EdgeConnectorStorageTypePropertiesRequest**](EdgeConnectorStorageTypePropertiesRequest.md) |  | [optional] 

## Methods

### NewPatchedEdgeConnectorStorageRequest

`func NewPatchedEdgeConnectorStorageRequest() *PatchedEdgeConnectorStorageRequest`

NewPatchedEdgeConnectorStorageRequest instantiates a new PatchedEdgeConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeConnectorStorageRequestWithDefaults

`func NewPatchedEdgeConnectorStorageRequestWithDefaults() *PatchedEdgeConnectorStorageRequest`

NewPatchedEdgeConnectorStorageRequestWithDefaults instantiates a new PatchedEdgeConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeConnectorStorageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchedEdgeConnectorStorageRequest) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchedEdgeConnectorStorageRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchedEdgeConnectorStorageRequest) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchedEdgeConnectorStorageRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedEdgeConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedEdgeConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedEdgeConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedEdgeConnectorStorageRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddresses

`func (o *PatchedEdgeConnectorStorageRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchedEdgeConnectorStorageRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchedEdgeConnectorStorageRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchedEdgeConnectorStorageRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *PatchedEdgeConnectorStorageRequest) GetTls() TLSEdgeConnectorRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *PatchedEdgeConnectorStorageRequest) GetTlsOk() (*TLSEdgeConnectorRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *PatchedEdgeConnectorStorageRequest) SetTls(v TLSEdgeConnectorRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *PatchedEdgeConnectorStorageRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageRequest) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *PatchedEdgeConnectorStorageRequest) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageRequest) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageRequest) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *PatchedEdgeConnectorStorageRequest) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *PatchedEdgeConnectorStorageRequest) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *PatchedEdgeConnectorStorageRequest) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *PatchedEdgeConnectorStorageRequest) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *PatchedEdgeConnectorStorageRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *PatchedEdgeConnectorStorageRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *PatchedEdgeConnectorStorageRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *PatchedEdgeConnectorStorageRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *PatchedEdgeConnectorStorageRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *PatchedEdgeConnectorStorageRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatchedEdgeConnectorStorageRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatchedEdgeConnectorStorageRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatchedEdgeConnectorStorageRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *PatchedEdgeConnectorStorageRequest) GetTypeProperties() EdgeConnectorStorageTypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *PatchedEdgeConnectorStorageRequest) GetTypePropertiesOk() (*EdgeConnectorStorageTypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *PatchedEdgeConnectorStorageRequest) SetTypeProperties(v EdgeConnectorStorageTypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.

### HasTypeProperties

`func (o *PatchedEdgeConnectorStorageRequest) HasTypeProperties() bool`

HasTypeProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


