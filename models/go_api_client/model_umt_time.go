/*
Nudm_PP

Nudm Parameter Provision Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UmtTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UmtTime{}

// UmtTime struct for UmtTime
type UmtTime struct {
	// String with format partial-time or full-time as defined in clause 5.6 of IETF RFC 3339. Examples, 20:15:00, 20:15:00-08:00 (for 8 hours behind UTC).  
	TimeOfDay string `json:"timeOfDay"`
	// integer between and including 1 and 7 denoting a weekday. 1 shall indicate Monday, and the subsequent weekdays shall be indicated with the next higher numbers. 7 shall indicate Sunday. 
	DayOfWeek int32 `json:"dayOfWeek"`
}

type _UmtTime UmtTime

// NewUmtTime instantiates a new UmtTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmtTime(timeOfDay string, dayOfWeek int32) *UmtTime {
	this := UmtTime{}
	this.TimeOfDay = timeOfDay
	this.DayOfWeek = dayOfWeek
	return &this
}

// NewUmtTimeWithDefaults instantiates a new UmtTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmtTimeWithDefaults() *UmtTime {
	this := UmtTime{}
	return &this
}

// GetTimeOfDay returns the TimeOfDay field value
func (o *UmtTime) GetTimeOfDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeOfDay
}

// GetTimeOfDayOk returns a tuple with the TimeOfDay field value
// and a boolean to check if the value has been set.
func (o *UmtTime) GetTimeOfDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOfDay, true
}

// SetTimeOfDay sets field value
func (o *UmtTime) SetTimeOfDay(v string) {
	o.TimeOfDay = v
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *UmtTime) GetDayOfWeek() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *UmtTime) GetDayOfWeekOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *UmtTime) SetDayOfWeek(v int32) {
	o.DayOfWeek = v
}

func (o UmtTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmtTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timeOfDay"] = o.TimeOfDay
	toSerialize["dayOfWeek"] = o.DayOfWeek
	return toSerialize, nil
}

func (o *UmtTime) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timeOfDay",
		"dayOfWeek",
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

	varUmtTime := _UmtTime{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUmtTime)

	if err != nil {
		return err
	}

	*o = UmtTime(varUmtTime)

	return err
}

type NullableUmtTime struct {
	value *UmtTime
	isSet bool
}

func (v NullableUmtTime) Get() *UmtTime {
	return v.value
}

func (v *NullableUmtTime) Set(val *UmtTime) {
	v.value = val
	v.isSet = true
}

func (v NullableUmtTime) IsSet() bool {
	return v.isSet
}

func (v *NullableUmtTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmtTime(val *UmtTime) *NullableUmtTime {
	return &NullableUmtTime{value: val, isSet: true}
}

func (v NullableUmtTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmtTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


