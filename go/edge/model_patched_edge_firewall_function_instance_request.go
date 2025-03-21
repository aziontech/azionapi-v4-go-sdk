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

// checks if the PatchedEdgeFirewallFunctionInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedEdgeFirewallFunctionInstanceRequest{}

// PatchedEdgeFirewallFunctionInstanceRequest struct for PatchedEdgeFirewallFunctionInstanceRequest
type PatchedEdgeFirewallFunctionInstanceRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	JsonArgs interface{} `json:"json_args,omitempty"`
	EdgeFunction *int64 `json:"edge_function,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewPatchedEdgeFirewallFunctionInstanceRequest instantiates a new PatchedEdgeFirewallFunctionInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEdgeFirewallFunctionInstanceRequest() *PatchedEdgeFirewallFunctionInstanceRequest {
	this := PatchedEdgeFirewallFunctionInstanceRequest{}
	return &this
}

// NewPatchedEdgeFirewallFunctionInstanceRequestWithDefaults instantiates a new PatchedEdgeFirewallFunctionInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEdgeFirewallFunctionInstanceRequestWithDefaults() *PatchedEdgeFirewallFunctionInstanceRequest {
	this := PatchedEdgeFirewallFunctionInstanceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetName(v string) {
	o.Name = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given interface{} and assigns it to the JsonArgs field.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetEdgeFunction returns the EdgeFunction field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetEdgeFunction() int64 {
	if o == nil || IsNil(o.EdgeFunction) {
		var ret int64
		return ret
	}
	return *o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetEdgeFunctionOk() (*int64, bool) {
	if o == nil || IsNil(o.EdgeFunction) {
		return nil, false
	}
	return o.EdgeFunction, true
}

// HasEdgeFunction returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasEdgeFunction() bool {
	if o != nil && !IsNil(o.EdgeFunction) {
		return true
	}

	return false
}

// SetEdgeFunction gets a reference to the given int64 and assigns it to the EdgeFunction field.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetEdgeFunction(v int64) {
	o.EdgeFunction = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedEdgeFirewallFunctionInstanceRequest) SetActive(v bool) {
	o.Active = &v
}

func (o PatchedEdgeFirewallFunctionInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedEdgeFirewallFunctionInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.EdgeFunction) {
		toSerialize["edge_function"] = o.EdgeFunction
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullablePatchedEdgeFirewallFunctionInstanceRequest struct {
	value *PatchedEdgeFirewallFunctionInstanceRequest
	isSet bool
}

func (v NullablePatchedEdgeFirewallFunctionInstanceRequest) Get() *PatchedEdgeFirewallFunctionInstanceRequest {
	return v.value
}

func (v *NullablePatchedEdgeFirewallFunctionInstanceRequest) Set(val *PatchedEdgeFirewallFunctionInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEdgeFirewallFunctionInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEdgeFirewallFunctionInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEdgeFirewallFunctionInstanceRequest(val *PatchedEdgeFirewallFunctionInstanceRequest) *NullablePatchedEdgeFirewallFunctionInstanceRequest {
	return &NullablePatchedEdgeFirewallFunctionInstanceRequest{value: val, isSet: true}
}

func (v NullablePatchedEdgeFirewallFunctionInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEdgeFirewallFunctionInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


