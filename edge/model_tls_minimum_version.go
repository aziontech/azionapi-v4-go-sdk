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

// TLSMinimumVersion - struct for TLSMinimumVersion
type TLSMinimumVersion struct {
	String *string
}

// stringAsTLSMinimumVersion is a convenience function that returns string wrapped in TLSMinimumVersion
func StringAsTLSMinimumVersion(v *string) TLSMinimumVersion {
	return TLSMinimumVersion{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TLSMinimumVersion) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
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
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TLSMinimumVersion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TLSMinimumVersion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TLSMinimumVersion) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TLSMinimumVersion) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj TLSMinimumVersion) GetActualInstanceValue() (interface{}) {
	if obj.String != nil {
		return *obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTLSMinimumVersion struct {
	value *TLSMinimumVersion
	isSet bool
}

func (v NullableTLSMinimumVersion) Get() *TLSMinimumVersion {
	return v.value
}

func (v *NullableTLSMinimumVersion) Set(val *TLSMinimumVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSMinimumVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSMinimumVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSMinimumVersion(val *TLSMinimumVersion) *NullableTLSMinimumVersion {
	return &NullableTLSMinimumVersion{value: val, isSet: true}
}

func (v NullableTLSMinimumVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSMinimumVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


