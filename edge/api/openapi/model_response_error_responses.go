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

// checks if the ResponseErrorResponses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseErrorResponses{}

// ResponseErrorResponses struct for ResponseErrorResponses
type ResponseErrorResponses struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data ErrorResponses `json:"data"`
}

type _ResponseErrorResponses ResponseErrorResponses

// NewResponseErrorResponses instantiates a new ResponseErrorResponses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseErrorResponses(state string, data ErrorResponses) *ResponseErrorResponses {
	this := ResponseErrorResponses{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseErrorResponsesWithDefaults instantiates a new ResponseErrorResponses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseErrorResponsesWithDefaults() *ResponseErrorResponses {
	this := ResponseErrorResponses{}
	return &this
}

// GetState returns the State field value
func (o *ResponseErrorResponses) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseErrorResponses) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseErrorResponses) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
func (o *ResponseErrorResponses) GetData() ErrorResponses {
	if o == nil {
		var ret ErrorResponses
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseErrorResponses) GetDataOk() (*ErrorResponses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseErrorResponses) SetData(v ErrorResponses) {
	o.Data = v
}

func (o ResponseErrorResponses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseErrorResponses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseErrorResponses) UnmarshalJSON(data []byte) (err error) {
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

	varResponseErrorResponses := _ResponseErrorResponses{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseErrorResponses)

	if err != nil {
		return err
	}

	*o = ResponseErrorResponses(varResponseErrorResponses)

	return err
}

type NullableResponseErrorResponses struct {
	value *ResponseErrorResponses
	isSet bool
}

func (v NullableResponseErrorResponses) Get() *ResponseErrorResponses {
	return v.value
}

func (v *NullableResponseErrorResponses) Set(val *ResponseErrorResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseErrorResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseErrorResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseErrorResponses(val *ResponseErrorResponses) *NullableResponseErrorResponses {
	return &NullableResponseErrorResponses{value: val, isSet: true}
}

func (v NullableResponseErrorResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseErrorResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


