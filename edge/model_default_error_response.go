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

// checks if the DefaultErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultErrorResponse{}

// DefaultErrorResponse struct for DefaultErrorResponse
type DefaultErrorResponse struct {
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewDefaultErrorResponse instantiates a new DefaultErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultErrorResponse() *DefaultErrorResponse {
	this := DefaultErrorResponse{}
	return &this
}

// NewDefaultErrorResponseWithDefaults instantiates a new DefaultErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultErrorResponseWithDefaults() *DefaultErrorResponse {
	this := DefaultErrorResponse{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DefaultErrorResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponse) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DefaultErrorResponse) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DefaultErrorResponse) SetDetail(v string) {
	o.Detail = &v
}

func (o DefaultErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableDefaultErrorResponse struct {
	value *DefaultErrorResponse
	isSet bool
}

func (v NullableDefaultErrorResponse) Get() *DefaultErrorResponse {
	return v.value
}

func (v *NullableDefaultErrorResponse) Set(val *DefaultErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultErrorResponse(val *DefaultErrorResponse) *NullableDefaultErrorResponse {
	return &NullableDefaultErrorResponse{value: val, isSet: true}
}

func (v NullableDefaultErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


