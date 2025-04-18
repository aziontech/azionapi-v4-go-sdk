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

// checks if the ResponseRetrieveBaseEdgeConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveBaseEdgeConnector{}

// ResponseRetrieveBaseEdgeConnector struct for ResponseRetrieveBaseEdgeConnector
type ResponseRetrieveBaseEdgeConnector struct {
	Data BaseEdgeConnector `json:"data"`
}

type _ResponseRetrieveBaseEdgeConnector ResponseRetrieveBaseEdgeConnector

// NewResponseRetrieveBaseEdgeConnector instantiates a new ResponseRetrieveBaseEdgeConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveBaseEdgeConnector(data BaseEdgeConnector) *ResponseRetrieveBaseEdgeConnector {
	this := ResponseRetrieveBaseEdgeConnector{}
	this.Data = data
	return &this
}

// NewResponseRetrieveBaseEdgeConnectorWithDefaults instantiates a new ResponseRetrieveBaseEdgeConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveBaseEdgeConnectorWithDefaults() *ResponseRetrieveBaseEdgeConnector {
	this := ResponseRetrieveBaseEdgeConnector{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveBaseEdgeConnector) GetData() BaseEdgeConnector {
	if o == nil {
		var ret BaseEdgeConnector
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveBaseEdgeConnector) GetDataOk() (*BaseEdgeConnector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveBaseEdgeConnector) SetData(v BaseEdgeConnector) {
	o.Data = v
}

func (o ResponseRetrieveBaseEdgeConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveBaseEdgeConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveBaseEdgeConnector) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveBaseEdgeConnector := _ResponseRetrieveBaseEdgeConnector{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveBaseEdgeConnector)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveBaseEdgeConnector(varResponseRetrieveBaseEdgeConnector)

	return err
}

type NullableResponseRetrieveBaseEdgeConnector struct {
	value *ResponseRetrieveBaseEdgeConnector
	isSet bool
}

func (v NullableResponseRetrieveBaseEdgeConnector) Get() *ResponseRetrieveBaseEdgeConnector {
	return v.value
}

func (v *NullableResponseRetrieveBaseEdgeConnector) Set(val *ResponseRetrieveBaseEdgeConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveBaseEdgeConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveBaseEdgeConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveBaseEdgeConnector(val *ResponseRetrieveBaseEdgeConnector) *NullableResponseRetrieveBaseEdgeConnector {
	return &NullableResponseRetrieveBaseEdgeConnector{value: val, isSet: true}
}

func (v NullableResponseRetrieveBaseEdgeConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveBaseEdgeConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


