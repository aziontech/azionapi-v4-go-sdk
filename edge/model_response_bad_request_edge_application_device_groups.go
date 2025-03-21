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

// checks if the ResponseBadRequestEdgeApplicationDeviceGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestEdgeApplicationDeviceGroups{}

// ResponseBadRequestEdgeApplicationDeviceGroups struct for ResponseBadRequestEdgeApplicationDeviceGroups
type ResponseBadRequestEdgeApplicationDeviceGroups struct {
	Name []string `json:"name,omitempty"`
	Id []string `json:"id,omitempty"`
	UserAgent []string `json:"user_agent,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestEdgeApplicationDeviceGroups instantiates a new ResponseBadRequestEdgeApplicationDeviceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestEdgeApplicationDeviceGroups() *ResponseBadRequestEdgeApplicationDeviceGroups {
	this := ResponseBadRequestEdgeApplicationDeviceGroups{}
	return &this
}

// NewResponseBadRequestEdgeApplicationDeviceGroupsWithDefaults instantiates a new ResponseBadRequestEdgeApplicationDeviceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestEdgeApplicationDeviceGroupsWithDefaults() *ResponseBadRequestEdgeApplicationDeviceGroups {
	this := ResponseBadRequestEdgeApplicationDeviceGroups{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) SetName(v []string) {
	o.Name = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetId() []string {
	if o == nil || IsNil(o.Id) {
		var ret []string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetIdOk() ([]string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []string and assigns it to the Id field.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) SetId(v []string) {
	o.Id = v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetUserAgent() []string {
	if o == nil || IsNil(o.UserAgent) {
		var ret []string
		return ret
	}
	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetUserAgentOk() ([]string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given []string and assigns it to the UserAgent field.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) SetUserAgent(v []string) {
	o.UserAgent = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestEdgeApplicationDeviceGroups) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestEdgeApplicationDeviceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestEdgeApplicationDeviceGroups struct {
	value *ResponseBadRequestEdgeApplicationDeviceGroups
	isSet bool
}

func (v NullableResponseBadRequestEdgeApplicationDeviceGroups) Get() *ResponseBadRequestEdgeApplicationDeviceGroups {
	return v.value
}

func (v *NullableResponseBadRequestEdgeApplicationDeviceGroups) Set(val *ResponseBadRequestEdgeApplicationDeviceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestEdgeApplicationDeviceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestEdgeApplicationDeviceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestEdgeApplicationDeviceGroups(val *ResponseBadRequestEdgeApplicationDeviceGroups) *NullableResponseBadRequestEdgeApplicationDeviceGroups {
	return &NullableResponseBadRequestEdgeApplicationDeviceGroups{value: val, isSet: true}
}

func (v NullableResponseBadRequestEdgeApplicationDeviceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestEdgeApplicationDeviceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


