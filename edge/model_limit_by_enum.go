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

// checks if the LimitByEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitByEnum{}

// LimitByEnum * `client_ip` - client_ip * `global` - global
type LimitByEnum struct {
}

// NewLimitByEnum instantiates a new LimitByEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitByEnum() *LimitByEnum {
	this := LimitByEnum{}
	return &this
}

// NewLimitByEnumWithDefaults instantiates a new LimitByEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitByEnumWithDefaults() *LimitByEnum {
	this := LimitByEnum{}
	return &this
}

func (o LimitByEnum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitByEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableLimitByEnum struct {
	value LimitByEnum
	isSet bool
}

func (v NullableLimitByEnum) Get() LimitByEnum {
	return v.value
}

func (v *NullableLimitByEnum) Set(val LimitByEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitByEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitByEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitByEnum(val LimitByEnum) *NullableLimitByEnum {
	return &NullableLimitByEnum{value: val, isSet: true}
}

func (v NullableLimitByEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitByEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


