# EdgeConnectorStorageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Modules** | [**EdgeConnectorModulesRequest**](EdgeConnectorModulesRequest.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** | * &#x60;http&#x60; - HTTP * &#x60;s3&#x60; - S3 * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Tls** | Pointer to [**TLSRequest**](TLSRequest.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | [**EdgeConnectorStorageTypePropertiesRequest**](EdgeConnectorStorageTypePropertiesRequest.md) |  | 

## Methods

### NewEdgeConnectorStorageRequest

`func NewEdgeConnectorStorageRequest(name string, modules EdgeConnectorModulesRequest, type_ string, typeProperties EdgeConnectorStorageTypePropertiesRequest, ) *EdgeConnectorStorageRequest`

NewEdgeConnectorStorageRequest instantiates a new EdgeConnectorStorageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorStorageRequestWithDefaults

`func NewEdgeConnectorStorageRequestWithDefaults() *EdgeConnectorStorageRequest`

NewEdgeConnectorStorageRequestWithDefaults instantiates a new EdgeConnectorStorageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeConnectorStorageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorStorageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorStorageRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *EdgeConnectorStorageRequest) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorStorageRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorStorageRequest) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.


### GetActive

`func (o *EdgeConnectorStorageRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorStorageRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorStorageRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorStorageRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *EdgeConnectorStorageRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorStorageRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorStorageRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAddresses

`func (o *EdgeConnectorStorageRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorStorageRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorStorageRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *EdgeConnectorStorageRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *EdgeConnectorStorageRequest) GetTls() TLSRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *EdgeConnectorStorageRequest) GetTlsOk() (*TLSRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *EdgeConnectorStorageRequest) SetTls(v TLSRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *EdgeConnectorStorageRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *EdgeConnectorStorageRequest) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *EdgeConnectorStorageRequest) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *EdgeConnectorStorageRequest) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *EdgeConnectorStorageRequest) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *EdgeConnectorStorageRequest) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *EdgeConnectorStorageRequest) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *EdgeConnectorStorageRequest) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *EdgeConnectorStorageRequest) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EdgeConnectorStorageRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EdgeConnectorStorageRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EdgeConnectorStorageRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EdgeConnectorStorageRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *EdgeConnectorStorageRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *EdgeConnectorStorageRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *EdgeConnectorStorageRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *EdgeConnectorStorageRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EdgeConnectorStorageRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EdgeConnectorStorageRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EdgeConnectorStorageRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EdgeConnectorStorageRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *EdgeConnectorStorageRequest) GetTypeProperties() EdgeConnectorStorageTypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *EdgeConnectorStorageRequest) GetTypePropertiesOk() (*EdgeConnectorStorageTypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *EdgeConnectorStorageRequest) SetTypeProperties(v EdgeConnectorStorageTypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


