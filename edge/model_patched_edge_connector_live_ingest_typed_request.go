/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
)

// checks if the PatchedEdgeConnectorLiveIngestTypedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedEdgeConnectorLiveIngestTypedRequest{}

// PatchedEdgeConnectorLiveIngestTypedRequest struct for PatchedEdgeConnectorLiveIngestTypedRequest
type PatchedEdgeConnectorLiveIngestTypedRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	Modules *EdgeConnectorModulesRequest `json:"modules,omitempty"`
	Active *bool `json:"active,omitempty"`
	Type *string `json:"type,omitempty"`
	Addresses []AddressRequest `json:"addresses,omitempty"`
	Tls *TLSEdgeConnectorRequest `json:"tls,omitempty"`
	// * `off` - Off * `ip_hash` - IP Hash * `least_connections` - Least Connections * `round_robin` - Round Robin
	LoadBalanceMethod *string `json:"load_balance_method,omitempty"`
	ConnectionPreference []string `json:"connection_preference,omitempty"`
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`
	ReadWriteTimeout *int64 `json:"read_write_timeout,omitempty"`
	MaxRetries *int64 `json:"max_retries,omitempty"`
	TypeProperties *EdgeConnectorLiveIngestTypePropertiesRequest `json:"type_properties,omitempty"`
}

// NewPatchedEdgeConnectorLiveIngestTypedRequest instantiates a new PatchedEdgeConnectorLiveIngestTypedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEdgeConnectorLiveIngestTypedRequest() *PatchedEdgeConnectorLiveIngestTypedRequest {
	this := PatchedEdgeConnectorLiveIngestTypedRequest{}
	return &this
}

// NewPatchedEdgeConnectorLiveIngestTypedRequestWithDefaults instantiates a new PatchedEdgeConnectorLiveIngestTypedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEdgeConnectorLiveIngestTypedRequestWithDefaults() *PatchedEdgeConnectorLiveIngestTypedRequest {
	this := PatchedEdgeConnectorLiveIngestTypedRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetName(v string) {
	o.Name = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetModules() EdgeConnectorModulesRequest {
	if o == nil || IsNil(o.Modules) {
		var ret EdgeConnectorModulesRequest
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetModulesOk() (*EdgeConnectorModulesRequest, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given EdgeConnectorModulesRequest and assigns it to the Modules field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetModules(v EdgeConnectorModulesRequest) {
	o.Modules = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetActive(v bool) {
	o.Active = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetType(v string) {
	o.Type = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetAddresses() []AddressRequest {
	if o == nil || IsNil(o.Addresses) {
		var ret []AddressRequest
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetAddressesOk() ([]AddressRequest, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []AddressRequest and assigns it to the Addresses field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetAddresses(v []AddressRequest) {
	o.Addresses = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetTls() TLSEdgeConnectorRequest {
	if o == nil || IsNil(o.Tls) {
		var ret TLSEdgeConnectorRequest
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetTlsOk() (*TLSEdgeConnectorRequest, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given TLSEdgeConnectorRequest and assigns it to the Tls field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetTls(v TLSEdgeConnectorRequest) {
	o.Tls = &v
}

// GetLoadBalanceMethod returns the LoadBalanceMethod field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetLoadBalanceMethod() string {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		var ret string
		return ret
	}
	return *o.LoadBalanceMethod
}

// GetLoadBalanceMethodOk returns a tuple with the LoadBalanceMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetLoadBalanceMethodOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceMethod) {
		return nil, false
	}
	return o.LoadBalanceMethod, true
}

// HasLoadBalanceMethod returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasLoadBalanceMethod() bool {
	if o != nil && !IsNil(o.LoadBalanceMethod) {
		return true
	}

	return false
}

// SetLoadBalanceMethod gets a reference to the given string and assigns it to the LoadBalanceMethod field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetLoadBalanceMethod(v string) {
	o.LoadBalanceMethod = &v
}

// GetConnectionPreference returns the ConnectionPreference field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetConnectionPreference() []string {
	if o == nil || IsNil(o.ConnectionPreference) {
		var ret []string
		return ret
	}
	return o.ConnectionPreference
}

// GetConnectionPreferenceOk returns a tuple with the ConnectionPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetConnectionPreferenceOk() ([]string, bool) {
	if o == nil || IsNil(o.ConnectionPreference) {
		return nil, false
	}
	return o.ConnectionPreference, true
}

// HasConnectionPreference returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasConnectionPreference() bool {
	if o != nil && !IsNil(o.ConnectionPreference) {
		return true
	}

	return false
}

// SetConnectionPreference gets a reference to the given []string and assigns it to the ConnectionPreference field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetConnectionPreference(v []string) {
	o.ConnectionPreference = v
}

// GetConnectionTimeout returns the ConnectionTimeout field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetConnectionTimeout() int64 {
	if o == nil || IsNil(o.ConnectionTimeout) {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ConnectionTimeout) {
		return nil, false
	}
	return o.ConnectionTimeout, true
}

// HasConnectionTimeout returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasConnectionTimeout() bool {
	if o != nil && !IsNil(o.ConnectionTimeout) {
		return true
	}

	return false
}

// SetConnectionTimeout gets a reference to the given int64 and assigns it to the ConnectionTimeout field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = &v
}

// GetReadWriteTimeout returns the ReadWriteTimeout field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetReadWriteTimeout() int64 {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		var ret int64
		return ret
	}
	return *o.ReadWriteTimeout
}

// GetReadWriteTimeoutOk returns a tuple with the ReadWriteTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetReadWriteTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.ReadWriteTimeout) {
		return nil, false
	}
	return o.ReadWriteTimeout, true
}

// HasReadWriteTimeout returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasReadWriteTimeout() bool {
	if o != nil && !IsNil(o.ReadWriteTimeout) {
		return true
	}

	return false
}

// SetReadWriteTimeout gets a reference to the given int64 and assigns it to the ReadWriteTimeout field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetReadWriteTimeout(v int64) {
	o.ReadWriteTimeout = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetMaxRetries() int64 {
	if o == nil || IsNil(o.MaxRetries) {
		var ret int64
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetMaxRetriesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRetries) {
		return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasMaxRetries() bool {
	if o != nil && !IsNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int64 and assigns it to the MaxRetries field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetMaxRetries(v int64) {
	o.MaxRetries = &v
}

// GetTypeProperties returns the TypeProperties field value if set, zero value otherwise.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetTypeProperties() EdgeConnectorLiveIngestTypePropertiesRequest {
	if o == nil || IsNil(o.TypeProperties) {
		var ret EdgeConnectorLiveIngestTypePropertiesRequest
		return ret
	}
	return *o.TypeProperties
}

// GetTypePropertiesOk returns a tuple with the TypeProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) GetTypePropertiesOk() (*EdgeConnectorLiveIngestTypePropertiesRequest, bool) {
	if o == nil || IsNil(o.TypeProperties) {
		return nil, false
	}
	return o.TypeProperties, true
}

// HasTypeProperties returns a boolean if a field has been set.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) HasTypeProperties() bool {
	if o != nil && !IsNil(o.TypeProperties) {
		return true
	}

	return false
}

// SetTypeProperties gets a reference to the given EdgeConnectorLiveIngestTypePropertiesRequest and assigns it to the TypeProperties field.
func (o *PatchedEdgeConnectorLiveIngestTypedRequest) SetTypeProperties(v EdgeConnectorLiveIngestTypePropertiesRequest) {
	o.TypeProperties = &v
}

func (o PatchedEdgeConnectorLiveIngestTypedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedEdgeConnectorLiveIngestTypedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
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
	if !IsNil(o.TypeProperties) {
		toSerialize["type_properties"] = o.TypeProperties
	}
	return toSerialize, nil
}

type NullablePatchedEdgeConnectorLiveIngestTypedRequest struct {
	value *PatchedEdgeConnectorLiveIngestTypedRequest
	isSet bool
}

func (v NullablePatchedEdgeConnectorLiveIngestTypedRequest) Get() *PatchedEdgeConnectorLiveIngestTypedRequest {
	return v.value
}

func (v *NullablePatchedEdgeConnectorLiveIngestTypedRequest) Set(val *PatchedEdgeConnectorLiveIngestTypedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEdgeConnectorLiveIngestTypedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEdgeConnectorLiveIngestTypedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEdgeConnectorLiveIngestTypedRequest(val *PatchedEdgeConnectorLiveIngestTypedRequest) *NullablePatchedEdgeConnectorLiveIngestTypedRequest {
	return &NullablePatchedEdgeConnectorLiveIngestTypedRequest{value: val, isSet: true}
}

func (v NullablePatchedEdgeConnectorLiveIngestTypedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEdgeConnectorLiveIngestTypedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


