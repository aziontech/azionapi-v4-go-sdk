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

// checks if the OperatorEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatorEnum{}

// OperatorEnum * `does_not_exist` - does_not_exist * `does_not_match` - does_not_match * `does_not_start_with` - does_not_start_with * `exists` - exists * `is_equal` - is_equal * `is_in_list` - is_in_list * `is_not_equal` - is_not_equal * `is_not_in_list` - is_not_in_list * `matches` - matches * `starts_with` - starts_with
type OperatorEnum struct {
}

// NewOperatorEnum instantiates a new OperatorEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorEnum() *OperatorEnum {
	this := OperatorEnum{}
	return &this
}

// NewOperatorEnumWithDefaults instantiates a new OperatorEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorEnumWithDefaults() *OperatorEnum {
	this := OperatorEnum{}
	return &this
}

func (o OperatorEnum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatorEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableOperatorEnum struct {
	value OperatorEnum
	isSet bool
}

func (v NullableOperatorEnum) Get() OperatorEnum {
	return v.value
}

func (v *NullableOperatorEnum) Set(val OperatorEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorEnum(val OperatorEnum) *NullableOperatorEnum {
	return &NullableOperatorEnum{value: val, isSet: true}
}

func (v NullableOperatorEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


