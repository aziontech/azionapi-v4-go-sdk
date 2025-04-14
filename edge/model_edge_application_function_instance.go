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

// checks if the EdgeApplicationFunctionInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationFunctionInstance{}

// EdgeApplicationFunctionInstance struct for EdgeApplicationFunctionInstance
type EdgeApplicationFunctionInstance struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	JsonArgs *EdgeApplicationFunctionInstanceJsonArgs `json:"json_args,omitempty"`
	EdgeFunction int64 `json:"edge_function"`
	Active *bool `json:"active,omitempty"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
}

type _EdgeApplicationFunctionInstance EdgeApplicationFunctionInstance

// NewEdgeApplicationFunctionInstance instantiates a new EdgeApplicationFunctionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationFunctionInstance(id int64, name string, edgeFunction int64, lastEditor string, lastModified time.Time) *EdgeApplicationFunctionInstance {
	this := EdgeApplicationFunctionInstance{}
	this.Id = id
	this.Name = name
	this.EdgeFunction = edgeFunction
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	return &this
}

// NewEdgeApplicationFunctionInstanceWithDefaults instantiates a new EdgeApplicationFunctionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationFunctionInstanceWithDefaults() *EdgeApplicationFunctionInstance {
	this := EdgeApplicationFunctionInstance{}
	return &this
}

// GetId returns the Id field value
func (o *EdgeApplicationFunctionInstance) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeApplicationFunctionInstance) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EdgeApplicationFunctionInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeApplicationFunctionInstance) SetName(v string) {
	o.Name = v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise.
func (o *EdgeApplicationFunctionInstance) GetJsonArgs() EdgeApplicationFunctionInstanceJsonArgs {
	if o == nil || IsNil(o.JsonArgs) {
		var ret EdgeApplicationFunctionInstanceJsonArgs
		return ret
	}
	return *o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetJsonArgsOk() (*EdgeApplicationFunctionInstanceJsonArgs, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *EdgeApplicationFunctionInstance) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given EdgeApplicationFunctionInstanceJsonArgs and assigns it to the JsonArgs field.
func (o *EdgeApplicationFunctionInstance) SetJsonArgs(v EdgeApplicationFunctionInstanceJsonArgs) {
	o.JsonArgs = &v
}

// GetEdgeFunction returns the EdgeFunction field value
func (o *EdgeApplicationFunctionInstance) GetEdgeFunction() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetEdgeFunctionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunction, true
}

// SetEdgeFunction sets field value
func (o *EdgeApplicationFunctionInstance) SetEdgeFunction(v int64) {
	o.EdgeFunction = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeApplicationFunctionInstance) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeApplicationFunctionInstance) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeApplicationFunctionInstance) SetActive(v bool) {
	o.Active = &v
}

// GetLastEditor returns the LastEditor field value
func (o *EdgeApplicationFunctionInstance) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *EdgeApplicationFunctionInstance) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *EdgeApplicationFunctionInstance) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationFunctionInstance) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *EdgeApplicationFunctionInstance) SetLastModified(v time.Time) {
	o.LastModified = v
}

func (o EdgeApplicationFunctionInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationFunctionInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.JsonArgs) {
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

func (o *EdgeApplicationFunctionInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varEdgeApplicationFunctionInstance := _EdgeApplicationFunctionInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeApplicationFunctionInstance)

	if err != nil {
		return err
	}

	*o = EdgeApplicationFunctionInstance(varEdgeApplicationFunctionInstance)

	return err
}

type NullableEdgeApplicationFunctionInstance struct {
	value *EdgeApplicationFunctionInstance
	isSet bool
}

func (v NullableEdgeApplicationFunctionInstance) Get() *EdgeApplicationFunctionInstance {
	return v.value
}

func (v *NullableEdgeApplicationFunctionInstance) Set(val *EdgeApplicationFunctionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationFunctionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationFunctionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationFunctionInstance(val *EdgeApplicationFunctionInstance) *NullableEdgeApplicationFunctionInstance {
	return &NullableEdgeApplicationFunctionInstance{value: val, isSet: true}
}

func (v NullableEdgeApplicationFunctionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationFunctionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


