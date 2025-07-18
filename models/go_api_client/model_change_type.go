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

// ChangeType Indicates the type of change to be performed.
type ChangeType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ChangeType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ChangeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ChangeType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableChangeType struct {
	value *ChangeType
	isSet bool
}

func (v NullableChangeType) Get() *ChangeType {
	return v.value
}

func (v *NullableChangeType) Set(val *ChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeType(val *ChangeType) *NullableChangeType {
	return &NullableChangeType{value: val, isSet: true}
}

func (v NullableChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


