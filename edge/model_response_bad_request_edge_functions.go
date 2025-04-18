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

// checks if the ResponseBadRequestEdgeFunctions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestEdgeFunctions{}

// ResponseBadRequestEdgeFunctions struct for ResponseBadRequestEdgeFunctions
type ResponseBadRequestEdgeFunctions struct {
	Id []string `json:"id,omitempty"`
	Name []string `json:"name,omitempty"`
	Language []string `json:"language,omitempty"`
	Code []string `json:"code,omitempty"`
	InitiatorType []string `json:"initiator_type,omitempty"`
	JsonArgs interface{} `json:"json_args,omitempty"`
	LastEditor []string `json:"last_editor,omitempty"`
	LastModified []string `json:"last_modified,omitempty"`
	Active []string `json:"active,omitempty"`
	ReferenceCount []string `json:"reference_count,omitempty"`
	Version []string `json:"version,omitempty"`
	Vendor []string `json:"vendor,omitempty"`
	ProductVersion []string `json:"product_version,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestEdgeFunctions instantiates a new ResponseBadRequestEdgeFunctions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestEdgeFunctions() *ResponseBadRequestEdgeFunctions {
	this := ResponseBadRequestEdgeFunctions{}
	return &this
}

// NewResponseBadRequestEdgeFunctionsWithDefaults instantiates a new ResponseBadRequestEdgeFunctions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestEdgeFunctionsWithDefaults() *ResponseBadRequestEdgeFunctions {
	this := ResponseBadRequestEdgeFunctions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetId() []string {
	if o == nil || IsNil(o.Id) {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *ResponseBadRequestEdgeFunctions) SetId(v []string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestEdgeFunctions) SetName(v []string) {
	o.Name = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetLanguage() []string {
	if o == nil || IsNil(o.Language) {
		var ret []string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetLanguageOk() ([]string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []string and assigns it to the Language field.
func (o *ResponseBadRequestEdgeFunctions) SetLanguage(v []string) {
	o.Language = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetCode() []string {
	if o == nil || IsNil(o.Code) {
		var ret []string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetCodeOk() ([]string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given []string and assigns it to the Code field.
func (o *ResponseBadRequestEdgeFunctions) SetCode(v []string) {
	o.Code = v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetInitiatorType() []string {
	if o == nil || IsNil(o.InitiatorType) {
		var ret []string
		return ret
	}
	return o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetInitiatorTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given []string and assigns it to the InitiatorType field.
func (o *ResponseBadRequestEdgeFunctions) SetInitiatorType(v []string) {
	o.InitiatorType = v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseBadRequestEdgeFunctions) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseBadRequestEdgeFunctions) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given interface{} and assigns it to the JsonArgs field.
func (o *ResponseBadRequestEdgeFunctions) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetLastEditor() []string {
	if o == nil || IsNil(o.LastEditor) {
		var ret []string
		return ret
	}
	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetLastEditorOk() ([]string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given []string and assigns it to the LastEditor field.
func (o *ResponseBadRequestEdgeFunctions) SetLastEditor(v []string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetLastModified() []string {
	if o == nil || IsNil(o.LastModified) {
		var ret []string
		return ret
	}
	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetLastModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given []string and assigns it to the LastModified field.
func (o *ResponseBadRequestEdgeFunctions) SetLastModified(v []string) {
	o.LastModified = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetActive() []string {
	if o == nil || IsNil(o.Active) {
		var ret []string
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetActiveOk() ([]string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []string and assigns it to the Active field.
func (o *ResponseBadRequestEdgeFunctions) SetActive(v []string) {
	o.Active = v
}

// GetReferenceCount returns the ReferenceCount field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetReferenceCount() []string {
	if o == nil || IsNil(o.ReferenceCount) {
		var ret []string
		return ret
	}
	return o.ReferenceCount
}

// GetReferenceCountOk returns a tuple with the ReferenceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetReferenceCountOk() ([]string, bool) {
	if o == nil || IsNil(o.ReferenceCount) {
		return nil, false
	}
	return o.ReferenceCount, true
}

// HasReferenceCount returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasReferenceCount() bool {
	if o != nil && !IsNil(o.ReferenceCount) {
		return true
	}

	return false
}

// SetReferenceCount gets a reference to the given []string and assigns it to the ReferenceCount field.
func (o *ResponseBadRequestEdgeFunctions) SetReferenceCount(v []string) {
	o.ReferenceCount = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetVersion() []string {
	if o == nil || IsNil(o.Version) {
		var ret []string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given []string and assigns it to the Version field.
func (o *ResponseBadRequestEdgeFunctions) SetVersion(v []string) {
	o.Version = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetVendor() []string {
	if o == nil || IsNil(o.Vendor) {
		var ret []string
		return ret
	}
	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetVendorOk() ([]string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given []string and assigns it to the Vendor field.
func (o *ResponseBadRequestEdgeFunctions) SetVendor(v []string) {
	o.Vendor = v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetProductVersion() []string {
	if o == nil || IsNil(o.ProductVersion) {
		var ret []string
		return ret
	}
	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetProductVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductVersion) {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasProductVersion() bool {
	if o != nil && !IsNil(o.ProductVersion) {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given []string and assigns it to the ProductVersion field.
func (o *ResponseBadRequestEdgeFunctions) SetProductVersion(v []string) {
	o.ProductVersion = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeFunctions) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeFunctions) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeFunctions) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestEdgeFunctions) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestEdgeFunctions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestEdgeFunctions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ReferenceCount) {
		toSerialize["reference_count"] = o.ReferenceCount
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.ProductVersion) {
		toSerialize["product_version"] = o.ProductVersion
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestEdgeFunctions struct {
	value *ResponseBadRequestEdgeFunctions
	isSet bool
}

func (v NullableResponseBadRequestEdgeFunctions) Get() *ResponseBadRequestEdgeFunctions {
	return v.value
}

func (v *NullableResponseBadRequestEdgeFunctions) Set(val *ResponseBadRequestEdgeFunctions) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestEdgeFunctions) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestEdgeFunctions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestEdgeFunctions(val *ResponseBadRequestEdgeFunctions) *NullableResponseBadRequestEdgeFunctions {
	return &NullableResponseBadRequestEdgeFunctions{value: val, isSet: true}
}

func (v NullableResponseBadRequestEdgeFunctions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestEdgeFunctions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


