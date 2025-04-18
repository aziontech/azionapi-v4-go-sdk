/*
Accounts API

REST API OpenAPI documentation for the Accounts API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounts-api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// AccountRequest - struct for AccountRequest
type AccountRequest struct {
	BrandRequest *BrandRequest
	OrganizationRequest *OrganizationRequest
	ResellerRequest *ResellerRequest
	WorkspaceRequest *WorkspaceRequest
}

// BrandRequestAsAccountRequest is a convenience function that returns BrandRequest wrapped in AccountRequest
func BrandRequestAsAccountRequest(v *BrandRequest) AccountRequest {
	return AccountRequest{
		BrandRequest: v,
	}
}

// OrganizationRequestAsAccountRequest is a convenience function that returns OrganizationRequest wrapped in AccountRequest
func OrganizationRequestAsAccountRequest(v *OrganizationRequest) AccountRequest {
	return AccountRequest{
		OrganizationRequest: v,
	}
}

// ResellerRequestAsAccountRequest is a convenience function that returns ResellerRequest wrapped in AccountRequest
func ResellerRequestAsAccountRequest(v *ResellerRequest) AccountRequest {
	return AccountRequest{
		ResellerRequest: v,
	}
}

// WorkspaceRequestAsAccountRequest is a convenience function that returns WorkspaceRequest wrapped in AccountRequest
func WorkspaceRequestAsAccountRequest(v *WorkspaceRequest) AccountRequest {
	return AccountRequest{
		WorkspaceRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AccountRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BrandRequest
	err = newStrictDecoder(data).Decode(&dst.BrandRequest)
	if err == nil {
		jsonBrandRequest, _ := json.Marshal(dst.BrandRequest)
		if string(jsonBrandRequest) == "{}" { // empty struct
			dst.BrandRequest = nil
		} else {
			if err = validator.Validate(dst.BrandRequest); err != nil {
				dst.BrandRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.BrandRequest = nil
	}

	// try to unmarshal data into OrganizationRequest
	err = newStrictDecoder(data).Decode(&dst.OrganizationRequest)
	if err == nil {
		jsonOrganizationRequest, _ := json.Marshal(dst.OrganizationRequest)
		if string(jsonOrganizationRequest) == "{}" { // empty struct
			dst.OrganizationRequest = nil
		} else {
			if err = validator.Validate(dst.OrganizationRequest); err != nil {
				dst.OrganizationRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.OrganizationRequest = nil
	}

	// try to unmarshal data into ResellerRequest
	err = newStrictDecoder(data).Decode(&dst.ResellerRequest)
	if err == nil {
		jsonResellerRequest, _ := json.Marshal(dst.ResellerRequest)
		if string(jsonResellerRequest) == "{}" { // empty struct
			dst.ResellerRequest = nil
		} else {
			if err = validator.Validate(dst.ResellerRequest); err != nil {
				dst.ResellerRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ResellerRequest = nil
	}

	// try to unmarshal data into WorkspaceRequest
	err = newStrictDecoder(data).Decode(&dst.WorkspaceRequest)
	if err == nil {
		jsonWorkspaceRequest, _ := json.Marshal(dst.WorkspaceRequest)
		if string(jsonWorkspaceRequest) == "{}" { // empty struct
			dst.WorkspaceRequest = nil
		} else {
			if err = validator.Validate(dst.WorkspaceRequest); err != nil {
				dst.WorkspaceRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.WorkspaceRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BrandRequest = nil
		dst.OrganizationRequest = nil
		dst.ResellerRequest = nil
		dst.WorkspaceRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AccountRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AccountRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AccountRequest) MarshalJSON() ([]byte, error) {
	if src.BrandRequest != nil {
		return json.Marshal(&src.BrandRequest)
	}

	if src.OrganizationRequest != nil {
		return json.Marshal(&src.OrganizationRequest)
	}

	if src.ResellerRequest != nil {
		return json.Marshal(&src.ResellerRequest)
	}

	if src.WorkspaceRequest != nil {
		return json.Marshal(&src.WorkspaceRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AccountRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BrandRequest != nil {
		return obj.BrandRequest
	}

	if obj.OrganizationRequest != nil {
		return obj.OrganizationRequest
	}

	if obj.ResellerRequest != nil {
		return obj.ResellerRequest
	}

	if obj.WorkspaceRequest != nil {
		return obj.WorkspaceRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj AccountRequest) GetActualInstanceValue() (interface{}) {
	if obj.BrandRequest != nil {
		return *obj.BrandRequest
	}

	if obj.OrganizationRequest != nil {
		return *obj.OrganizationRequest
	}

	if obj.ResellerRequest != nil {
		return *obj.ResellerRequest
	}

	if obj.WorkspaceRequest != nil {
		return *obj.WorkspaceRequest
	}

	// all schemas are nil
	return nil
}

type NullableAccountRequest struct {
	value *AccountRequest
	isSet bool
}

func (v NullableAccountRequest) Get() *AccountRequest {
	return v.value
}

func (v *NullableAccountRequest) Set(val *AccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRequest(val *AccountRequest) *NullableAccountRequest {
	return &NullableAccountRequest{value: val, isSet: true}
}

func (v NullableAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


