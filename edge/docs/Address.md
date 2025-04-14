# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPv4/IPv6 address or CNAME to resolve | 
**PlainPort** | Pointer to **int64** |  | [optional] 
**TlsPort** | Pointer to **int64** |  | [optional] 
**ServerRole** | Pointer to **string** | Role of the address in load balancing  * &#x60;primary&#x60; - Primary * &#x60;backup&#x60; - Backup | [optional] 
**Weight** | Pointer to **int64** | Weight used in load balancing strategy | [optional] 
**Active** | Pointer to **bool** | Indicates if the address is active for use | [optional] 
**MaxConns** | Pointer to **int64** | Maximum number of open connections per Edge Application instance | [optional] 
**MaxFails** | Pointer to **int64** | Maximum number of communication attempts before marking as unavailable | [optional] 
**FailTimeout** | Pointer to **int64** | Timeout for communication attempts | [optional] 

## Methods

### NewAddress

`func NewAddress(address string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPlainPort

`func (o *Address) GetPlainPort() int64`

GetPlainPort returns the PlainPort field if non-nil, zero value otherwise.

### GetPlainPortOk

`func (o *Address) GetPlainPortOk() (*int64, bool)`

GetPlainPortOk returns a tuple with the PlainPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlainPort

`func (o *Address) SetPlainPort(v int64)`

SetPlainPort sets PlainPort field to given value.

### HasPlainPort

`func (o *Address) HasPlainPort() bool`

HasPlainPort returns a boolean if a field has been set.

### GetTlsPort

`func (o *Address) GetTlsPort() int64`

GetTlsPort returns the TlsPort field if non-nil, zero value otherwise.

### GetTlsPortOk

`func (o *Address) GetTlsPortOk() (*int64, bool)`

GetTlsPortOk returns a tuple with the TlsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPort

`func (o *Address) SetTlsPort(v int64)`

SetTlsPort sets TlsPort field to given value.

### HasTlsPort

`func (o *Address) HasTlsPort() bool`

HasTlsPort returns a boolean if a field has been set.

### GetServerRole

`func (o *Address) GetServerRole() string`

GetServerRole returns the ServerRole field if non-nil, zero value otherwise.

### GetServerRoleOk

`func (o *Address) GetServerRoleOk() (*string, bool)`

GetServerRoleOk returns a tuple with the ServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRole

`func (o *Address) SetServerRole(v string)`

SetServerRole sets ServerRole field to given value.

### HasServerRole

`func (o *Address) HasServerRole() bool`

HasServerRole returns a boolean if a field has been set.

### GetWeight

`func (o *Address) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Address) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Address) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Address) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetActive

`func (o *Address) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Address) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Address) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Address) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMaxConns

`func (o *Address) GetMaxConns() int64`

GetMaxConns returns the MaxConns field if non-nil, zero value otherwise.

### GetMaxConnsOk

`func (o *Address) GetMaxConnsOk() (*int64, bool)`

GetMaxConnsOk returns a tuple with the MaxConns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConns

`func (o *Address) SetMaxConns(v int64)`

SetMaxConns sets MaxConns field to given value.

### HasMaxConns

`func (o *Address) HasMaxConns() bool`

HasMaxConns returns a boolean if a field has been set.

### GetMaxFails

`func (o *Address) GetMaxFails() int64`

GetMaxFails returns the MaxFails field if non-nil, zero value otherwise.

### GetMaxFailsOk

`func (o *Address) GetMaxFailsOk() (*int64, bool)`

GetMaxFailsOk returns a tuple with the MaxFails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFails

`func (o *Address) SetMaxFails(v int64)`

SetMaxFails sets MaxFails field to given value.

### HasMaxFails

`func (o *Address) HasMaxFails() bool`

HasMaxFails returns a boolean if a field has been set.

### GetFailTimeout

`func (o *Address) GetFailTimeout() int64`

GetFailTimeout returns the FailTimeout field if non-nil, zero value otherwise.

### GetFailTimeoutOk

`func (o *Address) GetFailTimeoutOk() (*int64, bool)`

GetFailTimeoutOk returns a tuple with the FailTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailTimeout

`func (o *Address) SetFailTimeout(v int64)`

SetFailTimeout sets FailTimeout field to given value.

### HasFailTimeout

`func (o *Address) HasFailTimeout() bool`

HasFailTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


