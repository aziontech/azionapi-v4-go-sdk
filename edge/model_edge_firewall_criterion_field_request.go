/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EdgeFirewallCriterionFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFirewallCriterionFieldRequest{}

// EdgeFirewallCriterionFieldRequest The criterion which will be evaluated to define if the configured behaviors for this rule can be executed.  | Variable | Description | Operators | Argument | | -------- | ----------- | --------- | ---------| | ${header_accept} |  | matches, does_not_match | string | | ${header_accept_encoding} |  | matches, does_not_match | string | | ${header_accept_language} |  | matches, does_not_match | string | | ${header_cookie} |  | matches, does_not_match | string | | ${header_origin} |  | matches, does_not_match | string | | ${header_referer} |  | matches, does_not_match | string | | ${header_user_agent} |  | matches, does_not_match | string | | ${host} |  | is_equal, is_not_equal, matches, does_not_match | string | | ${network} |  | is_in_list, is_not_in_list | string | | ${request_args} |  | is_equal, is_not_equal, matches, does_not_match, exists, does_not_exist | string | | ${request_method} |  | is_equal, is_not_equal | string | | ${request_uri} |  | starts_with, does_not_starts_with, is_equal, is_not_equal, matches, does_not_match | string | | ${scheme} |  | is_equal, is_not_equal | string | | ${ssl_verification_status} |  | is_equal, is_not_equal | SUCCESS, CERTIFICATE_VERIFICATION_ERROR, MISSING_CLIENT_CERTIFICATE | | ${client_certificate_validation} |  | is_equal, is_not_equal | string |   About `operator` field: it's the operator to be used to evaluate the current criterion. When used in the first criterion of a block it should be always the `if` operator.
type EdgeFirewallCriterionFieldRequest struct {
	Argument NullableEdgeFirewallCriterionPolymorphicArgumentRequest `json:"argument,omitempty"`
	// * `${header_accept}` - ${header_accept} * `${header_accept_encoding}` - ${header_accept_encoding} * `${header_accept_language}` - ${header_accept_language} * `${header_cookie}` - ${header_cookie} * `${header_origin}` - ${header_origin} * `${header_referer}` - ${header_referer} * `${header_user_agent}` - ${header_user_agent} * `${host}` - ${host} * `${network}` - ${network} * `${request_args}` - ${request_args} * `${request_method}` - ${request_method} * `${request_uri}` - ${request_uri} * `${scheme}` - ${scheme} * `${ssl_verification_status}` - ${ssl_verification_status} * `${client_certificate_validation}` - ${client_certificate_validation}
	Variable string `json:"variable"`
	// * `if` - if * `or` - or * `and` - and
	Conditional string `json:"conditional"`
	// * `does_not_exist` - does_not_exist * `does_not_match` - does_not_match * `does_not_start_with` - does_not_start_with * `exists` - exists * `is_equal` - is_equal * `is_in_list` - is_in_list * `is_not_equal` - is_not_equal * `is_not_in_list` - is_not_in_list * `matches` - matches * `starts_with` - starts_with
	Operator string `json:"operator"`
}

type _EdgeFirewallCriterionFieldRequest EdgeFirewallCriterionFieldRequest

// NewEdgeFirewallCriterionFieldRequest instantiates a new EdgeFirewallCriterionFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFirewallCriterionFieldRequest(variable string, conditional string, operator string) *EdgeFirewallCriterionFieldRequest {
	this := EdgeFirewallCriterionFieldRequest{}
	this.Variable = variable
	this.Conditional = conditional
	this.Operator = operator
	return &this
}

// NewEdgeFirewallCriterionFieldRequestWithDefaults instantiates a new EdgeFirewallCriterionFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFirewallCriterionFieldRequestWithDefaults() *EdgeFirewallCriterionFieldRequest {
	this := EdgeFirewallCriterionFieldRequest{}
	return &this
}

