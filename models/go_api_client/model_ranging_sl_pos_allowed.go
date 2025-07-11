/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// RangingSlPosAllowed Indicates the Ranging/SL positioning services that can be authorised to  use in the given PLMN for the UE. 
type RangingSlPosAllowed struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RangingSlPosAllowed) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RangingSlPosAllowed)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RangingSlPosAllowed) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRangingSlPosAllowed struct {
	value *RangingSlPosAllowed
	isSet bool
}

func (v NullableRangingSlPosAllowed) Get() *RangingSlPosAllowed {
	return v.value
}

func (v *NullableRangingSlPosAllowed) Set(val *RangingSlPosAllowed) {
	v.value = val
	v.isSet = true
}

func (v NullableRangingSlPosAllowed) IsSet() bool {
	return v.isSet
}

func (v *NullableRangingSlPosAllowed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangingSlPosAllowed(val *RangingSlPosAllowed) *NullableRangingSlPosAllowed {
	return &NullableRangingSlPosAllowed{value: val, isSet: true}
}

func (v NullableRangingSlPosAllowed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangingSlPosAllowed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


