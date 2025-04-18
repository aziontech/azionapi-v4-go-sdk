/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
)

// checks if the ResponseBadRequestDataStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestDataStream{}

// ResponseBadRequestDataStream struct for ResponseBadRequestDataStream
type ResponseBadRequestDataStream struct {
	Name []string `json:"name,omitempty"`
	DataSource []string `json:"data_source,omitempty"`
	DataSetId []string `json:"data_set_id,omitempty"`
	Active []string `json:"active,omitempty"`
	Filters *ResponseBadRequestSerializerMetaclassFiltersField `json:"filters,omitempty"`
	ProductVersion []string `json:"product_version,omitempty"`
	LastEditor []string `json:"last_editor,omitempty"`
	LastModified []string `json:"last_modified,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestDataStream instantiates a new ResponseBadRequestDataStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestDataStream() *ResponseBadRequestDataStream {
	this := ResponseBadRequestDataStream{}
	return &this
}

// NewResponseBadRequestDataStreamWithDefaults instantiates a new ResponseBadRequestDataStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestDataStreamWithDefaults() *ResponseBadRequestDataStream {
	this := ResponseBadRequestDataStream{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestDataStream) SetName(v []string) {
	o.Name = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetDataSource() []string {
	if o == nil || IsNil(o.DataSource) {
		var ret []string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetDataSourceOk() ([]string, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given []string and assigns it to the DataSource field.
func (o *ResponseBadRequestDataStream) SetDataSource(v []string) {
	o.DataSource = v
}

// GetDataSetId returns the DataSetId field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetDataSetId() []string {
	if o == nil || IsNil(o.DataSetId) {
		var ret []string
		return ret
	}
	return o.DataSetId
}

// GetDataSetIdOk returns a tuple with the DataSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetDataSetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.DataSetId) {
		return nil, false
	}
	return o.DataSetId, true
}

// HasDataSetId returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasDataSetId() bool {
	if o != nil && !IsNil(o.DataSetId) {
		return true
	}

	return false
}

// SetDataSetId gets a reference to the given []string and assigns it to the DataSetId field.
func (o *ResponseBadRequestDataStream) SetDataSetId(v []string) {
	o.DataSetId = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetActive() []string {
	if o == nil || IsNil(o.Active) {
		var ret []string
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetActiveOk() ([]string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []string and assigns it to the Active field.
func (o *ResponseBadRequestDataStream) SetActive(v []string) {
	o.Active = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetFilters() ResponseBadRequestSerializerMetaclassFiltersField {
	if o == nil || IsNil(o.Filters) {
		var ret ResponseBadRequestSerializerMetaclassFiltersField
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetFiltersOk() (*ResponseBadRequestSerializerMetaclassFiltersField, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ResponseBadRequestSerializerMetaclassFiltersField and assigns it to the Filters field.
func (o *ResponseBadRequestDataStream) SetFilters(v ResponseBadRequestSerializerMetaclassFiltersField) {
	o.Filters = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetProductVersion() []string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret []string
		return ret
	}
	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetProductVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given []string and assigns it to the ProductVersion field.
func (o *ResponseBadRequestDataStream) SetProductVersion(v []string) {
	o.ProductVersion = v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetLastEditor() []string {
	if o == nil || IsNil(o.LastEditor) {
		var ret []string
		return ret
	}
	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetLastEditorOk() ([]string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given []string and assigns it to the LastEditor field.
func (o *ResponseBadRequestDataStream) SetLastEditor(v []string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetLastModified() []string {
	if o == nil || IsNil(o.LastModified) {
		var ret []string
		return ret
	}
	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetLastModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given []string and assigns it to the LastModified field.
func (o *ResponseBadRequestDataStream) SetLastModified(v []string) {
	o.LastModified = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestDataStream) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestDataStream) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestDataStream) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestDataStream) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestDataStream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestDataStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	if !IsNil(o.DataSetId) {
		toSerialize["data_set_id"] = o.DataSetId
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["product_version"] = o.ProductVersion
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestDataStream struct {
	value *ResponseBadRequestDataStream
	isSet bool
}

func (v NullableResponseBadRequestDataStream) Get() *ResponseBadRequestDataStream {
	return v.value
}

func (v *NullableResponseBadRequestDataStream) Set(val *ResponseBadRequestDataStream) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestDataStream) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestDataStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestDataStream(val *ResponseBadRequestDataStream) *NullableResponseBadRequestDataStream {
	return &NullableResponseBadRequestDataStream{value: val, isSet: true}
}

func (v NullableResponseBadRequestDataStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestDataStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


