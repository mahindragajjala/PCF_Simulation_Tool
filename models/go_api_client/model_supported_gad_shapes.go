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

// SupportedGADShapes Indicates supported GAD shapes.
type SupportedGADShapes struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SupportedGADShapes) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(SupportedGADShapes)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SupportedGADShapes) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSupportedGADShapes struct {
	value *SupportedGADShapes
	isSet bool
}

func (v NullableSupportedGADShapes) Get() *SupportedGADShapes {
	return v.value
}

func (v *NullableSupportedGADShapes) Set(val *SupportedGADShapes) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedGADShapes) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedGADShapes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedGADShapes(val *SupportedGADShapes) *NullableSupportedGADShapes {
	return &NullableSupportedGADShapes{value: val, isSet: true}
}

func (v NullableSupportedGADShapes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedGADShapes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


