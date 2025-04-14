/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressRequest{}

// AddressRequest struct for AddressRequest
type AddressRequest struct {
	// IPv4/IPv6 address or CNAME to resolve
	Address string `json:"address" validate:"regexp=.*"`
	PlainPort *int64 `json:"plain_port,omitempty"`
	TlsPort *int64 `json:"tls_port,omitempty"`
	// Role of the address in load balancing  * `primary` - Primary * `backup` - Backup
	ServerRole *string `json:"server_role,omitempty"`
	// Weight used in load balancing strategy
	Weight *int64 `json:"weight,omitempty"`
	// Indicates if the address is active for use
	Active *bool `json:"active,omitempty"`
	// Maximum number of open connections per Edge Application instance
	MaxConns *int64 `json:"max_conns,omitempty"`
	// Maximum number of communication attempts before marking as unavailable
	MaxFails *int64 `json:"max_fails,omitempty"`
	// Timeout for communication attempts
	FailTimeout *int64 `json:"fail_timeout,omitempty"`
}

type _AddressRequest AddressRequest

// NewAddressRequest instantiates a new AddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRequest(address string) *AddressRequest {
	this := AddressRequest{}
	this.Address = address
	return &this
}

// NewAddressRequestWithDefaults instantiates a new AddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRequestWithDefaults() *AddressRequest {
	this := AddressRequest{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddressRequest) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddressRequest) SetAddress(v string) {
	o.Address = v
}

// GetPlainPort returns the PlainPort field value if set, zero value otherwise.
func (o *AddressRequest) GetPlainPort() int64 {
	if o == nil || IsNil(o.PlainPort) {
		var ret int64
		return ret
	}
	return *o.PlainPort
}

// GetPlainPortOk returns a tuple with the PlainPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetPlainPortOk() (*int64, bool) {
	if o == nil || IsNil(o.PlainPort) {
		return nil, false
	}
	return o.PlainPort, true
}

// HasPlainPort returns a boolean if a field has been set.
func (o *AddressRequest) HasPlainPort() bool {
	if o != nil && !IsNil(o.PlainPort) {
		return true
	}

	return false
}

// SetPlainPort gets a reference to the given int64 and assigns it to the PlainPort field.
func (o *AddressRequest) SetPlainPort(v int64) {
	o.PlainPort = &v
}

// GetTlsPort returns the TlsPort field value if set, zero value otherwise.
func (o *AddressRequest) GetTlsPort() int64 {
	if o == nil || IsNil(o.TlsPort) {
		var ret int64
		return ret
	}
	return *o.TlsPort
}

// GetTlsPortOk returns a tuple with the TlsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetTlsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TlsPort) {
		return nil, false
	}
	return o.TlsPort, true
}

// HasTlsPort returns a boolean if a field has been set.
func (o *AddressRequest) HasTlsPort() bool {
	if o != nil && !IsNil(o.TlsPort) {
		return true
	}

	return false
}

// SetTlsPort gets a reference to the given int64 and assigns it to the TlsPort field.
func (o *AddressRequest) SetTlsPort(v int64) {
	o.TlsPort = &v
}

// GetServerRole returns the ServerRole field value if set, zero value otherwise.
func (o *AddressRequest) GetServerRole() string {
	if o == nil || IsNil(o.ServerRole) {
		var ret string
		return ret
	}
	return *o.ServerRole
}

// GetServerRoleOk returns a tuple with the ServerRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetServerRoleOk() (*string, bool) {
	if o == nil || IsNil(o.ServerRole) {
		return nil, false
	}
	return o.ServerRole, true
}

// HasServerRole returns a boolean if a field has been set.
func (o *AddressRequest) HasServerRole() bool {
	if o != nil && !IsNil(o.ServerRole) {
		return true
	}

	return false
}

