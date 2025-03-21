/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge/api/openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the EdgeFirewallRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallRuleEngine{}

// EdgeFirewallRuleEngine struct for EdgeFirewallRuleEngine
type EdgeFirewallRuleEngine struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	LastEditor string `json:"last_editor" validate:"regexp=.*"`
	LastModified time.Time `json:"last_modified"`
	Active *bool `json:"active,omitempty"`
	Behaviors []EdgeFirewallBehaviorField `json:"behaviors"`
	Criteria [][]EdgeFirewallCriterionField `json:"criteria"`
	Description *string `json:"description,omitempty" validate:"regexp=.*"`
	Order int64 `json:"order"`
}

type _EdgeFirewallRuleEngine EdgeFirewallRuleEngine

// NewEdgeFirewallRuleEngine instantiates a new EdgeFirewallRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallRuleEngine(id int64, name string, lastEditor string, lastModified time.Time, behaviors []EdgeFirewallBehaviorField, criteria [][]EdgeFirewallCriterionField, order int64) *EdgeFirewallRuleEngine {
	this := EdgeFirewallRuleEngine{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	var active bool = true
	this.Active = &active
	this.Behaviors = behaviors
	this.Criteria = criteria
	this.Order = order
	return &this
}

// NewEdgeFirewallRuleEngineWithDefaults instantiates a new EdgeFirewallRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallRuleEngineWithDefaults() *EdgeFirewallRuleEngine {
	this := EdgeFirewallRuleEngine{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetId returns the Id field value
func (o *EdgeFirewallRuleEngine) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeFirewallRuleEngine) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EdgeFirewallRuleEngine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeFirewallRuleEngine) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *EdgeFirewallRuleEngine) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *EdgeFirewallRuleEngine) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *EdgeFirewallRuleEngine) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *EdgeFirewallRuleEngine) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EdgeFirewallRuleEngine) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EdgeFirewallRuleEngine) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EdgeFirewallRuleEngine) SetActive(v bool) {
	o.Active = &v
}

// GetBehaviors returns the Behaviors field value
func (o *EdgeFirewallRuleEngine) GetBehaviors() []EdgeFirewallBehaviorField {
	if o == nil {
		var ret []EdgeFirewallBehaviorField
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetBehaviorsOk() ([]EdgeFirewallBehaviorField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *EdgeFirewallRuleEngine) SetBehaviors(v []EdgeFirewallBehaviorField) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *EdgeFirewallRuleEngine) GetCriteria() [][]EdgeFirewallCriterionField {
	if o == nil {
		var ret [][]EdgeFirewallCriterionField
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetCriteriaOk() ([][]EdgeFirewallCriterionField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *EdgeFirewallRuleEngine) SetCriteria(v [][]EdgeFirewallCriterionField) {
	o.Criteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EdgeFirewallRuleEngine) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EdgeFirewallRuleEngine) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EdgeFirewallRuleEngine) SetDescription(v string) {
	o.Description = &v
}

// GetOrder returns the Order field value
func (o *EdgeFirewallRuleEngine) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallRuleEngine) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *EdgeFirewallRuleEngine) SetOrder(v int64) {
	o.Order = v
}

func (o EdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallRuleEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["last_editor"] = o.LastEditor
	toSerialize["last_modified"] = o.LastModified
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["behaviors"] = o.Behaviors
	toSerialize["criteria"] = o.Criteria
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

func (o *EdgeFirewallRuleEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"last_editor",
		"last_modified",
		"behaviors",
		"criteria",
		"order",
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

	varEdgeFirewallRuleEngine := _EdgeFirewallRuleEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeFirewallRuleEngine)

	if err != nil {
		return err
	}

	*o = EdgeFirewallRuleEngine(varEdgeFirewallRuleEngine)

	return err
}

type NullableEdgeFirewallRuleEngine struct {
	value *EdgeFirewallRuleEngine
	isSet bool
}

func (v NullableEdgeFirewallRuleEngine) Get() *EdgeFirewallRuleEngine {
	return v.value
}

func (v *NullableEdgeFirewallRuleEngine) Set(val *EdgeFirewallRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallRuleEngine(val *EdgeFirewallRuleEngine) *NullableEdgeFirewallRuleEngine {
	return &NullableEdgeFirewallRuleEngine{value: val, isSet: true}
}

func (v NullableEdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


