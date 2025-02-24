/*
edgesql-api

REST API OpenAPI documentation for the EdgeSQL API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgesql

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDatabase{}

// ResponseDatabase struct for ResponseDatabase
type ResponseDatabase struct {
	State StateEnum `json:"state"`
	Data Database `json:"data"`
}

type _ResponseDatabase ResponseDatabase

// NewResponseDatabase instantiates a new ResponseDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDatabase(state StateEnum, data Database) *ResponseDatabase {
	this := ResponseDatabase{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseDatabaseWithDefaults instantiates a new ResponseDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDatabaseWithDefaults() *ResponseDatabase {
	this := ResponseDatabase{}
	return &this
}

// GetState returns the State field value
func (o *ResponseDatabase) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseDatabase) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseDatabase) SetState(v StateEnum) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseDatabase) GetData() Database {
	if o == nil {
		var ret Database
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseDatabase) GetDataOk() (*Database, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseDatabase) SetData(v Database) {
	o.Data = v
}

func (o ResponseDatabase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseDatabase) UnmarshalJSON(data []byte) (err error) {
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

	varResponseDatabase := _ResponseDatabase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseDatabase)

	if err != nil {
		return err
	}

	*o = ResponseDatabase(varResponseDatabase)

	return err
}

type NullableResponseDatabase struct {
	value *ResponseDatabase
	isSet bool
}

func (v NullableResponseDatabase) Get() *ResponseDatabase {
	return v.value
}

func (v *NullableResponseDatabase) Set(val *ResponseDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDatabase(val *ResponseDatabase) *NullableResponseDatabase {
	return &NullableResponseDatabase{value: val, isSet: true}
}

func (v NullableResponseDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


