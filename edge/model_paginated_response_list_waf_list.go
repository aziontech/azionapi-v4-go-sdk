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

// checks if the PaginatedResponseListWAFList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseListWAFList{}

// PaginatedResponseListWAFList struct for PaginatedResponseListWAFList
type PaginatedResponseListWAFList struct {
	Count *int32 `json:"count,omitempty"`
	Results []ResponseListWAF `json:"results,omitempty"`
}

// NewPaginatedResponseListWAFList instantiates a new PaginatedResponseListWAFList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseListWAFList() *PaginatedResponseListWAFList {
	this := PaginatedResponseListWAFList{}
	return &this
}

// NewPaginatedResponseListWAFListWithDefaults instantiates a new PaginatedResponseListWAFList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseListWAFListWithDefaults() *PaginatedResponseListWAFList {
	this := PaginatedResponseListWAFList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedResponseListWAFList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListWAFList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedResponseListWAFList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedResponseListWAFList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedResponseListWAFList) GetResults() []ResponseListWAF {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseListWAF
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListWAFList) GetResultsOk() ([]ResponseListWAF, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedResponseListWAFList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseListWAF and assigns it to the Results field.
func (o *PaginatedResponseListWAFList) SetResults(v []ResponseListWAF) {
	o.Results = v
}

func (o PaginatedResponseListWAFList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseListWAFList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedResponseListWAFList struct {
	value *PaginatedResponseListWAFList
	isSet bool
}

func (v NullablePaginatedResponseListWAFList) Get() *PaginatedResponseListWAFList {
	return v.value
}

func (v *NullablePaginatedResponseListWAFList) Set(val *PaginatedResponseListWAFList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseListWAFList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseListWAFList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseListWAFList(val *PaginatedResponseListWAFList) *NullablePaginatedResponseListWAFList {
	return &NullablePaginatedResponseListWAFList{value: val, isSet: true}
}

func (v NullablePaginatedResponseListWAFList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseListWAFList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


