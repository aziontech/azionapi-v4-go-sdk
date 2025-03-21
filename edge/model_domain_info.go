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

// checks if the DomainInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainInfo{}

// DomainInfo struct for DomainInfo
type DomainInfo struct {
	Domain NullableString `json:"domain,omitempty" validate:"regexp=.*"`
	AllowAccess *bool `json:"allow_access,omitempty"`
}

// NewDomainInfo instantiates a new DomainInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInfo() *DomainInfo {
	this := DomainInfo{}
	var allowAccess bool = true
	this.AllowAccess = &allowAccess
	return &this
}

// NewDomainInfoWithDefaults instantiates a new DomainInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInfoWithDefaults() *DomainInfo {
	this := DomainInfo{}
	var allowAccess bool = true
	this.AllowAccess = &allowAccess
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainInfo) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainInfo) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainInfo) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *DomainInfo) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *DomainInfo) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *DomainInfo) UnsetDomain() {
	o.Domain.Unset()
}

// GetAllowAccess returns the AllowAccess field value if set, zero value otherwise.
func (o *DomainInfo) GetAllowAccess() bool {
	if o == nil || IsNil(o.AllowAccess) {
		var ret bool
		return ret
	}
	return *o.AllowAccess
}

// GetAllowAccessOk returns a tuple with the AllowAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInfo) GetAllowAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAccess) {
		return nil, false
	}
	return o.AllowAccess, true
}

// HasAllowAccess returns a boolean if a field has been set.
func (o *DomainInfo) HasAllowAccess() bool {
	if o != nil && !IsNil(o.AllowAccess) {
		return true
	}

	return false
}

// SetAllowAccess gets a reference to the given bool and assigns it to the AllowAccess field.
func (o *DomainInfo) SetAllowAccess(v bool) {
	o.AllowAccess = &v
}

func (o DomainInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if !IsNil(o.AllowAccess) {
		toSerialize["allow_access"] = o.AllowAccess
	}
	return toSerialize, nil
}

type NullableDomainInfo struct {
	value *DomainInfo
	isSet bool
}

func (v NullableDomainInfo) Get() *DomainInfo {
	return v.value
}

func (v *NullableDomainInfo) Set(val *DomainInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainInfo(val *DomainInfo) *NullableDomainInfo {
	return &NullableDomainInfo{value: val, isSet: true}
}

func (v NullableDomainInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


