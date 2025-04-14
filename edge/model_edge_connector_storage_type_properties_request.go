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

// checks if the EdgeConnectorStorageTypePropertiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeConnectorStorageTypePropertiesRequest{}

// EdgeConnectorStorageTypePropertiesRequest struct for EdgeConnectorStorageTypePropertiesRequest
type EdgeConnectorStorageTypePropertiesRequest struct {
	Bucket *string `json:"bucket,omitempty" validate:"regexp=.*"`
	Prefix *string `json:"prefix,omitempty" validate:"regexp=.*"`
}

// NewEdgeConnectorStorageTypePropertiesRequest instantiates a new EdgeConnectorStorageTypePropertiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeConnectorStorageTypePropertiesRequest() *EdgeConnectorStorageTypePropertiesRequest {
	this := EdgeConnectorStorageTypePropertiesRequest{}
	return &this
}

// NewEdgeConnectorStorageTypePropertiesRequestWithDefaults instantiates a new EdgeConnectorStorageTypePropertiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeConnectorStorageTypePropertiesRequestWithDefaults() *EdgeConnectorStorageTypePropertiesRequest {
	this := EdgeConnectorStorageTypePropertiesRequest{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *EdgeConnectorStorageTypePropertiesRequest) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageTypePropertiesRequest) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *EdgeConnectorStorageTypePropertiesRequest) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *EdgeConnectorStorageTypePropertiesRequest) SetBucket(v string) {
	o.Bucket = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *EdgeConnectorStorageTypePropertiesRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeConnectorStorageTypePropertiesRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *EdgeConnectorStorageTypePropertiesRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *EdgeConnectorStorageTypePropertiesRequest) SetPrefix(v string) {
	o.Prefix = &v
}

func (o EdgeConnectorStorageTypePropertiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeConnectorStorageTypePropertiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return toSerialize, nil
}

type NullableEdgeConnectorStorageTypePropertiesRequest struct {
	value *EdgeConnectorStorageTypePropertiesRequest
	isSet bool
}

func (v NullableEdgeConnectorStorageTypePropertiesRequest) Get() *EdgeConnectorStorageTypePropertiesRequest {
	return v.value
}

func (v *NullableEdgeConnectorStorageTypePropertiesRequest) Set(val *EdgeConnectorStorageTypePropertiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeConnectorStorageTypePropertiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeConnectorStorageTypePropertiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeConnectorStorageTypePropertiesRequest(val *EdgeConnectorStorageTypePropertiesRequest) *NullableEdgeConnectorStorageTypePropertiesRequest {
	return &NullableEdgeConnectorStorageTypePropertiesRequest{value: val, isSet: true}
}

func (v NullableEdgeConnectorStorageTypePropertiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeConnectorStorageTypePropertiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


