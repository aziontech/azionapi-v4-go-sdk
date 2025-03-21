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

// checks if the ResponseRetrieveNetworkList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseRetrieveNetworkList{}

// ResponseRetrieveNetworkList struct for ResponseRetrieveNetworkList
type ResponseRetrieveNetworkList struct {
	Data NetworkList `json:"data"`
}

type _ResponseRetrieveNetworkList ResponseRetrieveNetworkList

// NewResponseRetrieveNetworkList instantiates a new ResponseRetrieveNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseRetrieveNetworkList(data NetworkList) *ResponseRetrieveNetworkList {
	this := ResponseRetrieveNetworkList{}
	this.Data = data
	return &this
}

// NewResponseRetrieveNetworkListWithDefaults instantiates a new ResponseRetrieveNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseRetrieveNetworkListWithDefaults() *ResponseRetrieveNetworkList {
	this := ResponseRetrieveNetworkList{}
	return &this
}

// GetData returns the Data field value
func (o *ResponseRetrieveNetworkList) GetData() NetworkList {
	if o == nil {
		var ret NetworkList
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResponseRetrieveNetworkList) GetDataOk() (*NetworkList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ResponseRetrieveNetworkList) SetData(v NetworkList) {
	o.Data = v
}

func (o ResponseRetrieveNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseRetrieveNetworkList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *ResponseRetrieveNetworkList) UnmarshalJSON(data []byte) (err error) {
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

	varResponseRetrieveNetworkList := _ResponseRetrieveNetworkList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseRetrieveNetworkList)

	if err != nil {
		return err
	}

	*o = ResponseRetrieveNetworkList(varResponseRetrieveNetworkList)

	return err
}

type NullableResponseRetrieveNetworkList struct {
	value *ResponseRetrieveNetworkList
	isSet bool
}

func (v NullableResponseRetrieveNetworkList) Get() *ResponseRetrieveNetworkList {
	return v.value
}

func (v *NullableResponseRetrieveNetworkList) Set(val *ResponseRetrieveNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseRetrieveNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseRetrieveNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseRetrieveNetworkList(val *ResponseRetrieveNetworkList) *NullableResponseRetrieveNetworkList {
	return &NullableResponseRetrieveNetworkList{value: val, isSet: true}
}

func (v NullableResponseRetrieveNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseRetrieveNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