// SetServerRole gets a reference to the given string and assigns it to the ServerRole field.
func (o *AddressRequest) SetServerRole(v string) {
	o.ServerRole = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *AddressRequest) GetWeight() int64 {
	if o == nil || IsNil(o.Weight) {
		var ret int64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetWeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *AddressRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int64 and assigns it to the Weight field.
func (o *AddressRequest) SetWeight(v int64) {
	o.Weight = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AddressRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AddressRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AddressRequest) SetActive(v bool) {
	o.Active = &v
}

// GetMaxConns returns the MaxConns field value if set, zero value otherwise.
func (o *AddressRequest) GetMaxConns() int64 {
	if o == nil || IsNil(o.MaxConns) {
		var ret int64
		return ret
	}
	return *o.MaxConns
}

// GetMaxConnsOk returns a tuple with the MaxConns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetMaxConnsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxConns) {
		return nil, false
	}
	return o.MaxConns, true
}

// HasMaxConns returns a boolean if a field has been set.
func (o *AddressRequest) HasMaxConns() bool {
	if o != nil && !IsNil(o.MaxConns) {
		return true
	}

	return false
}

// SetMaxConns gets a reference to the given int64 and assigns it to the MaxConns field.
func (o *AddressRequest) SetMaxConns(v int64) {
	o.MaxConns = &v
}

// GetMaxFails returns the MaxFails field value if set, zero value otherwise.
func (o *AddressRequest) GetMaxFails() int64 {
	if o == nil || IsNil(o.MaxFails) {
		var ret int64
		return ret
	}
	return *o.MaxFails
}

// GetMaxFailsOk returns a tuple with the MaxFails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetMaxFailsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxFails) {
		return nil, false
	}
	return o.MaxFails, true
}

// HasMaxFails returns a boolean if a field has been set.
func (o *AddressRequest) HasMaxFails() bool {
	if o != nil && !IsNil(o.MaxFails) {
		return true
	}

	return false
}

// SetMaxFails gets a reference to the given int64 and assigns it to the MaxFails field.
func (o *AddressRequest) SetMaxFails(v int64) {
	o.MaxFails = &v
}

// GetFailTimeout returns the FailTimeout field value if set, zero value otherwise.
func (o *AddressRequest) GetFailTimeout() int64 {
	if o == nil || IsNil(o.FailTimeout) {
		var ret int64
		return ret
	}
	return *o.FailTimeout
}

// GetFailTimeoutOk returns a tuple with the FailTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetFailTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.FailTimeout) {
		return nil, false
	}
	return o.FailTimeout, true
}

// HasFailTimeout returns a boolean if a field has been set.
func (o *AddressRequest) HasFailTimeout() bool {
	if o != nil && !IsNil(o.FailTimeout) {
		return true
	}

	return false
}

// SetFailTimeout gets a reference to the given int64 and assigns it to the FailTimeout field.
func (o *AddressRequest) SetFailTimeout(v int64) {
	o.FailTimeout = &v
}

func (o AddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.PlainPort) {
		toSerialize["plain_port"] = o.PlainPort
	}
	if !IsNil(o.TlsPort) {
		toSerialize["tls_port"] = o.TlsPort
	}
	if !IsNil(o.ServerRole) {
		toSerialize["server_role"] = o.ServerRole
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.MaxConns) {
		toSerialize["max_conns"] = o.MaxConns
	}
	if !IsNil(o.MaxFails) {
		toSerialize["max_fails"] = o.MaxFails
	}
	if !IsNil(o.FailTimeout) {
		toSerialize["fail_timeout"] = o.FailTimeout
	}
	return toSerialize, nil
}

func (o *AddressRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAddressRequest := _AddressRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressRequest)

	if err != nil {
		return err
	}

	*o = AddressRequest(varAddressRequest)

	return err
}

type NullableAddressRequest struct {
	value *AddressRequest
	isSet bool
}

func (v NullableAddressRequest) Get() *AddressRequest {
	return v.value
}

func (v *NullableAddressRequest) Set(val *AddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRequest(val *AddressRequest) *NullableAddressRequest {
	return &NullableAddressRequest{value: val, isSet: true}
}

func (v NullableAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


