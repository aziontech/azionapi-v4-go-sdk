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

// checks if the ResponseListEdgeApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListEdgeApplication{}

// ResponseListEdgeApplication struct for ResponseListEdgeApplication
type ResponseListEdgeApplication struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	Modules *EdgeApplicationModules `json:"modules,omitempty"`
	Active *bool `json:"active,omitempty"`
	Debug *bool `json:"debug,omitempty"`
	ProductVersion string `json:"product_version"`
}

type _ResponseListEdgeApplication ResponseListEdgeApplication

// NewResponseListEdgeApplication instantiates a new ResponseListEdgeApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListEdgeApplication(id int64, name string, lastEditor string, lastModified time.Time, productVersion string) *ResponseListEdgeApplication {
	this := ResponseListEdgeApplication{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	return &this
}

// NewResponseListEdgeApplicationWithDefaults instantiates a new ResponseListEdgeApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListEdgeApplicationWithDefaults() *ResponseListEdgeApplication {
	this := ResponseListEdgeApplication{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListEdgeApplication) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListEdgeApplication) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListEdgeApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListEdgeApplication) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListEdgeApplication) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListEdgeApplication) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListEdgeApplication) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListEdgeApplication) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ResponseListEdgeApplication) GetModules() EdgeApplicationModules {
	if o == nil || IsNil(o.Modules) {
		var ret EdgeApplicationModules
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetModulesOk() (*EdgeApplicationModules, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ResponseListEdgeApplication) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given EdgeApplicationModules and assigns it to the Modules field.
func (o *ResponseListEdgeApplication) SetModules(v EdgeApplicationModules) {
	o.Modules = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListEdgeApplication) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListEdgeApplication) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListEdgeApplication) SetActive(v bool) {
	o.Active = &v
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *ResponseListEdgeApplication) GetDebug() bool {
	if o == nil || IsNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *ResponseListEdgeApplication) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *ResponseListEdgeApplication) SetDebug(v bool) {
	o.Debug = &v
}

// GetProductVersion returns the ProductVersion field value
func (o *ResponseListEdgeApplication) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplication) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *ResponseListEdgeApplication) SetProductVersion(v string) {
	o.ProductVersion = v
}

func (o ResponseListEdgeApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListEdgeApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	toSerialize["product_version"] = o.ProductVersion
	return toSerialize, nil
}

func (o *ResponseListEdgeApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"product_version",
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

	varResponseListEdgeApplication := _ResponseListEdgeApplication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListEdgeApplication)

	if err != nil {
		return err
	}

	*o = ResponseListEdgeApplication(varResponseListEdgeApplication)

	return err
}

type NullableResponseListEdgeApplication struct {
	value *ResponseListEdgeApplication
	isSet bool
}

func (v NullableResponseListEdgeApplication) Get() *ResponseListEdgeApplication {
	return v.value
}

func (v *NullableResponseListEdgeApplication) Set(val *ResponseListEdgeApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListEdgeApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListEdgeApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListEdgeApplication(val *ResponseListEdgeApplication) *NullableResponseListEdgeApplication {
	return &NullableResponseListEdgeApplication{value: val, isSet: true}
}

func (v NullableResponseListEdgeApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListEdgeApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


