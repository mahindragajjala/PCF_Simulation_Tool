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

// PdnConnectivityStatus PDN Connectivity Status.
type PdnConnectivityStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PdnConnectivityStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PdnConnectivityStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PdnConnectivityStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePdnConnectivityStatus struct {
	value *PdnConnectivityStatus
	isSet bool
}

func (v NullablePdnConnectivityStatus) Get() *PdnConnectivityStatus {
	return v.value
}

func (v *NullablePdnConnectivityStatus) Set(val *PdnConnectivityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnConnectivityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnConnectivityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnConnectivityStatus(val *PdnConnectivityStatus) *NullablePdnConnectivityStatus {
	return &NullablePdnConnectivityStatus{value: val, isSet: true}
}

func (v NullablePdnConnectivityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnConnectivityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


