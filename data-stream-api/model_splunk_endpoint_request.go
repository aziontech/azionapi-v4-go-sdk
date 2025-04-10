/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SplunkEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SplunkEndpointRequest{}

// SplunkEndpointRequest struct for SplunkEndpointRequest
type SplunkEndpointRequest struct {
	Url string `json:"url"`
	ApiKey string `json:"api_key"`
}

type _SplunkEndpointRequest SplunkEndpointRequest

// NewSplunkEndpointRequest instantiates a new SplunkEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplunkEndpointRequest(url string, apiKey string) *SplunkEndpointRequest {
	this := SplunkEndpointRequest{}
	this.Url = url
	this.ApiKey = apiKey
	return &this
}

// NewSplunkEndpointRequestWithDefaults instantiates a new SplunkEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplunkEndpointRequestWithDefaults() *SplunkEndpointRequest {
	this := SplunkEndpointRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *SplunkEndpointRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SplunkEndpointRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SplunkEndpointRequest) SetUrl(v string) {
	o.Url = v
}

// GetApiKey returns the ApiKey field value
func (o *SplunkEndpointRequest) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *SplunkEndpointRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *SplunkEndpointRequest) SetApiKey(v string) {
	o.ApiKey = v
}

func (o SplunkEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplunkEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["api_key"] = o.ApiKey
	return toSerialize, nil
}

func (o *SplunkEndpointRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"api_key",
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

	varSplunkEndpointRequest := _SplunkEndpointRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSplunkEndpointRequest)

	if err != nil {
		return err
	}

	*o = SplunkEndpointRequest(varSplunkEndpointRequest)

	return err
}

type NullableSplunkEndpointRequest struct {
	value *SplunkEndpointRequest
	isSet bool
}

func (v NullableSplunkEndpointRequest) Get() *SplunkEndpointRequest {
	return v.value
}

func (v *NullableSplunkEndpointRequest) Set(val *SplunkEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSplunkEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSplunkEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplunkEndpointRequest(val *SplunkEndpointRequest) *NullableSplunkEndpointRequest {
	return &NullableSplunkEndpointRequest{value: val, isSet: true}
}

func (v NullableSplunkEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplunkEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


