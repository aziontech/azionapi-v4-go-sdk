# EdgeConnectorHTTPAttributesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]AddressRequest**](AddressRequest.md) |  | 
**ConnectionOptions** | Pointer to [**HTTPConnectionOptionsRequest**](HTTPConnectionOptionsRequest.md) |  | [optional] 
**Modules** | Pointer to [**HTTPModulesRequest**](HTTPModulesRequest.md) |  | [optional] 

## Methods

### NewEdgeConnectorHTTPAttributesRequest

`func NewEdgeConnectorHTTPAttributesRequest(addresses []AddressRequest, ) *EdgeConnectorHTTPAttributesRequest`

NewEdgeConnectorHTTPAttributesRequest instantiates a new EdgeConnectorHTTPAttributesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorHTTPAttributesRequestWithDefaults

`func NewEdgeConnectorHTTPAttributesRequestWithDefaults() *EdgeConnectorHTTPAttributesRequest`

NewEdgeConnectorHTTPAttributesRequestWithDefaults instantiates a new EdgeConnectorHTTPAttributesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *EdgeConnectorHTTPAttributesRequest) GetAddresses() []AddressRequest`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorHTTPAttributesRequest) GetAddressesOk() (*[]AddressRequest, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorHTTPAttributesRequest) SetAddresses(v []AddressRequest)`

SetAddresses sets Addresses field to given value.


### GetConnectionOptions

`func (o *EdgeConnectorHTTPAttributesRequest) GetConnectionOptions() HTTPConnectionOptionsRequest`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *EdgeConnectorHTTPAttributesRequest) GetConnectionOptionsOk() (*HTTPConnectionOptionsRequest, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *EdgeConnectorHTTPAttributesRequest) SetConnectionOptions(v HTTPConnectionOptionsRequest)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *EdgeConnectorHTTPAttributesRequest) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### GetModules

`func (o *EdgeConnectorHTTPAttributesRequest) GetModules() HTTPModulesRequest`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorHTTPAttributesRequest) GetModulesOk() (*HTTPModulesRequest, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorHTTPAttributesRequest) SetModules(v HTTPModulesRequest)`

SetModules sets Modules field to given value.

### HasModules

`func (o *EdgeConnectorHTTPAttributesRequest) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


