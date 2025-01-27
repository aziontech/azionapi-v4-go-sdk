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

// checks if the EdgeFunctions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFunctions{}

// EdgeFunctions struct for EdgeFunctions
type EdgeFunctions struct {
	Id int32 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	Language *LanguageEnum `json:"language,omitempty"`
	Code string `json:"code" validate:"regexp=.*"`
	JsonArgs interface{} `json:"json_args,omitempty"`
	InitiatorType *InitiatorTypeEnum `json:"initiator_type,omitempty"`
	Active *bool `json:"active,omitempty"`
	ReferenceCount int32 `json:"reference_count"`
	// Installed version, which may not be the latest if the vendor has released updates since installation.
	Version string `json:"version"`
	Vendor string `json:"vendor"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	ProductVersion string `json:"product_version" validate:"regexp=.*"`
}

type _EdgeFunctions EdgeFunctions

// NewEdgeFunctions instantiates a new EdgeFunctions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctions(id int32, name string, code string, referenceCount int32, version string, vendor string, lastEditor string, lastModified time.Time, productVersion string) *EdgeFunctions {
	this := EdgeFunctions{}
	this.Id = id
	this.Name = name
	var language LanguageEnum = javascript
	this.Language = &language
	this.Code = code
	var initiatorType InitiatorTypeEnum = edge_application
	this.InitiatorType = &initiatorType
	var active bool = true
	this.Active = &active
	this.ReferenceCount = referenceCount
	this.Version = version
	this.Vendor = vendor
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.ProductVersion = productVersion
	return &this
}

// NewEdgeFunctionsWithDefaults instantiates a new EdgeFunctions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionsWithDefaults() *EdgeFunctions {
	this := EdgeFunctions{}
	var language LanguageEnum = javascript
	this.Language = &language
	var initiatorType InitiatorTypeEnum = edge_application
	this.InitiatorType = &initiatorType
	var active bool = true
	this.Active = &active
	return &this
}

// GetId returns the Id field value
func (o *EdgeFunctions) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeFunctions) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EdgeFunctions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeFunctions) SetName(v string) {
	o.Name = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *EdgeFunctions) GetLanguage() LanguageEnum {
	if o == nil || IsNil(o.Language) {
		var ret LanguageEnum
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetLanguageOk() (*LanguageEnum, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EdgeFunctions) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given LanguageEnum and assigns it to the Language field.
func (o *EdgeFunctions) SetLanguage(v LanguageEnum) {
	o.Language = &v
}

// GetCode returns the Code field value
func (o *EdgeFunctions) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EdgeFunctions) SetCode(v string) {
	o.Code = v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EdgeFunctions) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EdgeFunctions) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *EdgeFunctions) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given interface{} and assigns it to the JsonArgs field.
func (o *EdgeFunctions) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *EdgeFunctions) GetInitiatorType() InitiatorTypeEnum {
	if o == nil || IsNil(o.InitiatorType) {
		var ret InitiatorTypeEnum
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetInitiatorTypeOk() (*InitiatorTypeEnum, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *EdgeFunctions) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given InitiatorTypeEnum and assigns it to the InitiatorType field.
func (o *EdgeFunctions) SetInitiatorType(v InitiatorTypeEnum) {
	o.InitiatorType = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeFunctions) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeFunctions) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeFunctions) SetActive(v bool) {
	o.Active = &v
}

// GetReferenceCount returns the ReferenceCount field value
func (o *EdgeFunctions) GetReferenceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceCount
}

// GetReferenceCountOk returns a tuple with the ReferenceCount field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetReferenceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceCount, true
}

// SetReferenceCount sets field value
func (o *EdgeFunctions) SetReferenceCount(v int32) {
	o.ReferenceCount = v
}

// GetVersion returns the Version field value
func (o *EdgeFunctions) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EdgeFunctions) SetVersion(v string) {
	o.Version = v
}

// GetVendor returns the Vendor field value
func (o *EdgeFunctions) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *EdgeFunctions) SetVendor(v string) {
	o.Vendor = v
}

// GetLastEditor returns the LastEditor field value
func (o *EdgeFunctions) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *EdgeFunctions) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *EdgeFunctions) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *EdgeFunctions) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetProductVersion returns the ProductVersion field value
func (o *EdgeFunctions) GetProductVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value
// and a boolean to check if the value has been set.
func (o *EdgeFunctions) GetProductVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVersion, true
}

// SetProductVersion sets field value
func (o *EdgeFunctions) SetProductVersion(v string) {
	o.ProductVersion = v
}

func (o EdgeFunctions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFunctions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	toSerialize["code"] = o.Code
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["reference_count"] = o.ReferenceCount
	toSerialize["version"] = o.Version
	toSerialize["vendor"] = o.Vendor
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	toSerialize["product_version"] = o.ProductVersion
	return toSerialize, nil
}

func (o *EdgeFunctions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"code",
		"reference_count",
		"version",
		"vendor",
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

	varEdgeFunctions := _EdgeFunctions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeFunctions)

	if err != nil {
		return err
	}

	*o = EdgeFunctions(varEdgeFunctions)

	return err
}

type NullableEdgeFunctions struct {
	value *EdgeFunctions
	isSet bool
}

func (v NullableEdgeFunctions) Get() *EdgeFunctions {
	return v.value
}

func (v *NullableEdgeFunctions) Set(val *EdgeFunctions) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctions) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctions(val *EdgeFunctions) *NullableEdgeFunctions {
	return &NullableEdgeFunctions{value: val, isSet: true}
}

func (v NullableEdgeFunctions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


