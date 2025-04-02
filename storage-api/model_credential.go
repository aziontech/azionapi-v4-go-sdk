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

// checks if the Credential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credential{}

// Credential struct for Credential
type Credential struct {
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]+$"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Capabilities []string `json:"capabilities"`
	Bucket *string `json:"bucket,omitempty" validate:"regexp=.{0,63}"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type _Credential Credential

// NewCredential instantiates a new Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredential(name string, accessKey string, secretKey string, capabilities []string, createdAt time.Time) *Credential {
	this := Credential{}
	this.Name = name
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	this.Capabilities = capabilities
	this.CreatedAt = createdAt
	return &this
}

// NewCredentialWithDefaults instantiates a new Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialWithDefaults() *Credential {
	this := Credential{}
	return &this
}

// GetName returns the Name field value
func (o *Credential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Credential) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Credential) SetName(v string) {
	o.Name = v
}

// GetAccessKey returns the AccessKey field value
func (o *Credential) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *Credential) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *Credential) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *Credential) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *Credential) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *Credential) SetSecretKey(v string) {
	o.SecretKey = v
}

// GetCapabilities returns the Capabilities field value
func (o *Credential) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *Credential) GetCapabilitiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *Credential) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *Credential) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *Credential) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *Credential) SetBucket(v string) {
	o.Bucket = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *Credential) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Credential) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *Credential) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Credential) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Credential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Credential) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Credential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["access_key"] = o.AccessKey
	toSerialize["secret_key"] = o.SecretKey
	toSerialize["capabilities"] = o.Capabilities
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Credential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"access_key",
		"secret_key",
		"capabilities",
		"created_at",
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

	varCredential := _Credential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCredential)

	if err != nil {
		return err
	}

	*o = Credential(varCredential)

	return err
}

type NullableCredential struct {
	value *Credential
	isSet bool
}

func (v NullableCredential) Get() *Credential {
	return v.value
}

func (v *NullableCredential) Set(val *Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredential(val *Credential) *NullableCredential {
	return &NullableCredential{value: val, isSet: true}
}

func (v NullableCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


