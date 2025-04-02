/*
object-storage-api

REST API OpenAPI documentation for the Object Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage-api

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CredentialCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialCreateRequest{}

// CredentialCreateRequest struct for CredentialCreateRequest
type CredentialCreateRequest struct {
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]+$"`
	Capabilities []string `json:"capabilities"`
	Bucket *string `json:"bucket,omitempty" validate:"regexp=.{0,63}"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
}

type _CredentialCreateRequest CredentialCreateRequest

// NewCredentialCreateRequest instantiates a new CredentialCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialCreateRequest(name string, capabilities []string) *CredentialCreateRequest {
	this := CredentialCreateRequest{}
	this.Name = name
	this.Capabilities = capabilities
	return &this
}

// NewCredentialCreateRequestWithDefaults instantiates a new CredentialCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialCreateRequestWithDefaults() *CredentialCreateRequest {
	this := CredentialCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CredentialCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CredentialCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CredentialCreateRequest) SetName(v string) {
	o.Name = v
}

// GetCapabilities returns the Capabilities field value
func (o *CredentialCreateRequest) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *CredentialCreateRequest) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *CredentialCreateRequest) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *CredentialCreateRequest) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialCreateRequest) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *CredentialCreateRequest) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *CredentialCreateRequest) SetBucket(v string) {
	o.Bucket = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CredentialCreateRequest) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialCreateRequest) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CredentialCreateRequest) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CredentialCreateRequest) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

func (o CredentialCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	return toSerialize, nil
}

func (o *CredentialCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"capabilities",
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

	varCredentialCreateRequest := _CredentialCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCredentialCreateRequest)

	if err != nil {
		return err
	}

	*o = CredentialCreateRequest(varCredentialCreateRequest)

	return err
}

type NullableCredentialCreateRequest struct {
	value *CredentialCreateRequest
	isSet bool
}

func (v NullableCredentialCreateRequest) Get() *CredentialCreateRequest {
	return v.value
}

func (v *NullableCredentialCreateRequest) Set(val *CredentialCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialCreateRequest(val *CredentialCreateRequest) *NullableCredentialCreateRequest {
	return &NullableCredentialCreateRequest{value: val, isSet: true}
}

func (v NullableCredentialCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


