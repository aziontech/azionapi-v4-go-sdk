/*
Accounts API

REST API OpenAPI documentation for the Accounts API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounts-api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResponseRetrieveAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveAccount{}

// ResponseRetrieveAccount struct for ResponseRetrieveAccount
type ResponseRetrieveAccount struct {
	Data Account `json:"data"`
}

type _ResponseRetrieveAccount ResponseRetrieveAccount

// NewResponseRetrieveAccount instantiates a new ResponseRetrieveAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveAccount(data Account) *ResponseRetrieveAccount {
	this := ResponseRetrieveAccount{}
	this.Data = data
	return &this
}

// NewResponseRetrieveAccountWithDefaults instantiates a new ResponseRetrieveAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveAccountWithDefaults() *ResponseRetrieveAccount {
	this := ResponseRetrieveAccount{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveAccount) GetData() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveAccount) GetDataOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveAccount) SetData(v Account) {
	o.Data = v
}

func (o ResponseRetrieveAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varResponseRetrieveAccount := _ResponseRetrieveAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveAccount)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveAccount(varResponseRetrieveAccount)

	return err
}

type NullableResponseRetrieveAccount struct {
	value *ResponseRetrieveAccount
	isSet bool
}

func (v NullableResponseRetrieveAccount) Get() *ResponseRetrieveAccount {
	return v.value
}

func (v *NullableResponseRetrieveAccount) Set(val *ResponseRetrieveAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveAccount(val *ResponseRetrieveAccount) *NullableResponseRetrieveAccount {
	return &NullableResponseRetrieveAccount{value: val, isSet: true}
}

func (v NullableResponseRetrieveAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


