# EdgeConnectorPolymorphicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | * &#x60;http&#x60; - HTTP * &#x60;s3&#x60; - S3 * &#x60;edge_storage&#x60; - Edge Storage * &#x60;live_ingest&#x60; - Live Ingest | 
**Name** | **string** |  | 
**Modules** | [**EdgeConnectorModulesRequest**](EdgeConnectorModulesRequest.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Tls** | Pointer to [**TLSEdgeConnectorRequest**](TLSEdgeConnectorRequest.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | [**EdgeConnectorLiveIngestTypePropertiesRequest**](EdgeConnectorLiveIngestTypePropertiesRequest.md) |  | 

## Methods

### NewEdgeConnectorPolymorphicRequest

`func NewEdgeConnectorPolymorphicRequest(type_ string, name string, modules EdgeConnectorModulesRequest, typeProperties EdgeConnectorLiveIngestTypePropertiesRequest, ) *EdgeConnectorPolymorphicRequest`

NewEdgeConnectorPolymorphicRequest instantiates a new EdgeConnectorPolymorphicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorPolymorphicRequestWithDefaults

`func NewEdgeConnectorPolymorphicRequestWithDefaults() *EdgeConnectorPolymorphicRequest`

NewEdgeConnectorPolymorphicRequestWithDefaults instantiates a new EdgeConnectorPolymorphicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EdgeConnectorPolymorphicRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorPolymorphicRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorPolymorphicRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *EdgeConnectorPolymorphicRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorPolymorphicRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorPolymorphicRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *EdgeConnectorPolymorphicRequest) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorPolymorphicRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorPolymorphicRequest) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.


### GetActive

`func (o *EdgeConnectorPolymorphicRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorPolymorphicRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorPolymorphicRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorPolymorphicRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddresses

`func (o *EdgeConnectorPolymorphicRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorPolymorphicRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorPolymorphicRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *EdgeConnectorPolymorphicRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *EdgeConnectorPolymorphicRequest) GetTls() TLSEdgeConnectorRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *EdgeConnectorPolymorphicRequest) GetTlsOk() (*TLSEdgeConnectorRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *EdgeConnectorPolymorphicRequest) SetTls(v TLSEdgeConnectorRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *EdgeConnectorPolymorphicRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *EdgeConnectorPolymorphicRequest) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *EdgeConnectorPolymorphicRequest) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *EdgeConnectorPolymorphicRequest) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *EdgeConnectorPolymorphicRequest) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *EdgeConnectorPolymorphicRequest) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *EdgeConnectorPolymorphicRequest) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *EdgeConnectorPolymorphicRequest) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *EdgeConnectorPolymorphicRequest) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EdgeConnectorPolymorphicRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EdgeConnectorPolymorphicRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EdgeConnectorPolymorphicRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EdgeConnectorPolymorphicRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *EdgeConnectorPolymorphicRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *EdgeConnectorPolymorphicRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *EdgeConnectorPolymorphicRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *EdgeConnectorPolymorphicRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EdgeConnectorPolymorphicRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EdgeConnectorPolymorphicRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EdgeConnectorPolymorphicRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EdgeConnectorPolymorphicRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *EdgeConnectorPolymorphicRequest) GetTypeProperties() EdgeConnectorLiveIngestTypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *EdgeConnectorPolymorphicRequest) GetTypePropertiesOk() (*EdgeConnectorLiveIngestTypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *EdgeConnectorPolymorphicRequest) SetTypeProperties(v EdgeConnectorLiveIngestTypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


