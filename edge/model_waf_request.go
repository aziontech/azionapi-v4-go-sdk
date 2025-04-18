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

// checks if the WAFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFRequest{}

// WAFRequest struct for WAFRequest
type WAFRequest struct {
	Active *bool `json:"active,omitempty"`
	Name string `json:"name" validate:"regexp=.*"`
	ThreatsConfiguration *WAFThreatsConfigurationRequest `json:"threats_configuration,omitempty"`
}

type _WAFRequest WAFRequest

// NewWAFRequest instantiates a new WAFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFRequest(name string) *WAFRequest {
	this := WAFRequest{}
	this.Name = name
	return &this
}

// NewWAFRequestWithDefaults instantiates a new WAFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFRequestWithDefaults() *WAFRequest {
	this := WAFRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WAFRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WAFRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WAFRequest) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value
func (o *WAFRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WAFRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WAFRequest) SetName(v string) {
	o.Name = v
}

// GetThreatsConfiguration returns the ThreatsConfiguration field value if set, zero value otherwise.
func (o *WAFRequest) GetThreatsConfiguration() WAFThreatsConfigurationRequest {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		var ret WAFThreatsConfigurationRequest
		return ret
	}
	return *o.ThreatsConfiguration
}

// GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFRequest) GetThreatsConfigurationOk() (*WAFThreatsConfigurationRequest, bool) {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		return nil, false
	}
	return o.ThreatsConfiguration, true
}

// HasThreatsConfiguration returns a boolean if a field has been set.
func (o *WAFRequest) HasThreatsConfiguration() bool {
	if o != nil && !IsNil(o.ThreatsConfiguration) {
		return true
	}

	return false
}

// SetThreatsConfiguration gets a reference to the given WAFThreatsConfigurationRequest and assigns it to the ThreatsConfiguration field.
func (o *WAFRequest) SetThreatsConfiguration(v WAFThreatsConfigurationRequest) {
	o.ThreatsConfiguration = &v
}

func (o WAFRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ThreatsConfiguration) {
		toSerialize["threats_configuration"] = o.ThreatsConfiguration
	}
	return toSerialize, nil
}

func (o *WAFRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varWAFRequest := _WAFRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWAFRequest)

	if err != nil {
		return err
	}

	*o = WAFRequest(varWAFRequest)

	return err
}

type NullableWAFRequest struct {
	value *WAFRequest
	isSet bool
}

func (v NullableWAFRequest) Get() *WAFRequest {
	return v.value
}

func (v *NullableWAFRequest) Set(val *WAFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFRequest(val *WAFRequest) *NullableWAFRequest {
	return &NullableWAFRequest{value: val, isSet: true}
}

func (v NullableWAFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


