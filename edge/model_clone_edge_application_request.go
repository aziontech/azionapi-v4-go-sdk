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

// checks if the CloneEdgeApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloneEdgeApplicationRequest{}

// CloneEdgeApplicationRequest struct for CloneEdgeApplicationRequest
type CloneEdgeApplicationRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
}

type _CloneEdgeApplicationRequest CloneEdgeApplicationRequest

// NewCloneEdgeApplicationRequest instantiates a new CloneEdgeApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloneEdgeApplicationRequest(name string) *CloneEdgeApplicationRequest {
	this := CloneEdgeApplicationRequest{}
	this.Name = name
	return &this
}

// NewCloneEdgeApplicationRequestWithDefaults instantiates a new CloneEdgeApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloneEdgeApplicationRequestWithDefaults() *CloneEdgeApplicationRequest {
	this := CloneEdgeApplicationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CloneEdgeApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloneEdgeApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloneEdgeApplicationRequest) SetName(v string) {
	o.Name = v
}

func (o CloneEdgeApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloneEdgeApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CloneEdgeApplicationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCloneEdgeApplicationRequest := _CloneEdgeApplicationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloneEdgeApplicationRequest)

	if err != nil {
		return err
	}

	*o = CloneEdgeApplicationRequest(varCloneEdgeApplicationRequest)

	return err
}

type NullableCloneEdgeApplicationRequest struct {
	value *CloneEdgeApplicationRequest
	isSet bool
}

func (v NullableCloneEdgeApplicationRequest) Get() *CloneEdgeApplicationRequest {
	return v.value
}

func (v *NullableCloneEdgeApplicationRequest) Set(val *CloneEdgeApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneEdgeApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneEdgeApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneEdgeApplicationRequest(val *CloneEdgeApplicationRequest) *NullableCloneEdgeApplicationRequest {
	return &NullableCloneEdgeApplicationRequest{value: val, isSet: true}
}

func (v NullableCloneEdgeApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneEdgeApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


