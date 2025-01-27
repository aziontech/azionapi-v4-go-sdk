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

// checks if the RemoteFileInclusionSensitivityEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteFileInclusionSensitivityEnum{}

// RemoteFileInclusionSensitivityEnum * `lowest` - lowest * `low` - low * `medium` - medium * `high` - high * `highest` - highest
type RemoteFileInclusionSensitivityEnum struct {
}

// NewRemoteFileInclusionSensitivityEnum instantiates a new RemoteFileInclusionSensitivityEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteFileInclusionSensitivityEnum() *RemoteFileInclusionSensitivityEnum {
	this := RemoteFileInclusionSensitivityEnum{}
	return &this
}

// NewRemoteFileInclusionSensitivityEnumWithDefaults instantiates a new RemoteFileInclusionSensitivityEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteFileInclusionSensitivityEnumWithDefaults() *RemoteFileInclusionSensitivityEnum {
	this := RemoteFileInclusionSensitivityEnum{}
	return &this
}

func (o RemoteFileInclusionSensitivityEnum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteFileInclusionSensitivityEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableRemoteFileInclusionSensitivityEnum struct {
	value RemoteFileInclusionSensitivityEnum
	isSet bool
}

func (v NullableRemoteFileInclusionSensitivityEnum) Get() RemoteFileInclusionSensitivityEnum {
	return v.value
}

func (v *NullableRemoteFileInclusionSensitivityEnum) Set(val RemoteFileInclusionSensitivityEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteFileInclusionSensitivityEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteFileInclusionSensitivityEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteFileInclusionSensitivityEnum(val RemoteFileInclusionSensitivityEnum) *NullableRemoteFileInclusionSensitivityEnum {
	return &NullableRemoteFileInclusionSensitivityEnum{value: val, isSet: true}
}

func (v NullableRemoteFileInclusionSensitivityEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteFileInclusionSensitivityEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


