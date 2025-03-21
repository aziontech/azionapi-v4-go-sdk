/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ResponseListWAF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListWAF{}

// ResponseListWAF struct for ResponseListWAF
type ResponseListWAF struct {
	Id int64 `json:"id"`
	Active *bool `json:"active,omitempty"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	ProductVersion NullableString `json:"product_version" validate:"regexp=\\\\d+\\\\.\\\\d+"`
	ThreatsConfiguration *WAFThreatsConfiguration `json:"threats_configuration,omitempty"`
}

type _ResponseListWAF ResponseListWAF

// NewResponseListWAF instantiates a new ResponseListWAF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListWAF(id int64, name string, lastEditor string, lastModified time.Time, productVersion NullableString) *ResponseListWAF {
	this := ResponseListWAF{}
	this.Id = id
	var active bool = true
	this.Active = &active
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	return &this
}

// NewResponseListWAFWithDefaults instantiates a new ResponseListWAF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListWAFWithDefaults() *ResponseListWAF {
	this := ResponseListWAF{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetId returns the Id field value
func (o *ResponseListWAF) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListWAF) SetId(v int64) {
	o.Id = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListWAF) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListWAF) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListWAF) SetActive(v bool) {
	o.Active = &v
}

// GetName returns the Name field value
func (o *ResponseListWAF) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListWAF) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListWAF) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListWAF) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListWAF) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListWAF) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetProductVersion returns the ProductVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponseListWAF) GetProductVersion() string {
	if o == nil || o.ProductVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProductVersion.Get()
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseListWAF) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductVersion.Get(), o.ProductVersion.IsSet()
}

// SetProductVersion sets field value
func (o *ResponseListWAF) SetProductVersion(v string) {
	o.ProductVersion.Set(&v)
}

// GetThreatsConfiguration returns the ThreatsConfiguration field value if set, zero value otherwise.
func (o *ResponseListWAF) GetThreatsConfiguration() WAFThreatsConfiguration {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		var ret WAFThreatsConfiguration
		return ret
	}
	return *o.ThreatsConfiguration
}

// GetThreatsConfigurationOk returns a tuple with the ThreatsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListWAF) GetThreatsConfigurationOk() (*WAFThreatsConfiguration, bool) {
	if o == nil || IsNil(o.ThreatsConfiguration) {
		return nil, false
	}
	return o.ThreatsConfiguration, true
}

// HasThreatsConfiguration returns a boolean if a field has been set.
func (o *ResponseListWAF) HasThreatsConfiguration() bool {
	if o != nil && !IsNil(o.ThreatsConfiguration) {
		return true
	}

	return false
}

// SetThreatsConfiguration gets a reference to the given WAFThreatsConfiguration and assigns it to the ThreatsConfiguration field.
func (o *ResponseListWAF) SetThreatsConfiguration(v WAFThreatsConfiguration) {
	o.ThreatsConfiguration = &v
}

func (o ResponseListWAF) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListWAF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["product_version"] = o.ProductVersion.Get()
	if !IsNil(o.ThreatsConfiguration) {
		toSerialize["threats_configuration"] = o.ThreatsConfiguration
	}
	return toSerialize, nil
}

func (o *ResponseListWAF) UnmarshalJSON(data []byte) (err error) {
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

	varResponseListWAF := _ResponseListWAF{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListWAF)

	if err != nil {
		return err
	}

	*o = ResponseListWAF(varResponseListWAF)

	return err
}

type NullableResponseListWAF struct {
	value *ResponseListWAF
	isSet bool
}

func (v NullableResponseListWAF) Get() *ResponseListWAF {
	return v.value
}

func (v *NullableResponseListWAF) Set(val *ResponseListWAF) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListWAF) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListWAF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListWAF(val *ResponseListWAF) *NullableResponseListWAF {
	return &NullableResponseListWAF{value: val, isSet: true}
}

func (v NullableResponseListWAF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListWAF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


