# EdgeConnectorHTTPTypedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Modules** | [**EdgeConnectorModulesRequest**](EdgeConnectorModulesRequest.md) |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | 
**Addresses** | Pointer to [**[]AddressRequest**](AddressRequest.md) |  | [optional] 
**Tls** | Pointer to [**TLSRequest**](TLSRequest.md) |  | [optional] 
**LoadBalanceMethod** | Pointer to **string** | * &#x60;off&#x60; - Off * &#x60;ip_hash&#x60; - IP Hash * &#x60;least_connections&#x60; - Least Connections * &#x60;round_robin&#x60; - Round Robin | [optional] 
**ConnectionPreference** | Pointer to **[]string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**ReadWriteTimeout** | Pointer to **int64** |  | [optional] 
**MaxRetries** | Pointer to **int64** |  | [optional] 
**TypeProperties** | [**EdgeConnectorHTTPTypePropertiesRequest**](EdgeConnectorHTTPTypePropertiesRequest.md) |  | 

## Methods

### NewEdgeConnectorHTTPTypedRequest

`func NewEdgeConnectorHTTPTypedRequest(name string, modules EdgeConnectorModulesRequest, type_ string, typeProperties EdgeConnectorHTTPTypePropertiesRequest, ) *EdgeConnectorHTTPTypedRequest`

NewEdgeConnectorHTTPTypedRequest instantiates a new EdgeConnectorHTTPTypedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorHTTPTypedRequestWithDefaults

`func NewEdgeConnectorHTTPTypedRequestWithDefaults() *EdgeConnectorHTTPTypedRequest`

NewEdgeConnectorHTTPTypedRequestWithDefaults instantiates a new EdgeConnectorHTTPTypedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EdgeConnectorHTTPTypedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeConnectorHTTPTypedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeConnectorHTTPTypedRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModules

`func (o *EdgeConnectorHTTPTypedRequest) GetModules() EdgeConnectorModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorHTTPTypedRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorHTTPTypedRequest) SetModules(v EdgeConnectorModulesRequest)`

SetModules sets Modules field to given value.


### GetActive

`func (o *EdgeConnectorHTTPTypedRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *EdgeConnectorHTTPTypedRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *EdgeConnectorHTTPTypedRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *EdgeConnectorHTTPTypedRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetType

`func (o *EdgeConnectorHTTPTypedRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeConnectorHTTPTypedRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeConnectorHTTPTypedRequest) SetType(v string)`

SetType sets Type field to given value.


### GetAddresses

`func (o *EdgeConnectorHTTPTypedRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorHTTPTypedRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorHTTPTypedRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *EdgeConnectorHTTPTypedRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetTls

`func (o *EdgeConnectorHTTPTypedRequest) GetTls() TLSRequest`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *EdgeConnectorHTTPTypedRequest) GetTlsOk() (*TLSRequest, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *EdgeConnectorHTTPTypedRequest) SetTls(v TLSRequest)`

SetTls sets Tls field to given value.

### HasTls

`func (o *EdgeConnectorHTTPTypedRequest) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetLoadBalanceMethod

`func (o *EdgeConnectorHTTPTypedRequest) GetLoadBalanceMethod() string`

GetLoadBalanceMethod returns the LoadBalanceMethod field if non-nil, zero value otherwise.

### GetLoadBalanceMethodOk

`func (o *EdgeConnectorHTTPTypedRequest) GetLoadBalanceMethodOk() (*string, bool)`

GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMethod

`func (o *EdgeConnectorHTTPTypedRequest) SetLoadBalanceMethod(v string)`

SetLoadBalanceMethod sets LoadBalanceMethod field to given value.

### HasLoadBalanceMethod

`func (o *EdgeConnectorHTTPTypedRequest) HasLoadBalanceMethod() bool`

HasLoadBalanceMethod returns a boolean if a field has been set.

### GetConnectionPreference

`func (o *EdgeConnectorHTTPTypedRequest) GetConnectionPreference() []string`

GetConnectionPreference returns the ConnectionPreference field if non-nil, zero value otherwise.

### GetConnectionPreferenceOk

`func (o *EdgeConnectorHTTPTypedRequest) GetConnectionPreferenceOk() (*[]string, bool)`

GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreference

`func (o *EdgeConnectorHTTPTypedRequest) SetConnectionPreference(v []string)`

SetConnectionPreference sets ConnectionPreference field to given value.

### HasConnectionPreference

`func (o *EdgeConnectorHTTPTypedRequest) HasConnectionPreference() bool`

HasConnectionPreference returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *EdgeConnectorHTTPTypedRequest) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *EdgeConnectorHTTPTypedRequest) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *EdgeConnectorHTTPTypedRequest) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *EdgeConnectorHTTPTypedRequest) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetReadWriteTimeout

`func (o *EdgeConnectorHTTPTypedRequest) GetReadWriteTimeout() int64`

GetReadWriteTimeout returns the ReadWriteTimeout field if non-nil, zero value otherwise.

### GetReadWriteTimeoutOk

`func (o *EdgeConnectorHTTPTypedRequest) GetReadWriteTimeoutOk() (*int64, bool)`

GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteTimeout

`func (o *EdgeConnectorHTTPTypedRequest) SetReadWriteTimeout(v int64)`

SetReadWriteTimeout sets ReadWriteTimeout field to given value.

### HasReadWriteTimeout

`func (o *EdgeConnectorHTTPTypedRequest) HasReadWriteTimeout() bool`

HasReadWriteTimeout returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EdgeConnectorHTTPTypedRequest) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EdgeConnectorHTTPTypedRequest) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EdgeConnectorHTTPTypedRequest) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EdgeConnectorHTTPTypedRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetTypeProperties

`func (o *EdgeConnectorHTTPTypedRequest) GetTypeProperties() EdgeConnectorHTTPTypePropertiesRequest`

GetTypeProperties returns the TypeProperties field if non-nil, zero value otherwise.

### GetTypePropertiesOk

`func (o *EdgeConnectorHTTPTypedRequest) GetTypePropertiesOk() (*EdgeConnectorHTTPTypePropertiesRequest, bool)`

GetTypePropertiesOk returns a tuple with the TypeProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeProperties

`func (o *EdgeConnectorHTTPTypedRequest) SetTypeProperties(v EdgeConnectorHTTPTypePropertiesRequest)`

SetTypeProperties sets TypeProperties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


