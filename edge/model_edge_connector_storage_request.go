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

// checks if the EdgeConnectorStorageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeConnectorStorageRequest{}

// EdgeConnectorStorageRequest struct for EdgeConnectorStorageRequest
type EdgeConnectorStorageRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	Modules EdgeConnectorModulesRequest `json:"modules"`
	Active *bool `json:"active,omitempty"`
	// * `http` - HTTP * `s3` - S3 * `edge_storage` - Edge Storage * `live_ingest` - Live Ingest
	Type string `json:"type"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	Tls *TLSEdgeConnectorRequest `json:"tls,omitempty"`
	// * `off` - Off * `ip_hash` - IP Hash * `least_connections` - Least Connections * `round_robin` - Round Robin
	LoadBalanceMethod *string `json:"load_balance_method,omitempty"`
	ConnectionPreference []string `json:"connection_preference,omitempty"`
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	ReadWriteTimeout *int64 `json:"read_write_timeout,omitempty"`
	MaxRetries *int64 `json:"max_retries,omitempty"`
	TypeProperties EdgeConnectorStorageTypePropertiesRequest `json:"type_properties"`
}

type _EdgeConnectorStorageRequest EdgeConnectorStorageRequest

// NewEdgeConnectorStorageRequest instantiates a new EdgeConnectorStorageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeConnectorStorageRequest(name string, modules EdgeConnectorModulesRequest, type_ string, typeProperties EdgeConnectorStorageTypePropertiesRequest) *EdgeConnectorStorageRequest {
	this := EdgeConnectorStorageRequest{}
	this.Name = name
	this.Modules = modules
	this.Type = type_
	this.TypeProperties = typeProperties
	return &this
}

// NewEdgeConnectorStorageRequestWithDefaults instantiates a new EdgeConnectorStorageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeConnectorStorageRequestWithDefaults() *EdgeConnectorStorageRequest {
	this := EdgeConnectorStorageRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EdgeConnectorStorageRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeConnectorStorageRequest) SetName(v string) {
	o.Name = v
}

// GetModules returns the Modules field value
func (o *EdgeConnectorStorageRequest) GetModules() EdgeConnectorModulesRequest {
	if o == nil {
		var ret EdgeConnectorModulesRequest
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modules, true
}

// SetModules sets field value
func (o *EdgeConnectorStorageRequest) SetModules(v EdgeConnectorModulesRequest) {
	o.Modules = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeConnectorStorageRequest) SetActive(v bool) {
	o.Active = &v
}

// GetType returns the Type field value
func (o *EdgeConnectorStorageRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EdgeConnectorStorageRequest) SetType(v string) {
	o.Type = v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetAddresses() []AddressRequest {
	if o == nil || IsNil(o.Addresses) {
		var ret []AddressRequest
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetAddressesOk() ([]AddressRequest, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *EdgeConnectorStorageRequest) SetAddresses(v []AddressRequest) {
	o.Addresses = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetTls() TLSEdgeConnectorRequest {
	if o == nil || IsNil(o.Tls) {
		var ret TLSEdgeConnectorRequest
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetTlsOk() (*TLSEdgeConnectorRequest, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given TLSEdgeConnectorRequest and assigns it to the Tls field.
func (o *EdgeConnectorStorageRequest) SetTls(v TLSEdgeConnectorRequest) {
	o.Tls = &v
}

// GetLoadBalanceMethod returns the LoadBalanceMethod field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetLoadBalanceMethod() string {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		var ret string
		return ret
	}
	return *o.LoadBalanceMethod
}

// GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetLoadBalanceMethodOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		return nil, false
	}
	return o.LoadBalanceMethod, true
}

// HasLoadBalanceMethod returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasLoadBalanceMethod() bool {
	if o != nil && !IsNil(o.LoadBalanceMethod) {
		return true
	}

	return false
}

// SetLoadBalanceMethod gets a reference to the given string and assigns it to the LoadBalanceMethod field.
func (o *EdgeConnectorStorageRequest) SetLoadBalanceMethod(v string) {
	o.LoadBalanceMethod = &v
}

// GetConnectionPreference returns the ConnectionPreference field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetConnectionPreference() []string {
	if o == nil || IsNil(o.ConnectionPreference) {
		var ret []string
		return ret
	}
	return o.ConnectionPreference
}

// GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetConnectionPreferenceOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectionPreference) {
		return nil, false
	}
	return o.ConnectionPreference, true
}

// HasConnectionPreference returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasConnectionPreference() bool {
	if o != nil && !IsNil(o.ConnectionPreference) {
		return true
	}

	return false
}

// SetConnectionPreference gets a reference to the given []string and assigns it to the ConnectionPreference field.
func (o *EdgeConnectorStorageRequest) SetConnectionPreference(v []string) {
	o.ConnectionPreference = v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *EdgeConnectorStorageRequest) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetReadWriteTimeout returns the ReadWriteTimeout field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetReadWriteTimeout() int64 {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		var ret int64
		return ret
	}
	return *o.ReadWriteTimeout
}

// GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetReadWriteTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		return nil, false
	}
	return o.ReadWriteTimeout, true
}

// HasReadWriteTimeout returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasReadWriteTimeout() bool {
	if o != nil && !IsNil(o.ReadWriteTimeout) {
		return true
	}

	return false
}

// SetReadWriteTimeout gets a reference to the given int64 and assigns it to the ReadWriteTimeout field.
func (o *EdgeConnectorStorageRequest) SetReadWriteTimeout(v int64) {
	o.ReadWriteTimeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *EdgeConnectorStorageRequest) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int64
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetMaxRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *EdgeConnectorStorageRequest) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int64 and assigns it to the MaxRetries field.
func (o *EdgeConnectorStorageRequest) SetMaxRetries(v int64) {
	o.MaxRetries = &v
}

// GetTypeProperties returns the TypeProperties field value
func (o *EdgeConnectorStorageRequest) GetTypeProperties() EdgeConnectorStorageTypePropertiesRequest {
	if o == nil {
		var ret EdgeConnectorStorageTypePropertiesRequest
		return ret
	}

	return o.TypeProperties
}

// GetTypePropertiesOk returns a tuple with the TypeProperties field value
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageRequest) GetTypePropertiesOk() (*EdgeConnectorStorageTypePropertiesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeProperties, true
}

// SetTypeProperties sets field value
func (o *EdgeConnectorStorageRequest) SetTypeProperties(v EdgeConnectorStorageTypePropertiesRequest) {
	o.TypeProperties = v
}

func (o EdgeConnectorStorageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeConnectorStorageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["modules"] = o.Modules
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	if !IsNil(o.LoadBalanceMethod) {
		toSerialize["load_balance_method"] = o.LoadBalanceMethod
	}
	if !IsNil(o.ConnectionPreference) {
		toSerialize["connection_preference"] = o.ConnectionPreference
	}
	if !IsNil(o.ConnectionTimeout) {
		toSerialize["connection_timeout"] = o.ConnectionTimeout
	}
	if !IsNil(o.ReadWriteTimeout) {
		toSerialize["read_write_timeout"] = o.ReadWriteTimeout
	}
	if !IsNil(o.MaxRetries) {
		toSerialize["max_retries"] = o.MaxRetries
	}
	toSerialize["type_properties"] = o.TypeProperties
	return toSerialize, nil
}

func (o *EdgeConnectorStorageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"modules",
		"type",
		"type_properties",
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

	varEdgeConnectorStorageRequest := _EdgeConnectorStorageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeConnectorStorageRequest)

	if err != nil {
		return err
	}

	*o = EdgeConnectorStorageRequest(varEdgeConnectorStorageRequest)

	return err
}

type NullableEdgeConnectorStorageRequest struct {
	value *EdgeConnectorStorageRequest
	isSet bool
}

func (v NullableEdgeConnectorStorageRequest) Get() *EdgeConnectorStorageRequest {
	return v.value
}

func (v *NullableEdgeConnectorStorageRequest) Set(val *EdgeConnectorStorageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeConnectorStorageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeConnectorStorageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeConnectorStorageRequest(val *EdgeConnectorStorageRequest) *NullableEdgeConnectorStorageRequest {
	return &NullableEdgeConnectorStorageRequest{value: val, isSet: true}
}

func (v NullableEdgeConnectorStorageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeConnectorStorageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


