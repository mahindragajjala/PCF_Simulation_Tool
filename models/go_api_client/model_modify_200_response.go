/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// Modify200Response - struct for Modify200Response
type Modify200Response struct {
	PatchResult *PatchResult
	SdmSubscription *SdmSubscription
}

// PatchResultAsModify200Response is a convenience function that returns PatchResult wrapped in Modify200Response
func PatchResultAsModify200Response(v *PatchResult) Modify200Response {
	return Modify200Response{
		PatchResult: v,
	}
}

// SdmSubscriptionAsModify200Response is a convenience function that returns SdmSubscription wrapped in Modify200Response
func SdmSubscriptionAsModify200Response(v *SdmSubscription) Modify200Response {
	return Modify200Response{
		SdmSubscription: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Modify200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PatchResult
	err = newStrictDecoder(data).Decode(&dst.PatchResult)
	if err == nil {
		jsonPatchResult, _ := json.Marshal(dst.PatchResult)
		if string(jsonPatchResult) == "{}" { // empty struct
			dst.PatchResult = nil
		} else {
			if err = validator.Validate(dst.PatchResult); err != nil {
				dst.PatchResult = nil
			} else {
				match++
			}
		}
	} else {
		dst.PatchResult = nil
	}

	// try to unmarshal data into SdmSubscription
	err = newStrictDecoder(data).Decode(&dst.SdmSubscription)
	if err == nil {
		jsonSdmSubscription, _ := json.Marshal(dst.SdmSubscription)
		if string(jsonSdmSubscription) == "{}" { // empty struct
			dst.SdmSubscription = nil
		} else {
			if err = validator.Validate(dst.SdmSubscription); err != nil {
				dst.SdmSubscription = nil
			} else {
				match++
			}
		}
	} else {
		dst.SdmSubscription = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PatchResult = nil
		dst.SdmSubscription = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Modify200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Modify200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Modify200Response) MarshalJSON() ([]byte, error) {
	if src.PatchResult != nil {
		return json.Marshal(&src.PatchResult)
	}

	if src.SdmSubscription != nil {
		return json.Marshal(&src.SdmSubscription)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Modify200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PatchResult != nil {
		return obj.PatchResult
	}

	if obj.SdmSubscription != nil {
		return obj.SdmSubscription
	}

	// all schemas are nil
	return nil
}

type NullableModify200Response struct {
	value *Modify200Response
	isSet bool
}

func (v NullableModify200Response) Get() *Modify200Response {
	return v.value
}

func (v *NullableModify200Response) Set(val *Modify200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModify200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModify200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModify200Response(val *Modify200Response) *NullableModify200Response {
	return &NullableModify200Response{value: val, isSet: true}
}

func (v NullableModify200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModify200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


