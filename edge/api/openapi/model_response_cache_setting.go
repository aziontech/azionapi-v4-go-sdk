/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseCacheSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCacheSetting{}

// ResponseCacheSetting struct for ResponseCacheSetting
type ResponseCacheSetting struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data CacheSetting `json:"data"`
}

type _ResponseCacheSetting ResponseCacheSetting

// NewResponseCacheSetting instantiates a new ResponseCacheSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCacheSetting(state string, data CacheSetting) *ResponseCacheSetting {
	this := ResponseCacheSetting{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseCacheSettingWithDefaults instantiates a new ResponseCacheSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCacheSettingWithDefaults() *ResponseCacheSetting {
	this := ResponseCacheSetting{}
	return &this
}

// GetState returns the State field value
func (o *ResponseCacheSetting) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseCacheSetting) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseCacheSetting) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseCacheSetting) GetData() CacheSetting {
	if o == nil {
		var ret CacheSetting
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseCacheSetting) GetDataOk() (*CacheSetting, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseCacheSetting) SetData(v CacheSetting) {
	o.Data = v
}

func (o ResponseCacheSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCacheSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseCacheSetting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
		"data",
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

	varResponseCacheSetting := _ResponseCacheSetting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseCacheSetting)

	if err != nil {
		return err
	}

	*o = ResponseCacheSetting(varResponseCacheSetting)

	return err
}

type NullableResponseCacheSetting struct {
	value *ResponseCacheSetting
	isSet bool
}

func (v NullableResponseCacheSetting) Get() *ResponseCacheSetting {
	return v.value
}

func (v *NullableResponseCacheSetting) Set(val *ResponseCacheSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCacheSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCacheSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCacheSetting(val *ResponseCacheSetting) *NullableResponseCacheSetting {
	return &NullableResponseCacheSetting{value: val, isSet: true}
}

func (v NullableResponseCacheSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCacheSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


