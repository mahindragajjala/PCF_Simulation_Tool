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

// SscMode represents the service and session continuity mode It shall comply with the provisions defined in table 5.4.3.6-1.  
type SscMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SscMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(SscMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SscMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSscMode struct {
	value *SscMode
	isSet bool
}

func (v NullableSscMode) Get() *SscMode {
	return v.value
}

func (v *NullableSscMode) Set(val *SscMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSscMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSscMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSscMode(val *SscMode) *NullableSscMode {
	return &NullableSscMode{value: val, isSet: true}
}

func (v NullableSscMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSscMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


