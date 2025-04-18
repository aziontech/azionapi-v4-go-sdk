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

// checks if the EdgeApplicationModulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationModulesRequest{}

// EdgeApplicationModulesRequest struct for EdgeApplicationModulesRequest
type EdgeApplicationModulesRequest struct {
	EdgeCacheEnabled *bool `json:"edge_cache_enabled,omitempty"`
	EdgeFunctionsEnabled *bool `json:"edge_functions_enabled,omitempty"`
	ApplicationAcceleratorEnabled *bool `json:"application_accelerator_enabled,omitempty"`
	ImageProcessorEnabled *bool `json:"image_processor_enabled,omitempty"`
	TieredCacheEnabled *bool `json:"tiered_cache_enabled,omitempty"`
}

// NewEdgeApplicationModulesRequest instantiates a new EdgeApplicationModulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationModulesRequest() *EdgeApplicationModulesRequest {
	this := EdgeApplicationModulesRequest{}
	return &this
}

// NewEdgeApplicationModulesRequestWithDefaults instantiates a new EdgeApplicationModulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationModulesRequestWithDefaults() *EdgeApplicationModulesRequest {
	this := EdgeApplicationModulesRequest{}
	return &this
}

// GetEdgeCacheEnabled returns the EdgeCacheEnabled field value if set, zero value otherwise.
func (o *EdgeApplicationModulesRequest) GetEdgeCacheEnabled() bool {
	if o == nil || IsNil(o.EdgeCacheEnabled) {
		var ret bool
		return ret
	}
	return *o.EdgeCacheEnabled
}

// GetEdgeCacheEnabledOk returns a tuple with the EdgeCacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationModulesRequest) GetEdgeCacheEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeCacheEnabled) {
		return nil, false
	}
	return o.EdgeCacheEnabled, true
}

// HasEdgeCacheEnabled returns a boolean if a field has been set.
func (o *EdgeApplicationModulesRequest) HasEdgeCacheEnabled() bool {
	if o != nil && !IsNil(o.EdgeCacheEnabled) {
		return true
	}

	return false
}

// SetEdgeCacheEnabled gets a reference to the given bool and assigns it to the EdgeCacheEnabled field.
func (o *EdgeApplicationModulesRequest) SetEdgeCacheEnabled(v bool) {
	o.EdgeCacheEnabled = &v
}

// GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field value if set, zero value otherwise.
func (o *EdgeApplicationModulesRequest) GetEdgeFunctionsEnabled() bool {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		var ret bool
		return ret
	}
	return *o.EdgeFunctionsEnabled
}

// GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationModulesRequest) GetEdgeFunctionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		return nil, false
	}
	return o.EdgeFunctionsEnabled, true
}

// HasEdgeFunctionsEnabled returns a boolean if a field has been set.
func (o *EdgeApplicationModulesRequest) HasEdgeFunctionsEnabled() bool {
	if o != nil && !IsNil(o.EdgeFunctionsEnabled) {
		return true
	}

	return false
}

// SetEdgeFunctionsEnabled gets a reference to the given bool and assigns it to the EdgeFunctionsEnabled field.
func (o *EdgeApplicationModulesRequest) SetEdgeFunctionsEnabled(v bool) {
	o.EdgeFunctionsEnabled = &v
}

// GetApplicationAcceleratorEnabled returns the ApplicationAcceleratorEnabled field value if set, zero value otherwise.
func (o *EdgeApplicationModulesRequest) GetApplicationAcceleratorEnabled() bool {
	if o == nil || IsNil(o.ApplicationAcceleratorEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationAcceleratorEnabled
}

// GetApplicationAcceleratorEnabledOk returns a tuple with the ApplicationAcceleratorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationModulesRequest) GetApplicationAcceleratorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationAcceleratorEnabled) {
		return nil, false
	}
	return o.ApplicationAcceleratorEnabled, true
}

// HasApplicationAcceleratorEnabled returns a boolean if a field has been set.
func (o *EdgeApplicationModulesRequest) HasApplicationAcceleratorEnabled() bool {
	if o != nil && !IsNil(o.ApplicationAcceleratorEnabled) {
		return true
	}

	return false
}

