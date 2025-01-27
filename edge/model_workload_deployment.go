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

// checks if the WorkloadDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadDeployment{}

// WorkloadDeployment struct for WorkloadDeployment
type WorkloadDeployment struct {
	Id int32 `json:"id"`
	Tag string `json:"tag" validate:"regexp=.*"`
	Binds WorkloadDeploymentBinds `json:"binds"`
	Current *bool `json:"current,omitempty"`
}

type _WorkloadDeployment WorkloadDeployment

// NewWorkloadDeployment instantiates a new WorkloadDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadDeployment(id int32, tag string, binds WorkloadDeploymentBinds) *WorkloadDeployment {
	this := WorkloadDeployment{}
	this.Id = id
	this.Tag = tag
	this.Binds = binds
	var current bool = true
	this.Current = &current
	return &this
}

// NewWorkloadDeploymentWithDefaults instantiates a new WorkloadDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadDeploymentWithDefaults() *WorkloadDeployment {
	this := WorkloadDeployment{}
	var current bool = true
	this.Current = &current
	return &this
}

// GetId returns the Id field value
func (o *WorkloadDeployment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkloadDeployment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkloadDeployment) SetId(v int32) {
	o.Id = v
}

// GetTag returns the Tag field value
func (o *WorkloadDeployment) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *WorkloadDeployment) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *WorkloadDeployment) SetTag(v string) {
	o.Tag = v
}

// GetBinds returns the Binds field value
func (o *WorkloadDeployment) GetBinds() WorkloadDeploymentBinds {
	if o == nil {
		var ret WorkloadDeploymentBinds
		return ret
	}

	return o.Binds
}

// GetBindsOk returns a tuple with the Binds field value
// and a boolean to check if the value has been set.
func (o *WorkloadDeployment) GetBindsOk() (*WorkloadDeploymentBinds, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Binds, true
}

// SetBinds sets field value
func (o *WorkloadDeployment) SetBinds(v WorkloadDeploymentBinds) {
	o.Binds = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *WorkloadDeployment) GetCurrent() bool {
	if o == nil || IsNil(o.Current) {
		var ret bool
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadDeployment) GetCurrentOk() (*bool, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *WorkloadDeployment) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given bool and assigns it to the Current field.
func (o *WorkloadDeployment) SetCurrent(v bool) {
	o.Current = &v
}

func (o WorkloadDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tag"] = o.Tag
	toSerialize["binds"] = o.Binds
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	return toSerialize, nil
}

func (o *WorkloadDeployment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tag",
		"binds",
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

	varWorkloadDeployment := _WorkloadDeployment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkloadDeployment)

	if err != nil {
		return err
	}

	*o = WorkloadDeployment(varWorkloadDeployment)

	return err
}

type NullableWorkloadDeployment struct {
	value *WorkloadDeployment
	isSet bool
}

func (v NullableWorkloadDeployment) Get() *WorkloadDeployment {
	return v.value
}

func (v *NullableWorkloadDeployment) Set(val *WorkloadDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadDeployment(val *WorkloadDeployment) *NullableWorkloadDeployment {
	return &NullableWorkloadDeployment{value: val, isSet: true}
}

func (v NullableWorkloadDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


