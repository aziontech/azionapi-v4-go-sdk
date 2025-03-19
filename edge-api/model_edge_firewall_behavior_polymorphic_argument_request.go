/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// EdgeFirewallBehaviorPolymorphicArgumentRequest - struct for EdgeFirewallBehaviorPolymorphicArgumentRequest
type EdgeFirewallBehaviorPolymorphicArgumentRequest struct {
	SetCustomResponseArgumentRequest *SetCustomResponseArgumentRequest
	SetRateLimitArgumentRequest *SetRateLimitArgumentRequest
	SetWafRuleSetArgumentRequest *SetWafRuleSetArgumentRequest
	Int64 *int64
	String *string
}

// SetCustomResponseArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest is a convenience function that returns SetCustomResponseArgumentRequest wrapped in EdgeFirewallBehaviorPolymorphicArgumentRequest
func SetCustomResponseArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest(v *SetCustomResponseArgumentRequest) EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return EdgeFirewallBehaviorPolymorphicArgumentRequest{
		SetCustomResponseArgumentRequest: v,
	}
}

// SetRateLimitArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest is a convenience function that returns SetRateLimitArgumentRequest wrapped in EdgeFirewallBehaviorPolymorphicArgumentRequest
func SetRateLimitArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest(v *SetRateLimitArgumentRequest) EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return EdgeFirewallBehaviorPolymorphicArgumentRequest{
		SetRateLimitArgumentRequest: v,
	}
}

// SetWafRuleSetArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest is a convenience function that returns SetWafRuleSetArgumentRequest wrapped in EdgeFirewallBehaviorPolymorphicArgumentRequest
func SetWafRuleSetArgumentRequestAsEdgeFirewallBehaviorPolymorphicArgumentRequest(v *SetWafRuleSetArgumentRequest) EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return EdgeFirewallBehaviorPolymorphicArgumentRequest{
		SetWafRuleSetArgumentRequest: v,
	}
}

// int64AsEdgeFirewallBehaviorPolymorphicArgumentRequest is a convenience function that returns int64 wrapped in EdgeFirewallBehaviorPolymorphicArgumentRequest
func Int64AsEdgeFirewallBehaviorPolymorphicArgumentRequest(v *int64) EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return EdgeFirewallBehaviorPolymorphicArgumentRequest{
		Int64: v,
	}
}

// stringAsEdgeFirewallBehaviorPolymorphicArgumentRequest is a convenience function that returns string wrapped in EdgeFirewallBehaviorPolymorphicArgumentRequest
func StringAsEdgeFirewallBehaviorPolymorphicArgumentRequest(v *string) EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return EdgeFirewallBehaviorPolymorphicArgumentRequest{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EdgeFirewallBehaviorPolymorphicArgumentRequest) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into SetCustomResponseArgumentRequest
	err = newStrictDecoder(data).Decode(&dst.SetCustomResponseArgumentRequest)
	if err == nil {
		jsonSetCustomResponseArgumentRequest, _ := json.Marshal(dst.SetCustomResponseArgumentRequest)
		if string(jsonSetCustomResponseArgumentRequest) == "{}" { // empty struct
			dst.SetCustomResponseArgumentRequest = nil
		} else {
			if err = validator.Validate(dst.SetCustomResponseArgumentRequest); err != nil {
				dst.SetCustomResponseArgumentRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetCustomResponseArgumentRequest = nil
	}

	// try to unmarshal data into SetRateLimitArgumentRequest
	err = newStrictDecoder(data).Decode(&dst.SetRateLimitArgumentRequest)
	if err == nil {
		jsonSetRateLimitArgumentRequest, _ := json.Marshal(dst.SetRateLimitArgumentRequest)
		if string(jsonSetRateLimitArgumentRequest) == "{}" { // empty struct
			dst.SetRateLimitArgumentRequest = nil
		} else {
			if err = validator.Validate(dst.SetRateLimitArgumentRequest); err != nil {
				dst.SetRateLimitArgumentRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetRateLimitArgumentRequest = nil
	}

	// try to unmarshal data into SetWafRuleSetArgumentRequest
	err = newStrictDecoder(data).Decode(&dst.SetWafRuleSetArgumentRequest)
	if err == nil {
		jsonSetWafRuleSetArgumentRequest, _ := json.Marshal(dst.SetWafRuleSetArgumentRequest)
		if string(jsonSetWafRuleSetArgumentRequest) == "{}" { // empty struct
			dst.SetWafRuleSetArgumentRequest = nil
		} else {
			if err = validator.Validate(dst.SetWafRuleSetArgumentRequest); err != nil {
				dst.SetWafRuleSetArgumentRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetWafRuleSetArgumentRequest = nil
	}

	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			if err = validator.Validate(dst.Int64); err != nil {
				dst.Int64 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SetCustomResponseArgumentRequest = nil
		dst.SetRateLimitArgumentRequest = nil
		dst.SetWafRuleSetArgumentRequest = nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EdgeFirewallBehaviorPolymorphicArgumentRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EdgeFirewallBehaviorPolymorphicArgumentRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EdgeFirewallBehaviorPolymorphicArgumentRequest) MarshalJSON() ([]byte, error) {
	if src.SetCustomResponseArgumentRequest != nil {
		return json.Marshal(&src.SetCustomResponseArgumentRequest)
	}

	if src.SetRateLimitArgumentRequest != nil {
		return json.Marshal(&src.SetRateLimitArgumentRequest)
	}

	if src.SetWafRuleSetArgumentRequest != nil {
		return json.Marshal(&src.SetWafRuleSetArgumentRequest)
	}

	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EdgeFirewallBehaviorPolymorphicArgumentRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SetCustomResponseArgumentRequest != nil {
		return obj.SetCustomResponseArgumentRequest
	}

	if obj.SetRateLimitArgumentRequest != nil {
		return obj.SetRateLimitArgumentRequest
	}

	if obj.SetWafRuleSetArgumentRequest != nil {
		return obj.SetWafRuleSetArgumentRequest
	}

	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableEdgeFirewallBehaviorPolymorphicArgumentRequest struct {
	value *EdgeFirewallBehaviorPolymorphicArgumentRequest
	isSet bool
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) Get() *EdgeFirewallBehaviorPolymorphicArgumentRequest {
	return v.value
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) Set(val *EdgeFirewallBehaviorPolymorphicArgumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallBehaviorPolymorphicArgumentRequest(val *EdgeFirewallBehaviorPolymorphicArgumentRequest) *NullableEdgeFirewallBehaviorPolymorphicArgumentRequest {
	return &NullableEdgeFirewallBehaviorPolymorphicArgumentRequest{value: val, isSet: true}
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


