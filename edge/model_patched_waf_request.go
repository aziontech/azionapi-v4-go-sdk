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

// checks if the PatchedWAFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWAFRequest{}

// PatchedWAFRequest struct for PatchedWAFRequest
type PatchedWAFRequest struct {
	Active *bool `json:"active,omitempty"`
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	ThreatsConfiguration *WAFThreatsConfigurationRequest `json:"threats_configuration,omitempty"`
}

// NewPatchedWAFRequest instantiates a new PatchedWAFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWAFRequest() *PatchedWAFRequest {
	this := PatchedWAFRequest{}
	return &this
}

// NewPatchedWAFRequestWithDefaults instantiates a new PatchedWAFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWAFRequestWithDefaults() *PatchedWAFRequest {
	this := PatchedWAFRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedWAFRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedWAFRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedWAFRequest) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWAFRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWAFRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWAFRequest) SetName(v string) {
	o.Name = &v
}

// GetThreatsConfiguration returns the ThreatsConfiguration field value if set, zero value otherwise.
func (o *PatchedWAFRequest) GetThreatsConfiguration() WAFThreatsConfigurationRequest {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		var ret WAFThreatsConfigurationRequest
		return ret
	}
	return *o.ThreatsConfiguration
}

// GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWAFRequest) GetThreatsConfigurationOk() (*WAFThreatsConfigurationRequest, bool) {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		return nil, false
	}
	return o.ThreatsConfiguration, true
}

// HasThreatsConfiguration returns a boolean if a field has been set.
func (o *PatchedWAFRequest) HasThreatsConfiguration() bool {
	if o != nil && !IsNil(o.ThreatsConfiguration) {
		return true
	}

	return false
}

// SetThreatsConfiguration gets a reference to the given WAFThreatsConfigurationRequest and assigns it to the ThreatsConfiguration field.
func (o *PatchedWAFRequest) SetThreatsConfiguration(v WAFThreatsConfigurationRequest) {
	o.ThreatsConfiguration = &v
}

func (o PatchedWAFRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWAFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ThreatsConfiguration) {
		toSerialize["threats_configuration"] = o.ThreatsConfiguration
	}
	return toSerialize, nil
}

type NullablePatchedWAFRequest struct {
	value *PatchedWAFRequest
	isSet bool
}

func (v NullablePatchedWAFRequest) Get() *PatchedWAFRequest {
	return v.value
}

func (v *NullablePatchedWAFRequest) Set(val *PatchedWAFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWAFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWAFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWAFRequest(val *PatchedWAFRequest) *NullablePatchedWAFRequest {
	return &NullablePatchedWAFRequest{value: val, isSet: true}
}

func (v NullablePatchedWAFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWAFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


