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

// checks if the SliceControlsModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceControlsModuleRequest{}

// SliceControlsModuleRequest struct for SliceControlsModuleRequest
type SliceControlsModuleRequest struct {
	SliceConfigurationEnabled bool `json:"slice_configuration_enabled"`
	SliceEdgeCachingEnabled bool `json:"slice_edge_caching_enabled"`
	SliceTieredCachingEnabled bool `json:"slice_tiered_caching_enabled"`
	SliceConfigurationRange *int64 `json:"slice_configuration_range,omitempty"`
}

type _SliceControlsModuleRequest SliceControlsModuleRequest

// NewSliceControlsModuleRequest instantiates a new SliceControlsModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceControlsModuleRequest(sliceConfigurationEnabled bool, sliceEdgeCachingEnabled bool, sliceTieredCachingEnabled bool) *SliceControlsModuleRequest {
	this := SliceControlsModuleRequest{}
	this.SliceConfigurationEnabled = sliceConfigurationEnabled
	this.SliceEdgeCachingEnabled = sliceEdgeCachingEnabled
	this.SliceTieredCachingEnabled = sliceTieredCachingEnabled
	return &this
}

// NewSliceControlsModuleRequestWithDefaults instantiates a new SliceControlsModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceControlsModuleRequestWithDefaults() *SliceControlsModuleRequest {
	this := SliceControlsModuleRequest{}
	return &this
}

// GetSliceConfigurationEnabled returns the SliceConfigurationEnabled field value
func (o *SliceControlsModuleRequest) GetSliceConfigurationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SliceConfigurationEnabled
}

// GetSliceConfigurationEnabledOk returns a tuple with the SliceConfigurationEnabled field value
// and a boolean to check if the value has been set.
func (o *SliceControlsModuleRequest) GetSliceConfigurationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliceConfigurationEnabled, true
}

// SetSliceConfigurationEnabled sets field value
func (o *SliceControlsModuleRequest) SetSliceConfigurationEnabled(v bool) {
	o.SliceConfigurationEnabled = v
}

// GetSliceEdgeCachingEnabled returns the SliceEdgeCachingEnabled field value
func (o *SliceControlsModuleRequest) GetSliceEdgeCachingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SliceEdgeCachingEnabled
}

// GetSliceEdgeCachingEnabledOk returns a tuple with the SliceEdgeCachingEnabled field value
// and a boolean to check if the value has been set.
func (o *SliceControlsModuleRequest) GetSliceEdgeCachingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliceEdgeCachingEnabled, true
}

// SetSliceEdgeCachingEnabled sets field value
func (o *SliceControlsModuleRequest) SetSliceEdgeCachingEnabled(v bool) {
	o.SliceEdgeCachingEnabled = v
}

// GetSliceTieredCachingEnabled returns the SliceTieredCachingEnabled field value
func (o *SliceControlsModuleRequest) GetSliceTieredCachingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SliceTieredCachingEnabled
}

// GetSliceTieredCachingEnabledOk returns a tuple with the SliceTieredCachingEnabled field value
// and a boolean to check if the value has been set.
func (o *SliceControlsModuleRequest) GetSliceTieredCachingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliceTieredCachingEnabled, true
}

// SetSliceTieredCachingEnabled sets field value
func (o *SliceControlsModuleRequest) SetSliceTieredCachingEnabled(v bool) {
	o.SliceTieredCachingEnabled = v
}

// GetSliceConfigurationRange returns the SliceConfigurationRange field value if set, zero value otherwise.
func (o *SliceControlsModuleRequest) GetSliceConfigurationRange() int64 {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		var ret int64
		return ret
	}
	return *o.SliceConfigurationRange
}

// GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceControlsModuleRequest) GetSliceConfigurationRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		return nil, false
	}
	return o.SliceConfigurationRange, true
}

// HasSliceConfigurationRange returns a boolean if a field has been set.
func (o *SliceControlsModuleRequest) HasSliceConfigurationRange() bool {
	if o != nil && !IsNil(o.SliceConfigurationRange) {
		return true
	}

	return false
}

// SetSliceConfigurationRange gets a reference to the given int64 and assigns it to the SliceConfigurationRange field.
func (o *SliceControlsModuleRequest) SetSliceConfigurationRange(v int64) {
	o.SliceConfigurationRange = &v
}

func (o SliceControlsModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceControlsModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["slice_configuration_enabled"] = o.SliceConfigurationEnabled
	toSerialize["slice_edge_caching_enabled"] = o.SliceEdgeCachingEnabled
	toSerialize["slice_tiered_caching_enabled"] = o.SliceTieredCachingEnabled
	if !IsNil(o.SliceConfigurationRange) {
		toSerialize["slice_configuration_range"] = o.SliceConfigurationRange
	}
	return toSerialize, nil
}

func (o *SliceControlsModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"slice_configuration_enabled",
		"slice_edge_caching_enabled",
		"slice_tiered_caching_enabled",
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

	varSliceControlsModuleRequest := _SliceControlsModuleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSliceControlsModuleRequest)

	if err != nil {
		return err
	}

	*o = SliceControlsModuleRequest(varSliceControlsModuleRequest)

	return err
}

type NullableSliceControlsModuleRequest struct {
	value *SliceControlsModuleRequest
	isSet bool
}

func (v NullableSliceControlsModuleRequest) Get() *SliceControlsModuleRequest {
	return v.value
}

func (v *NullableSliceControlsModuleRequest) Set(val *SliceControlsModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceControlsModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceControlsModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceControlsModuleRequest(val *SliceControlsModuleRequest) *NullableSliceControlsModuleRequest {
	return &NullableSliceControlsModuleRequest{value: val, isSet: true}
}

func (v NullableSliceControlsModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceControlsModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