// GetArgument returns the Argument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EdgeFirewallCriterionFieldRequest) GetArgument() EdgeFirewallCriterionPolymorphicArgumentRequest {
	if o == nil || IsNil(o.Argument.Get()) {
		var ret EdgeFirewallCriterionPolymorphicArgumentRequest
		return ret
	}
	return *o.Argument.Get()
}

// GetArgumentOk returns a tuple with the Argument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EdgeFirewallCriterionFieldRequest) GetArgumentOk() (*EdgeFirewallCriterionPolymorphicArgumentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Argument.Get(), o.Argument.IsSet()
}

// HasArgument returns a boolean if a field has been set.
func (o *EdgeFirewallCriterionFieldRequest) HasArgument() bool {
	if o != nil && o.Argument.IsSet() {
		return true
	}

	return false
}

// SetArgument gets a reference to the given NullableEdgeFirewallCriterionPolymorphicArgumentRequest and assigns it to the Argument field.
func (o *EdgeFirewallCriterionFieldRequest) SetArgument(v EdgeFirewallCriterionPolymorphicArgumentRequest) {
	o.Argument.Set(&v)
}
// SetArgumentNil sets the value for Argument to be an explicit nil
func (o *EdgeFirewallCriterionFieldRequest) SetArgumentNil() {
	o.Argument.Set(nil)
}

// UnsetArgument ensures that no value is present for Argument, not even an explicit nil
func (o *EdgeFirewallCriterionFieldRequest) UnsetArgument() {
	o.Argument.Unset()
}

// GetVariable returns the Variable field value
func (o *EdgeFirewallCriterionFieldRequest) GetVariable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variable
}

// GetVariableOk returns a tuple with the Variable field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallCriterionFieldRequest) GetVariableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variable, true
}

// SetVariable sets field value
func (o *EdgeFirewallCriterionFieldRequest) SetVariable(v string) {
	o.Variable = v
}

// GetConditional returns the Conditional field value
func (o *EdgeFirewallCriterionFieldRequest) GetConditional() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Conditional
}

// GetConditionalOk returns a tuple with the Conditional field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallCriterionFieldRequest) GetConditionalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditional, true
}

// SetConditional sets field value
func (o *EdgeFirewallCriterionFieldRequest) SetConditional(v string) {
	o.Conditional = v
}

// GetOperator returns the Operator field value
func (o *EdgeFirewallCriterionFieldRequest) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *EdgeFirewallCriterionFieldRequest) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *EdgeFirewallCriterionFieldRequest) SetOperator(v string) {
	o.Operator = v
}

func (o EdgeFirewallCriterionFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFirewallCriterionFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Argument.IsSet() {
		toSerialize["argument"] = o.Argument.Get()
	}
	toSerialize["variable"] = o.Variable
	toSerialize["conditional"] = o.Conditional
	toSerialize["operator"] = o.Operator
	return toSerialize, nil
}

func (o *EdgeFirewallCriterionFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"variable",
		"conditional",
		"operator",
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

	varEdgeFirewallCriterionFieldRequest := _EdgeFirewallCriterionFieldRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEdgeFirewallCriterionFieldRequest)

	if err != nil {
		return err
	}

	*o = EdgeFirewallCriterionFieldRequest(varEdgeFirewallCriterionFieldRequest)

	return err
}

type NullableEdgeFirewallCriterionFieldRequest struct {
	value *EdgeFirewallCriterionFieldRequest
	isSet bool
}

func (v NullableEdgeFirewallCriterionFieldRequest) Get() *EdgeFirewallCriterionFieldRequest {
	return v.value
}

func (v *NullableEdgeFirewallCriterionFieldRequest) Set(val *EdgeFirewallCriterionFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallCriterionFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallCriterionFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallCriterionFieldRequest(val *EdgeFirewallCriterionFieldRequest) *NullableEdgeFirewallCriterionFieldRequest {
	return &NullableEdgeFirewallCriterionFieldRequest{value: val, isSet: true}
}

func (v NullableEdgeFirewallCriterionFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallCriterionFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


