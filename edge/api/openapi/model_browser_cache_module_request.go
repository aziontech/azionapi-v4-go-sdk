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

// checks if the BrowserCacheModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrowserCacheModuleRequest{}

// BrowserCacheModuleRequest struct for BrowserCacheModuleRequest
type BrowserCacheModuleRequest struct {
	// * `honor` - Honor Origin Cache Headers * `override` - Override Cache Settings * `no-cache` - No Cache
	Behavior string `json:"behavior"`
	MaxAge int64 `json:"max_age"`
}

type _BrowserCacheModuleRequest BrowserCacheModuleRequest

// NewBrowserCacheModuleRequest instantiates a new BrowserCacheModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserCacheModuleRequest(behavior string, maxAge int64) *BrowserCacheModuleRequest {
	this := BrowserCacheModuleRequest{}
	this.Behavior = behavior
	this.MaxAge = maxAge
	return &this
}

// NewBrowserCacheModuleRequestWithDefaults instantiates a new BrowserCacheModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserCacheModuleRequestWithDefaults() *BrowserCacheModuleRequest {
	this := BrowserCacheModuleRequest{}
	return &this
}

// GetBehavior returns the Behavior field value
func (o *BrowserCacheModuleRequest) GetBehavior() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Behavior
}

// GetBehaviorOk returns a tuple with the Behavior field value
// and a boolean to check if the value has been set.
func (o *BrowserCacheModuleRequest) GetBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Behavior, true
}

// SetBehavior sets field value
func (o *BrowserCacheModuleRequest) SetBehavior(v string) {
	o.Behavior = v
}

// GetMaxAge returns the MaxAge field value
func (o *BrowserCacheModuleRequest) GetMaxAge() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value
// and a boolean to check if the value has been set.
func (o *BrowserCacheModuleRequest) GetMaxAgeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAge, true
}

// SetMaxAge sets field value
func (o *BrowserCacheModuleRequest) SetMaxAge(v int64) {
	o.MaxAge = v
}

func (o BrowserCacheModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrowserCacheModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["behavior"] = o.Behavior
	toSerialize["max_age"] = o.MaxAge
	return toSerialize, nil
}

func (o *BrowserCacheModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"behavior",
		"max_age",
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

	varBrowserCacheModuleRequest := _BrowserCacheModuleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBrowserCacheModuleRequest)

	if err != nil {
		return err
	}

	*o = BrowserCacheModuleRequest(varBrowserCacheModuleRequest)

	return err
}

type NullableBrowserCacheModuleRequest struct {
	value *BrowserCacheModuleRequest
	isSet bool
}

func (v NullableBrowserCacheModuleRequest) Get() *BrowserCacheModuleRequest {
	return v.value
}

func (v *NullableBrowserCacheModuleRequest) Set(val *BrowserCacheModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserCacheModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserCacheModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserCacheModuleRequest(val *BrowserCacheModuleRequest) *NullableBrowserCacheModuleRequest {
	return &NullableBrowserCacheModuleRequest{value: val, isSet: true}
}

func (v NullableBrowserCacheModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserCacheModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


