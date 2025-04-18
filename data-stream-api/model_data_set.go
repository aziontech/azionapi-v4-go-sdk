/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the DataSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataSet{}

// DataSet struct for DataSet
type DataSet struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor"`
	LastModified time.Time `json:"last_modified"`
	Custom bool `json:"custom"`
	Active *bool `json:"active,omitempty"`
	DataSet string `json:"data_set"`
}

type _DataSet DataSet

// NewDataSet instantiates a new DataSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSet(id int64, name string, lastEditor string, lastModified time.Time, custom bool, dataSet string) *DataSet {
	this := DataSet{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.Custom = custom
	this.DataSet = dataSet
	return &this
}

// NewDataSetWithDefaults instantiates a new DataSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSetWithDefaults() *DataSet {
	this := DataSet{}
	return &this
}

// GetId returns the Id field value
func (o *DataSet) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataSet) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DataSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataSet) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *DataSet) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *DataSet) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *DataSet) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *DataSet) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetCustom returns the Custom field value
func (o *DataSet) GetCustom() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetCustomOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *DataSet) SetCustom(v bool) {
	o.Custom = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DataSet) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSet) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DataSet) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DataSet) SetActive(v bool) {
	o.Active = &v
}

// GetDataSet returns the DataSet field value
func (o *DataSet) GetDataSet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataSet
}

// GetDataSetOk returns a tuple with the DataSet field value
// and a boolean to check if the value has been set.
func (o *DataSet) GetDataSetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSet, true
}

// SetDataSet sets field value
func (o *DataSet) SetDataSet(v string) {
	o.DataSet = v
}

func (o DataSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["custom"] = o.Custom
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["data_set"] = o.DataSet
	return toSerialize, nil
}

func (o *DataSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"custom",
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

	varDataSet := _DataSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataSet)

	if err != nil {
		return err
	}

	*o = DataSet(varDataSet)

	return err
}

type NullableDataSet struct {
	value *DataSet
	isSet bool
}

func (v NullableDataSet) Get() *DataSet {
	return v.value
}

func (v *NullableDataSet) Set(val *DataSet) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSet) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSet(val *DataSet) *NullableDataSet {
	return &NullableDataSet{value: val, isSet: true}
}

func (v NullableDataSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


