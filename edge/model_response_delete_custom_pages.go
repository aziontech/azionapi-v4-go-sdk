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

// checks if the ResponseDeleteCustomPages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDeleteCustomPages{}

// ResponseDeleteCustomPages struct for ResponseDeleteCustomPages
type ResponseDeleteCustomPages struct {
	// * `pending` - pending * `executed` - executed
	State string `json:"state"`
	Data NullableCustomPages `json:"data"`
}

type _ResponseDeleteCustomPages ResponseDeleteCustomPages

// NewResponseDeleteCustomPages instantiates a new ResponseDeleteCustomPages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDeleteCustomPages(state string, data NullableCustomPages) *ResponseDeleteCustomPages {
	this := ResponseDeleteCustomPages{}
	this.State = state
	this.Data = data
	return &this
}

// NewResponseDeleteCustomPagesWithDefaults instantiates a new ResponseDeleteCustomPages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDeleteCustomPagesWithDefaults() *ResponseDeleteCustomPages {
	this := ResponseDeleteCustomPages{}
	return &this
}

// GetState returns the State field value
func (o *ResponseDeleteCustomPages) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ResponseDeleteCustomPages) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ResponseDeleteCustomPages) SetState(v string) {
	o.State = v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for CustomPages will be returned
func (o *ResponseDeleteCustomPages) GetData() CustomPages {
	if o == nil || o.Data.Get() == nil {
		var ret CustomPages
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseDeleteCustomPages) GetDataOk() (*CustomPages, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *ResponseDeleteCustomPages) SetData(v CustomPages) {
	o.Data.Set(&v)
}

func (o ResponseDeleteCustomPages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDeleteCustomPages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data.Get()
	return toSerialize, nil
}

func (o *ResponseDeleteCustomPages) UnmarshalJSON(data []byte) (err error) {
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

	varResponseDeleteCustomPages := _ResponseDeleteCustomPages{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseDeleteCustomPages)

	if err != nil {
		return err
	}

	*o = ResponseDeleteCustomPages(varResponseDeleteCustomPages)

	return err
}

type NullableResponseDeleteCustomPages struct {
	value *ResponseDeleteCustomPages
	isSet bool
}

func (v NullableResponseDeleteCustomPages) Get() *ResponseDeleteCustomPages {
	return v.value
}

func (v *NullableResponseDeleteCustomPages) Set(val *ResponseDeleteCustomPages) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDeleteCustomPages) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDeleteCustomPages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDeleteCustomPages(val *ResponseDeleteCustomPages) *NullableResponseDeleteCustomPages {
	return &NullableResponseDeleteCustomPages{value: val, isSet: true}
}

func (v NullableResponseDeleteCustomPages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDeleteCustomPages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


