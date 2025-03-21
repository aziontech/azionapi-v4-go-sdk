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

// checks if the PaginatedResponseListEdgeFirewallRuleEngineList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseListEdgeFirewallRuleEngineList{}

// PaginatedResponseListEdgeFirewallRuleEngineList struct for PaginatedResponseListEdgeFirewallRuleEngineList
type PaginatedResponseListEdgeFirewallRuleEngineList struct {
	Count *int64 `json:"count,omitempty"`
	Results []ResponseListEdgeFirewallRuleEngine `json:"results,omitempty"`
}

// NewPaginatedResponseListEdgeFirewallRuleEngineList instantiates a new PaginatedResponseListEdgeFirewallRuleEngineList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseListEdgeFirewallRuleEngineList() *PaginatedResponseListEdgeFirewallRuleEngineList {
	this := PaginatedResponseListEdgeFirewallRuleEngineList{}
	return &this
}

// NewPaginatedResponseListEdgeFirewallRuleEngineListWithDefaults instantiates a new PaginatedResponseListEdgeFirewallRuleEngineList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseListEdgeFirewallRuleEngineListWithDefaults() *PaginatedResponseListEdgeFirewallRuleEngineList {
	this := PaginatedResponseListEdgeFirewallRuleEngineList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) SetCount(v int64) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) GetResults() []ResponseListEdgeFirewallRuleEngine {
	if o == nil || IsNil(o.Results) {
		var ret []ResponseListEdgeFirewallRuleEngine
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) GetResultsOk() ([]ResponseListEdgeFirewallRuleEngine, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ResponseListEdgeFirewallRuleEngine and assigns it to the Results field.
func (o *PaginatedResponseListEdgeFirewallRuleEngineList) SetResults(v []ResponseListEdgeFirewallRuleEngine) {
	o.Results = v
}

func (o PaginatedResponseListEdgeFirewallRuleEngineList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseListEdgeFirewallRuleEngineList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedResponseListEdgeFirewallRuleEngineList struct {
	value *PaginatedResponseListEdgeFirewallRuleEngineList
	isSet bool
}

func (v NullablePaginatedResponseListEdgeFirewallRuleEngineList) Get() *PaginatedResponseListEdgeFirewallRuleEngineList {
	return v.value
}

func (v *NullablePaginatedResponseListEdgeFirewallRuleEngineList) Set(val *PaginatedResponseListEdgeFirewallRuleEngineList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseListEdgeFirewallRuleEngineList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseListEdgeFirewallRuleEngineList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseListEdgeFirewallRuleEngineList(val *PaginatedResponseListEdgeFirewallRuleEngineList) *NullablePaginatedResponseListEdgeFirewallRuleEngineList {
	return &NullablePaginatedResponseListEdgeFirewallRuleEngineList{value: val, isSet: true}
}

func (v NullablePaginatedResponseListEdgeFirewallRuleEngineList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseListEdgeFirewallRuleEngineList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


