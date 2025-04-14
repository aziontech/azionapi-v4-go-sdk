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

// checks if the ResponseBadRequestSerializerMetaclassApplicationControlsField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestSerializerMetaclassApplicationControlsField{}

// ResponseBadRequestSerializerMetaclassApplicationControlsField struct for ResponseBadRequestSerializerMetaclassApplicationControlsField
type ResponseBadRequestSerializerMetaclassApplicationControlsField struct {
	CacheByQueryString []string `json:"cache_by_query_string,omitempty"`
	QueryStringFields *ResponseBadRequestBaseEdgeConnectorConnectionPreference `json:"query_string_fields,omitempty"`
	QueryStringSortEnabled []string `json:"query_string_sort_enabled,omitempty"`
	CacheByCookies []string `json:"cache_by_cookies,omitempty"`
	CookieNames *ResponseBadRequestBaseEdgeConnectorConnectionPreference `json:"cookie_names,omitempty"`
	AdaptiveDeliveryAction []string `json:"adaptive_delivery_action,omitempty"`
	DeviceGroup []string `json:"device_group,omitempty"`
}

// NewResponseBadRequestSerializerMetaclassApplicationControlsField instantiates a new ResponseBadRequestSerializerMetaclassApplicationControlsField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestSerializerMetaclassApplicationControlsField() *ResponseBadRequestSerializerMetaclassApplicationControlsField {
	this := ResponseBadRequestSerializerMetaclassApplicationControlsField{}
	return &this
}

// NewResponseBadRequestSerializerMetaclassApplicationControlsFieldWithDefaults instantiates a new ResponseBadRequestSerializerMetaclassApplicationControlsField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestSerializerMetaclassApplicationControlsFieldWithDefaults() *ResponseBadRequestSerializerMetaclassApplicationControlsField {
	this := ResponseBadRequestSerializerMetaclassApplicationControlsField{}
	return &this
}

// GetCacheByQueryString returns the CacheByQueryString field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCacheByQueryString() []string {
	if o == nil || IsNil(o.CacheByQueryString) {
		var ret []string
		return ret
	}
	return o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCacheByQueryStringOk() ([]string, bool) {
	if o == nil || IsNil(o.CacheByQueryString) {
		return nil, false
	}
	return o.CacheByQueryString, true
}

// HasCacheByQueryString returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasCacheByQueryString() bool {
	if o != nil && !IsNil(o.CacheByQueryString) {
		return true
	}

	return false
}

// SetCacheByQueryString gets a reference to the given []string and assigns it to the CacheByQueryString field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetCacheByQueryString(v []string) {
	o.CacheByQueryString = v
}

// GetQueryStringFields returns the QueryStringFields field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetQueryStringFields() ResponseBadRequestBaseEdgeConnectorConnectionPreference {
	if o == nil || IsNil(o.QueryStringFields) {
		var ret ResponseBadRequestBaseEdgeConnectorConnectionPreference
		return ret
	}
	return *o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetQueryStringFieldsOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool) {
	if o == nil || IsNil(o.QueryStringFields) {
		return nil, false
	}
	return o.QueryStringFields, true
}

// HasQueryStringFields returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasQueryStringFields() bool {
	if o != nil && !IsNil(o.QueryStringFields) {
		return true
	}

	return false
}

// SetQueryStringFields gets a reference to the given ResponseBadRequestBaseEdgeConnectorConnectionPreference and assigns it to the QueryStringFields field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetQueryStringFields(v ResponseBadRequestBaseEdgeConnectorConnectionPreference) {
	o.QueryStringFields = &v
}

// GetQueryStringSortEnabled returns the QueryStringSortEnabled field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetQueryStringSortEnabled() []string {
	if o == nil || IsNil(o.QueryStringSortEnabled) {
		var ret []string
		return ret
	}
	return o.QueryStringSortEnabled
}

// GetQueryStringSortEnabledOk returns a tuple with the QueryStringSortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetQueryStringSortEnabledOk() ([]string, bool) {
	if o == nil || IsNil(o.QueryStringSortEnabled) {
		return nil, false
	}
	return o.QueryStringSortEnabled, true
}

// HasQueryStringSortEnabled returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasQueryStringSortEnabled() bool {
	if o != nil && !IsNil(o.QueryStringSortEnabled) {
		return true
	}

	return false
}

// SetQueryStringSortEnabled gets a reference to the given []string and assigns it to the QueryStringSortEnabled field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetQueryStringSortEnabled(v []string) {
	o.QueryStringSortEnabled = v
}

// GetCacheByCookies returns the CacheByCookies field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCacheByCookies() []string {
	if o == nil || IsNil(o.CacheByCookies) {
		var ret []string
		return ret
	}
	return o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCacheByCookiesOk() ([]string, bool) {
	if o == nil || IsNil(o.CacheByCookies) {
		return nil, false
	}
	return o.CacheByCookies, true
}

