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

// checks if the CaptureMatchGroupsArgumentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureMatchGroupsArgumentRequest{}

// CaptureMatchGroupsArgumentRequest struct for CaptureMatchGroupsArgumentRequest
type CaptureMatchGroupsArgumentRequest struct {
	Subject string `json:"subject" validate:"regexp=^(\\\\$\\\\{[a-z_]+\\\\})$"`
	Regex string `json:"regex" validate:"regexp=.*"`
	CapturedArray string `json:"captured_array" validate:"regexp=^[a-zA-Z][a-zA-Z_]{0,9}$"`
}

type _CaptureMatchGroupsArgumentRequest CaptureMatchGroupsArgumentRequest

// NewCaptureMatchGroupsArgumentRequest instantiates a new CaptureMatchGroupsArgumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureMatchGroupsArgumentRequest(subject string, regex string, capturedArray string) *CaptureMatchGroupsArgumentRequest {
	this := CaptureMatchGroupsArgumentRequest{}
	this.Subject = subject
	this.Regex = regex
	this.CapturedArray = capturedArray
	return &this
}

// NewCaptureMatchGroupsArgumentRequestWithDefaults instantiates a new CaptureMatchGroupsArgumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureMatchGroupsArgumentRequestWithDefaults() *CaptureMatchGroupsArgumentRequest {
	this := CaptureMatchGroupsArgumentRequest{}
	return &this
}

// GetSubject returns the Subject field value
func (o *CaptureMatchGroupsArgumentRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CaptureMatchGroupsArgumentRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CaptureMatchGroupsArgumentRequest) SetSubject(v string) {
	o.Subject = v
}

// GetRegex returns the Regex field value
func (o *CaptureMatchGroupsArgumentRequest) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *CaptureMatchGroupsArgumentRequest) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value
func (o *CaptureMatchGroupsArgumentRequest) SetRegex(v string) {
	o.Regex = v
}

// GetCapturedArray returns the CapturedArray field value
func (o *CaptureMatchGroupsArgumentRequest) GetCapturedArray() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CapturedArray
}

// GetCapturedArrayOk returns a tuple with the CapturedArray field value
// and a boolean to check if the value has been set.
func (o *CaptureMatchGroupsArgumentRequest) GetCapturedArrayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedArray, true
}

// SetCapturedArray sets field value
func (o *CaptureMatchGroupsArgumentRequest) SetCapturedArray(v string) {
	o.CapturedArray = v
}

func (o CaptureMatchGroupsArgumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureMatchGroupsArgumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subject"] = o.Subject
	toSerialize["regex"] = o.Regex
	toSerialize["captured_array"] = o.CapturedArray
	return toSerialize, nil
}

func (o *CaptureMatchGroupsArgumentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subject",
		"regex",
		"captured_array",
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

	varCaptureMatchGroupsArgumentRequest := _CaptureMatchGroupsArgumentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaptureMatchGroupsArgumentRequest)

	if err != nil {
		return err
	}

	*o = CaptureMatchGroupsArgumentRequest(varCaptureMatchGroupsArgumentRequest)

	return err
}

type NullableCaptureMatchGroupsArgumentRequest struct {
	value *CaptureMatchGroupsArgumentRequest
	isSet bool
}

func (v NullableCaptureMatchGroupsArgumentRequest) Get() *CaptureMatchGroupsArgumentRequest {
	return v.value
}

func (v *NullableCaptureMatchGroupsArgumentRequest) Set(val *CaptureMatchGroupsArgumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureMatchGroupsArgumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureMatchGroupsArgumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureMatchGroupsArgumentRequest(val *CaptureMatchGroupsArgumentRequest) *NullableCaptureMatchGroupsArgumentRequest {
	return &NullableCaptureMatchGroupsArgumentRequest{value: val, isSet: true}
}

func (v NullableCaptureMatchGroupsArgumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureMatchGroupsArgumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


