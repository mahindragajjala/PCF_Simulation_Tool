/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ImmediateReport - struct for ImmediateReport
type ImmediateReport struct {
	SubscriptionDataSets *SubscriptionDataSets
	ArrayOfSharedData *[]SharedData
}

// SubscriptionDataSetsAsImmediateReport is a convenience function that returns SubscriptionDataSets wrapped in ImmediateReport
func SubscriptionDataSetsAsImmediateReport(v *SubscriptionDataSets) ImmediateReport {
	return ImmediateReport{
		SubscriptionDataSets: v,
	}
}

// []SharedDataAsImmediateReport is a convenience function that returns []SharedData wrapped in ImmediateReport
func ArrayOfSharedDataAsImmediateReport(v *[]SharedData) ImmediateReport {
	return ImmediateReport{
		ArrayOfSharedData: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ImmediateReport) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubscriptionDataSets
	err = newStrictDecoder(data).Decode(&dst.SubscriptionDataSets)
	if err == nil {
		jsonSubscriptionDataSets, _ := json.Marshal(dst.SubscriptionDataSets)
		if string(jsonSubscriptionDataSets) == "{}" { // empty struct
			dst.SubscriptionDataSets = nil
		} else {
			if err = validator.Validate(dst.SubscriptionDataSets); err != nil {
				dst.SubscriptionDataSets = nil
			} else {
				match++
			}
		}
	} else {
		dst.SubscriptionDataSets = nil
	}

	// try to unmarshal data into ArrayOfSharedData
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSharedData)
	if err == nil {
		jsonArrayOfSharedData, _ := json.Marshal(dst.ArrayOfSharedData)
		if string(jsonArrayOfSharedData) == "{}" { // empty struct
			dst.ArrayOfSharedData = nil
		} else {
			if err = validator.Validate(dst.ArrayOfSharedData); err != nil {
				dst.ArrayOfSharedData = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfSharedData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubscriptionDataSets = nil
		dst.ArrayOfSharedData = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ImmediateReport)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ImmediateReport)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ImmediateReport) MarshalJSON() ([]byte, error) {
	if src.SubscriptionDataSets != nil {
		return json.Marshal(&src.SubscriptionDataSets)
	}

	if src.ArrayOfSharedData != nil {
		return json.Marshal(&src.ArrayOfSharedData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ImmediateReport) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubscriptionDataSets != nil {
		return obj.SubscriptionDataSets
	}

	if obj.ArrayOfSharedData != nil {
		return obj.ArrayOfSharedData
	}

	// all schemas are nil
	return nil
}

type NullableImmediateReport struct {
	value *ImmediateReport
	isSet bool
}

func (v NullableImmediateReport) Get() *ImmediateReport {
	return v.value
}

func (v *NullableImmediateReport) Set(val *ImmediateReport) {
	v.value = val
	v.isSet = true
}

func (v NullableImmediateReport) IsSet() bool {
	return v.isSet
}

func (v *NullableImmediateReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImmediateReport(val *ImmediateReport) *NullableImmediateReport {
	return &NullableImmediateReport{value: val, isSet: true}
}

func (v NullableImmediateReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImmediateReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


