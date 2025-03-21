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

// checks if the ResponseWAFRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseWAFRule{}

// ResponseWAFRule struct for ResponseWAFRule
type ResponseWAFRule struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data WAFRule `json:"data"`
}

type _ResponseWAFRule ResponseWAFRule

// NewResponseWAFRule instantiates a new ResponseWAFRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseWAFRule(state string, data WAFRule) *ResponseWAFRule {
	this := ResponseWAFRule{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseWAFRuleWithDefaults instantiates a new ResponseWAFRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWAFRuleWithDefaults() *ResponseWAFRule {
	this := ResponseWAFRule{}
	return &this
}

// GetState returns the State field value
func (o *ResponseWAFRule) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseWAFRule) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseWAFRule) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseWAFRule) GetData() WAFRule {
	if o == nil {
		var ret WAFRule
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseWAFRule) GetDataOk() (*WAFRule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseWAFRule) SetData(v WAFRule) {
	o.Data = v
}

func (o ResponseWAFRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseWAFRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseWAFRule) UnmarshalJSON(data []byte) (err error) {
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

	varResponseWAFRule := _ResponseWAFRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseWAFRule)

	if err != nil {
		return err
	}

	*o = ResponseWAFRule(varResponseWAFRule)

	return err
}

type NullableResponseWAFRule struct {
	value *ResponseWAFRule
	isSet bool
}

func (v NullableResponseWAFRule) Get() *ResponseWAFRule {
	return v.value
}

func (v *NullableResponseWAFRule) Set(val *ResponseWAFRule) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseWAFRule) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseWAFRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseWAFRule(val *ResponseWAFRule) *NullableResponseWAFRule {
	return &NullableResponseWAFRule{value: val, isSet: true}
}

func (v NullableResponseWAFRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseWAFRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


