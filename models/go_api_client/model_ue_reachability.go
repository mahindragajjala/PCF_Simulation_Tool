/*
Nudm_EE

Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UeReachability Describes the reachability of the UE
type UeReachability struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeReachability) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeReachability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeReachability) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeReachability struct {
	value *UeReachability
	isSet bool
}

func (v NullableUeReachability) Get() *UeReachability {
	return v.value
}

func (v *NullableUeReachability) Set(val *UeReachability) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachability) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachability(val *UeReachability) *NullableUeReachability {
	return &NullableUeReachability{value: val, isSet: true}
}

func (v NullableUeReachability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


