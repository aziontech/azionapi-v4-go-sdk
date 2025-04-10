/*
data-stream-api

REST API OpenAPI documentation for the Data Stream

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data-stream-api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AWSKinesisFirehoseEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSKinesisFirehoseEndpointRequest{}

// AWSKinesisFirehoseEndpointRequest struct for AWSKinesisFirehoseEndpointRequest
type AWSKinesisFirehoseEndpointRequest struct {
	AccessKey string `json:"access_key"`
	StreamName string `json:"stream_name"`
	Region string `json:"region"`
	SecretKey string `json:"secret_key"`
}

type _AWSKinesisFirehoseEndpointRequest AWSKinesisFirehoseEndpointRequest

// NewAWSKinesisFirehoseEndpointRequest instantiates a new AWSKinesisFirehoseEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSKinesisFirehoseEndpointRequest(accessKey string, streamName string, region string, secretKey string) *AWSKinesisFirehoseEndpointRequest {
	this := AWSKinesisFirehoseEndpointRequest{}
	this.AccessKey = accessKey
	this.StreamName = streamName
	this.Region = region
	this.SecretKey = secretKey
	return &this
}

// NewAWSKinesisFirehoseEndpointRequestWithDefaults instantiates a new AWSKinesisFirehoseEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSKinesisFirehoseEndpointRequestWithDefaults() *AWSKinesisFirehoseEndpointRequest {
	this := AWSKinesisFirehoseEndpointRequest{}
	return &this
}

// GetAccessKey returns the AccessKey field value
func (o *AWSKinesisFirehoseEndpointRequest) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *AWSKinesisFirehoseEndpointRequest) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *AWSKinesisFirehoseEndpointRequest) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetStreamName returns the StreamName field value
func (o *AWSKinesisFirehoseEndpointRequest) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *AWSKinesisFirehoseEndpointRequest) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *AWSKinesisFirehoseEndpointRequest) SetStreamName(v string) {
	o.StreamName = v
}

// GetRegion returns the Region field value
func (o *AWSKinesisFirehoseEndpointRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSKinesisFirehoseEndpointRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AWSKinesisFirehoseEndpointRequest) SetRegion(v string) {
	o.Region = v
}

// GetSecretKey returns the SecretKey field value
func (o *AWSKinesisFirehoseEndpointRequest) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *AWSKinesisFirehoseEndpointRequest) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *AWSKinesisFirehoseEndpointRequest) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o AWSKinesisFirehoseEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSKinesisFirehoseEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_key"] = o.AccessKey
	toSerialize["stream_name"] = o.StreamName
	toSerialize["region"] = o.Region
	toSerialize["secret_key"] = o.SecretKey
	return toSerialize, nil
}

func (o *AWSKinesisFirehoseEndpointRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_key",
		"stream_name",
		"region",
		"secret_key",
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

	varAWSKinesisFirehoseEndpointRequest := _AWSKinesisFirehoseEndpointRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAWSKinesisFirehoseEndpointRequest)

	if err != nil {
		return err
	}

	*o = AWSKinesisFirehoseEndpointRequest(varAWSKinesisFirehoseEndpointRequest)

	return err
}

type NullableAWSKinesisFirehoseEndpointRequest struct {
	value *AWSKinesisFirehoseEndpointRequest
	isSet bool
}

func (v NullableAWSKinesisFirehoseEndpointRequest) Get() *AWSKinesisFirehoseEndpointRequest {
	return v.value
}

func (v *NullableAWSKinesisFirehoseEndpointRequest) Set(val *AWSKinesisFirehoseEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSKinesisFirehoseEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSKinesisFirehoseEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSKinesisFirehoseEndpointRequest(val *AWSKinesisFirehoseEndpointRequest) *NullableAWSKinesisFirehoseEndpointRequest {
	return &NullableAWSKinesisFirehoseEndpointRequest{value: val, isSet: true}
}

func (v NullableAWSKinesisFirehoseEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSKinesisFirehoseEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


