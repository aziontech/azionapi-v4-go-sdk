/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
)

// checks if the EdgeFirewallModulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallModulesRequest{}

// EdgeFirewallModulesRequest struct for EdgeFirewallModulesRequest
type EdgeFirewallModulesRequest struct {
	DdosProtectionEnabled *bool `json:"ddos_protection_enabled,omitempty"`
	EdgeFunctionsEnabled *bool `json:"edge_functions_enabled,omitempty"`
	NetworkProtectionEnabled *bool `json:"network_protection_enabled,omitempty"`
	WafEnabled *bool `json:"waf_enabled,omitempty"`
}

// NewEdgeFirewallModulesRequest instantiates a new EdgeFirewallModulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallModulesRequest() *EdgeFirewallModulesRequest {
	this := EdgeFirewallModulesRequest{}
	var ddosProtectionEnabled bool = true
	this.DdosProtectionEnabled = &ddosProtectionEnabled
	var edgeFunctionsEnabled bool = true
	this.EdgeFunctionsEnabled = &edgeFunctionsEnabled
	var networkProtectionEnabled bool = true
	this.NetworkProtectionEnabled = &networkProtectionEnabled
	var wafEnabled bool = true
	this.WafEnabled = &wafEnabled
	return &this
}

// NewEdgeFirewallModulesRequestWithDefaults instantiates a new EdgeFirewallModulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallModulesRequestWithDefaults() *EdgeFirewallModulesRequest {
	this := EdgeFirewallModulesRequest{}
	var ddosProtectionEnabled bool = true
	this.DdosProtectionEnabled = &ddosProtectionEnabled
	var edgeFunctionsEnabled bool = true
	this.EdgeFunctionsEnabled = &edgeFunctionsEnabled
	var networkProtectionEnabled bool = true
	this.NetworkProtectionEnabled = &networkProtectionEnabled
	var wafEnabled bool = true
	this.WafEnabled = &wafEnabled
	return &this
}

// GetDdosProtectionEnabled returns the DdosProtectionEnabled field value if set, zero value otherwise.
func (o *EdgeFirewallModulesRequest) GetDdosProtectionEnabled() bool {
	if o == nil || IsNil(o.DdosProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.DdosProtectionEnabled
}

// GetDdosProtectionEnabledOk returns a tuple with the DdosProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModulesRequest) GetDdosProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DdosProtectionEnabled) {
		return nil, false
	}
	return o.DdosProtectionEnabled, true
}

// HasDdosProtectionEnabled returns a boolean if a field has been set.
func (o *EdgeFirewallModulesRequest) HasDdosProtectionEnabled() bool {
	if o != nil && !IsNil(o.DdosProtectionEnabled) {
		return true
	}

	return false
}

// SetDdosProtectionEnabled gets a reference to the given bool and assigns it to the DdosProtectionEnabled field.
func (o *EdgeFirewallModulesRequest) SetDdosProtectionEnabled(v bool) {
	o.DdosProtectionEnabled = &v
}

// GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field value if set, zero value otherwise.
func (o *EdgeFirewallModulesRequest) GetEdgeFunctionsEnabled() bool {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		var ret bool
		return ret
	}
	return *o.EdgeFunctionsEnabled
}

// GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModulesRequest) GetEdgeFunctionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		return nil, false
	}
	return o.EdgeFunctionsEnabled, true
}

// HasEdgeFunctionsEnabled returns a boolean if a field has been set.
func (o *EdgeFirewallModulesRequest) HasEdgeFunctionsEnabled() bool {
	if o != nil && !IsNil(o.EdgeFunctionsEnabled) {
		return true
	}

	return false
}

// SetEdgeFunctionsEnabled gets a reference to the given bool and assigns it to the EdgeFunctionsEnabled field.
func (o *EdgeFirewallModulesRequest) SetEdgeFunctionsEnabled(v bool) {
	o.EdgeFunctionsEnabled = &v
}

// GetNetworkProtectionEnabled returns the NetworkProtectionEnabled field value if set, zero value otherwise.
func (o *EdgeFirewallModulesRequest) GetNetworkProtectionEnabled() bool {
	if o == nil || IsNil(o.NetworkProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.NetworkProtectionEnabled
}

// GetNetworkProtectionEnabledOk returns a tuple with the NetworkProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModulesRequest) GetNetworkProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NetworkProtectionEnabled) {
		return nil, false
	}
	return o.NetworkProtectionEnabled, true
}

// HasNetworkProtectionEnabled returns a boolean if a field has been set.
func (o *EdgeFirewallModulesRequest) HasNetworkProtectionEnabled() bool {
	if o != nil && !IsNil(o.NetworkProtectionEnabled) {
		return true
	}

	return false
}

// SetNetworkProtectionEnabled gets a reference to the given bool and assigns it to the NetworkProtectionEnabled field.
func (o *EdgeFirewallModulesRequest) SetNetworkProtectionEnabled(v bool) {
	o.NetworkProtectionEnabled = &v
}

// GetWafEnabled returns the WafEnabled field value if set, zero value otherwise.
func (o *EdgeFirewallModulesRequest) GetWafEnabled() bool {
	if o == nil || IsNil(o.WafEnabled) {
		var ret bool
		return ret
	}
	return *o.WafEnabled
}

// GetWafEnabledOk returns a tuple with the WafEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallModulesRequest) GetWafEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WafEnabled) {
		return nil, false
	}
	return o.WafEnabled, true
}

// HasWafEnabled returns a boolean if a field has been set.
func (o *EdgeFirewallModulesRequest) HasWafEnabled() bool {
	if o != nil && !IsNil(o.WafEnabled) {
		return true
	}

	return false
}

// SetWafEnabled gets a reference to the given bool and assigns it to the WafEnabled field.
func (o *EdgeFirewallModulesRequest) SetWafEnabled(v bool) {
	o.WafEnabled = &v
}

func (o EdgeFirewallModulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallModulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdosProtectionEnabled) {
		toSerialize["ddos_protection_enabled"] = o.DdosProtectionEnabled
	}
	if !IsNil(o.EdgeFunctionsEnabled) {
		toSerialize["edge_functions_enabled"] = o.EdgeFunctionsEnabled
	}
	if !IsNil(o.NetworkProtectionEnabled) {
		toSerialize["network_protection_enabled"] = o.NetworkProtectionEnabled
	}
	if !IsNil(o.WafEnabled) {
		toSerialize["waf_enabled"] = o.WafEnabled
	}
	return toSerialize, nil
}

type NullableEdgeFirewallModulesRequest struct {
	value *EdgeFirewallModulesRequest
	isSet bool
}

func (v NullableEdgeFirewallModulesRequest) Get() *EdgeFirewallModulesRequest {
	return v.value
}

func (v *NullableEdgeFirewallModulesRequest) Set(val *EdgeFirewallModulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallModulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallModulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallModulesRequest(val *EdgeFirewallModulesRequest) *NullableEdgeFirewallModulesRequest {
	return &NullableEdgeFirewallModulesRequest{value: val, isSet: true}
}

func (v NullableEdgeFirewallModulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallModulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


