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

// checks if the HttpPostEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpPostEndpointRequest{}

// HttpPostEndpointRequest struct for HttpPostEndpointRequest
type HttpPostEndpointRequest struct {
	Url string `json:"url"`
	LogLineSeparator *string `json:"log_line_separator,omitempty"`
	PayloadFormat *string `json:"payload_format,omitempty"`
	MaxSize NullableInt64 `json:"max_size,omitempty"`
	Headers map[string]string `json:"headers"`
}

type _HttpPostEndpointRequest HttpPostEndpointRequest

// NewHttpPostEndpointRequest instantiates a new HttpPostEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpPostEndpointRequest(url string, headers map[string]string) *HttpPostEndpointRequest {
	this := HttpPostEndpointRequest{}
	this.Url = url
	this.Headers = headers
	return &this
}

// NewHttpPostEndpointRequestWithDefaults instantiates a new HttpPostEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpPostEndpointRequestWithDefaults() *HttpPostEndpointRequest {
	this := HttpPostEndpointRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *HttpPostEndpointRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *HttpPostEndpointRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *HttpPostEndpointRequest) SetUrl(v string) {
	o.Url = v
}

// GetLogLineSeparator returns the LogLineSeparator field value if set, zero value otherwise.
func (o *HttpPostEndpointRequest) GetLogLineSeparator() string {
	if o == nil || IsNil(o.LogLineSeparator) {
		var ret string
		return ret
	}
	return *o.LogLineSeparator
}

// GetLogLineSeparatorOk returns a tuple with the LogLineSeparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpPostEndpointRequest) GetLogLineSeparatorOk() (*string, bool) {
	if o == nil || IsNil(o.LogLineSeparator) {
		return nil, false
	}
	return o.LogLineSeparator, true
}

// HasLogLineSeparator returns a boolean if a field has been set.
func (o *HttpPostEndpointRequest) HasLogLineSeparator() bool {
	if o != nil && !IsNil(o.LogLineSeparator) {
		return true
	}

	return false
}

// SetLogLineSeparator gets a reference to the given string and assigns it to the LogLineSeparator field.
func (o *HttpPostEndpointRequest) SetLogLineSeparator(v string) {
	o.LogLineSeparator = &v
}

// GetPayloadFormat returns the PayloadFormat field value if set, zero value otherwise.
func (o *HttpPostEndpointRequest) GetPayloadFormat() string {
	if o == nil || IsNil(o.PayloadFormat) {
		var ret string
		return ret
	}
	return *o.PayloadFormat
}

// GetPayloadFormatOk returns a tuple with the PayloadFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpPostEndpointRequest) GetPayloadFormatOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadFormat) {
		return nil, false
	}
	return o.PayloadFormat, true
}

// HasPayloadFormat returns a boolean if a field has been set.
func (o *HttpPostEndpointRequest) HasPayloadFormat() bool {
	if o != nil && !IsNil(o.PayloadFormat) {
		return true
	}

	return false
}

// SetPayloadFormat gets a reference to the given string and assigns it to the PayloadFormat field.
func (o *HttpPostEndpointRequest) SetPayloadFormat(v string) {
	o.PayloadFormat = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpPostEndpointRequest) GetMaxSize() int64 {
	if o == nil || IsNil(o.MaxSize.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxSize.Get()
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpPostEndpointRequest) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSize.Get(), o.MaxSize.IsSet()
}

// HasMaxSize returns a boolean if a field has been set.
func (o *HttpPostEndpointRequest) HasMaxSize() bool {
	if o != nil && o.MaxSize.IsSet() {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given NullableInt64 and assigns it to the MaxSize field.
func (o *HttpPostEndpointRequest) SetMaxSize(v int64) {
	o.MaxSize.Set(&v)
}
// SetMaxSizeNil sets the value for MaxSize to be an explicit nil
func (o *HttpPostEndpointRequest) SetMaxSizeNil() {
	o.MaxSize.Set(nil)
}

// UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
func (o *HttpPostEndpointRequest) UnsetMaxSize() {
	o.MaxSize.Unset()
}

// GetHeaders returns the Headers field value
func (o *HttpPostEndpointRequest) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *HttpPostEndpointRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *HttpPostEndpointRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

func (o HttpPostEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpPostEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.LogLineSeparator) {
		toSerialize["log_line_separator"] = o.LogLineSeparator
	}
	if !IsNil(o.PayloadFormat) {
		toSerialize["payload_format"] = o.PayloadFormat
	}
	if o.MaxSize.IsSet() {
		toSerialize["max_size"] = o.MaxSize.Get()
	}
	toSerialize["headers"] = o.Headers
	return toSerialize, nil
}

func (o *HttpPostEndpointRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"headers",
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

	varHttpPostEndpointRequest := _HttpPostEndpointRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHttpPostEndpointRequest)

	if err != nil {
		return err
	}

	*o = HttpPostEndpointRequest(varHttpPostEndpointRequest)

	return err
}

type NullableHttpPostEndpointRequest struct {
	value *HttpPostEndpointRequest
	isSet bool
}

func (v NullableHttpPostEndpointRequest) Get() *HttpPostEndpointRequest {
	return v.value
}

func (v *NullableHttpPostEndpointRequest) Set(val *HttpPostEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpPostEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpPostEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpPostEndpointRequest(val *HttpPostEndpointRequest) *NullableHttpPostEndpointRequest {
	return &NullableHttpPostEndpointRequest{value: val, isSet: true}
}

func (v NullableHttpPostEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpPostEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


