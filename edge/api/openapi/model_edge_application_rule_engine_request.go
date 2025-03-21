/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EdgeApplicationRuleEngineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeApplicationRuleEngineRequest{}

// EdgeApplicationRuleEngineRequest struct for EdgeApplicationRuleEngineRequest
type EdgeApplicationRuleEngineRequest struct {
	Name string `json:"name" validate:"regexp=.*"`
	// * `default` - default * `request` - request * `response` - response
	Phase string `json:"phase"`
	Active *bool `json:"active,omitempty"`
	Behaviors []EdgeApplicationBehaviorFieldRequest `json:"behaviors"`
	Criteria [][]EdgeApplicationCriterionFieldRequest `json:"criteria"`
	Description *string `json:"description,omitempty" validate:"regexp=.*"`
}

type _EdgeApplicationRuleEngineRequest EdgeApplicationRuleEngineRequest

// NewEdgeApplicationRuleEngineRequest instantiates a new EdgeApplicationRuleEngineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeApplicationRuleEngineRequest(name string, phase string, behaviors []EdgeApplicationBehaviorFieldRequest, criteria [][]EdgeApplicationCriterionFieldRequest) *EdgeApplicationRuleEngineRequest {
	this := EdgeApplicationRuleEngineRequest{}
	this.Name = name
	this.Phase = phase
	var active bool = true
	this.Active = &active
	this.Behaviors = behaviors
	this.Criteria = criteria
	return &this
}

// NewEdgeApplicationRuleEngineRequestWithDefaults instantiates a new EdgeApplicationRuleEngineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeApplicationRuleEngineRequestWithDefaults() *EdgeApplicationRuleEngineRequest {
	this := EdgeApplicationRuleEngineRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetName returns the Name field value
func (o *EdgeApplicationRuleEngineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeApplicationRuleEngineRequest) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value
func (o *EdgeApplicationRuleEngineRequest) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *EdgeApplicationRuleEngineRequest) SetPhase(v string) {
	o.Phase = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeApplicationRuleEngineRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeApplicationRuleEngineRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeApplicationRuleEngineRequest) SetActive(v bool) {
	o.Active = &v
}

// GetBehaviors returns the Behaviors field value
func (o *EdgeApplicationRuleEngineRequest) GetBehaviors() []EdgeApplicationBehaviorFieldRequest {
	if o == nil {
		var ret []EdgeApplicationBehaviorFieldRequest
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetBehaviorsOk() ([]EdgeApplicationBehaviorFieldRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *EdgeApplicationRuleEngineRequest) SetBehaviors(v []EdgeApplicationBehaviorFieldRequest) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *EdgeApplicationRuleEngineRequest) GetCriteria() [][]EdgeApplicationCriterionFieldRequest {
	if o == nil {
		var ret [][]EdgeApplicationCriterionFieldRequest
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetCriteriaOk() ([][]EdgeApplicationCriterionFieldRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *EdgeApplicationRuleEngineRequest) SetCriteria(v [][]EdgeApplicationCriterionFieldRequest) {
	o.Criteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EdgeApplicationRuleEngineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeApplicationRuleEngineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EdgeApplicationRuleEngineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EdgeApplicationRuleEngineRequest) SetDescription(v string) {
	o.Description = &v
}

func (o EdgeApplicationRuleEngineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeApplicationRuleEngineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["phase"] = o.Phase
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["behaviors"] = o.Behaviors
	toSerialize["criteria"] = o.Criteria
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *EdgeApplicationRuleEngineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"phase",
		"behaviors",
		"criteria",
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

	varEdgeApplicationRuleEngineRequest := _EdgeApplicationRuleEngineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeApplicationRuleEngineRequest)

	if err != nil {
		return err
	}

	*o = EdgeApplicationRuleEngineRequest(varEdgeApplicationRuleEngineRequest)

	return err
}

type NullableEdgeApplicationRuleEngineRequest struct {
	value *EdgeApplicationRuleEngineRequest
	isSet bool
}

func (v NullableEdgeApplicationRuleEngineRequest) Get() *EdgeApplicationRuleEngineRequest {
	return v.value
}

func (v *NullableEdgeApplicationRuleEngineRequest) Set(val *EdgeApplicationRuleEngineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeApplicationRuleEngineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeApplicationRuleEngineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeApplicationRuleEngineRequest(val *EdgeApplicationRuleEngineRequest) *NullableEdgeApplicationRuleEngineRequest {
	return &NullableEdgeApplicationRuleEngineRequest{value: val, isSet: true}
}

func (v NullableEdgeApplicationRuleEngineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeApplicationRuleEngineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


