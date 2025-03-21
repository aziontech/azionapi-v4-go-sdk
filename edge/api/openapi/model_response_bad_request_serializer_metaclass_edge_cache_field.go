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

// checks if the ResponseBadRequestSerializerMetaclassEdgeCacheField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestSerializerMetaclassEdgeCacheField{}

// ResponseBadRequestSerializerMetaclassEdgeCacheField struct for ResponseBadRequestSerializerMetaclassEdgeCacheField
type ResponseBadRequestSerializerMetaclassEdgeCacheField struct {
	Behavior []string `json:"behavior,omitempty"`
	MaxAge []string `json:"max_age,omitempty"`
	CachingForPostEnabled []string `json:"caching_for_post_enabled,omitempty"`
	CachingForOptionsEnabled []string `json:"caching_for_options_enabled,omitempty"`
	StaleCacheEnabled []string `json:"stale_cache_enabled,omitempty"`
	TieredCacheEnabled []string `json:"tiered_cache_enabled,omitempty"`
	TieredCacheRegion []string `json:"tiered_cache_region,omitempty"`
}

// NewResponseBadRequestSerializerMetaclassEdgeCacheField instantiates a new ResponseBadRequestSerializerMetaclassEdgeCacheField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestSerializerMetaclassEdgeCacheField() *ResponseBadRequestSerializerMetaclassEdgeCacheField {
	this := ResponseBadRequestSerializerMetaclassEdgeCacheField{}
	return &this
}

// NewResponseBadRequestSerializerMetaclassEdgeCacheFieldWithDefaults instantiates a new ResponseBadRequestSerializerMetaclassEdgeCacheField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestSerializerMetaclassEdgeCacheFieldWithDefaults() *ResponseBadRequestSerializerMetaclassEdgeCacheField {
	this := ResponseBadRequestSerializerMetaclassEdgeCacheField{}
	return &this
}

// GetBehavior returns the Behavior field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetBehavior() []string {
	if o == nil || IsNil(o.Behavior) {
		var ret []string
		return ret
	}
	return o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetBehaviorOk() ([]string, bool) {
	if o == nil || IsNil(o.Behavior) {
		return nil, false
	}
	return o.Behavior, true
}

// HasBehavior returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasBehavior() bool {
	if o != nil && !IsNil(o.Behavior) {
		return true
	}

	return false
}

// SetBehavior gets a reference to the given []string and assigns it to the Behavior field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetBehavior(v []string) {
	o.Behavior = v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetMaxAge() []string {
	if o == nil || IsNil(o.MaxAge) {
		var ret []string
		return ret
	}
	return o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetMaxAgeOk() ([]string, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given []string and assigns it to the MaxAge field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetMaxAge(v []string) {
	o.MaxAge = v
}

// GetCachingForPostEnabled returns the CachingForPostEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetCachingForPostEnabled() []string {
	if o == nil || IsNil(o.CachingForPostEnabled) {
		var ret []string
		return ret
	}
	return o.CachingForPostEnabled
}

// GetCachingForPostEnabledOk returns a tuple with the CachingForPostEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetCachingForPostEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.CachingForPostEnabled) {
		return nil, false
	}
	return o.CachingForPostEnabled, true
}

// HasCachingForPostEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasCachingForPostEnabled() bool {
	if o != nil && !IsNil(o.CachingForPostEnabled) {
		return true
	}

	return false
}

// SetCachingForPostEnabled gets a reference to the given []string and assigns it to the CachingForPostEnabled field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetCachingForPostEnabled(v []string) {
	o.CachingForPostEnabled = v
}

// GetCachingForOptionsEnabled returns the CachingForOptionsEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetCachingForOptionsEnabled() []string {
	if o == nil || IsNil(o.CachingForOptionsEnabled) {
		var ret []string
		return ret
	}
	return o.CachingForOptionsEnabled
}

// GetCachingForOptionsEnabledOk returns a tuple with the CachingForOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetCachingForOptionsEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.CachingForOptionsEnabled) {
		return nil, false
	}
	return o.CachingForOptionsEnabled, true
}

// HasCachingForOptionsEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasCachingForOptionsEnabled() bool {
	if o != nil && !IsNil(o.CachingForOptionsEnabled) {
		return true
	}

	return false
}

// SetCachingForOptionsEnabled gets a reference to the given []string and assigns it to the CachingForOptionsEnabled field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetCachingForOptionsEnabled(v []string) {
	o.CachingForOptionsEnabled = v
}

// GetStaleCacheEnabled returns the StaleCacheEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetStaleCacheEnabled() []string {
	if o == nil || IsNil(o.StaleCacheEnabled) {
		var ret []string
		return ret
	}
	return o.StaleCacheEnabled
}

