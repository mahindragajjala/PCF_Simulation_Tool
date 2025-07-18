/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MbsSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsSubscriptionData{}

// MbsSubscriptionData Contains the 5MBS Subscription Data.
type MbsSubscriptionData struct {
	MbsAllowed *bool `json:"mbsAllowed,omitempty"`
	MbsSessionIdList []MbsSessionId `json:"mbsSessionIdList,omitempty"`
	UeMbsAssistanceInfo []MbsSessionId `json:"ueMbsAssistanceInfo,omitempty"`
}

// NewMbsSubscriptionData instantiates a new MbsSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsSubscriptionData() *MbsSubscriptionData {
	this := MbsSubscriptionData{}
	var mbsAllowed bool = false
	this.MbsAllowed = &mbsAllowed
	return &this
}

// NewMbsSubscriptionDataWithDefaults instantiates a new MbsSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsSubscriptionDataWithDefaults() *MbsSubscriptionData {
	this := MbsSubscriptionData{}
	var mbsAllowed bool = false
	this.MbsAllowed = &mbsAllowed
	return &this
}

// GetMbsAllowed returns the MbsAllowed field value if set, zero value otherwise.
func (o *MbsSubscriptionData) GetMbsAllowed() bool {
	if o == nil || IsNil(o.MbsAllowed) {
		var ret bool
		return ret
	}
	return *o.MbsAllowed
}

// GetMbsAllowedOk returns a tuple with the MbsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSubscriptionData) GetMbsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.MbsAllowed) {
		return nil, false
	}
	return o.MbsAllowed, true
}

// HasMbsAllowed returns a boolean if a field has been set.
func (o *MbsSubscriptionData) HasMbsAllowed() bool {
	if o != nil && !IsNil(o.MbsAllowed) {
		return true
	}

	return false
}

// SetMbsAllowed gets a reference to the given bool and assigns it to the MbsAllowed field.
func (o *MbsSubscriptionData) SetMbsAllowed(v bool) {
	o.MbsAllowed = &v
}

// GetMbsSessionIdList returns the MbsSessionIdList field value if set, zero value otherwise.
func (o *MbsSubscriptionData) GetMbsSessionIdList() []MbsSessionId {
	if o == nil || IsNil(o.MbsSessionIdList) {
		var ret []MbsSessionId
		return ret
	}
	return o.MbsSessionIdList
}

// GetMbsSessionIdListOk returns a tuple with the MbsSessionIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSubscriptionData) GetMbsSessionIdListOk() ([]MbsSessionId, bool) {
	if o == nil || IsNil(o.MbsSessionIdList) {
		return nil, false
	}
	return o.MbsSessionIdList, true
}

// HasMbsSessionIdList returns a boolean if a field has been set.
func (o *MbsSubscriptionData) HasMbsSessionIdList() bool {
	if o != nil && !IsNil(o.MbsSessionIdList) {
		return true
	}

	return false
}

// SetMbsSessionIdList gets a reference to the given []MbsSessionId and assigns it to the MbsSessionIdList field.
func (o *MbsSubscriptionData) SetMbsSessionIdList(v []MbsSessionId) {
	o.MbsSessionIdList = v
}

// GetUeMbsAssistanceInfo returns the UeMbsAssistanceInfo field value if set, zero value otherwise.
func (o *MbsSubscriptionData) GetUeMbsAssistanceInfo() []MbsSessionId {
	if o == nil || IsNil(o.UeMbsAssistanceInfo) {
		var ret []MbsSessionId
		return ret
	}
	return o.UeMbsAssistanceInfo
}

// GetUeMbsAssistanceInfoOk returns a tuple with the UeMbsAssistanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsSubscriptionData) GetUeMbsAssistanceInfoOk() ([]MbsSessionId, bool) {
	if o == nil || IsNil(o.UeMbsAssistanceInfo) {
		return nil, false
	}
	return o.UeMbsAssistanceInfo, true
}

// HasUeMbsAssistanceInfo returns a boolean if a field has been set.
func (o *MbsSubscriptionData) HasUeMbsAssistanceInfo() bool {
	if o != nil && !IsNil(o.UeMbsAssistanceInfo) {
		return true
	}

	return false
}

// SetUeMbsAssistanceInfo gets a reference to the given []MbsSessionId and assigns it to the UeMbsAssistanceInfo field.
func (o *MbsSubscriptionData) SetUeMbsAssistanceInfo(v []MbsSessionId) {
	o.UeMbsAssistanceInfo = v
}

func (o MbsSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsAllowed) {
		toSerialize["mbsAllowed"] = o.MbsAllowed
	}
	if !IsNil(o.MbsSessionIdList) {
		toSerialize["mbsSessionIdList"] = o.MbsSessionIdList
	}
	if !IsNil(o.UeMbsAssistanceInfo) {
		toSerialize["ueMbsAssistanceInfo"] = o.UeMbsAssistanceInfo
	}
	return toSerialize, nil
}

type NullableMbsSubscriptionData struct {
	value *MbsSubscriptionData
	isSet bool
}

func (v NullableMbsSubscriptionData) Get() *MbsSubscriptionData {
	return v.value
}

func (v *NullableMbsSubscriptionData) Set(val *MbsSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSubscriptionData(val *MbsSubscriptionData) *NullableMbsSubscriptionData {
	return &NullableMbsSubscriptionData{value: val, isSet: true}
}

func (v NullableMbsSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


