/*
Nudm_PP

Nudm Parameter Provision Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BatteryIndication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatteryIndication{}

// BatteryIndication Parameters \"replaceableInd\" and \"rechargeableInd\" are only included if the value of Parameter \"batteryInd\" is true. 
type BatteryIndication struct {
	// This IE shall indicate whether the UE is battery powered or not. true: the UE is battery powered; false or absent: the UE is not battery powered 
	BatteryInd *bool `json:"batteryInd,omitempty"`
	// This IE shall indicate whether the battery of the UE is replaceable or not. true: the battery of the UE is replaceable; false or absent: the battery of the UE is not replaceable. 
	ReplaceableInd *bool `json:"replaceableInd,omitempty"`
	// This IE shall indicate whether the battery of the UE is rechargeable or not. true: the battery of UE is rechargeable; false or absent: the battery of the UE is not rechargeable. 
	RechargeableInd *bool `json:"rechargeableInd,omitempty"`
}

// NewBatteryIndication instantiates a new BatteryIndication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatteryIndication() *BatteryIndication {
	this := BatteryIndication{}
	return &this
}

// NewBatteryIndicationWithDefaults instantiates a new BatteryIndication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatteryIndicationWithDefaults() *BatteryIndication {
	this := BatteryIndication{}
	return &this
}

// GetBatteryInd returns the BatteryInd field value if set, zero value otherwise.
func (o *BatteryIndication) GetBatteryInd() bool {
	if o == nil || IsNil(o.BatteryInd) {
		var ret bool
		return ret
	}
	return *o.BatteryInd
}

// GetBatteryIndOk returns a tuple with the BatteryInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatteryIndication) GetBatteryIndOk() (*bool, bool) {
	if o == nil || IsNil(o.BatteryInd) {
		return nil, false
	}
	return o.BatteryInd, true
}

// HasBatteryInd returns a boolean if a field has been set.
func (o *BatteryIndication) HasBatteryInd() bool {
	if o != nil && !IsNil(o.BatteryInd) {
		return true
	}

	return false
}

// SetBatteryInd gets a reference to the given bool and assigns it to the BatteryInd field.
func (o *BatteryIndication) SetBatteryInd(v bool) {
	o.BatteryInd = &v
}

// GetReplaceableInd returns the ReplaceableInd field value if set, zero value otherwise.
func (o *BatteryIndication) GetReplaceableInd() bool {
	if o == nil || IsNil(o.ReplaceableInd) {
		var ret bool
		return ret
	}
	return *o.ReplaceableInd
}

// GetReplaceableIndOk returns a tuple with the ReplaceableInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatteryIndication) GetReplaceableIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceableInd) {
		return nil, false
	}
	return o.ReplaceableInd, true
}

// HasReplaceableInd returns a boolean if a field has been set.
func (o *BatteryIndication) HasReplaceableInd() bool {
	if o != nil && !IsNil(o.ReplaceableInd) {
		return true
	}

	return false
}

// SetReplaceableInd gets a reference to the given bool and assigns it to the ReplaceableInd field.
func (o *BatteryIndication) SetReplaceableInd(v bool) {
	o.ReplaceableInd = &v
}

// GetRechargeableInd returns the RechargeableInd field value if set, zero value otherwise.
func (o *BatteryIndication) GetRechargeableInd() bool {
	if o == nil || IsNil(o.RechargeableInd) {
		var ret bool
		return ret
	}
	return *o.RechargeableInd
}

// GetRechargeableIndOk returns a tuple with the RechargeableInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatteryIndication) GetRechargeableIndOk() (*bool, bool) {
	if o == nil || IsNil(o.RechargeableInd) {
		return nil, false
	}
	return o.RechargeableInd, true
}

// HasRechargeableInd returns a boolean if a field has been set.
func (o *BatteryIndication) HasRechargeableInd() bool {
	if o != nil && !IsNil(o.RechargeableInd) {
		return true
	}

	return false
}

// SetRechargeableInd gets a reference to the given bool and assigns it to the RechargeableInd field.
func (o *BatteryIndication) SetRechargeableInd(v bool) {
	o.RechargeableInd = &v
}

func (o BatteryIndication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatteryIndication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatteryInd) {
		toSerialize["batteryInd"] = o.BatteryInd
	}
	if !IsNil(o.ReplaceableInd) {
		toSerialize["replaceableInd"] = o.ReplaceableInd
	}
	if !IsNil(o.RechargeableInd) {
		toSerialize["rechargeableInd"] = o.RechargeableInd
	}
	return toSerialize, nil
}

type NullableBatteryIndication struct {
	value *BatteryIndication
	isSet bool
}

func (v NullableBatteryIndication) Get() *BatteryIndication {
	return v.value
}

func (v *NullableBatteryIndication) Set(val *BatteryIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableBatteryIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableBatteryIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatteryIndication(val *BatteryIndication) *NullableBatteryIndication {
	return &NullableBatteryIndication{value: val, isSet: true}
}

func (v NullableBatteryIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatteryIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


