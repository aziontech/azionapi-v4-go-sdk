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

// checks if the ResponseRetrieveWAF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveWAF{}

// ResponseRetrieveWAF struct for ResponseRetrieveWAF
type ResponseRetrieveWAF struct {
	Data WAF `json:"data"`
}

type _ResponseRetrieveWAF ResponseRetrieveWAF

// NewResponseRetrieveWAF instantiates a new ResponseRetrieveWAF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveWAF(data WAF) *ResponseRetrieveWAF {
	this := ResponseRetrieveWAF{}
	this.Data = data
	return &this
}

// NewResponseRetrieveWAFWithDefaults instantiates a new ResponseRetrieveWAF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveWAFWithDefaults() *ResponseRetrieveWAF {
	this := ResponseRetrieveWAF{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveWAF) GetData() WAF {
	if o == nil {
		var ret WAF
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveWAF) GetDataOk() (*WAF, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveWAF) SetData(v WAF) {
	o.Data = v
}

func (o ResponseRetrieveWAF) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveWAF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveWAF) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveWAF := _ResponseRetrieveWAF{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveWAF)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveWAF(varResponseRetrieveWAF)

	return err
}

type NullableResponseRetrieveWAF struct {
	value *ResponseRetrieveWAF
	isSet bool
}

func (v NullableResponseRetrieveWAF) Get() *ResponseRetrieveWAF {
	return v.value
}

func (v *NullableResponseRetrieveWAF) Set(val *ResponseRetrieveWAF) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveWAF) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveWAF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveWAF(val *ResponseRetrieveWAF) *NullableResponseRetrieveWAF {
	return &NullableResponseRetrieveWAF{value: val, isSet: true}
}

func (v NullableResponseRetrieveWAF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveWAF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


