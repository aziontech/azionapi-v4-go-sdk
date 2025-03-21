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

// checks if the ResponseRetrieveEdgeApplicationDeviceGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveEdgeApplicationDeviceGroups{}

// ResponseRetrieveEdgeApplicationDeviceGroups struct for ResponseRetrieveEdgeApplicationDeviceGroups
type ResponseRetrieveEdgeApplicationDeviceGroups struct {
	Data EdgeApplicationDeviceGroups `json:"data"`
}

type _ResponseRetrieveEdgeApplicationDeviceGroups ResponseRetrieveEdgeApplicationDeviceGroups

// NewResponseRetrieveEdgeApplicationDeviceGroups instantiates a new ResponseRetrieveEdgeApplicationDeviceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveEdgeApplicationDeviceGroups(data EdgeApplicationDeviceGroups) *ResponseRetrieveEdgeApplicationDeviceGroups {
	this := ResponseRetrieveEdgeApplicationDeviceGroups{}
	this.Data = data
	return &this
}

// NewResponseRetrieveEdgeApplicationDeviceGroupsWithDefaults instantiates a new ResponseRetrieveEdgeApplicationDeviceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveEdgeApplicationDeviceGroupsWithDefaults() *ResponseRetrieveEdgeApplicationDeviceGroups {
	this := ResponseRetrieveEdgeApplicationDeviceGroups{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveEdgeApplicationDeviceGroups) GetData() EdgeApplicationDeviceGroups {
	if o == nil {
		var ret EdgeApplicationDeviceGroups
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveEdgeApplicationDeviceGroups) GetDataOk() (*EdgeApplicationDeviceGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveEdgeApplicationDeviceGroups) SetData(v EdgeApplicationDeviceGroups) {
	o.Data = v
}

func (o ResponseRetrieveEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveEdgeApplicationDeviceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveEdgeApplicationDeviceGroups) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveEdgeApplicationDeviceGroups := _ResponseRetrieveEdgeApplicationDeviceGroups{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveEdgeApplicationDeviceGroups)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveEdgeApplicationDeviceGroups(varResponseRetrieveEdgeApplicationDeviceGroups)

	return err
}

type NullableResponseRetrieveEdgeApplicationDeviceGroups struct {
	value *ResponseRetrieveEdgeApplicationDeviceGroups
	isSet bool
}

func (v NullableResponseRetrieveEdgeApplicationDeviceGroups) Get() *ResponseRetrieveEdgeApplicationDeviceGroups {
	return v.value
}

func (v *NullableResponseRetrieveEdgeApplicationDeviceGroups) Set(val *ResponseRetrieveEdgeApplicationDeviceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveEdgeApplicationDeviceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveEdgeApplicationDeviceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveEdgeApplicationDeviceGroups(val *ResponseRetrieveEdgeApplicationDeviceGroups) *NullableResponseRetrieveEdgeApplicationDeviceGroups {
	return &NullableResponseRetrieveEdgeApplicationDeviceGroups{value: val, isSet: true}
}

func (v NullableResponseRetrieveEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveEdgeApplicationDeviceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


