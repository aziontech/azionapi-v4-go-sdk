/*
edge-api

REST API OpenAPI documentation for the edge-api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edge

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// ResponseBadRequestEdgeApplicationRuleEngineBehaviors - struct for ResponseBadRequestEdgeApplicationRuleEngineBehaviors
type ResponseBadRequestEdgeApplicationRuleEngineBehaviors struct {
	ArrayOfString *[]string
	MapmapOfStringarrayOfString *map[string][]string
}

// []stringAsResponseBadRequestEdgeApplicationRuleEngineBehaviors is a convenience function that returns []string wrapped in ResponseBadRequestEdgeApplicationRuleEngineBehaviors
func ArrayOfStringAsResponseBadRequestEdgeApplicationRuleEngineBehaviors(v *[]string) ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	return ResponseBadRequestEdgeApplicationRuleEngineBehaviors{
		ArrayOfString: v,
	}
}

// map[string][]stringAsResponseBadRequestEdgeApplicationRuleEngineBehaviors is a convenience function that returns map[string][]string wrapped in ResponseBadRequestEdgeApplicationRuleEngineBehaviors
func MapmapOfStringarrayOfStringAsResponseBadRequestEdgeApplicationRuleEngineBehaviors(v *map[string][]string) ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	return ResponseBadRequestEdgeApplicationRuleEngineBehaviors{
		MapmapOfStringarrayOfString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ResponseBadRequestEdgeApplicationRuleEngineBehaviors) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into MapmapOfStringarrayOfString
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringarrayOfString)
	if err == nil {
		jsonMapmapOfStringarrayOfString, _ := json.Marshal(dst.MapmapOfStringarrayOfString)
		if string(jsonMapmapOfStringarrayOfString) == "{}" { // empty struct
			dst.MapmapOfStringarrayOfString = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringarrayOfString); err != nil {
				dst.MapmapOfStringarrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringarrayOfString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.MapmapOfStringarrayOfString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ResponseBadRequestEdgeApplicationRuleEngineBehaviors)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ResponseBadRequestEdgeApplicationRuleEngineBehaviors)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ResponseBadRequestEdgeApplicationRuleEngineBehaviors) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.MapmapOfStringarrayOfString != nil {
		return json.Marshal(&src.MapmapOfStringarrayOfString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ResponseBadRequestEdgeApplicationRuleEngineBehaviors) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.MapmapOfStringarrayOfString != nil {
		return obj.MapmapOfStringarrayOfString
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ResponseBadRequestEdgeApplicationRuleEngineBehaviors) GetActualInstanceValue() (interface{}) {
	if obj.ArrayOfString != nil {
		return *obj.ArrayOfString
	}

	if obj.MapmapOfStringarrayOfString != nil {
		return *obj.MapmapOfStringarrayOfString
	}

	// all schemas are nil
	return nil
}

type NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors struct {
	value *ResponseBadRequestEdgeApplicationRuleEngineBehaviors
	isSet bool
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) Get() *ResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	return v.value
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) Set(val *ResponseBadRequestEdgeApplicationRuleEngineBehaviors) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBadRequestEdgeApplicationRuleEngineBehaviors(val *ResponseBadRequestEdgeApplicationRuleEngineBehaviors) *NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors {
	return &NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors{value: val, isSet: true}
}

func (v NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBadRequestEdgeApplicationRuleEngineBehaviors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


