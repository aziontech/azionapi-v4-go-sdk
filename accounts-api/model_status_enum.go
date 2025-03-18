/*
Accounts API

REST API OpenAPI documentation for the Accounts API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounts

import (
	"encoding/json"
)

// checks if the StatusEnum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusEnum{}

// StatusEnum * `active` - Active account status, can be used for regular operations. * `suspended` - Suspended account status, for accounts with limited access to support and payments only. * `disabled` - Disabled account status, services are offline, user can only access support. * `closed` - Closed account status, services are offline but can be reactivated.
type StatusEnum struct {
}

// NewStatusEnum instantiates a new StatusEnum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusEnum() *StatusEnum {
	this := StatusEnum{}
	return &this
}

// NewStatusEnumWithDefaults instantiates a new StatusEnum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusEnumWithDefaults() *StatusEnum {
	this := StatusEnum{}
	return &this
}

func (o StatusEnum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusEnum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableStatusEnum struct {
	value StatusEnum
	isSet bool
}

func (v NullableStatusEnum) Get() StatusEnum {
	return v.value
}

func (v *NullableStatusEnum) Set(val StatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusEnum) Unset() {
	v.isSet = false
}

func NewNullableStatusEnum(val StatusEnum) *NullableStatusEnum {
	return &NullableStatusEnum{value: val, isSet: true}
}

func (v NullableStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


