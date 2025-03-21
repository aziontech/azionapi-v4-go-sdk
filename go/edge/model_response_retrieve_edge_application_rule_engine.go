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

// checks if the ResponseRetrieveEdgeApplicationRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveEdgeApplicationRuleEngine{}

// ResponseRetrieveEdgeApplicationRuleEngine struct for ResponseRetrieveEdgeApplicationRuleEngine
type ResponseRetrieveEdgeApplicationRuleEngine struct {
	Data EdgeApplicationRuleEngine `json:"data"`
}

type _ResponseRetrieveEdgeApplicationRuleEngine ResponseRetrieveEdgeApplicationRuleEngine

// NewResponseRetrieveEdgeApplicationRuleEngine instantiates a new ResponseRetrieveEdgeApplicationRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveEdgeApplicationRuleEngine(data EdgeApplicationRuleEngine) *ResponseRetrieveEdgeApplicationRuleEngine {
	this := ResponseRetrieveEdgeApplicationRuleEngine{}
	this.Data = data
	return &this
}

// NewResponseRetrieveEdgeApplicationRuleEngineWithDefaults instantiates a new ResponseRetrieveEdgeApplicationRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveEdgeApplicationRuleEngineWithDefaults() *ResponseRetrieveEdgeApplicationRuleEngine {
	this := ResponseRetrieveEdgeApplicationRuleEngine{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveEdgeApplicationRuleEngine) GetData() EdgeApplicationRuleEngine {
	if o == nil {
		var ret EdgeApplicationRuleEngine
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveEdgeApplicationRuleEngine) GetDataOk() (*EdgeApplicationRuleEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveEdgeApplicationRuleEngine) SetData(v EdgeApplicationRuleEngine) {
	o.Data = v
}

func (o ResponseRetrieveEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveEdgeApplicationRuleEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveEdgeApplicationRuleEngine) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveEdgeApplicationRuleEngine := _ResponseRetrieveEdgeApplicationRuleEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveEdgeApplicationRuleEngine)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveEdgeApplicationRuleEngine(varResponseRetrieveEdgeApplicationRuleEngine)

	return err
}

type NullableResponseRetrieveEdgeApplicationRuleEngine struct {
	value *ResponseRetrieveEdgeApplicationRuleEngine
	isSet bool
}

func (v NullableResponseRetrieveEdgeApplicationRuleEngine) Get() *ResponseRetrieveEdgeApplicationRuleEngine {
	return v.value
}

func (v *NullableResponseRetrieveEdgeApplicationRuleEngine) Set(val *ResponseRetrieveEdgeApplicationRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveEdgeApplicationRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveEdgeApplicationRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveEdgeApplicationRuleEngine(val *ResponseRetrieveEdgeApplicationRuleEngine) *NullableResponseRetrieveEdgeApplicationRuleEngine {
	return &NullableResponseRetrieveEdgeApplicationRuleEngine{value: val, isSet: true}
}

func (v NullableResponseRetrieveEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveEdgeApplicationRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


