# PatchedEdgeConnectorS3Request

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
**TypeProperties** | Pointer to [**EdgeConnectorS3TypePropertiesRequest**](EdgeConnectorS3TypePropertiesRequest.md) |  | [optional] 

## Methods

### NewPatchedEdgeConnectorS3Request

`func NewPatchedEdgeConnectorS3Request() *PatchedEdgeConnectorS3Request`

NewPatchedEdgeConnectorS3Request instantiates a new PatchedEdgeConnectorS3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEdgeConnectorS3RequestWithDefaults

`func NewPatchedEdgeConnectorS3RequestWithDefaults() *PatchedEdgeConnectorS3Request`

NewPatchedEdgeConnectorS3RequestWithDefaults instantiates a new PatchedEdgeConnectorS3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedEdgeConnectorS3Request) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedEdgeConnectorS3Request) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedEdgeConnectorS3Request) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedEdgeConnectorS3Request) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModules

`func (o *PatchedEdgeConnectorS3Request) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *PatchedEdgeConnectorS3Request) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *PatchedEdgeConnectorS3Request) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *PatchedEdgeConnectorS3Request) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetActive

`func (o *PatchedEdgeConnectorS3Request) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PatchedEdgeConnectorS3Request) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PatchedEdgeConnectorS3Request) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PatchedEdgeConnectorS3Request) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *PatchedEdgeConnectorS3Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedEdgeConnectorS3Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedEdgeConnectorS3Request) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedEdgeConnectorS3Request) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddresses

`func (o *PatchedEdgeConnectorS3Request) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchedEdgeConnectorS3Request) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchedEdgeConnectorS3Request) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchedEdgeConnectorS3Request) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *PatchedEdgeConnectorS3Request) GetTls() TLSEdgeConnectorRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *PatchedEdgeConnectorS3Request) GetTlsOk() (*TLSEdgeConnectorRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *PatchedEdgeConnectorS3Request) SetTls(v TLSEdgeConnectorRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *PatchedEdgeConnectorS3Request) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *PatchedEdgeConnectorS3Request) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *PatchedEdgeConnectorS3Request) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *PatchedEdgeConnectorS3Request) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *PatchedEdgeConnectorS3Request) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *PatchedEdgeConnectorS3Request) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *PatchedEdgeConnectorS3Request) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *PatchedEdgeConnectorS3Request) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *PatchedEdgeConnectorS3Request) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *PatchedEdgeConnectorS3Request) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *PatchedEdgeConnectorS3Request) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *PatchedEdgeConnectorS3Request) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *PatchedEdgeConnectorS3Request) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *PatchedEdgeConnectorS3Request) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *PatchedEdgeConnectorS3Request) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *PatchedEdgeConnectorS3Request) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *PatchedEdgeConnectorS3Request) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *PatchedEdgeConnectorS3Request) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PatchedEdgeConnectorS3Request) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PatchedEdgeConnectorS3Request) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PatchedEdgeConnectorS3Request) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *PatchedEdgeConnectorS3Request) GetTypeProperties() EdgeConnectorS3TypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *PatchedEdgeConnectorS3Request) GetTypePropertiesOk() (*EdgeConnectorS3TypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *PatchedEdgeConnectorS3Request) SetTypeProperties(v EdgeConnectorS3TypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.

### HasTypeProperties

`func (o *PatchedEdgeConnectorS3Request) HasTypeProperties() bool`

HasTypeProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


