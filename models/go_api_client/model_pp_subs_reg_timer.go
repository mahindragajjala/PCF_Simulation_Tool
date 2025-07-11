/*
Nudm_PP

Nudm Parameter Provision Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PpSubsRegTimer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PpSubsRegTimer{}

// PpSubsRegTimer struct for PpSubsRegTimer
type PpSubsRegTimer struct {
	// indicating a time in seconds.
	SubsRegTimer int32 `json:"subsRegTimer"`
	AfInstanceId string `json:"afInstanceId"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	ReferenceId int32 `json:"referenceId"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

type _PpSubsRegTimer PpSubsRegTimer

// NewPpSubsRegTimer instantiates a new PpSubsRegTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpSubsRegTimer(subsRegTimer int32, afInstanceId string, referenceId int32) *PpSubsRegTimer {
	this := PpSubsRegTimer{}
	this.SubsRegTimer = subsRegTimer
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewPpSubsRegTimerWithDefaults instantiates a new PpSubsRegTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpSubsRegTimerWithDefaults() *PpSubsRegTimer {
	this := PpSubsRegTimer{}
	return &this
}

// GetSubsRegTimer returns the SubsRegTimer field value
func (o *PpSubsRegTimer) GetSubsRegTimer() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubsRegTimer
}

// GetSubsRegTimerOk returns a tuple with the SubsRegTimer field value
// and a boolean to check if the value has been set.
func (o *PpSubsRegTimer) GetSubsRegTimerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubsRegTimer, true
}

// SetSubsRegTimer sets field value
func (o *PpSubsRegTimer) SetSubsRegTimer(v int32) {
	o.SubsRegTimer = v
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *PpSubsRegTimer) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *PpSubsRegTimer) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *PpSubsRegTimer) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *PpSubsRegTimer) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *PpSubsRegTimer) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *PpSubsRegTimer) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *PpSubsRegTimer) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpSubsRegTimer) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *PpSubsRegTimer) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *PpSubsRegTimer) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *PpSubsRegTimer) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpSubsRegTimer) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *PpSubsRegTimer) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *PpSubsRegTimer) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o PpSubsRegTimer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PpSubsRegTimer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subsRegTimer"] = o.SubsRegTimer
	toSerialize["afInstanceId"] = o.AfInstanceId
	toSerialize["referenceId"] = o.ReferenceId
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return toSerialize, nil
}

func (o *PpSubsRegTimer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subsRegTimer",
		"afInstanceId",
		"referenceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPpSubsRegTimer := _PpSubsRegTimer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPpSubsRegTimer)

	if err != nil {
		return err
	}

	*o = PpSubsRegTimer(varPpSubsRegTimer)

	return err
}

type NullablePpSubsRegTimer struct {
	value *PpSubsRegTimer
	isSet bool
}

func (v NullablePpSubsRegTimer) Get() *PpSubsRegTimer {
	return v.value
}

func (v *NullablePpSubsRegTimer) Set(val *PpSubsRegTimer) {
	v.value = val
	v.isSet = true
}

func (v NullablePpSubsRegTimer) IsSet() bool {
	return v.isSet
}

func (v *NullablePpSubsRegTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpSubsRegTimer(val *PpSubsRegTimer) *NullablePpSubsRegTimer {
	return &NullablePpSubsRegTimer{value: val, isSet: true}
}

func (v NullablePpSubsRegTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpSubsRegTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


