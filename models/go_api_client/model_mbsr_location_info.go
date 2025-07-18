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

// checks if the MbsrLocationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsrLocationInfo{}

// MbsrLocationInfo struct for MbsrLocationInfo
type MbsrLocationInfo struct {
	MbsrLocation []Tai `json:"mbsrLocation,omitempty"`
	MbsrLocationAreas []string `json:"mbsrLocationAreas,omitempty"`
}

// NewMbsrLocationInfo instantiates a new MbsrLocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsrLocationInfo() *MbsrLocationInfo {
	this := MbsrLocationInfo{}
	return &this
}

// NewMbsrLocationInfoWithDefaults instantiates a new MbsrLocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsrLocationInfoWithDefaults() *MbsrLocationInfo {
	this := MbsrLocationInfo{}
	return &this
}

// GetMbsrLocation returns the MbsrLocation field value if set, zero value otherwise.
func (o *MbsrLocationInfo) GetMbsrLocation() []Tai {
	if o == nil || IsNil(o.MbsrLocation) {
		var ret []Tai
		return ret
	}
	return o.MbsrLocation
}

// GetMbsrLocationOk returns a tuple with the MbsrLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsrLocationInfo) GetMbsrLocationOk() ([]Tai, bool) {
	if o == nil || IsNil(o.MbsrLocation) {
		return nil, false
	}
	return o.MbsrLocation, true
}

// HasMbsrLocation returns a boolean if a field has been set.
func (o *MbsrLocationInfo) HasMbsrLocation() bool {
	if o != nil && !IsNil(o.MbsrLocation) {
		return true
	}

	return false
}

// SetMbsrLocation gets a reference to the given []Tai and assigns it to the MbsrLocation field.
func (o *MbsrLocationInfo) SetMbsrLocation(v []Tai) {
	o.MbsrLocation = v
}

// GetMbsrLocationAreas returns the MbsrLocationAreas field value if set, zero value otherwise.
func (o *MbsrLocationInfo) GetMbsrLocationAreas() []string {
	if o == nil || IsNil(o.MbsrLocationAreas) {
		var ret []string
		return ret
	}
	return o.MbsrLocationAreas
}

// GetMbsrLocationAreasOk returns a tuple with the MbsrLocationAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsrLocationInfo) GetMbsrLocationAreasOk() ([]string, bool) {
	if o == nil || IsNil(o.MbsrLocationAreas) {
		return nil, false
	}
	return o.MbsrLocationAreas, true
}

// HasMbsrLocationAreas returns a boolean if a field has been set.
func (o *MbsrLocationInfo) HasMbsrLocationAreas() bool {
	if o != nil && !IsNil(o.MbsrLocationAreas) {
		return true
	}

	return false
}

// SetMbsrLocationAreas gets a reference to the given []string and assigns it to the MbsrLocationAreas field.
func (o *MbsrLocationInfo) SetMbsrLocationAreas(v []string) {
	o.MbsrLocationAreas = v
}

func (o MbsrLocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsrLocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsrLocation) {
		toSerialize["mbsrLocation"] = o.MbsrLocation
	}
	if !IsNil(o.MbsrLocationAreas) {
		toSerialize["mbsrLocationAreas"] = o.MbsrLocationAreas
	}
	return toSerialize, nil
}

type NullableMbsrLocationInfo struct {
	value *MbsrLocationInfo
	isSet bool
}

func (v NullableMbsrLocationInfo) Get() *MbsrLocationInfo {
	return v.value
}

func (v *NullableMbsrLocationInfo) Set(val *MbsrLocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsrLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsrLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsrLocationInfo(val *MbsrLocationInfo) *NullableMbsrLocationInfo {
	return &NullableMbsrLocationInfo{value: val, isSet: true}
}

func (v NullableMbsrLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsrLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


