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

// checks if the TLSRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TLSRequest{}

// TLSRequest struct for TLSRequest
type TLSRequest struct {
	Certificate NullableInt64 `json:"certificate,omitempty"`
	Ciphers NullableTLSCiphers `json:"ciphers,omitempty"`
	MinimumVersion NullableTLSMinimumVersion `json:"minimum_version,omitempty"`
}

// NewTLSRequest instantiates a new TLSRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSRequest() *TLSRequest {
	this := TLSRequest{}
	var minimumVersion TLSMinimumVersion = tls_1_2
	this.MinimumVersion = *NewNullableTLSMinimumVersion(&minimumVersion)
	return &this
}

// NewTLSRequestWithDefaults instantiates a new TLSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSRequestWithDefaults() *TLSRequest {
	this := TLSRequest{}
	var minimumVersion TLSMinimumVersion = tls_1_2
	this.MinimumVersion = *NewNullableTLSMinimumVersion(&minimumVersion)
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSRequest) GetCertificate() int64 {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret int64
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSRequest) GetCertificateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *TLSRequest) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableInt64 and assigns it to the Certificate field.
func (o *TLSRequest) SetCertificate(v int64) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *TLSRequest) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *TLSRequest) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCiphers returns the Ciphers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSRequest) GetCiphers() TLSCiphers {
	if o == nil || IsNil(o.Ciphers.Get()) {
		var ret TLSCiphers
		return ret
	}
	return *o.Ciphers.Get()
}

// GetCiphersOk returns a tuple with the Ciphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSRequest) GetCiphersOk() (*TLSCiphers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ciphers.Get(), o.Ciphers.IsSet()
}

// HasCiphers returns a boolean if a field has been set.
func (o *TLSRequest) HasCiphers() bool {
	if o != nil && o.Ciphers.IsSet() {
		return true
	}

	return false
}

// SetCiphers gets a reference to the given NullableTLSCiphers and assigns it to the Ciphers field.
func (o *TLSRequest) SetCiphers(v TLSCiphers) {
	o.Ciphers.Set(&v)
}
// SetCiphersNil sets the value for Ciphers to be an explicit nil
func (o *TLSRequest) SetCiphersNil() {
	o.Ciphers.Set(nil)
}

// UnsetCiphers ensures that no value is present for Ciphers, not even an explicit nil
func (o *TLSRequest) UnsetCiphers() {
	o.Ciphers.Unset()
}

// GetMinimumVersion returns the MinimumVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSRequest) GetMinimumVersion() TLSMinimumVersion {
	if o == nil || IsNil(o.MinimumVersion.Get()) {
		var ret TLSMinimumVersion
		return ret
	}
	return *o.MinimumVersion.Get()
}

// GetMinimumVersionOk returns a tuple with the MinimumVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSRequest) GetMinimumVersionOk() (*TLSMinimumVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumVersion.Get(), o.MinimumVersion.IsSet()
}

// HasMinimumVersion returns a boolean if a field has been set.
func (o *TLSRequest) HasMinimumVersion() bool {
	if o != nil && o.MinimumVersion.IsSet() {
		return true
	}

	return false
}

// SetMinimumVersion gets a reference to the given NullableTLSMinimumVersion and assigns it to the MinimumVersion field.
func (o *TLSRequest) SetMinimumVersion(v TLSMinimumVersion) {
	o.MinimumVersion.Set(&v)
}
// SetMinimumVersionNil sets the value for MinimumVersion to be an explicit nil
func (o *TLSRequest) SetMinimumVersionNil() {
	o.MinimumVersion.Set(nil)
}

// UnsetMinimumVersion ensures that no value is present for MinimumVersion, not even an explicit nil
func (o *TLSRequest) UnsetMinimumVersion() {
	o.MinimumVersion.Unset()
}

func (o TLSRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TLSRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.Ciphers.IsSet() {
		toSerialize["ciphers"] = o.Ciphers.Get()
	}
	if o.MinimumVersion.IsSet() {
		toSerialize["minimum_version"] = o.MinimumVersion.Get()
	}
	return toSerialize, nil
}

type NullableTLSRequest struct {
	value *TLSRequest
	isSet bool
}

func (v NullableTLSRequest) Get() *TLSRequest {
	return v.value
}

func (v *NullableTLSRequest) Set(val *TLSRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSRequest(val *TLSRequest) *NullableTLSRequest {
	return &NullableTLSRequest{value: val, isSet: true}
}

func (v NullableTLSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


