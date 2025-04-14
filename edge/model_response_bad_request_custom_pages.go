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

// checks if the ResponseBadRequestCustomPages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestCustomPages{}

// ResponseBadRequestCustomPages struct for ResponseBadRequestCustomPages
type ResponseBadRequestCustomPages struct {
	Name []string `json:"name,omitempty"`
	LastEditor []string `json:"last_editor,omitempty"`
	LastModified []string `json:"last_modified,omitempty"`
	ConnectorCustomPages *ResponseBadRequestSerializerMetaclassConnectorCustomPagesField `json:"connector_custom_pages,omitempty"`
	ProductVersion []string `json:"product_version,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestCustomPages instantiates a new ResponseBadRequestCustomPages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestCustomPages() *ResponseBadRequestCustomPages {
	this := ResponseBadRequestCustomPages{}
	return &this
}

// NewResponseBadRequestCustomPagesWithDefaults instantiates a new ResponseBadRequestCustomPages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestCustomPagesWithDefaults() *ResponseBadRequestCustomPages {
	this := ResponseBadRequestCustomPages{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestCustomPages) SetName(v []string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetLastEditor() []string {
	if o == nil || IsNil(o.LastEditor) {
		var ret []string
		return ret
	}
	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetLastEditorOk() ([]string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given []string and assigns it to the LastEditor field.
func (o *ResponseBadRequestCustomPages) SetLastEditor(v []string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetLastModified() []string {
	if o == nil || IsNil(o.LastModified) {
		var ret []string
		return ret
	}
	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetLastModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given []string and assigns it to the LastModified field.
func (o *ResponseBadRequestCustomPages) SetLastModified(v []string) {
	o.LastModified = v
}

// GetConnectorCustomPages returns the ConnectorCustomPages field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetConnectorCustomPages() ResponseBadRequestSerializerMetaclassConnectorCustomPagesField {
	if o == nil || IsNil(o.ConnectorCustomPages) {
		var ret ResponseBadRequestSerializerMetaclassConnectorCustomPagesField
		return ret
	}
	return *o.ConnectorCustomPages
}

// GetConnectorCustomPagesOk returns a tuple with the ConnectorCustomPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetConnectorCustomPagesOk() (*ResponseBadRequestSerializerMetaclassConnectorCustomPagesField, bool) {
	if o == nil || IsNil(o.ConnectorCustomPages) {
		return nil, false
	}
	return o.ConnectorCustomPages, true
}

// HasConnectorCustomPages returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasConnectorCustomPages() bool {
	if o != nil && !IsNil(o.ConnectorCustomPages) {
		return true
	}

	return false
}

// SetConnectorCustomPages gets a reference to the given ResponseBadRequestSerializerMetaclassConnectorCustomPagesField and assigns it to the ConnectorCustomPages field.
func (o *ResponseBadRequestCustomPages) SetConnectorCustomPages(v ResponseBadRequestSerializerMetaclassConnectorCustomPagesField) {
	o.ConnectorCustomPages = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetProductVersion() []string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret []string
		return ret
	}
	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetProductVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given []string and assigns it to the ProductVersion field.
func (o *ResponseBadRequestCustomPages) SetProductVersion(v []string) {
	o.ProductVersion = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestCustomPages) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestCustomPages) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestCustomPages) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestCustomPages) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestCustomPages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestCustomPages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.ConnectorCustomPages) {
		toSerialize["connector_custom_pages"] = o.ConnectorCustomPages
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["product_version"] = o.ProductVersion
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestCustomPages struct {
	value *ResponseBadRequestCustomPages
	isSet bool
}

func (v NullableResponseBadRequestCustomPages) Get() *ResponseBadRequestCustomPages {
	return v.value
}

func (v *NullableResponseBadRequestCustomPages) Set(val *ResponseBadRequestCustomPages) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestCustomPages) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestCustomPages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestCustomPages(val *ResponseBadRequestCustomPages) *NullableResponseBadRequestCustomPages {
	return &NullableResponseBadRequestCustomPages{value: val, isSet: true}
}

func (v NullableResponseBadRequestCustomPages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestCustomPages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


