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

// checks if the ResponseBadRequestEdgeApplicationRuleEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBadRequestEdgeApplicationRuleEngine{}

// ResponseBadRequestEdgeApplicationRuleEngine struct for ResponseBadRequestEdgeApplicationRuleEngine
type ResponseBadRequestEdgeApplicationRuleEngine struct {
	Name []string `json:"name,omitempty"`
	Phase []string `json:"phase,omitempty"`
	Active []string `json:"active,omitempty"`
	Description []string `json:"description,omitempty"`
	Order []string `json:"order,omitempty"`
	Behaviors *ResponseBadRequestBaseEdgeConnectorConnectionPreference `json:"behaviors,omitempty"`
	Criteria *ResponseBadRequestBaseEdgeConnectorConnectionPreference `json:"criteria,omitempty"`
	LastEditor []string `json:"last_editor,omitempty"`
	LastModified []string `json:"last_modified,omitempty"`
	Detail *string `json:"detail,omitempty" validate:"regexp=.*"`
}

// NewResponseBadRequestEdgeApplicationRuleEngine instantiates a new ResponseBadRequestEdgeApplicationRuleEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBadRequestEdgeApplicationRuleEngine() *ResponseBadRequestEdgeApplicationRuleEngine {
	this := ResponseBadRequestEdgeApplicationRuleEngine{}
	return &this
}

// NewResponseBadRequestEdgeApplicationRuleEngineWithDefaults instantiates a new ResponseBadRequestEdgeApplicationRuleEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBadRequestEdgeApplicationRuleEngineWithDefaults() *ResponseBadRequestEdgeApplicationRuleEngine {
	this := ResponseBadRequestEdgeApplicationRuleEngine{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetName(v []string) {
	o.Name = v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetPhase() []string {
	if o == nil || IsNil(o.Phase) {
		var ret []string
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetPhaseOk() ([]string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given []string and assigns it to the Phase field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetPhase(v []string) {
	o.Phase = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetActive() []string {
	if o == nil || IsNil(o.Active) {
		var ret []string
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetActiveOk() ([]string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []string and assigns it to the Active field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetActive(v []string) {
	o.Active = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetDescription() []string {
	if o == nil || IsNil(o.Description) {
		var ret []string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetDescriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []string and assigns it to the Description field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetDescription(v []string) {
	o.Description = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetOrder() []string {
	if o == nil || IsNil(o.Order) {
		var ret []string
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given []string and assigns it to the Order field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetOrder(v []string) {
	o.Order = v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetBehaviors() ResponseBadRequestBaseEdgeConnectorConnectionPreference {
	if o == nil || IsNil(o.Behaviors) {
		var ret ResponseBadRequestBaseEdgeConnectorConnectionPreference
		return ret
	}
	return *o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetBehaviorsOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given ResponseBadRequestBaseEdgeConnectorConnectionPreference and assigns it to the Behaviors field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetBehaviors(v ResponseBadRequestBaseEdgeConnectorConnectionPreference) {
	o.Behaviors = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetCriteria() ResponseBadRequestBaseEdgeConnectorConnectionPreference {
	if o == nil || IsNil(o.Criteria) {
		var ret ResponseBadRequestBaseEdgeConnectorConnectionPreference
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetCriteriaOk() (*ResponseBadRequestBaseEdgeConnectorConnectionPreference, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given ResponseBadRequestBaseEdgeConnectorConnectionPreference and assigns it to the Criteria field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetCriteria(v ResponseBadRequestBaseEdgeConnectorConnectionPreference) {
	o.Criteria = &v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetLastEditor() []string {
	if o == nil || IsNil(o.LastEditor) {
		var ret []string
		return ret
	}
	return o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetLastEditorOk() ([]string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given []string and assigns it to the LastEditor field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetLastEditor(v []string) {
	o.LastEditor = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetLastModified() []string {
	if o == nil || IsNil(o.LastModified) {
		var ret []string
		return ret
	}
	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetLastModifiedOk() ([]string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given []string and assigns it to the LastModified field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetLastModified(v []string) {
	o.LastModified = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ResponseBadRequestEdgeApplicationRuleEngine) SetDetail(v string) {
	o.Detail = &v
}

func (o ResponseBadRequestEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBadRequestEdgeApplicationRuleEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableResponseBadRequestEdgeApplicationRuleEngine struct {
	value *ResponseBadRequestEdgeApplicationRuleEngine
	isSet bool
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngine) Get() *ResponseBadRequestEdgeApplicationRuleEngine {
	return v.value
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngine) Set(val *ResponseBadRequestEdgeApplicationRuleEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestEdgeApplicationRuleEngine(val *ResponseBadRequestEdgeApplicationRuleEngine) *NullableResponseBadRequestEdgeApplicationRuleEngine {
	return &NullableResponseBadRequestEdgeApplicationRuleEngine{value: val, isSet: true}
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


