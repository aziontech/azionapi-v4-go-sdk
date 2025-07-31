# EdgeConnectorHTTPAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]Address**](Address.md) |  | 
**ConnectionOptions** | Pointer to [**HTTPConnectionOptions**](HTTPConnectionOptions.md) |  | [optional] 
**Modules** | Pointer to [**HTTPModules**](HTTPModules.md) |  | [optional] 

## Methods

### NewEdgeConnectorHTTPAttributes

`func NewEdgeConnectorHTTPAttributes(addresses []Address, ) *EdgeConnectorHTTPAttributes`

NewEdgeConnectorHTTPAttributes instantiates a new EdgeConnectorHTTPAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeConnectorHTTPAttributesWithDefaults

`func NewEdgeConnectorHTTPAttributesWithDefaults() *EdgeConnectorHTTPAttributes`

NewEdgeConnectorHTTPAttributesWithDefaults instantiates a new EdgeConnectorHTTPAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *EdgeConnectorHTTPAttributes) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *EdgeConnectorHTTPAttributes) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *EdgeConnectorHTTPAttributes) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.


### GetConnectionOptions

`func (o *EdgeConnectorHTTPAttributes) GetConnectionOptions() HTTPConnectionOptions`

GetConnectionOptions returns the ConnectionOptions field if non-nil, zero value otherwise.

### GetConnectionOptionsOk

`func (o *EdgeConnectorHTTPAttributes) GetConnectionOptionsOk() (*HTTPConnectionOptions, bool)`

GetConnectionOptionsOk returns a tuple with the ConnectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionOptions

`func (o *EdgeConnectorHTTPAttributes) SetConnectionOptions(v HTTPConnectionOptions)`

SetConnectionOptions sets ConnectionOptions field to given value.

### HasConnectionOptions

`func (o *EdgeConnectorHTTPAttributes) HasConnectionOptions() bool`

HasConnectionOptions returns a boolean if a field has been set.

### GetModules

`func (o *EdgeConnectorHTTPAttributes) GetModules() HTTPModules`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *EdgeConnectorHTTPAttributes) GetModulesOk() (*HTTPModules, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *EdgeConnectorHTTPAttributes) SetModules(v HTTPModules)`

SetModules sets Modules field to given value.

### HasModules

`func (o *EdgeConnectorHTTPAttributes) HasModules() bool`

HasModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


