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

// checks if the ResponseListCustomPages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListCustomPages{}

// ResponseListCustomPages struct for ResponseListCustomPages
type ResponseListCustomPages struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	Active *bool `json:"active,omitempty"`
	ProductVersion string `json:"product_version" validate:"regexp=\\\\d+\\\\.\\\\d+"`
	ConnectorCustomPages ConnectorCustomPages `json:"connector_custom_pages"`
}

type _ResponseListCustomPages ResponseListCustomPages

// NewResponseListCustomPages instantiates a new ResponseListCustomPages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListCustomPages(id int64, name string, lastEditor string, lastModified time.Time, productVersion string, connectorCustomPages ConnectorCustomPages) *ResponseListCustomPages {
	this := ResponseListCustomPages{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	this.ConnectorCustomPages = connectorCustomPages
	return &this
}

// NewResponseListCustomPagesWithDefaults instantiates a new ResponseListCustomPages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListCustomPagesWithDefaults() *ResponseListCustomPages {
	this := ResponseListCustomPages{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListCustomPages) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListCustomPages) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListCustomPages) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListCustomPages) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListCustomPages) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListCustomPages) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListCustomPages) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListCustomPages) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListCustomPages) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListCustomPages) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListCustomPages) SetActive(v bool) {
	o.Active = &v
}

// GetProductVersion returns the ProductVersion field value
func (o *ResponseListCustomPages) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *ResponseListCustomPages) SetProductVersion(v string) {
	o.ProductVersion = v
}

// GetConnectorCustomPages returns the ConnectorCustomPages field value
func (o *ResponseListCustomPages) GetConnectorCustomPages() ConnectorCustomPages {
	if o == nil {
		var ret ConnectorCustomPages
		return ret
	}

	return o.ConnectorCustomPages
}

// GetConnectorCustomPagesOk returns a tuple with the ConnectorCustomPages field value
// and a boolean to check if the value has been set.
func (o *ResponseListCustomPages) GetConnectorCustomPagesOk() (*ConnectorCustomPages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorCustomPages, true
}

// SetConnectorCustomPages sets field value
func (o *ResponseListCustomPages) SetConnectorCustomPages(v ConnectorCustomPages) {
	o.ConnectorCustomPages = v
}

func (o ResponseListCustomPages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListCustomPages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["product_version"] = o.ProductVersion
	toSerialize["connector_custom_pages"] = o.ConnectorCustomPages
	return toSerialize, nil
}

func (o *ResponseListCustomPages) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"product_version",
		"connector_custom_pages",
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

	varResponseListCustomPages := _ResponseListCustomPages{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListCustomPages)

	if err != nil {
		return err
	}

	*o = ResponseListCustomPages(varResponseListCustomPages)

	return err
}

type NullableResponseListCustomPages struct {
	value *ResponseListCustomPages
	isSet bool
}

func (v NullableResponseListCustomPages) Get() *ResponseListCustomPages {
	return v.value
}

func (v *NullableResponseListCustomPages) Set(val *ResponseListCustomPages) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListCustomPages) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListCustomPages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListCustomPages(val *ResponseListCustomPages) *NullableResponseListCustomPages {
	return &NullableResponseListCustomPages{value: val, isSet: true}
}

func (v NullableResponseListCustomPages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListCustomPages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


