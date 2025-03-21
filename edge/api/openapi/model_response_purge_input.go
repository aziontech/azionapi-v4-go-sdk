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

// checks if the ResponsePurgeInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsePurgeInput{}

// ResponsePurgeInput struct for ResponsePurgeInput
type ResponsePurgeInput struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data PurgeInput `json:"data"`
}

type _ResponsePurgeInput ResponsePurgeInput

// NewResponsePurgeInput instantiates a new ResponsePurgeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePurgeInput(state string, data PurgeInput) *ResponsePurgeInput {
	this := ResponsePurgeInput{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponsePurgeInputWithDefaults instantiates a new ResponsePurgeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePurgeInputWithDefaults() *ResponsePurgeInput {
	this := ResponsePurgeInput{}
	return &this
}

// GetState returns the State field value
func (o *ResponsePurgeInput) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponsePurgeInput) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponsePurgeInput) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponsePurgeInput) GetData() PurgeInput {
	if o == nil {
		var ret PurgeInput
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponsePurgeInput) GetDataOk() (*PurgeInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponsePurgeInput) SetData(v PurgeInput) {
	o.Data = v
}

func (o ResponsePurgeInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsePurgeInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponsePurgeInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
		"data",
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

	varResponsePurgeInput := _ResponsePurgeInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponsePurgeInput)

	if err != nil {
		return err
	}

	*o = ResponsePurgeInput(varResponsePurgeInput)

	return err
}

type NullableResponsePurgeInput struct {
	value *ResponsePurgeInput
	isSet bool
}

func (v NullableResponsePurgeInput) Get() *ResponsePurgeInput {
	return v.value
}

func (v *NullableResponsePurgeInput) Set(val *ResponsePurgeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePurgeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePurgeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePurgeInput(val *ResponsePurgeInput) *NullableResponsePurgeInput {
	return &NullableResponsePurgeInput{value: val, isSet: true}
}

func (v NullableResponsePurgeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePurgeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


