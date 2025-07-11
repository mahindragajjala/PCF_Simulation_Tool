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

// TraceDepth The enumeration TraceDepth defines how detailed information should be recorded in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.1-1 
type TraceDepth struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TraceDepth) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TraceDepth)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TraceDepth) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTraceDepth struct {
	value *TraceDepth
	isSet bool
}

func (v NullableTraceDepth) Get() *TraceDepth {
	return v.value
}

func (v *NullableTraceDepth) Set(val *TraceDepth) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDepth) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDepth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDepth(val *TraceDepth) *NullableTraceDepth {
	return &NullableTraceDepth{value: val, isSet: true}
}

func (v NullableTraceDepth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDepth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


