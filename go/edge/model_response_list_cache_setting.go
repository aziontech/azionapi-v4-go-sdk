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

// checks if the ResponseListCacheSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListCacheSetting{}

// ResponseListCacheSetting struct for ResponseListCacheSetting
type ResponseListCacheSetting struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9 \\\\-\\\\.\\\\'\\\\,]+$"`
	BrowserCache BrowserCacheModule `json:"browser_cache"`
	EdgeCache EdgeCacheModule `json:"edge_cache"`
	ApplicationControls ApplicationControlsModule `json:"application_controls"`
	SliceControls SliceControlsModule `json:"slice_controls"`
}

type _ResponseListCacheSetting ResponseListCacheSetting

// NewResponseListCacheSetting instantiates a new ResponseListCacheSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListCacheSetting(id int64, name string, browserCache BrowserCacheModule, edgeCache EdgeCacheModule, applicationControls ApplicationControlsModule, sliceControls SliceControlsModule) *ResponseListCacheSetting {
	this := ResponseListCacheSetting{}
	this.Id = id
	this.Name = name
	this.BrowserCache = browserCache
	this.EdgeCache = edgeCache
	this.ApplicationControls = applicationControls
	this.SliceControls = sliceControls
	return &this
}

// NewResponseListCacheSettingWithDefaults instantiates a new ResponseListCacheSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListCacheSettingWithDefaults() *ResponseListCacheSetting {
	this := ResponseListCacheSetting{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListCacheSetting) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListCacheSetting) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListCacheSetting) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListCacheSetting) SetName(v string) {
	o.Name = v
}

// GetBrowserCache returns the BrowserCache field value
func (o *ResponseListCacheSetting) GetBrowserCache() BrowserCacheModule {
	if o == nil {
		var ret BrowserCacheModule
		return ret
	}

	return o.BrowserCache
}

// GetBrowserCacheOk returns a tuple with the BrowserCache field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetBrowserCacheOk() (*BrowserCacheModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserCache, true
}

// SetBrowserCache sets field value
func (o *ResponseListCacheSetting) SetBrowserCache(v BrowserCacheModule) {
	o.BrowserCache = v
}

// GetEdgeCache returns the EdgeCache field value
func (o *ResponseListCacheSetting) GetEdgeCache() EdgeCacheModule {
	if o == nil {
		var ret EdgeCacheModule
		return ret
	}

	return o.EdgeCache
}

// GetEdgeCacheOk returns a tuple with the EdgeCache field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetEdgeCacheOk() (*EdgeCacheModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeCache, true
}

// SetEdgeCache sets field value
func (o *ResponseListCacheSetting) SetEdgeCache(v EdgeCacheModule) {
	o.EdgeCache = v
}

// GetApplicationControls returns the ApplicationControls field value
func (o *ResponseListCacheSetting) GetApplicationControls() ApplicationControlsModule {
	if o == nil {
		var ret ApplicationControlsModule
		return ret
	}

	return o.ApplicationControls
}

// GetApplicationControlsOk returns a tuple with the ApplicationControls field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetApplicationControlsOk() (*ApplicationControlsModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationControls, true
}

// SetApplicationControls sets field value
func (o *ResponseListCacheSetting) SetApplicationControls(v ApplicationControlsModule) {
	o.ApplicationControls = v
}

// GetSliceControls returns the SliceControls field value
func (o *ResponseListCacheSetting) GetSliceControls() SliceControlsModule {
	if o == nil {
		var ret SliceControlsModule
		return ret
	}

	return o.SliceControls
}

// GetSliceControlsOk returns a tuple with the SliceControls field value
// and a boolean to check if the value has been set.
func (o *ResponseListCacheSetting) GetSliceControlsOk() (*SliceControlsModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SliceControls, true
}

// SetSliceControls sets field value
func (o *ResponseListCacheSetting) SetSliceControls(v SliceControlsModule) {
	o.SliceControls = v
}

func (o ResponseListCacheSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListCacheSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["browser_cache"] = o.BrowserCache
	toSerialize["edge_cache"] = o.EdgeCache
	toSerialize["application_controls"] = o.ApplicationControls
	toSerialize["slice_controls"] = o.SliceControls
	return toSerialize, nil
}

func (o *ResponseListCacheSetting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"browser_cache",
		"edge_cache",
		"application_controls",
		"slice_controls",
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

	varResponseListCacheSetting := _ResponseListCacheSetting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListCacheSetting)

	if err != nil {
		return err
	}

	*o = ResponseListCacheSetting(varResponseListCacheSetting)

	return err
}

type NullableResponseListCacheSetting struct {
	value *ResponseListCacheSetting
	isSet bool
}

func (v NullableResponseListCacheSetting) Get() *ResponseListCacheSetting {
	return v.value
}

func (v *NullableResponseListCacheSetting) Set(val *ResponseListCacheSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListCacheSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListCacheSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListCacheSetting(val *ResponseListCacheSetting) *NullableResponseListCacheSetting {
	return &NullableResponseListCacheSetting{value: val, isSet: true}
}

func (v NullableResponseListCacheSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListCacheSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


