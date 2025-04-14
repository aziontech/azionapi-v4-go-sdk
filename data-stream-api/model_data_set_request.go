/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DataSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSetRequest{}

// DataSetRequest struct for DataSetRequest
type DataSetRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	Active *bool `json:"active,omitempty"`
	DataSet string `json:"data_set"`
}

type _DataSetRequest DataSetRequest

// NewDataSetRequest instantiates a new DataSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSetRequest(name string, dataSet string) *DataSetRequest {
	this := DataSetRequest{}
	this.Name = name
	this.DataSet = dataSet
	return &this
}

// NewDataSetRequestWithDefaults instantiates a new DataSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSetRequestWithDefaults() *DataSetRequest {
	this := DataSetRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DataSetRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSetRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataSetRequest) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DataSetRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSetRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DataSetRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DataSetRequest) SetActive(v bool) {
	o.Active = &v
}

// GetDataSet returns the DataSet field value
func (o *DataSetRequest) GetDataSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSet
}

// GetDataSetOk returns a tuple with the DataSet field value
// and a boolean to check if the value has been set.
func (o *DataSetRequest) GetDataSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSet, true
}

// SetDataSet sets field value
func (o *DataSetRequest) SetDataSet(v string) {
	o.DataSet = v
}

func (o DataSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["data_set"] = o.DataSet
	return toSerialize, nil
}

func (o *DataSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"data_set",
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

	varDataSetRequest := _DataSetRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSetRequest)

	if err != nil {
		return err
	}

	*o = DataSetRequest(varDataSetRequest)

	return err
}

type NullableDataSetRequest struct {
	value *DataSetRequest
	isSet bool
}

func (v NullableDataSetRequest) Get() *DataSetRequest {
	return v.value
}

func (v *NullableDataSetRequest) Set(val *DataSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetRequest(val *DataSetRequest) *NullableDataSetRequest {
	return &NullableDataSetRequest{value: val, isSet: true}
}

func (v NullableDataSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


