/*
Nudm_PP

Nudm Parameter Provision Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AreaUsageInd Indicates one of the mutually exclusive global settings (whether the UE is allowed to generate and send the reports inside or outside the event report expected area). 
type AreaUsageInd struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AreaUsageInd) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AreaUsageInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AreaUsageInd) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAreaUsageInd struct {
	value *AreaUsageInd
	isSet bool
}

func (v NullableAreaUsageInd) Get() *AreaUsageInd {
	return v.value
}

func (v *NullableAreaUsageInd) Set(val *AreaUsageInd) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaUsageInd) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaUsageInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaUsageInd(val *AreaUsageInd) *NullableAreaUsageInd {
	return &NullableAreaUsageInd{value: val, isSet: true}
}

func (v NullableAreaUsageInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaUsageInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


