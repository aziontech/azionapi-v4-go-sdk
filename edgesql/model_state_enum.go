/*
edgesql-api

REST API OpenAPI documentation for the EdgeSQL API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgesql

import (
	"encoding/json"
)

// checks if the StateEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateEnum{}

// StateEnum struct for StateEnum
type StateEnum struct {
}

// NewStateEnum instantiates a new StateEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateEnum() *StateEnum {
	this := StateEnum{}
	return &this
}

// NewStateEnumWithDefaults instantiates a new StateEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateEnumWithDefaults() *StateEnum {
	this := StateEnum{}
	return &this
}

func (o StateEnum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableStateEnum struct {
	value StateEnum
	isSet bool
}

func (v NullableStateEnum) Get() StateEnum {
	return v.value
}

func (v *NullableStateEnum) Set(val StateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStateEnum) Unset() {
	v.isSet = false
}

func NewNullableStateEnum(val StateEnum) *NullableStateEnum {
	return &NullableStateEnum{value: val, isSet: true}
}

func (v NullableStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


