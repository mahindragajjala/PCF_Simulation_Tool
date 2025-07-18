/*
NudmUEAU

UDM UE Authentication Service. � 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AccessNetworkId struct for AccessNetworkId
type AccessNetworkId struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessNetworkId) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessNetworkId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessNetworkId) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessNetworkId struct {
	value *AccessNetworkId
	isSet bool
}

func (v NullableAccessNetworkId) Get() *AccessNetworkId {
	return v.value
}

func (v *NullableAccessNetworkId) Set(val *AccessNetworkId) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessNetworkId) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessNetworkId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessNetworkId(val *AccessNetworkId) *NullableAccessNetworkId {
	return &NullableAccessNetworkId{value: val, isSet: true}
}

func (v NullableAccessNetworkId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessNetworkId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