// HasCacheByCookies returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasCacheByCookies() bool {
	if o != nil && !IsNil(o.CacheByCookies) {
		return true
	}

	return false
}

// SetCacheByCookies gets a reference to the given []string and assigns it to the CacheByCookies field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetCacheByCookies(v []string) {
	o.CacheByCookies = v
}

// GetCookieNames returns the CookieNames field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCookieNames() ResponseBadRequestBaseEdgeConnectorConnectionPreference {
	if o == nil || IsNil(o.CookieNames) {
		var ret ResponseBadRequestBaseEdgeConnectorConnectionPreference
		return ret
	}
	return *o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetCookieNamesOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool) {
	if o == nil || IsNil(o.CookieNames) {
		return nil, false
	}
	return o.CookieNames, true
}

// HasCookieNames returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasCookieNames() bool {
	if o != nil && !IsNil(o.CookieNames) {
		return true
	}

	return false
}

// SetCookieNames gets a reference to the given ResponseBadRequestBaseEdgeConnectorConnectionPreference and assigns it to the CookieNames field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetCookieNames(v ResponseBadRequestBaseEdgeConnectorConnectionPreference) {
	o.CookieNames = &v
}

// GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetAdaptiveDeliveryAction() []string {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		var ret []string
		return ret
	}
	return o.AdaptiveDeliveryAction
}

// GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetAdaptiveDeliveryActionOk() ([]string, bool) {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		return nil, false
	}
	return o.AdaptiveDeliveryAction, true
}

// HasAdaptiveDeliveryAction returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasAdaptiveDeliveryAction() bool {
	if o != nil && !IsNil(o.AdaptiveDeliveryAction) {
		return true
	}

	return false
}

// SetAdaptiveDeliveryAction gets a reference to the given []string and assigns it to the AdaptiveDeliveryAction field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetAdaptiveDeliveryAction(v []string) {
	o.AdaptiveDeliveryAction = v
}

// GetDeviceGroup returns the DeviceGroup field value if set, zero value otherwise.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetDeviceGroup() []string {
	if o == nil || IsNil(o.DeviceGroup) {
		var ret []string
		return ret
	}
	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) GetDeviceGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceGroup) {
		return nil, false
	}
	return o.DeviceGroup, true
}

// HasDeviceGroup returns a boolean if a field has been set.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) HasDeviceGroup() bool {
	if o != nil && !IsNil(o.DeviceGroup) {
		return true
	}

	return false
}

// SetDeviceGroup gets a reference to the given []string and assigns it to the DeviceGroup field.
func (o *ResponseBadRequestSerializerMetaclassApplicationControlsField) SetDeviceGroup(v []string) {
	o.DeviceGroup = v
}

func (o ResponseBadRequestSerializerMetaclassApplicationControlsField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestSerializerMetaclassApplicationControlsField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheByQueryString) {
		toSerialize["cache_by_query_string"] = o.CacheByQueryString
	}
	if !IsNil(o.QueryStringFields) {
		toSerialize["query_string_fields"] = o.QueryStringFields
	}
	if !IsNil(o.QueryStringSortEnabled) {
		toSerialize["query_string_sort_enabled"] = o.QueryStringSortEnabled
	}
	if !IsNil(o.CacheByCookies) {
		toSerialize["cache_by_cookies"] = o.CacheByCookies
	}
	if !IsNil(o.CookieNames) {
		toSerialize["cookie_names"] = o.CookieNames
	}
	if !IsNil(o.AdaptiveDeliveryAction) {
		toSerialize["adaptive_delivery_action"] = o.AdaptiveDeliveryAction
	}
	if !IsNil(o.DeviceGroup) {
		toSerialize["device_group"] = o.DeviceGroup
	}
	return toSerialize, nil
}

type NullableResponseBadRequestSerializerMetaclassApplicationControlsField struct {
	value *ResponseBadRequestSerializerMetaclassApplicationControlsField
	isSet bool
}

func (v NullableResponseBadRequestSerializerMetaclassApplicationControlsField) Get() *ResponseBadRequestSerializerMetaclassApplicationControlsField {
	return v.value
}

func (v *NullableResponseBadRequestSerializerMetaclassApplicationControlsField) Set(val *ResponseBadRequestSerializerMetaclassApplicationControlsField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestSerializerMetaclassApplicationControlsField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestSerializerMetaclassApplicationControlsField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestSerializerMetaclassApplicationControlsField(val *ResponseBadRequestSerializerMetaclassApplicationControlsField) *NullableResponseBadRequestSerializerMetaclassApplicationControlsField {
	return &NullableResponseBadRequestSerializerMetaclassApplicationControlsField{value: val, isSet: true}
}

func (v NullableResponseBadRequestSerializerMetaclassApplicationControlsField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestSerializerMetaclassApplicationControlsField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


