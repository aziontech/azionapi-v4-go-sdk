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

// checks if the ResponseListErrorResponses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListErrorResponses{}

// ResponseListErrorResponses struct for ResponseListErrorResponses
type ResponseListErrorResponses struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	EdgeApplicationId int64 `json:"edge_application_id"`
	OriginId NullableInt64 `json:"origin_id,omitempty"`
	ErrorResponses []NestedErrorResponse `json:"error_responses"`
}

type _ResponseListErrorResponses ResponseListErrorResponses

// NewResponseListErrorResponses instantiates a new ResponseListErrorResponses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListErrorResponses(id int64, name string, edgeApplicationId int64, errorResponses []NestedErrorResponse) *ResponseListErrorResponses {
	this := ResponseListErrorResponses{}
	this.Id = id
	this.Name = name
	this.EdgeApplicationId = edgeApplicationId
	this.ErrorResponses = errorResponses
	return &this
}

// NewResponseListErrorResponsesWithDefaults instantiates a new ResponseListErrorResponses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListErrorResponsesWithDefaults() *ResponseListErrorResponses {
	this := ResponseListErrorResponses{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListErrorResponses) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListErrorResponses) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListErrorResponses) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListErrorResponses) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListErrorResponses) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListErrorResponses) SetName(v string) {
	o.Name = v
}

// GetEdgeApplicationId returns the EdgeApplicationId field value
func (o *ResponseListErrorResponses) GetEdgeApplicationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeApplicationId
}

// GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field value
// and a boolean to check if the value has been set.
func (o *ResponseListErrorResponses) GetEdgeApplicationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeApplicationId, true
}

// SetEdgeApplicationId sets field value
func (o *ResponseListErrorResponses) SetEdgeApplicationId(v int64) {
	o.EdgeApplicationId = v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseListErrorResponses) GetOriginId() int64 {
	if o == nil || IsNil(o.OriginId.Get()) {
		var ret int64
		return ret
	}
	return *o.OriginId.Get()
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseListErrorResponses) GetOriginIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginId.Get(), o.OriginId.IsSet()
}

// HasOriginId returns a boolean if a field has been set.
func (o *ResponseListErrorResponses) HasOriginId() bool {
	if o != nil && o.OriginId.IsSet() {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given NullableInt64 and assigns it to the OriginId field.
func (o *ResponseListErrorResponses) SetOriginId(v int64) {
	o.OriginId.Set(&v)
}
// SetOriginIdNil sets the value for OriginId to be an explicit nil
func (o *ResponseListErrorResponses) SetOriginIdNil() {
	o.OriginId.Set(nil)
}

// UnsetOriginId ensures that no value is present for OriginId, not even an explicit nil
func (o *ResponseListErrorResponses) UnsetOriginId() {
	o.OriginId.Unset()
}

// GetErrorResponses returns the ErrorResponses field value
func (o *ResponseListErrorResponses) GetErrorResponses() []NestedErrorResponse {
	if o == nil {
		var ret []NestedErrorResponse
		return ret
	}

	return o.ErrorResponses
}

// GetErrorResponsesOk returns a tuple with the ErrorResponses field value
// and a boolean to check if the value has been set.
func (o *ResponseListErrorResponses) GetErrorResponsesOk() ([]NestedErrorResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorResponses, true
}

// SetErrorResponses sets field value
func (o *ResponseListErrorResponses) SetErrorResponses(v []NestedErrorResponse) {
	o.ErrorResponses = v
}

func (o ResponseListErrorResponses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListErrorResponses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["edge_application_id"] = o.EdgeApplicationId
	if o.OriginId.IsSet() {
		toSerialize["origin_id"] = o.OriginId.Get()
	}
	toSerialize["error_responses"] = o.ErrorResponses
	return toSerialize, nil
}

func (o *ResponseListErrorResponses) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"edge_application_id",
		"error_responses",
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

	varResponseListErrorResponses := _ResponseListErrorResponses{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListErrorResponses)

	if err != nil {
		return err
	}

	*o = ResponseListErrorResponses(varResponseListErrorResponses)

	return err
}

type NullableResponseListErrorResponses struct {
	value *ResponseListErrorResponses
	isSet bool
}

func (v NullableResponseListErrorResponses) Get() *ResponseListErrorResponses {
	return v.value
}

func (v *NullableResponseListErrorResponses) Set(val *ResponseListErrorResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListErrorResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListErrorResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListErrorResponses(val *ResponseListErrorResponses) *NullableResponseListErrorResponses {
	return &NullableResponseListErrorResponses{value: val, isSet: true}
}

func (v NullableResponseListErrorResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListErrorResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


