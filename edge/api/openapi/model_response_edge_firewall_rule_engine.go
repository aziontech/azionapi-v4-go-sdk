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

// checks if the ResponseEdgeFirewallRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEdgeFirewallRuleEngine{}

// ResponseEdgeFirewallRuleEngine struct for ResponseEdgeFirewallRuleEngine
type ResponseEdgeFirewallRuleEngine struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data EdgeFirewallRuleEngine `json:"data"`
}

type _ResponseEdgeFirewallRuleEngine ResponseEdgeFirewallRuleEngine

// NewResponseEdgeFirewallRuleEngine instantiates a new ResponseEdgeFirewallRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEdgeFirewallRuleEngine(state string, data EdgeFirewallRuleEngine) *ResponseEdgeFirewallRuleEngine {
	this := ResponseEdgeFirewallRuleEngine{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseEdgeFirewallRuleEngineWithDefaults instantiates a new ResponseEdgeFirewallRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEdgeFirewallRuleEngineWithDefaults() *ResponseEdgeFirewallRuleEngine {
	this := ResponseEdgeFirewallRuleEngine{}
	return &this
}

// GetState returns the State field value
func (o *ResponseEdgeFirewallRuleEngine) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFirewallRuleEngine) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseEdgeFirewallRuleEngine) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseEdgeFirewallRuleEngine) GetData() EdgeFirewallRuleEngine {
	if o == nil {
		var ret EdgeFirewallRuleEngine
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeFirewallRuleEngine) GetDataOk() (*EdgeFirewallRuleEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseEdgeFirewallRuleEngine) SetData(v EdgeFirewallRuleEngine) {
	o.Data = v
}

func (o ResponseEdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEdgeFirewallRuleEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseEdgeFirewallRuleEngine) UnmarshalJSON(data []byte) (err error) {
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

	varResponseEdgeFirewallRuleEngine := _ResponseEdgeFirewallRuleEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseEdgeFirewallRuleEngine)

	if err != nil {
		return err
	}

	*o = ResponseEdgeFirewallRuleEngine(varResponseEdgeFirewallRuleEngine)

	return err
}

type NullableResponseEdgeFirewallRuleEngine struct {
	value *ResponseEdgeFirewallRuleEngine
	isSet bool
}

func (v NullableResponseEdgeFirewallRuleEngine) Get() *ResponseEdgeFirewallRuleEngine {
	return v.value
}

func (v *NullableResponseEdgeFirewallRuleEngine) Set(val *ResponseEdgeFirewallRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEdgeFirewallRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEdgeFirewallRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEdgeFirewallRuleEngine(val *ResponseEdgeFirewallRuleEngine) *NullableResponseEdgeFirewallRuleEngine {
	return &NullableResponseEdgeFirewallRuleEngine{value: val, isSet: true}
}

func (v NullableResponseEdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEdgeFirewallRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


