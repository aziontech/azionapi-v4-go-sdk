/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
)

// checks if the ResponseBadRequestProtocolsSerializerHttpField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestProtocolsSerializerHttpField{}

// ResponseBadRequestProtocolsSerializerHttpField struct for ResponseBadRequestProtocolsSerializerHttpField
type ResponseBadRequestProtocolsSerializerHttpField struct {
	Versions *ResponseBadRequestEdgeApplicationRuleEngineBehaviors `json:"versions,omitempty"`
	HttpPorts *ResponseBadRequestEdgeApplicationRuleEngineBehaviors `json:"http_ports,omitempty"`
	HttpsPorts *ResponseBadRequestEdgeApplicationRuleEngineBehaviors `json:"https_ports,omitempty"`
	QuicPorts *ResponseBadRequestEdgeApplicationRuleEngineBehaviors `json:"quic_ports,omitempty"`
}

// NewResponseBadRequestProtocolsSerializerHttpField instantiates a new ResponseBadRequestProtocolsSerializerHttpField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestProtocolsSerializerHttpField() *ResponseBadRequestProtocolsSerializerHttpField {
	this := ResponseBadRequestProtocolsSerializerHttpField{}
	return &this
}

// NewResponseBadRequestProtocolsSerializerHttpFieldWithDefaults instantiates a new ResponseBadRequestProtocolsSerializerHttpField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestProtocolsSerializerHttpFieldWithDefaults() *ResponseBadRequestProtocolsSerializerHttpField {
	this := ResponseBadRequestProtocolsSerializerHttpField{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetVersions() ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	if o == nil || IsNil(o.Versions) {
		var ret ResponseBadRequestEdgeApplicationRuleEngineBehaviors
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetVersionsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given ResponseBadRequestEdgeApplicationRuleEngineBehaviors and assigns it to the Versions field.
func (o *ResponseBadRequestProtocolsSerializerHttpField) SetVersions(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors) {
	o.Versions = &v
}

// GetHttpPorts returns the HttpPorts field value if set, zero value otherwise.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetHttpPorts() ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	if o == nil || IsNil(o.HttpPorts) {
		var ret ResponseBadRequestEdgeApplicationRuleEngineBehaviors
		return ret
	}
	return *o.HttpPorts
}

// GetHttpPortsOk returns a tuple with the HttpPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetHttpPortsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool) {
	if o == nil || IsNil(o.HttpPorts) {
		return nil, false
	}
	return o.HttpPorts, true
}

// HasHttpPorts returns a boolean if a field has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) HasHttpPorts() bool {
	if o != nil && !IsNil(o.HttpPorts) {
		return true
	}

	return false
}

// SetHttpPorts gets a reference to the given ResponseBadRequestEdgeApplicationRuleEngineBehaviors and assigns it to the HttpPorts field.
func (o *ResponseBadRequestProtocolsSerializerHttpField) SetHttpPorts(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors) {
	o.HttpPorts = &v
}

// GetHttpsPorts returns the HttpsPorts field value if set, zero value otherwise.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetHttpsPorts() ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	if o == nil || IsNil(o.HttpsPorts) {
		var ret ResponseBadRequestEdgeApplicationRuleEngineBehaviors
		return ret
	}
	return *o.HttpsPorts
}

// GetHttpsPortsOk returns a tuple with the HttpsPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetHttpsPortsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool) {
	if o == nil || IsNil(o.HttpsPorts) {
		return nil, false
	}
	return o.HttpsPorts, true
}

// HasHttpsPorts returns a boolean if a field has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) HasHttpsPorts() bool {
	if o != nil && !IsNil(o.HttpsPorts) {
		return true
	}

	return false
}

// SetHttpsPorts gets a reference to the given ResponseBadRequestEdgeApplicationRuleEngineBehaviors and assigns it to the HttpsPorts field.
func (o *ResponseBadRequestProtocolsSerializerHttpField) SetHttpsPorts(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors) {
	o.HttpsPorts = &v
}

// GetQuicPorts returns the QuicPorts field value if set, zero value otherwise.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetQuicPorts() ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	if o == nil || IsNil(o.QuicPorts) {
		var ret ResponseBadRequestEdgeApplicationRuleEngineBehaviors
		return ret
	}
	return *o.QuicPorts
}

// GetQuicPortsOk returns a tuple with the QuicPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) GetQuicPortsOk() (*ResponseBadRequestEdgeApplicationRuleEngineBehaviors, bool) {
	if o == nil || IsNil(o.QuicPorts) {
		return nil, false
	}
	return o.QuicPorts, true
}

// HasQuicPorts returns a boolean if a field has been set.
func (o *ResponseBadRequestProtocolsSerializerHttpField) HasQuicPorts() bool {
	if o != nil && !IsNil(o.QuicPorts) {
		return true
	}

	return false
}

// SetQuicPorts gets a reference to the given ResponseBadRequestEdgeApplicationRuleEngineBehaviors and assigns it to the QuicPorts field.
func (o *ResponseBadRequestProtocolsSerializerHttpField) SetQuicPorts(v ResponseBadRequestEdgeApplicationRuleEngineBehaviors) {
	o.QuicPorts = &v
}

func (o ResponseBadRequestProtocolsSerializerHttpField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestProtocolsSerializerHttpField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.HttpPorts) {
		toSerialize["http_ports"] = o.HttpPorts
	}
	if !IsNil(o.HttpsPorts) {
		toSerialize["https_ports"] = o.HttpsPorts
	}
	if !IsNil(o.QuicPorts) {
		toSerialize["quic_ports"] = o.QuicPorts
	}
	return toSerialize, nil
}

type NullableResponseBadRequestProtocolsSerializerHttpField struct {
	value *ResponseBadRequestProtocolsSerializerHttpField
	isSet bool
}

func (v NullableResponseBadRequestProtocolsSerializerHttpField) Get() *ResponseBadRequestProtocolsSerializerHttpField {
	return v.value
}

func (v *NullableResponseBadRequestProtocolsSerializerHttpField) Set(val *ResponseBadRequestProtocolsSerializerHttpField) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestProtocolsSerializerHttpField) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestProtocolsSerializerHttpField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestProtocolsSerializerHttpField(val *ResponseBadRequestProtocolsSerializerHttpField) *NullableResponseBadRequestProtocolsSerializerHttpField {
	return &NullableResponseBadRequestProtocolsSerializerHttpField{value: val, isSet: true}
}

func (v NullableResponseBadRequestProtocolsSerializerHttpField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestProtocolsSerializerHttpField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


