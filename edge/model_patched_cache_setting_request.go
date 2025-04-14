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

// checks if the PatchedCacheSettingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCacheSettingRequest{}

// PatchedCacheSettingRequest struct for PatchedCacheSettingRequest
type PatchedCacheSettingRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=^[a-zA-Z0-9 \\\\-\\\\.\\\\'\\\\,|]+$"`
	BrowserCache *BrowserCacheModuleRequest `json:"browser_cache,omitempty"`
	EdgeCache *EdgeCacheModuleRequest `json:"edge_cache,omitempty"`
	ApplicationControls *ApplicationControlsModuleRequest `json:"application_controls,omitempty"`
	SliceControls *SliceControlsModuleRequest `json:"slice_controls,omitempty"`
}

// NewPatchedCacheSettingRequest instantiates a new PatchedCacheSettingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCacheSettingRequest() *PatchedCacheSettingRequest {
	this := PatchedCacheSettingRequest{}
	return &this
}

// NewPatchedCacheSettingRequestWithDefaults instantiates a new PatchedCacheSettingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCacheSettingRequestWithDefaults() *PatchedCacheSettingRequest {
	this := PatchedCacheSettingRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCacheSettingRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCacheSettingRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCacheSettingRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCacheSettingRequest) SetName(v string) {
	o.Name = &v
}

// GetBrowserCache returns the BrowserCache field value if set, zero value otherwise.
func (o *PatchedCacheSettingRequest) GetBrowserCache() BrowserCacheModuleRequest {
	if o == nil || IsNil(o.BrowserCache) {
		var ret BrowserCacheModuleRequest
		return ret
	}
	return *o.BrowserCache
}

// GetBrowserCacheOk returns a tuple with the BrowserCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCacheSettingRequest) GetBrowserCacheOk() (*BrowserCacheModuleRequest, bool) {
	if o == nil || IsNil(o.BrowserCache) {
		return nil, false
	}
	return o.BrowserCache, true
}

// HasBrowserCache returns a boolean if a field has been set.
func (o *PatchedCacheSettingRequest) HasBrowserCache() bool {
	if o != nil && !IsNil(o.BrowserCache) {
		return true
	}

	return false
}

// SetBrowserCache gets a reference to the given BrowserCacheModuleRequest and assigns it to the BrowserCache field.
func (o *PatchedCacheSettingRequest) SetBrowserCache(v BrowserCacheModuleRequest) {
	o.BrowserCache = &v
}

// GetEdgeCache returns the EdgeCache field value if set, zero value otherwise.
func (o *PatchedCacheSettingRequest) GetEdgeCache() EdgeCacheModuleRequest {
	if o == nil || IsNil(o.EdgeCache) {
		var ret EdgeCacheModuleRequest
		return ret
	}
	return *o.EdgeCache
}

// GetEdgeCacheOk returns a tuple with the EdgeCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCacheSettingRequest) GetEdgeCacheOk() (*EdgeCacheModuleRequest, bool) {
	if o == nil || IsNil(o.EdgeCache) {
		return nil, false
	}
	return o.EdgeCache, true
}

// HasEdgeCache returns a boolean if a field has been set.
func (o *PatchedCacheSettingRequest) HasEdgeCache() bool {
	if o != nil && !IsNil(o.EdgeCache) {
		return true
	}

	return false
}

// SetEdgeCache gets a reference to the given EdgeCacheModuleRequest and assigns it to the EdgeCache field.
func (o *PatchedCacheSettingRequest) SetEdgeCache(v EdgeCacheModuleRequest) {
	o.EdgeCache = &v
}

// GetApplicationControls returns the ApplicationControls field value if set, zero value otherwise.
func (o *PatchedCacheSettingRequest) GetApplicationControls() ApplicationControlsModuleRequest {
	if o == nil || IsNil(o.ApplicationControls) {
		var ret ApplicationControlsModuleRequest
		return ret
	}
	return *o.ApplicationControls
}

// GetApplicationControlsOk returns a tuple with the ApplicationControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCacheSettingRequest) GetApplicationControlsOk() (*ApplicationControlsModuleRequest, bool) {
	if o == nil || IsNil(o.ApplicationControls) {
		return nil, false
	}
	return o.ApplicationControls, true
}

// HasApplicationControls returns a boolean if a field has been set.
func (o *PatchedCacheSettingRequest) HasApplicationControls() bool {
	if o != nil && !IsNil(o.ApplicationControls) {
		return true
	}

	return false
}

// SetApplicationControls gets a reference to the given ApplicationControlsModuleRequest and assigns it to the ApplicationControls field.
func (o *PatchedCacheSettingRequest) SetApplicationControls(v ApplicationControlsModuleRequest) {
	o.ApplicationControls = &v
}

// GetSliceControls returns the SliceControls field value if set, zero value otherwise.
func (o *PatchedCacheSettingRequest) GetSliceControls() SliceControlsModuleRequest {
	if o == nil || IsNil(o.SliceControls) {
		var ret SliceControlsModuleRequest
		return ret
	}
	return *o.SliceControls
}

// GetSliceControlsOk returns a tuple with the SliceControls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCacheSettingRequest) GetSliceControlsOk() (*SliceControlsModuleRequest, bool) {
	if o == nil || IsNil(o.SliceControls) {
		return nil, false
	}
	return o.SliceControls, true
}

// HasSliceControls returns a boolean if a field has been set.
func (o *PatchedCacheSettingRequest) HasSliceControls() bool {
	if o != nil && !IsNil(o.SliceControls) {
		return true
	}

	return false
}

// SetSliceControls gets a reference to the given SliceControlsModuleRequest and assigns it to the SliceControls field.
func (o *PatchedCacheSettingRequest) SetSliceControls(v SliceControlsModuleRequest) {
	o.SliceControls = &v
}

func (o PatchedCacheSettingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCacheSettingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BrowserCache) {
		toSerialize["browser_cache"] = o.BrowserCache
	}
	if !IsNil(o.EdgeCache) {
		toSerialize["edge_cache"] = o.EdgeCache
	}
	if !IsNil(o.ApplicationControls) {
		toSerialize["application_controls"] = o.ApplicationControls
	}
	if !IsNil(o.SliceControls) {
		toSerialize["slice_controls"] = o.SliceControls
	}
	return toSerialize, nil
}

type NullablePatchedCacheSettingRequest struct {
	value *PatchedCacheSettingRequest
	isSet bool
}

func (v NullablePatchedCacheSettingRequest) Get() *PatchedCacheSettingRequest {
	return v.value
}

func (v *NullablePatchedCacheSettingRequest) Set(val *PatchedCacheSettingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCacheSettingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCacheSettingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCacheSettingRequest(val *PatchedCacheSettingRequest) *NullablePatchedCacheSettingRequest {
	return &NullablePatchedCacheSettingRequest{value: val, isSet: true}
}

func (v NullablePatchedCacheSettingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCacheSettingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


