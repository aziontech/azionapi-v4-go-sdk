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

// TLSMinimumVersion - struct for TLSMinimumVersion
type TLSMinimumVersion struct {
	BlankEnum *BlankEnum
	MinimumVersionEnum *MinimumVersionEnum
}

// BlankEnumAsTLSMinimumVersion is a convenience function that returns BlankEnum wrapped in TLSMinimumVersion
func BlankEnumAsTLSMinimumVersion(v *BlankEnum) TLSMinimumVersion {
	return TLSMinimumVersion{
		BlankEnum: v,
	}
}

// MinimumVersionEnumAsTLSMinimumVersion is a convenience function that returns MinimumVersionEnum wrapped in TLSMinimumVersion
func MinimumVersionEnumAsTLSMinimumVersion(v *MinimumVersionEnum) TLSMinimumVersion {
	return TLSMinimumVersion{
		MinimumVersionEnum: v,
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
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			if err = validator.Validate(dst.BlankEnum); err != nil {
				dst.BlankEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into MinimumVersionEnum
	err = newStrictDecoder(data).Decode(&dst.MinimumVersionEnum)
	if err == nil {
		jsonMinimumVersionEnum, _ := json.Marshal(dst.MinimumVersionEnum)
		if string(jsonMinimumVersionEnum) == "{}" { // empty struct
			dst.MinimumVersionEnum = nil
		} else {
			if err = validator.Validate(dst.MinimumVersionEnum); err != nil {
				dst.MinimumVersionEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.MinimumVersionEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.MinimumVersionEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TLSMinimumVersion)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TLSMinimumVersion)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TLSMinimumVersion) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.MinimumVersionEnum != nil {
		return json.Marshal(&src.MinimumVersionEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TLSMinimumVersion) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.MinimumVersionEnum != nil {
		return obj.MinimumVersionEnum
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


