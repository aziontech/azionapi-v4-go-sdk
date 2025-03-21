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

// checks if the SetCustomResponseArgumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetCustomResponseArgumentRequest{}

// SetCustomResponseArgumentRequest struct for SetCustomResponseArgumentRequest
type SetCustomResponseArgumentRequest struct {
	StatusCode int64 `json:"status_code"`
	ContentType *string `json:"content_type,omitempty" validate:"regexp=.*"`
	ContentBody *string `json:"content_body,omitempty" validate:"regexp=.*"`
}

type _SetCustomResponseArgumentRequest SetCustomResponseArgumentRequest

// NewSetCustomResponseArgumentRequest instantiates a new SetCustomResponseArgumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCustomResponseArgumentRequest(statusCode int64) *SetCustomResponseArgumentRequest {
	this := SetCustomResponseArgumentRequest{}
	this.StatusCode = statusCode
	var contentType string = ""
	this.ContentType = &contentType
	var contentBody string = ""
	this.ContentBody = &contentBody
	return &this
}

// NewSetCustomResponseArgumentRequestWithDefaults instantiates a new SetCustomResponseArgumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCustomResponseArgumentRequestWithDefaults() *SetCustomResponseArgumentRequest {
	this := SetCustomResponseArgumentRequest{}
	var contentType string = ""
	this.ContentType = &contentType
	var contentBody string = ""
	this.ContentBody = &contentBody
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *SetCustomResponseArgumentRequest) GetStatusCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgumentRequest) GetStatusCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *SetCustomResponseArgumentRequest) SetStatusCode(v int64) {
	o.StatusCode = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SetCustomResponseArgumentRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgumentRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SetCustomResponseArgumentRequest) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SetCustomResponseArgumentRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetContentBody returns the ContentBody field value if set, zero value otherwise.
func (o *SetCustomResponseArgumentRequest) GetContentBody() string {
	if o == nil || IsNil(o.ContentBody) {
		var ret string
		return ret
	}
	return *o.ContentBody
}

// GetContentBodyOk returns a tuple with the ContentBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgumentRequest) GetContentBodyOk() (*string, bool) {
	if o == nil || IsNil(o.ContentBody) {
		return nil, false
	}
	return o.ContentBody, true
}

// HasContentBody returns a boolean if a field has been set.
func (o *SetCustomResponseArgumentRequest) HasContentBody() bool {
	if o != nil && !IsNil(o.ContentBody) {
		return true
	}

	return false
}

// SetContentBody gets a reference to the given string and assigns it to the ContentBody field.
func (o *SetCustomResponseArgumentRequest) SetContentBody(v string) {
	o.ContentBody = &v
}

func (o SetCustomResponseArgumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetCustomResponseArgumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status_code"] = o.StatusCode
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.ContentBody) {
		toSerialize["content_body"] = o.ContentBody
	}
	return toSerialize, nil
}

func (o *SetCustomResponseArgumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status_code",
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

	varSetCustomResponseArgumentRequest := _SetCustomResponseArgumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetCustomResponseArgumentRequest)

	if err != nil {
		return err
	}

	*o = SetCustomResponseArgumentRequest(varSetCustomResponseArgumentRequest)

	return err
}

type NullableSetCustomResponseArgumentRequest struct {
	value *SetCustomResponseArgumentRequest
	isSet bool
}

func (v NullableSetCustomResponseArgumentRequest) Get() *SetCustomResponseArgumentRequest {
	return v.value
}

func (v *NullableSetCustomResponseArgumentRequest) Set(val *SetCustomResponseArgumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCustomResponseArgumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCustomResponseArgumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCustomResponseArgumentRequest(val *SetCustomResponseArgumentRequest) *NullableSetCustomResponseArgumentRequest {
	return &NullableSetCustomResponseArgumentRequest{value: val, isSet: true}
}

func (v NullableSetCustomResponseArgumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCustomResponseArgumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


