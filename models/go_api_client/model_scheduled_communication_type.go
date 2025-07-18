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

// ScheduledCommunicationType Possible values are: -DOWNLINK_ONLY: Downlink only -UPLINK_ONLY: Uplink only -BIDIRECTIONA: Bi-directional 
type ScheduledCommunicationType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScheduledCommunicationType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ScheduledCommunicationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScheduledCommunicationType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScheduledCommunicationType struct {
	value *ScheduledCommunicationType
	isSet bool
}

func (v NullableScheduledCommunicationType) Get() *ScheduledCommunicationType {
	return v.value
}

func (v *NullableScheduledCommunicationType) Set(val *ScheduledCommunicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledCommunicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledCommunicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledCommunicationType(val *ScheduledCommunicationType) *NullableScheduledCommunicationType {
	return &NullableScheduledCommunicationType{value: val, isSet: true}
}

func (v NullableScheduledCommunicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledCommunicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


