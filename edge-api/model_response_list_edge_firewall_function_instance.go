/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ResponseListEdgeFirewallFunctionInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListEdgeFirewallFunctionInstance{}

// ResponseListEdgeFirewallFunctionInstance struct for ResponseListEdgeFirewallFunctionInstance
type ResponseListEdgeFirewallFunctionInstance struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	JsonArgs interface{} `json:"json_args"`
	EdgeFunction int64 `json:"edge_function"`
	Active *bool `json:"active,omitempty"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
}

type _ResponseListEdgeFirewallFunctionInstance ResponseListEdgeFirewallFunctionInstance

// NewResponseListEdgeFirewallFunctionInstance instantiates a new ResponseListEdgeFirewallFunctionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListEdgeFirewallFunctionInstance(id int64, name string, jsonArgs interface{}, edgeFunction int64, lastEditor string, lastModified time.Time) *ResponseListEdgeFirewallFunctionInstance {
	this := ResponseListEdgeFirewallFunctionInstance{}
	this.Id = id
	this.Name = name
	this.JsonArgs = jsonArgs
	this.EdgeFunction = edgeFunction
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	return &this
}

// NewResponseListEdgeFirewallFunctionInstanceWithDefaults instantiates a new ResponseListEdgeFirewallFunctionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListEdgeFirewallFunctionInstanceWithDefaults() *ResponseListEdgeFirewallFunctionInstance {
	this := ResponseListEdgeFirewallFunctionInstance{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListEdgeFirewallFunctionInstance) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListEdgeFirewallFunctionInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetName(v string) {
	o.Name = v
}

// GetJsonArgs returns the JsonArgs field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ResponseListEdgeFirewallFunctionInstance) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseListEdgeFirewallFunctionInstance) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// SetJsonArgs sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetEdgeFunction returns the EdgeFunction field value
func (o *ResponseListEdgeFirewallFunctionInstance) GetEdgeFunction() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetEdgeFunctionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunction, true
}

// SetEdgeFunction sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetEdgeFunction(v int64) {
	o.EdgeFunction = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListEdgeFirewallFunctionInstance) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListEdgeFirewallFunctionInstance) SetActive(v bool) {
	o.Active = &v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListEdgeFirewallFunctionInstance) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListEdgeFirewallFunctionInstance) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallFunctionInstance) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListEdgeFirewallFunctionInstance) SetLastModified(v time.Time) {
	o.LastModified = v
}

func (o ResponseListEdgeFirewallFunctionInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListEdgeFirewallFunctionInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	toSerialize["edge_function"] = o.EdgeFunction
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	return toSerialize, nil
}

func (o *ResponseListEdgeFirewallFunctionInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"json_args",
		"edge_function",
		"last_editor",
		"last_modified",
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

	varResponseListEdgeFirewallFunctionInstance := _ResponseListEdgeFirewallFunctionInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListEdgeFirewallFunctionInstance)

	if err != nil {
		return err
	}

	*o = ResponseListEdgeFirewallFunctionInstance(varResponseListEdgeFirewallFunctionInstance)

	return err
}

type NullableResponseListEdgeFirewallFunctionInstance struct {
	value *ResponseListEdgeFirewallFunctionInstance
	isSet bool
}

func (v NullableResponseListEdgeFirewallFunctionInstance) Get() *ResponseListEdgeFirewallFunctionInstance {
	return v.value
}

func (v *NullableResponseListEdgeFirewallFunctionInstance) Set(val *ResponseListEdgeFirewallFunctionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListEdgeFirewallFunctionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListEdgeFirewallFunctionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListEdgeFirewallFunctionInstance(val *ResponseListEdgeFirewallFunctionInstance) *NullableResponseListEdgeFirewallFunctionInstance {
	return &NullableResponseListEdgeFirewallFunctionInstance{value: val, isSet: true}
}

func (v NullableResponseListEdgeFirewallFunctionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListEdgeFirewallFunctionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


