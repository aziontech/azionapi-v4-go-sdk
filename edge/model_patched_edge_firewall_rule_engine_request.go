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

// checks if the PatchedEdgeFirewallRuleEngineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedEdgeFirewallRuleEngineRequest{}

// PatchedEdgeFirewallRuleEngineRequest struct for PatchedEdgeFirewallRuleEngineRequest
type PatchedEdgeFirewallRuleEngineRequest struct {
	Name *string `json:"name,omitempty" validate:"regexp=.*"`
	Active *bool `json:"active,omitempty"`
	Behaviors []EdgeFirewallBehaviorFieldRequest `json:"behaviors,omitempty"`
	Criteria [][]EdgeFirewallCriterionFieldRequest `json:"criteria,omitempty"`
	Description *string `json:"description,omitempty" validate:"regexp=.*"`
}

// NewPatchedEdgeFirewallRuleEngineRequest instantiates a new PatchedEdgeFirewallRuleEngineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedEdgeFirewallRuleEngineRequest() *PatchedEdgeFirewallRuleEngineRequest {
	this := PatchedEdgeFirewallRuleEngineRequest{}
	return &this
}

// NewPatchedEdgeFirewallRuleEngineRequestWithDefaults instantiates a new PatchedEdgeFirewallRuleEngineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedEdgeFirewallRuleEngineRequestWithDefaults() *PatchedEdgeFirewallRuleEngineRequest {
	this := PatchedEdgeFirewallRuleEngineRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedEdgeFirewallRuleEngineRequest) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchedEdgeFirewallRuleEngineRequest) SetActive(v bool) {
	o.Active = &v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetBehaviors() []EdgeFirewallBehaviorFieldRequest {
	if o == nil || IsNil(o.Behaviors) {
		var ret []EdgeFirewallBehaviorFieldRequest
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetBehaviorsOk() ([]EdgeFirewallBehaviorFieldRequest, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []EdgeFirewallBehaviorFieldRequest and assigns it to the Behaviors field.
func (o *PatchedEdgeFirewallRuleEngineRequest) SetBehaviors(v []EdgeFirewallBehaviorFieldRequest) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetCriteria() [][]EdgeFirewallCriterionFieldRequest {
	if o == nil || IsNil(o.Criteria) {
		var ret [][]EdgeFirewallCriterionFieldRequest
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetCriteriaOk() ([][]EdgeFirewallCriterionFieldRequest, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given [][]EdgeFirewallCriterionFieldRequest and assigns it to the Criteria field.
func (o *PatchedEdgeFirewallRuleEngineRequest) SetCriteria(v [][]EdgeFirewallCriterionFieldRequest) {
	o.Criteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedEdgeFirewallRuleEngineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedEdgeFirewallRuleEngineRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedEdgeFirewallRuleEngineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedEdgeFirewallRuleEngineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullablePatchedEdgeFirewallRuleEngineRequest struct {
	value *PatchedEdgeFirewallRuleEngineRequest
	isSet bool
}

func (v NullablePatchedEdgeFirewallRuleEngineRequest) Get() *PatchedEdgeFirewallRuleEngineRequest {
	return v.value
}

func (v *NullablePatchedEdgeFirewallRuleEngineRequest) Set(val *PatchedEdgeFirewallRuleEngineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedEdgeFirewallRuleEngineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedEdgeFirewallRuleEngineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedEdgeFirewallRuleEngineRequest(val *PatchedEdgeFirewallRuleEngineRequest) *NullablePatchedEdgeFirewallRuleEngineRequest {
	return &NullablePatchedEdgeFirewallRuleEngineRequest{value: val, isSet: true}
}

func (v NullablePatchedEdgeFirewallRuleEngineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedEdgeFirewallRuleEngineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


