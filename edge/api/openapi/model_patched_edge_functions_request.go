/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
)

// checks if the PatchedEdgeFunctionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedEdgeFunctionsRequest{}

// PatchedEdgeFunctionsRequest struct for PatchedEdgeFunctionsRequest
type PatchedEdgeFunctionsRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	// * `javascript` - JavaScript * `lua` - Lua
	Language *string `json:"language,omitempty"`
	Code *string `json:"code,omitempty" validate:"regexp=.*"`
	JsonArgs interface{} `json:"json_args,omitempty"`
	// * `edge_application` - Edge Application * `edge_firewall` - Edge Firewall
	InitiatorType *string `json:"initiator_type,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewPatchedEdgeFunctionsRequest instantiates a new PatchedEdgeFunctionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEdgeFunctionsRequest() *PatchedEdgeFunctionsRequest {
	this := PatchedEdgeFunctionsRequest{}
	var language string = "javascript"
	this.Language = &language
	var initiatorType string = "edge_application"
	this.InitiatorType = &initiatorType
	var active bool = true
	this.Active = &active
	return &this
}

// NewPatchedEdgeFunctionsRequestWithDefaults instantiates a new PatchedEdgeFunctionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEdgeFunctionsRequestWithDefaults() *PatchedEdgeFunctionsRequest {
	this := PatchedEdgeFunctionsRequest{}
	var language string = "javascript"
	this.Language = &language
	var initiatorType string = "edge_application"
	this.InitiatorType = &initiatorType
	var active bool = true
	this.Active = &active
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEdgeFunctionsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFunctionsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEdgeFunctionsRequest) SetName(v string) {
	o.Name = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PatchedEdgeFunctionsRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFunctionsRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PatchedEdgeFunctionsRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PatchedEdgeFunctionsRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFunctionsRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PatchedEdgeFunctionsRequest) SetCode(v string) {
	o.Code = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedEdgeFunctionsRequest) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedEdgeFunctionsRequest) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given interface{} and assigns it to the JsonArgs field.
func (o *PatchedEdgeFunctionsRequest) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *PatchedEdgeFunctionsRequest) GetInitiatorType() string {
	if o == nil || IsNil(o.InitiatorType) {
		var ret string
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFunctionsRequest) GetInitiatorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given string and assigns it to the InitiatorType field.
func (o *PatchedEdgeFunctionsRequest) SetInitiatorType(v string) {
	o.InitiatorType = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedEdgeFunctionsRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFunctionsRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedEdgeFunctionsRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedEdgeFunctionsRequest) SetActive(v bool) {
	o.Active = &v
}

func (o PatchedEdgeFunctionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedEdgeFunctionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullablePatchedEdgeFunctionsRequest struct {
	value *PatchedEdgeFunctionsRequest
	isSet bool
}

func (v NullablePatchedEdgeFunctionsRequest) Get() *PatchedEdgeFunctionsRequest {
	return v.value
}

func (v *NullablePatchedEdgeFunctionsRequest) Set(val *PatchedEdgeFunctionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEdgeFunctionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEdgeFunctionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEdgeFunctionsRequest(val *PatchedEdgeFunctionsRequest) *NullablePatchedEdgeFunctionsRequest {
	return &NullablePatchedEdgeFunctionsRequest{value: val, isSet: true}
}

func (v NullablePatchedEdgeFunctionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEdgeFunctionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


