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

// checks if the ResponseBadRequestPurgeInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestPurgeInput{}

// ResponseBadRequestPurgeInput struct for ResponseBadRequestPurgeInput
type ResponseBadRequestPurgeInput struct {
	Items *ResponseBadRequestBaseEdgeConnectorConnectionPreference `json:"items,omitempty"`
	Layer []string `json:"layer,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestPurgeInput instantiates a new ResponseBadRequestPurgeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestPurgeInput() *ResponseBadRequestPurgeInput {
	this := ResponseBadRequestPurgeInput{}
	return &this
}

// NewResponseBadRequestPurgeInputWithDefaults instantiates a new ResponseBadRequestPurgeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestPurgeInputWithDefaults() *ResponseBadRequestPurgeInput {
	this := ResponseBadRequestPurgeInput{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ResponseBadRequestPurgeInput) GetItems() ResponseBadRequestBaseEdgeConnectorConnectionPreference {
	if o == nil || IsNil(o.Items) {
		var ret ResponseBadRequestBaseEdgeConnectorConnectionPreference
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestPurgeInput) GetItemsOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ResponseBadRequestPurgeInput) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given ResponseBadRequestBaseEdgeConnectorConnectionPreference and assigns it to the Items field.
func (o *ResponseBadRequestPurgeInput) SetItems(v ResponseBadRequestBaseEdgeConnectorConnectionPreference) {
	o.Items = &v
}

// GetLayer returns the Layer field value if set, zero value otherwise.
func (o *ResponseBadRequestPurgeInput) GetLayer() []string {
	if o == nil || IsNil(o.Layer) {
		var ret []string
		return ret
	}
	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestPurgeInput) GetLayerOk() ([]string, bool) {
	if o == nil || IsNil(o.Layer) {
		return nil, false
	}
	return o.Layer, true
}

// HasLayer returns a boolean if a field has been set.
func (o *ResponseBadRequestPurgeInput) HasLayer() bool {
	if o != nil && !IsNil(o.Layer) {
		return true
	}

	return false
}

// SetLayer gets a reference to the given []string and assigns it to the Layer field.
func (o *ResponseBadRequestPurgeInput) SetLayer(v []string) {
	o.Layer = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestPurgeInput) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestPurgeInput) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestPurgeInput) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestPurgeInput) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestPurgeInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestPurgeInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Layer) {
		toSerialize["layer"] = o.Layer
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestPurgeInput struct {
	value *ResponseBadRequestPurgeInput
	isSet bool
}

func (v NullableResponseBadRequestPurgeInput) Get() *ResponseBadRequestPurgeInput {
	return v.value
}

func (v *NullableResponseBadRequestPurgeInput) Set(val *ResponseBadRequestPurgeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestPurgeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestPurgeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestPurgeInput(val *ResponseBadRequestPurgeInput) *NullableResponseBadRequestPurgeInput {
	return &NullableResponseBadRequestPurgeInput{value: val, isSet: true}
}

func (v NullableResponseBadRequestPurgeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestPurgeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


