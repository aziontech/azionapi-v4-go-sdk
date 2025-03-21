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

// checks if the ResponseRetrieveEdgeFunctions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveEdgeFunctions{}

// ResponseRetrieveEdgeFunctions struct for ResponseRetrieveEdgeFunctions
type ResponseRetrieveEdgeFunctions struct {
	Data EdgeFunctions `json:"data"`
}

type _ResponseRetrieveEdgeFunctions ResponseRetrieveEdgeFunctions

// NewResponseRetrieveEdgeFunctions instantiates a new ResponseRetrieveEdgeFunctions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveEdgeFunctions(data EdgeFunctions) *ResponseRetrieveEdgeFunctions {
	this := ResponseRetrieveEdgeFunctions{}
	this.Data = data
	return &this
}

// NewResponseRetrieveEdgeFunctionsWithDefaults instantiates a new ResponseRetrieveEdgeFunctions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveEdgeFunctionsWithDefaults() *ResponseRetrieveEdgeFunctions {
	this := ResponseRetrieveEdgeFunctions{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveEdgeFunctions) GetData() EdgeFunctions {
	if o == nil {
		var ret EdgeFunctions
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveEdgeFunctions) GetDataOk() (*EdgeFunctions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveEdgeFunctions) SetData(v EdgeFunctions) {
	o.Data = v
}

func (o ResponseRetrieveEdgeFunctions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveEdgeFunctions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveEdgeFunctions) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveEdgeFunctions := _ResponseRetrieveEdgeFunctions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveEdgeFunctions)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveEdgeFunctions(varResponseRetrieveEdgeFunctions)

	return err
}

type NullableResponseRetrieveEdgeFunctions struct {
	value *ResponseRetrieveEdgeFunctions
	isSet bool
}

func (v NullableResponseRetrieveEdgeFunctions) Get() *ResponseRetrieveEdgeFunctions {
	return v.value
}

func (v *NullableResponseRetrieveEdgeFunctions) Set(val *ResponseRetrieveEdgeFunctions) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveEdgeFunctions) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveEdgeFunctions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveEdgeFunctions(val *ResponseRetrieveEdgeFunctions) *NullableResponseRetrieveEdgeFunctions {
	return &NullableResponseRetrieveEdgeFunctions{value: val, isSet: true}
}

func (v NullableResponseRetrieveEdgeFunctions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveEdgeFunctions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


