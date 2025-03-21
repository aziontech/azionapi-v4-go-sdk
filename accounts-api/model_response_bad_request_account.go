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

// checks if the ResponseBadRequestAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestAccount{}

// ResponseBadRequestAccount struct for ResponseBadRequestAccount
type ResponseBadRequestAccount struct {
	Type []string `json:"type,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestAccount instantiates a new ResponseBadRequestAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestAccount() *ResponseBadRequestAccount {
	this := ResponseBadRequestAccount{}
	return &this
}

// NewResponseBadRequestAccountWithDefaults instantiates a new ResponseBadRequestAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestAccountWithDefaults() *ResponseBadRequestAccount {
	this := ResponseBadRequestAccount{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponseBadRequestAccount) GetType() []string {
	if o == nil || IsNil(o.Type) {
		var ret []string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestAccount) GetTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponseBadRequestAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *ResponseBadRequestAccount) SetType(v []string) {
	o.Type = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestAccount) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestAccount) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestAccount) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestAccount) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestAccount struct {
	value *ResponseBadRequestAccount
	isSet bool
}

func (v NullableResponseBadRequestAccount) Get() *ResponseBadRequestAccount {
	return v.value
}

func (v *NullableResponseBadRequestAccount) Set(val *ResponseBadRequestAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestAccount(val *ResponseBadRequestAccount) *NullableResponseBadRequestAccount {
	return &NullableResponseBadRequestAccount{value: val, isSet: true}
}

func (v NullableResponseBadRequestAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


