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

// checks if the PaginatedResponseListEdgeApplicationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseListEdgeApplicationList{}

// PaginatedResponseListEdgeApplicationList struct for PaginatedResponseListEdgeApplicationList
type PaginatedResponseListEdgeApplicationList struct {
	Count *int64 `json:"count,omitempty"`
	Results []ResponseListEdgeApplication `json:"results,omitempty"`
}

// NewPaginatedResponseListEdgeApplicationList instantiates a new PaginatedResponseListEdgeApplicationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseListEdgeApplicationList() *PaginatedResponseListEdgeApplicationList {
	this := PaginatedResponseListEdgeApplicationList{}
	return &this
}

// NewPaginatedResponseListEdgeApplicationListWithDefaults instantiates a new PaginatedResponseListEdgeApplicationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseListEdgeApplicationListWithDefaults() *PaginatedResponseListEdgeApplicationList {
	this := PaginatedResponseListEdgeApplicationList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedResponseListEdgeApplicationList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListEdgeApplicationList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedResponseListEdgeApplicationList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedResponseListEdgeApplicationList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedResponseListEdgeApplicationList) GetResults() []ResponseListEdgeApplication {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseListEdgeApplication
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListEdgeApplicationList) GetResultsOk() ([]ResponseListEdgeApplication, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedResponseListEdgeApplicationList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseListEdgeApplication and assigns it to the Results field.
func (o *PaginatedResponseListEdgeApplicationList) SetResults(v []ResponseListEdgeApplication) {
	o.Results = v
}

func (o PaginatedResponseListEdgeApplicationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseListEdgeApplicationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedResponseListEdgeApplicationList struct {
	value *PaginatedResponseListEdgeApplicationList
	isSet bool
}

func (v NullablePaginatedResponseListEdgeApplicationList) Get() *PaginatedResponseListEdgeApplicationList {
	return v.value
}

func (v *NullablePaginatedResponseListEdgeApplicationList) Set(val *PaginatedResponseListEdgeApplicationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseListEdgeApplicationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseListEdgeApplicationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseListEdgeApplicationList(val *PaginatedResponseListEdgeApplicationList) *NullablePaginatedResponseListEdgeApplicationList {
	return &NullablePaginatedResponseListEdgeApplicationList{value: val, isSet: true}
}

func (v NullablePaginatedResponseListEdgeApplicationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseListEdgeApplicationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


