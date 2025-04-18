/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ResponseListEdgeFirewallRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListEdgeFirewallRuleEngine{}

// ResponseListEdgeFirewallRuleEngine struct for ResponseListEdgeFirewallRuleEngine
type ResponseListEdgeFirewallRuleEngine struct {
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

type _ResponseListEdgeFirewallRuleEngine ResponseListEdgeFirewallRuleEngine

// NewResponseListEdgeFirewallRuleEngine instantiates a new ResponseListEdgeFirewallRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListEdgeFirewallRuleEngine(id int64, name string, lastEditor string, lastModified time.Time, behaviors []EdgeFirewallBehaviorField, criteria [][]EdgeFirewallCriterionField, order int64) *ResponseListEdgeFirewallRuleEngine {
	this := ResponseListEdgeFirewallRuleEngine{}
	this.Id = id
	this.Name = name
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	this.Behaviors = behaviors
	this.Criteria = criteria
	this.Order = order
	return &this
}

// NewResponseListEdgeFirewallRuleEngineWithDefaults instantiates a new ResponseListEdgeFirewallRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListEdgeFirewallRuleEngineWithDefaults() *ResponseListEdgeFirewallRuleEngine {
	this := ResponseListEdgeFirewallRuleEngine{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListEdgeFirewallRuleEngine) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListEdgeFirewallRuleEngine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetName(v string) {
	o.Name = v
}

// GetLastEditor returns the LastEditor field value
func (o *ResponseListEdgeFirewallRuleEngine) GetLastEditor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastEditor, true
}

// SetLastEditor sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetLastEditor(v string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value
func (o *ResponseListEdgeFirewallRuleEngine) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListEdgeFirewallRuleEngine) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListEdgeFirewallRuleEngine) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListEdgeFirewallRuleEngine) SetActive(v bool) {
	o.Active = &v
}

// GetBehaviors returns the Behaviors field value
func (o *ResponseListEdgeFirewallRuleEngine) GetBehaviors() []EdgeFirewallBehaviorField {
	if o == nil {
		var ret []EdgeFirewallBehaviorField
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetBehaviorsOk() ([]EdgeFirewallBehaviorField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetBehaviors(v []EdgeFirewallBehaviorField) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *ResponseListEdgeFirewallRuleEngine) GetCriteria() [][]EdgeFirewallCriterionField {
	if o == nil {
		var ret [][]EdgeFirewallCriterionField
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetCriteriaOk() ([][]EdgeFirewallCriterionField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetCriteria(v [][]EdgeFirewallCriterionField) {
	o.Criteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResponseListEdgeFirewallRuleEngine) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResponseListEdgeFirewallRuleEngine) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResponseListEdgeFirewallRuleEngine) SetDescription(v string) {
	o.Description = &v
}

// GetOrder returns the Order field value
func (o *ResponseListEdgeFirewallRuleEngine) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeFirewallRuleEngine) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ResponseListEdgeFirewallRuleEngine) SetOrder(v int64) {
	o.Order = v
}

func (o ResponseListEdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListEdgeFirewallRuleEngine) ToMap() (map[string]interface{}, error) {
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

func (o *ResponseListEdgeFirewallRuleEngine) UnmarshalJSON(data []byte) (err error) {
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

	varResponseListEdgeFirewallRuleEngine := _ResponseListEdgeFirewallRuleEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListEdgeFirewallRuleEngine)

	if err != nil {
		return err
	}

	*o = ResponseListEdgeFirewallRuleEngine(varResponseListEdgeFirewallRuleEngine)

	return err
}

type NullableResponseListEdgeFirewallRuleEngine struct {
	value *ResponseListEdgeFirewallRuleEngine
	isSet bool
}

func (v NullableResponseListEdgeFirewallRuleEngine) Get() *ResponseListEdgeFirewallRuleEngine {
	return v.value
}

func (v *NullableResponseListEdgeFirewallRuleEngine) Set(val *ResponseListEdgeFirewallRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListEdgeFirewallRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListEdgeFirewallRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListEdgeFirewallRuleEngine(val *ResponseListEdgeFirewallRuleEngine) *NullableResponseListEdgeFirewallRuleEngine {
	return &NullableResponseListEdgeFirewallRuleEngine{value: val, isSet: true}
}

func (v NullableResponseListEdgeFirewallRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListEdgeFirewallRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