// SetApplicationAcceleratorEnabled gets a reference to the given bool and assigns it to the ApplicationAcceleratorEnabled field.
func (o *EdgeApplicationModulesRequest) SetApplicationAcceleratorEnabled(v bool) {
	o.ApplicationAcceleratorEnabled = &v
}

// GetImageProcessorEnabled returns the ImageProcessorEnabled field value if set, zero value otherwise.
func (o *EdgeApplicationModulesRequest) GetImageProcessorEnabled() bool {
	if o == nil || IsNil(o.ImageProcessorEnabled) {
		var ret bool
		return ret
	}
	return *o.ImageProcessorEnabled
}

// GetImageProcessorEnabledOk returns a tuple with the ImageProcessorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationModulesRequest) GetImageProcessorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageProcessorEnabled) {
		return nil, false
	}
	return o.ImageProcessorEnabled, true
}

// HasImageProcessorEnabled returns a boolean if a field has been set.
func (o *EdgeApplicationModulesRequest) HasImageProcessorEnabled() bool {
	if o != nil && !IsNil(o.ImageProcessorEnabled) {
		return true
	}

	return false
}

// SetImageProcessorEnabled gets a reference to the given bool and assigns it to the ImageProcessorEnabled field.
func (o *EdgeApplicationModulesRequest) SetImageProcessorEnabled(v bool) {
	o.ImageProcessorEnabled = &v
}

// GetTieredCacheEnabled returns the TieredCacheEnabled field value if set, zero value otherwise.
func (o *EdgeApplicationModulesRequest) GetTieredCacheEnabled() bool {
	if o == nil || IsNil(o.TieredCacheEnabled) {
		var ret bool
		return ret
	}
	return *o.TieredCacheEnabled
}

// GetTieredCacheEnabledOk returns a tuple with the TieredCacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationModulesRequest) GetTieredCacheEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TieredCacheEnabled) {
		return nil, false
	}
	return o.TieredCacheEnabled, true
}

// HasTieredCacheEnabled returns a boolean if a field has been set.
func (o *EdgeApplicationModulesRequest) HasTieredCacheEnabled() bool {
	if o != nil && !IsNil(o.TieredCacheEnabled) {
		return true
	}

	return false
}

// SetTieredCacheEnabled gets a reference to the given bool and assigns it to the TieredCacheEnabled field.
func (o *EdgeApplicationModulesRequest) SetTieredCacheEnabled(v bool) {
	o.TieredCacheEnabled = &v
}

func (o EdgeApplicationModulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationModulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EdgeCacheEnabled) {
		toSerialize["edge_cache_enabled"] = o.EdgeCacheEnabled
	}
	if !IsNil(o.EdgeFunctionsEnabled) {
		toSerialize["edge_functions_enabled"] = o.EdgeFunctionsEnabled
	}
	if !IsNil(o.ApplicationAcceleratorEnabled) {
		toSerialize["application_accelerator_enabled"] = o.ApplicationAcceleratorEnabled
	}
	if !IsNil(o.ImageProcessorEnabled) {
		toSerialize["image_processor_enabled"] = o.ImageProcessorEnabled
	}
	if !IsNil(o.TieredCacheEnabled) {
		toSerialize["tiered_cache_enabled"] = o.TieredCacheEnabled
	}
	return toSerialize, nil
}

type NullableEdgeApplicationModulesRequest struct {
	value *EdgeApplicationModulesRequest
	isSet bool
}

func (v NullableEdgeApplicationModulesRequest) Get() *EdgeApplicationModulesRequest {
	return v.value
}

func (v *NullableEdgeApplicationModulesRequest) Set(val *EdgeApplicationModulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationModulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationModulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationModulesRequest(val *EdgeApplicationModulesRequest) *NullableEdgeApplicationModulesRequest {
	return &NullableEdgeApplicationModulesRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationModulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationModulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


