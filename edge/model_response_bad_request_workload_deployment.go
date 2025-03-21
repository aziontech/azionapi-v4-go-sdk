/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
)

// checks if the ResponseBadRequestWorkloadDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestWorkloadDeployment{}

// ResponseBadRequestWorkloadDeployment struct for ResponseBadRequestWorkloadDeployment
type ResponseBadRequestWorkloadDeployment struct {
	Id []string `json:"id,omitempty"`
	Tag []string `json:"tag,omitempty"`
	Binds *ResponseBadRequestSerializerMetaclassBindsField `json:"binds,omitempty"`
	Current []string `json:"current,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestWorkloadDeployment instantiates a new ResponseBadRequestWorkloadDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestWorkloadDeployment() *ResponseBadRequestWorkloadDeployment {
	this := ResponseBadRequestWorkloadDeployment{}
	return &this
}

// NewResponseBadRequestWorkloadDeploymentWithDefaults instantiates a new ResponseBadRequestWorkloadDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestWorkloadDeploymentWithDefaults() *ResponseBadRequestWorkloadDeployment {
	this := ResponseBadRequestWorkloadDeployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseBadRequestWorkloadDeployment) GetId() []string {
	if o == nil || IsNil(o.Id) {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestWorkloadDeployment) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseBadRequestWorkloadDeployment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *ResponseBadRequestWorkloadDeployment) SetId(v []string) {
	o.Id = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ResponseBadRequestWorkloadDeployment) GetTag() []string {
	if o == nil || IsNil(o.Tag) {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestWorkloadDeployment) GetTagOk() ([]string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ResponseBadRequestWorkloadDeployment) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *ResponseBadRequestWorkloadDeployment) SetTag(v []string) {
	o.Tag = v
}

// GetBinds returns the Binds field value if set, zero value otherwise.
func (o *ResponseBadRequestWorkloadDeployment) GetBinds() ResponseBadRequestSerializerMetaclassBindsField {
	if o == nil || IsNil(o.Binds) {
		var ret ResponseBadRequestSerializerMetaclassBindsField
		return ret
	}
	return *o.Binds
}

// GetBindsOk returns a tuple with the Binds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestWorkloadDeployment) GetBindsOk() (*ResponseBadRequestSerializerMetaclassBindsField, bool) {
	if o == nil || IsNil(o.Binds) {
		return nil, false
	}
	return o.Binds, true
}

// HasBinds returns a boolean if a field has been set.
func (o *ResponseBadRequestWorkloadDeployment) HasBinds() bool {
	if o != nil && !IsNil(o.Binds) {
		return true
	}

	return false
}

// SetBinds gets a reference to the given ResponseBadRequestSerializerMetaclassBindsField and assigns it to the Binds field.
func (o *ResponseBadRequestWorkloadDeployment) SetBinds(v ResponseBadRequestSerializerMetaclassBindsField) {
	o.Binds = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *ResponseBadRequestWorkloadDeployment) GetCurrent() []string {
	if o == nil || IsNil(o.Current) {
		var ret []string
		return ret
	}
	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestWorkloadDeployment) GetCurrentOk() ([]string, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *ResponseBadRequestWorkloadDeployment) HasCurrent() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given []string and assigns it to the Current field.
func (o *ResponseBadRequestWorkloadDeployment) SetCurrent(v []string) {
	o.Current = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestWorkloadDeployment) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestWorkloadDeployment) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestWorkloadDeployment) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestWorkloadDeployment) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestWorkloadDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestWorkloadDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Binds) {
		toSerialize["binds"] = o.Binds
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestWorkloadDeployment struct {
	value *ResponseBadRequestWorkloadDeployment
	isSet bool
}

func (v NullableResponseBadRequestWorkloadDeployment) Get() *ResponseBadRequestWorkloadDeployment {
	return v.value
}

func (v *NullableResponseBadRequestWorkloadDeployment) Set(val *ResponseBadRequestWorkloadDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestWorkloadDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestWorkloadDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestWorkloadDeployment(val *ResponseBadRequestWorkloadDeployment) *NullableResponseBadRequestWorkloadDeployment {
	return &NullableResponseBadRequestWorkloadDeployment{value: val, isSet: true}
}

func (v NullableResponseBadRequestWorkloadDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestWorkloadDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


