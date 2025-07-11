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

// RestrictionType It contains the restriction type ALLOWED_AREAS or NOT_ALLOWED_AREAS.
type RestrictionType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RestrictionType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RestrictionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RestrictionType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRestrictionType struct {
	value *RestrictionType
	isSet bool
}

func (v NullableRestrictionType) Get() *RestrictionType {
	return v.value
}

func (v *NullableRestrictionType) Set(val *RestrictionType) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictionType) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictionType(val *RestrictionType) *NullableRestrictionType {
	return &NullableRestrictionType{value: val, isSet: true}
}

func (v NullableRestrictionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


