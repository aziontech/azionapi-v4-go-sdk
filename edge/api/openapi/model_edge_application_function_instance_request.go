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

// checks if the EdgeApplicationFunctionInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationFunctionInstanceRequest{}

// EdgeApplicationFunctionInstanceRequest struct for EdgeApplicationFunctionInstanceRequest
type EdgeApplicationFunctionInstanceRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	JsonArgs interface{} `json:"json_args"`
	EdgeFunction int64 `json:"edge_function"`
	Active *bool `json:"active,omitempty"`
}

type _EdgeApplicationFunctionInstanceRequest EdgeApplicationFunctionInstanceRequest

// NewEdgeApplicationFunctionInstanceRequest instantiates a new EdgeApplicationFunctionInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationFunctionInstanceRequest(name string, jsonArgs interface{}, edgeFunction int64) *EdgeApplicationFunctionInstanceRequest {
	this := EdgeApplicationFunctionInstanceRequest{}
	this.Name = name
	this.JsonArgs = jsonArgs
	this.EdgeFunction = edgeFunction
	return &this
}

// NewEdgeApplicationFunctionInstanceRequestWithDefaults instantiates a new EdgeApplicationFunctionInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationFunctionInstanceRequestWithDefaults() *EdgeApplicationFunctionInstanceRequest {
	this := EdgeApplicationFunctionInstanceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *EdgeApplicationFunctionInstanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeApplicationFunctionInstanceRequest) SetName(v string) {
	o.Name = v
}

// GetJsonArgs returns the JsonArgs field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EdgeApplicationFunctionInstanceRequest) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EdgeApplicationFunctionInstanceRequest) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// SetJsonArgs sets field value
func (o *EdgeApplicationFunctionInstanceRequest) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetEdgeFunction returns the EdgeFunction field value
func (o *EdgeApplicationFunctionInstanceRequest) GetEdgeFunction() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunction, true
}

// SetEdgeFunction sets field value
func (o *EdgeApplicationFunctionInstanceRequest) SetEdgeFunction(v int64) {
	o.EdgeFunction = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeApplicationFunctionInstanceRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstanceRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeApplicationFunctionInstanceRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeApplicationFunctionInstanceRequest) SetActive(v bool) {
	o.Active = &v
}

func (o EdgeApplicationFunctionInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationFunctionInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	toSerialize["edge_function"] = o.EdgeFunction
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

func (o *EdgeApplicationFunctionInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"json_args",
		"edge_function",
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

	varEdgeApplicationFunctionInstanceRequest := _EdgeApplicationFunctionInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeApplicationFunctionInstanceRequest)

	if err != nil {
		return err
	}

	*o = EdgeApplicationFunctionInstanceRequest(varEdgeApplicationFunctionInstanceRequest)

	return err
}

type NullableEdgeApplicationFunctionInstanceRequest struct {
	value *EdgeApplicationFunctionInstanceRequest
	isSet bool
}

func (v NullableEdgeApplicationFunctionInstanceRequest) Get() *EdgeApplicationFunctionInstanceRequest {
	return v.value
}

func (v *NullableEdgeApplicationFunctionInstanceRequest) Set(val *EdgeApplicationFunctionInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationFunctionInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationFunctionInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationFunctionInstanceRequest(val *EdgeApplicationFunctionInstanceRequest) *NullableEdgeApplicationFunctionInstanceRequest {
	return &NullableEdgeApplicationFunctionInstanceRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationFunctionInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationFunctionInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


