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

// checks if the EdgeApplicationDeviceGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationDeviceGroups{}

// EdgeApplicationDeviceGroups struct for EdgeApplicationDeviceGroups
type EdgeApplicationDeviceGroups struct {
	Id int32 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	UserAgent string `json:"user_agent" validate:"regexp=.*"`
}

type _EdgeApplicationDeviceGroups EdgeApplicationDeviceGroups

// NewEdgeApplicationDeviceGroups instantiates a new EdgeApplicationDeviceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationDeviceGroups(id int32, name string, userAgent string) *EdgeApplicationDeviceGroups {
	this := EdgeApplicationDeviceGroups{}
	this.Id = id
	this.Name = name
	this.UserAgent = userAgent
	return &this
}

// NewEdgeApplicationDeviceGroupsWithDefaults instantiates a new EdgeApplicationDeviceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationDeviceGroupsWithDefaults() *EdgeApplicationDeviceGroups {
	this := EdgeApplicationDeviceGroups{}
	return &this
}

// GetId returns the Id field value
func (o *EdgeApplicationDeviceGroups) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationDeviceGroups) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeApplicationDeviceGroups) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EdgeApplicationDeviceGroups) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationDeviceGroups) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeApplicationDeviceGroups) SetName(v string) {
	o.Name = v
}

// GetUserAgent returns the UserAgent field value
func (o *EdgeApplicationDeviceGroups) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationDeviceGroups) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *EdgeApplicationDeviceGroups) SetUserAgent(v string) {
	o.UserAgent = v
}

func (o EdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationDeviceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["user_agent"] = o.UserAgent
	return toSerialize, nil
}

func (o *EdgeApplicationDeviceGroups) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"user_agent",
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

	varEdgeApplicationDeviceGroups := _EdgeApplicationDeviceGroups{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeApplicationDeviceGroups)

	if err != nil {
		return err
	}

	*o = EdgeApplicationDeviceGroups(varEdgeApplicationDeviceGroups)

	return err
}

type NullableEdgeApplicationDeviceGroups struct {
	value *EdgeApplicationDeviceGroups
	isSet bool
}

func (v NullableEdgeApplicationDeviceGroups) Get() *EdgeApplicationDeviceGroups {
	return v.value
}

func (v *NullableEdgeApplicationDeviceGroups) Set(val *EdgeApplicationDeviceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationDeviceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationDeviceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationDeviceGroups(val *EdgeApplicationDeviceGroups) *NullableEdgeApplicationDeviceGroups {
	return &NullableEdgeApplicationDeviceGroups{value: val, isSet: true}
}

func (v NullableEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationDeviceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


