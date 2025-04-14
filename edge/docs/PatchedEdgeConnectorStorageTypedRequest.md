# PatchedEdgeConnectorStorageTypedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Modules** | Pointer to [**EdgeConnectorModulesRequest**](EdgeConnectorModulesRequest.md) |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Tls** | Pointer to [**TLSRequest**](TLSRequest.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | Pointer to [**EdgeConnectorStorageTypePropertiesRequest**](EdgeConnectorStorageTypePropertiesRequest.md) |  | [optional] 

## Methods

### NewPatchedEdgeConnectorStorageTypedRequest

`func NewPatchedEdgeConnectorStorageTypedRequest() *PatchedEdgeConnectorStorageTypedRequest`

NewPatchedEdgeConnectorStorageTypedRequest instantiates a new PatchedEdgeConnectorStorageTypedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeConnectorStorageTypedRequestWithDefaults

`func NewPatchedEdgeConnectorStorageTypedRequestWithDefaults() *PatchedEdgeConnectorStorageTypedRequest`

NewPatchedEdgeConnectorStorageTypedRequestWithDefaults instantiates a new PatchedEdgeConnectorStorageTypedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddresses

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetTls() TLSRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetTlsOk() (*TLSRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetTls(v TLSRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetTypeProperties() EdgeConnectorStorageTypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *PatchedEdgeConnectorStorageTypedRequest) GetTypePropertiesOk() (*EdgeConnectorStorageTypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *PatchedEdgeConnectorStorageTypedRequest) SetTypeProperties(v EdgeConnectorStorageTypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.

### HasTypeProperties

`func (o *PatchedEdgeConnectorStorageTypedRequest) HasTypeProperties() bool`

HasTypeProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


