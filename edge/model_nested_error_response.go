/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NestedErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedErrorResponse{}

// NestedErrorResponse struct for NestedErrorResponse
type NestedErrorResponse struct {
	Code CodeEnum `json:"code"`
	Timeout int32 `json:"timeout"`
	Uri NullableString `json:"uri,omitempty" validate:"regexp=^\\/[\\/a-zA-Z0-9\\\\-_\\\\.\\\\~@:]*$"`
	CustomStatusCode NullableString `json:"custom_status_code,omitempty" validate:"regexp=^[1-5]\\\\d{2}$"`
}

type _NestedErrorResponse NestedErrorResponse

// NewNestedErrorResponse instantiates a new NestedErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedErrorResponse(code CodeEnum, timeout int32) *NestedErrorResponse {
	this := NestedErrorResponse{}
	this.Code = code
	this.Timeout = timeout
	return &this
}

// NewNestedErrorResponseWithDefaults instantiates a new NestedErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedErrorResponseWithDefaults() *NestedErrorResponse {
	this := NestedErrorResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *NestedErrorResponse) GetCode() CodeEnum {
	if o == nil {
		var ret CodeEnum
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *NestedErrorResponse) GetCodeOk() (*CodeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *NestedErrorResponse) SetCode(v CodeEnum) {
	o.Code = v
}

// GetTimeout returns the Timeout field value
func (o *NestedErrorResponse) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *NestedErrorResponse) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *NestedErrorResponse) SetTimeout(v int32) {
	o.Timeout = v
}

// GetUri returns the Uri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedErrorResponse) GetUri() string {
	if o == nil || IsNil(o.Uri.Get()) {
		var ret string
		return ret
	}
	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedErrorResponse) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// HasUri returns a boolean if a field has been set.
func (o *NestedErrorResponse) HasUri() bool {
	if o != nil && o.Uri.IsSet() {
		return true
	}

	return false
}

// SetUri gets a reference to the given NullableString and assigns it to the Uri field.
func (o *NestedErrorResponse) SetUri(v string) {
	o.Uri.Set(&v)
}
// SetUriNil sets the value for Uri to be an explicit nil
func (o *NestedErrorResponse) SetUriNil() {
	o.Uri.Set(nil)
}

// UnsetUri ensures that no value is present for Uri, not even an explicit nil
func (o *NestedErrorResponse) UnsetUri() {
	o.Uri.Unset()
}

// GetCustomStatusCode returns the CustomStatusCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedErrorResponse) GetCustomStatusCode() string {
	if o == nil || IsNil(o.CustomStatusCode.Get()) {
		var ret string
		return ret
	}
	return *o.CustomStatusCode.Get()
}

// GetCustomStatusCodeOk returns a tuple with the CustomStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedErrorResponse) GetCustomStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomStatusCode.Get(), o.CustomStatusCode.IsSet()
}

// HasCustomStatusCode returns a boolean if a field has been set.
func (o *NestedErrorResponse) HasCustomStatusCode() bool {
	if o != nil && o.CustomStatusCode.IsSet() {
		return true
	}

	return false
}

// SetCustomStatusCode gets a reference to the given NullableString and assigns it to the CustomStatusCode field.
func (o *NestedErrorResponse) SetCustomStatusCode(v string) {
	o.CustomStatusCode.Set(&v)
}
// SetCustomStatusCodeNil sets the value for CustomStatusCode to be an explicit nil
func (o *NestedErrorResponse) SetCustomStatusCodeNil() {
	o.CustomStatusCode.Set(nil)
}

// UnsetCustomStatusCode ensures that no value is present for CustomStatusCode, not even an explicit nil
func (o *NestedErrorResponse) UnsetCustomStatusCode() {
	o.CustomStatusCode.Unset()
}

func (o NestedErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["timeout"] = o.Timeout
	if o.Uri.IsSet() {
		toSerialize["uri"] = o.Uri.Get()
	}
	if o.CustomStatusCode.IsSet() {
		toSerialize["custom_status_code"] = o.CustomStatusCode.Get()
	}
	return toSerialize, nil
}

func (o *NestedErrorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"timeout",
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

	varNestedErrorResponse := _NestedErrorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNestedErrorResponse)

	if err != nil {
		return err
	}

	*o = NestedErrorResponse(varNestedErrorResponse)

	return err
}

type NullableNestedErrorResponse struct {
	value *NestedErrorResponse
	isSet bool
}

func (v NullableNestedErrorResponse) Get() *NestedErrorResponse {
	return v.value
}

func (v *NullableNestedErrorResponse) Set(val *NestedErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedErrorResponse(val *NestedErrorResponse) *NullableNestedErrorResponse {
	return &NullableNestedErrorResponse{value: val, isSet: true}
}

func (v NullableNestedErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