// GetStaleCacheEnabledOk returns a tuple with the StaleCacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetStaleCacheEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.StaleCacheEnabled) {
		return nil, false
	}
	return o.StaleCacheEnabled, true
}

// HasStaleCacheEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasStaleCacheEnabled() bool {
	if o != nil && !IsNil(o.StaleCacheEnabled) {
		return true
	}

	return false
}

// SetStaleCacheEnabled gets a reference to the given []string and assigns it to the StaleCacheEnabled field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetStaleCacheEnabled(v []string) {
	o.StaleCacheEnabled = v
}

// GetTieredCacheEnabled returns the TieredCacheEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetTieredCacheEnabled() []string {
	if o == nil || IsNil(o.TieredCacheEnabled) {
		var ret []string
		return ret
	}
	return o.TieredCacheEnabled
}

// GetTieredCacheEnabledOk returns a tuple with the TieredCacheEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetTieredCacheEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.TieredCacheEnabled) {
		return nil, false
	}
	return o.TieredCacheEnabled, true
}

// HasTieredCacheEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasTieredCacheEnabled() bool {
	if o != nil && !IsNil(o.TieredCacheEnabled) {
		return true
	}

	return false
}

// SetTieredCacheEnabled gets a reference to the given []string and assigns it to the TieredCacheEnabled field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetTieredCacheEnabled(v []string) {
	o.TieredCacheEnabled = v
}

// GetTieredCacheRegion returns the TieredCacheRegion field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetTieredCacheRegion() []string {
	if o == nil || IsNil(o.TieredCacheRegion) {
		var ret []string
		return ret
	}
	return o.TieredCacheRegion
}

// GetTieredCacheRegionOk returns a tuple with the TieredCacheRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) GetTieredCacheRegionOk() ([]string, bool) {
	if o == nil || IsNil(o.TieredCacheRegion) {
		return nil, false
	}
	return o.TieredCacheRegion, true
}

// HasTieredCacheRegion returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) HasTieredCacheRegion() bool {
	if o != nil && !IsNil(o.TieredCacheRegion) {
		return true
	}

	return false
}

// SetTieredCacheRegion gets a reference to the given []string and assigns it to the TieredCacheRegion field.
func (o *ResponseBadRequestSerializerMetaclassEdgeCacheField) SetTieredCacheRegion(v []string) {
	o.TieredCacheRegion = v
}

func (o ResponseBadRequestSerializerMetaclassEdgeCacheField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestSerializerMetaclassEdgeCacheField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Behavior) {
		toSerialize["behavior"] = o.Behavior
	}
	if !IsNil(o.MaxAge) {
		toSerialize["max_age"] = o.MaxAge
	}
	if !IsNil(o.CachingForPostEnabled) {
		toSerialize["caching_for_post_enabled"] = o.CachingForPostEnabled
	}
	if !IsNil(o.CachingForOptionsEnabled) {
		toSerialize["caching_for_options_enabled"] = o.CachingForOptionsEnabled
	}
	if !IsNil(o.StaleCacheEnabled) {
		toSerialize["stale_cache_enabled"] = o.StaleCacheEnabled
	}
	if !IsNil(o.TieredCacheEnabled) {
		toSerialize["tiered_cache_enabled"] = o.TieredCacheEnabled
	}
	if !IsNil(o.TieredCacheRegion) {
		toSerialize["tiered_cache_region"] = o.TieredCacheRegion
	}
	return toSerialize, nil
}

type NullableResponseBadRequestSerializerMetaclassEdgeCacheField struct {
	value *ResponseBadRequestSerializerMetaclassEdgeCacheField
	isSet bool
}

func (v NullableResponseBadRequestSerializerMetaclassEdgeCacheField) Get() *ResponseBadRequestSerializerMetaclassEdgeCacheField {
	return v.value
}

func (v *NullableResponseBadRequestSerializerMetaclassEdgeCacheField) Set(val *ResponseBadRequestSerializerMetaclassEdgeCacheField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestSerializerMetaclassEdgeCacheField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestSerializerMetaclassEdgeCacheField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestSerializerMetaclassEdgeCacheField(val *ResponseBadRequestSerializerMetaclassEdgeCacheField) *NullableResponseBadRequestSerializerMetaclassEdgeCacheField {
	return &NullableResponseBadRequestSerializerMetaclassEdgeCacheField{value: val, isSet: true}
}

func (v NullableResponseBadRequestSerializerMetaclassEdgeCacheField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestSerializerMetaclassEdgeCacheField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


