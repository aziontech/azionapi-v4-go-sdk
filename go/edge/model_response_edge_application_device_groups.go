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

// checks if the ResponseEdgeApplicationDeviceGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseEdgeApplicationDeviceGroups{}

// ResponseEdgeApplicationDeviceGroups struct for ResponseEdgeApplicationDeviceGroups
type ResponseEdgeApplicationDeviceGroups struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data EdgeApplicationDeviceGroups `json:"data"`
}

type _ResponseEdgeApplicationDeviceGroups ResponseEdgeApplicationDeviceGroups

// NewResponseEdgeApplicationDeviceGroups instantiates a new ResponseEdgeApplicationDeviceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseEdgeApplicationDeviceGroups(state string, data EdgeApplicationDeviceGroups) *ResponseEdgeApplicationDeviceGroups {
	this := ResponseEdgeApplicationDeviceGroups{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseEdgeApplicationDeviceGroupsWithDefaults instantiates a new ResponseEdgeApplicationDeviceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseEdgeApplicationDeviceGroupsWithDefaults() *ResponseEdgeApplicationDeviceGroups {
	this := ResponseEdgeApplicationDeviceGroups{}
	return &this
}

// GetState returns the State field value
func (o *ResponseEdgeApplicationDeviceGroups) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeApplicationDeviceGroups) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseEdgeApplicationDeviceGroups) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseEdgeApplicationDeviceGroups) GetData() EdgeApplicationDeviceGroups {
	if o == nil {
		var ret EdgeApplicationDeviceGroups
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseEdgeApplicationDeviceGroups) GetDataOk() (*EdgeApplicationDeviceGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseEdgeApplicationDeviceGroups) SetData(v EdgeApplicationDeviceGroups) {
	o.Data = v
}

func (o ResponseEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseEdgeApplicationDeviceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseEdgeApplicationDeviceGroups) UnmarshalJSON(data []byte) (err error) {
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

	varResponseEdgeApplicationDeviceGroups := _ResponseEdgeApplicationDeviceGroups{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseEdgeApplicationDeviceGroups)

	if err != nil {
		return err
	}

	*o = ResponseEdgeApplicationDeviceGroups(varResponseEdgeApplicationDeviceGroups)

	return err
}

type NullableResponseEdgeApplicationDeviceGroups struct {
	value *ResponseEdgeApplicationDeviceGroups
	isSet bool
}

func (v NullableResponseEdgeApplicationDeviceGroups) Get() *ResponseEdgeApplicationDeviceGroups {
	return v.value
}

func (v *NullableResponseEdgeApplicationDeviceGroups) Set(val *ResponseEdgeApplicationDeviceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseEdgeApplicationDeviceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseEdgeApplicationDeviceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseEdgeApplicationDeviceGroups(val *ResponseEdgeApplicationDeviceGroups) *NullableResponseEdgeApplicationDeviceGroups {
	return &NullableResponseEdgeApplicationDeviceGroups{value: val, isSet: true}
}

func (v NullableResponseEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseEdgeApplicationDeviceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


