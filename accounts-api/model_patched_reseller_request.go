/*
Accounts API

REST API OpenAPI documentation for the Accounts API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounts-api

import (
	"encoding/json"
)

// checks if the PatchedResellerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedResellerRequest{}

// PatchedResellerRequest struct for PatchedResellerRequest
type PatchedResellerRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	// * `USD` - USD * `BRL` - BRL
	CurrencyIsoCode *string `json:"currency_iso_code,omitempty"`
	TermsOfServiceUrl *string `json:"terms_of_service_url,omitempty" validate:"regexp=.*"`
}

// NewPatchedResellerRequest instantiates a new PatchedResellerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedResellerRequest() *PatchedResellerRequest {
	this := PatchedResellerRequest{}
	return &this
}

// NewPatchedResellerRequestWithDefaults instantiates a new PatchedResellerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedResellerRequestWithDefaults() *PatchedResellerRequest {
	this := PatchedResellerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedResellerRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedResellerRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedResellerRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedResellerRequest) SetName(v string) {
	o.Name = &v
}

// GetCurrencyIsoCode returns the CurrencyIsoCode field value if set, zero value otherwise.
func (o *PatchedResellerRequest) GetCurrencyIsoCode() string {
	if o == nil || IsNil(o.CurrencyIsoCode) {
		var ret string
		return ret
	}
	return *o.CurrencyIsoCode
}

// GetCurrencyIsoCodeOk returns a tuple with the CurrencyIsoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedResellerRequest) GetCurrencyIsoCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyIsoCode) {
		return nil, false
	}
	return o.CurrencyIsoCode, true
}

// HasCurrencyIsoCode returns a boolean if a field has been set.
func (o *PatchedResellerRequest) HasCurrencyIsoCode() bool {
	if o != nil && !IsNil(o.CurrencyIsoCode) {
		return true
	}

	return false
}

// SetCurrencyIsoCode gets a reference to the given string and assigns it to the CurrencyIsoCode field.
func (o *PatchedResellerRequest) SetCurrencyIsoCode(v string) {
	o.CurrencyIsoCode = &v
}

// GetTermsOfServiceUrl returns the TermsOfServiceUrl field value if set, zero value otherwise.
func (o *PatchedResellerRequest) GetTermsOfServiceUrl() string {
	if o == nil || IsNil(o.TermsOfServiceUrl) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceUrl
}

// GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedResellerRequest) GetTermsOfServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfServiceUrl) {
		return nil, false
	}
	return o.TermsOfServiceUrl, true
}

// HasTermsOfServiceUrl returns a boolean if a field has been set.
func (o *PatchedResellerRequest) HasTermsOfServiceUrl() bool {
	if o != nil && !IsNil(o.TermsOfServiceUrl) {
		return true
	}

	return false
}

// SetTermsOfServiceUrl gets a reference to the given string and assigns it to the TermsOfServiceUrl field.
func (o *PatchedResellerRequest) SetTermsOfServiceUrl(v string) {
	o.TermsOfServiceUrl = &v
}

func (o PatchedResellerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedResellerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CurrencyIsoCode) {
		toSerialize["currency_iso_code"] = o.CurrencyIsoCode
	}
	if !IsNil(o.TermsOfServiceUrl) {
		toSerialize["terms_of_service_url"] = o.TermsOfServiceUrl
	}
	return toSerialize, nil
}

type NullablePatchedResellerRequest struct {
	value *PatchedResellerRequest
	isSet bool
}

func (v NullablePatchedResellerRequest) Get() *PatchedResellerRequest {
	return v.value
}

func (v *NullablePatchedResellerRequest) Set(val *PatchedResellerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedResellerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedResellerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedResellerRequest(val *PatchedResellerRequest) *NullablePatchedResellerRequest {
	return &NullablePatchedResellerRequest{value: val, isSet: true}
}

func (v NullablePatchedResellerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedResellerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


