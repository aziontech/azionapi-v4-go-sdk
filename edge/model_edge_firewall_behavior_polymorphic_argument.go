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

// EdgeFirewallBehaviorPolymorphicArgument - struct for EdgeFirewallBehaviorPolymorphicArgument
type EdgeFirewallBehaviorPolymorphicArgument struct {
	SetCustomResponseArgument *SetCustomResponseArgument
	SetRateLimitArgument *SetRateLimitArgument
	SetWafRuleSetArgument *SetWafRuleSetArgument
	Int32 *int32
	String *string
}

// SetCustomResponseArgumentAsEdgeFirewallBehaviorPolymorphicArgument is a convenience function that returns SetCustomResponseArgument wrapped in EdgeFirewallBehaviorPolymorphicArgument
func SetCustomResponseArgumentAsEdgeFirewallBehaviorPolymorphicArgument(v *SetCustomResponseArgument) EdgeFirewallBehaviorPolymorphicArgument {
	return EdgeFirewallBehaviorPolymorphicArgument{
		SetCustomResponseArgument: v,
	}
}

// SetRateLimitArgumentAsEdgeFirewallBehaviorPolymorphicArgument is a convenience function that returns SetRateLimitArgument wrapped in EdgeFirewallBehaviorPolymorphicArgument
func SetRateLimitArgumentAsEdgeFirewallBehaviorPolymorphicArgument(v *SetRateLimitArgument) EdgeFirewallBehaviorPolymorphicArgument {
	return EdgeFirewallBehaviorPolymorphicArgument{
		SetRateLimitArgument: v,
	}
}

// SetWafRuleSetArgumentAsEdgeFirewallBehaviorPolymorphicArgument is a convenience function that returns SetWafRuleSetArgument wrapped in EdgeFirewallBehaviorPolymorphicArgument
func SetWafRuleSetArgumentAsEdgeFirewallBehaviorPolymorphicArgument(v *SetWafRuleSetArgument) EdgeFirewallBehaviorPolymorphicArgument {
	return EdgeFirewallBehaviorPolymorphicArgument{
		SetWafRuleSetArgument: v,
	}
}

// int32AsEdgeFirewallBehaviorPolymorphicArgument is a convenience function that returns int32 wrapped in EdgeFirewallBehaviorPolymorphicArgument
func Int32AsEdgeFirewallBehaviorPolymorphicArgument(v *int32) EdgeFirewallBehaviorPolymorphicArgument {
	return EdgeFirewallBehaviorPolymorphicArgument{
		Int32: v,
	}
}

// stringAsEdgeFirewallBehaviorPolymorphicArgument is a convenience function that returns string wrapped in EdgeFirewallBehaviorPolymorphicArgument
func StringAsEdgeFirewallBehaviorPolymorphicArgument(v *string) EdgeFirewallBehaviorPolymorphicArgument {
	return EdgeFirewallBehaviorPolymorphicArgument{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EdgeFirewallBehaviorPolymorphicArgument) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into SetCustomResponseArgument
	err = newStrictDecoder(data).Decode(&dst.SetCustomResponseArgument)
	if err == nil {
		jsonSetCustomResponseArgument, _ := json.Marshal(dst.SetCustomResponseArgument)
		if string(jsonSetCustomResponseArgument) == "{}" { // empty struct
			dst.SetCustomResponseArgument = nil
		} else {
			if err = validator.Validate(dst.SetCustomResponseArgument); err != nil {
				dst.SetCustomResponseArgument = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetCustomResponseArgument = nil
	}

	// try to unmarshal data into SetRateLimitArgument
	err = newStrictDecoder(data).Decode(&dst.SetRateLimitArgument)
	if err == nil {
		jsonSetRateLimitArgument, _ := json.Marshal(dst.SetRateLimitArgument)
		if string(jsonSetRateLimitArgument) == "{}" { // empty struct
			dst.SetRateLimitArgument = nil
		} else {
			if err = validator.Validate(dst.SetRateLimitArgument); err != nil {
				dst.SetRateLimitArgument = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetRateLimitArgument = nil
	}

	// try to unmarshal data into SetWafRuleSetArgument
	err = newStrictDecoder(data).Decode(&dst.SetWafRuleSetArgument)
	if err == nil {
		jsonSetWafRuleSetArgument, _ := json.Marshal(dst.SetWafRuleSetArgument)
		if string(jsonSetWafRuleSetArgument) == "{}" { // empty struct
			dst.SetWafRuleSetArgument = nil
		} else {
			if err = validator.Validate(dst.SetWafRuleSetArgument); err != nil {
				dst.SetWafRuleSetArgument = nil
			} else {
				match++
			}
		}
	} else {
		dst.SetWafRuleSetArgument = nil
	}

	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
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
		dst.SetCustomResponseArgument = nil
		dst.SetRateLimitArgument = nil
		dst.SetWafRuleSetArgument = nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EdgeFirewallBehaviorPolymorphicArgument)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EdgeFirewallBehaviorPolymorphicArgument)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EdgeFirewallBehaviorPolymorphicArgument) MarshalJSON() ([]byte, error) {
	if src.SetCustomResponseArgument != nil {
		return json.Marshal(&src.SetCustomResponseArgument)
	}

	if src.SetRateLimitArgument != nil {
		return json.Marshal(&src.SetRateLimitArgument)
	}

	if src.SetWafRuleSetArgument != nil {
		return json.Marshal(&src.SetWafRuleSetArgument)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EdgeFirewallBehaviorPolymorphicArgument) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SetCustomResponseArgument != nil {
		return obj.SetCustomResponseArgument
	}

	if obj.SetRateLimitArgument != nil {
		return obj.SetRateLimitArgument
	}

	if obj.SetWafRuleSetArgument != nil {
		return obj.SetWafRuleSetArgument
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableEdgeFirewallBehaviorPolymorphicArgument struct {
	value *EdgeFirewallBehaviorPolymorphicArgument
	isSet bool
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgument) Get() *EdgeFirewallBehaviorPolymorphicArgument {
	return v.value
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgument) Set(val *EdgeFirewallBehaviorPolymorphicArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFirewallBehaviorPolymorphicArgument(val *EdgeFirewallBehaviorPolymorphicArgument) *NullableEdgeFirewallBehaviorPolymorphicArgument {
	return &NullableEdgeFirewallBehaviorPolymorphicArgument{value: val, isSet: true}
}

func (v NullableEdgeFirewallBehaviorPolymorphicArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFirewallBehaviorPolymorphicArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


