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

// checks if the ResponseWorkload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseWorkload{}

// ResponseWorkload struct for ResponseWorkload
type ResponseWorkload struct {
	State StateEnum `json:"state"`
	Data Workload `json:"data"`
}

type _ResponseWorkload ResponseWorkload

// NewResponseWorkload instantiates a new ResponseWorkload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseWorkload(state StateEnum, data Workload) *ResponseWorkload {
	this := ResponseWorkload{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseWorkloadWithDefaults instantiates a new ResponseWorkload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWorkloadWithDefaults() *ResponseWorkload {
	this := ResponseWorkload{}
	return &this
}

// GetState returns the State field value
func (o *ResponseWorkload) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseWorkload) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseWorkload) SetState(v StateEnum) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseWorkload) GetData() Workload {
	if o == nil {
		var ret Workload
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseWorkload) GetDataOk() (*Workload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseWorkload) SetData(v Workload) {
	o.Data = v
}

func (o ResponseWorkload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseWorkload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseWorkload) UnmarshalJSON(data []byte) (err error) {
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

	varResponseWorkload := _ResponseWorkload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseWorkload)

	if err != nil {
		return err
	}

	*o = ResponseWorkload(varResponseWorkload)

	return err
}

type NullableResponseWorkload struct {
	value *ResponseWorkload
	isSet bool
}

func (v NullableResponseWorkload) Get() *ResponseWorkload {
	return v.value
}

func (v *NullableResponseWorkload) Set(val *ResponseWorkload) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseWorkload(val *ResponseWorkload) *NullableResponseWorkload {
	return &NullableResponseWorkload{value: val, isSet: true}
}

func (v NullableResponseWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


