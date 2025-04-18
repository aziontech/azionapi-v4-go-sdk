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

// checks if the ResponseListEdgeApplicationRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListEdgeApplicationRuleEngine{}

// ResponseListEdgeApplicationRuleEngine struct for ResponseListEdgeApplicationRuleEngine
type ResponseListEdgeApplicationRuleEngine struct {
	Id int64 `json:"id"`
	Name string `json:"name" validate:"regexp=.*"`
	// * `default` - default * `request` - request * `response` - response
	Phase string `json:"phase"`
	Active *bool `json:"active,omitempty"`
	Behaviors []EdgeApplicationBehaviorField `json:"behaviors"`
	Criteria [][]EdgeApplicationCriterionField `json:"criteria"`
	Description *string `json:"description,omitempty" validate:"regexp=.*"`
	Order int64 `json:"order"`
	LastEditor NullableString `json:"last_editor" validate:"regexp=.*"`
	LastModified NullableTime `json:"last_modified"`
}

type _ResponseListEdgeApplicationRuleEngine ResponseListEdgeApplicationRuleEngine

// NewResponseListEdgeApplicationRuleEngine instantiates a new ResponseListEdgeApplicationRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListEdgeApplicationRuleEngine(id int64, name string, phase string, behaviors []EdgeApplicationBehaviorField, criteria [][]EdgeApplicationCriterionField, order int64, lastEditor NullableString, lastModified NullableTime) *ResponseListEdgeApplicationRuleEngine {
	this := ResponseListEdgeApplicationRuleEngine{}
	this.Id = id
	this.Name = name
	this.Phase = phase
	this.Behaviors = behaviors
	this.Criteria = criteria
	this.Order = order
	this.LastEditor = lastEditor
	this.LastModified = lastModified
	return &this
}

// NewResponseListEdgeApplicationRuleEngineWithDefaults instantiates a new ResponseListEdgeApplicationRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListEdgeApplicationRuleEngineWithDefaults() *ResponseListEdgeApplicationRuleEngine {
	this := ResponseListEdgeApplicationRuleEngine{}
	return &this
}

// GetId returns the Id field value
func (o *ResponseListEdgeApplicationRuleEngine) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResponseListEdgeApplicationRuleEngine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value
func (o *ResponseListEdgeApplicationRuleEngine) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetPhase(v string) {
	o.Phase = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseListEdgeApplicationRuleEngine) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseListEdgeApplicationRuleEngine) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ResponseListEdgeApplicationRuleEngine) SetActive(v bool) {
	o.Active = &v
}

// GetBehaviors returns the Behaviors field value
func (o *ResponseListEdgeApplicationRuleEngine) GetBehaviors() []EdgeApplicationBehaviorField {
	if o == nil {
		var ret []EdgeApplicationBehaviorField
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetBehaviorsOk() ([]EdgeApplicationBehaviorField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetBehaviors(v []EdgeApplicationBehaviorField) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *ResponseListEdgeApplicationRuleEngine) GetCriteria() [][]EdgeApplicationCriterionField {
	if o == nil {
		var ret [][]EdgeApplicationCriterionField
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetCriteriaOk() ([][]EdgeApplicationCriterionField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetCriteria(v [][]EdgeApplicationCriterionField) {
	o.Criteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResponseListEdgeApplicationRuleEngine) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResponseListEdgeApplicationRuleEngine) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResponseListEdgeApplicationRuleEngine) SetDescription(v string) {
	o.Description = &v
}

// GetOrder returns the Order field value
func (o *ResponseListEdgeApplicationRuleEngine) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *ResponseListEdgeApplicationRuleEngine) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetOrder(v int64) {
	o.Order = v
}

// GetLastEditor returns the LastEditor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ResponseListEdgeApplicationRuleEngine) GetLastEditor() string {
	if o == nil || o.LastEditor.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastEditor.Get()
}

// GetLastEditorOk returns a tuple with the LastEditor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseListEdgeApplicationRuleEngine) GetLastEditorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastEditor.Get(), o.LastEditor.IsSet()
}

// SetLastEditor sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetLastEditor(v string) {
	o.LastEditor.Set(&v)
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ResponseListEdgeApplicationRuleEngine) GetLastModified() time.Time {
	if o == nil || o.LastModified.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseListEdgeApplicationRuleEngine) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// SetLastModified sets field value
func (o *ResponseListEdgeApplicationRuleEngine) SetLastModified(v time.Time) {
	o.LastModified.Set(&v)
}

func (o ResponseListEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListEdgeApplicationRuleEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	toSerialize["order"] = o.Order
	toSerialize["last_editor"] = o.LastEditor.Get()
	toSerialize["last_modified"] = o.LastModified.Get()
	return toSerialize, nil
}

func (o *ResponseListEdgeApplicationRuleEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"phase",
		"behaviors",
		"criteria",
		"order",
		"last_editor",
		"last_modified",
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

	varResponseListEdgeApplicationRuleEngine := _ResponseListEdgeApplicationRuleEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseListEdgeApplicationRuleEngine)

	if err != nil {
		return err
	}

	*o = ResponseListEdgeApplicationRuleEngine(varResponseListEdgeApplicationRuleEngine)

	return err
}

type NullableResponseListEdgeApplicationRuleEngine struct {
	value *ResponseListEdgeApplicationRuleEngine
	isSet bool
}

func (v NullableResponseListEdgeApplicationRuleEngine) Get() *ResponseListEdgeApplicationRuleEngine {
	return v.value
}

func (v *NullableResponseListEdgeApplicationRuleEngine) Set(val *ResponseListEdgeApplicationRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListEdgeApplicationRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListEdgeApplicationRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListEdgeApplicationRuleEngine(val *ResponseListEdgeApplicationRuleEngine) *NullableResponseListEdgeApplicationRuleEngine {
	return &NullableResponseListEdgeApplicationRuleEngine{value: val, isSet: true}
}

func (v NullableResponseListEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListEdgeApplicationRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


