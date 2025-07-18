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

// checks if the LocationArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationArea{}

// LocationArea struct for LocationArea
type LocationArea struct {
	// Identifies a list of geographic area of the user where the UE is located.
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	// Identifies a list of civic addresses of the user where the UE is located.
	CivicAddresses []CivicAddress `json:"civicAddresses,omitempty"`
	NwAreaInfo *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	UmtTime *UmtTime `json:"umtTime,omitempty"`
}

// NewLocationArea instantiates a new LocationArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationArea() *LocationArea {
	this := LocationArea{}
	return &this
}

// NewLocationAreaWithDefaults instantiates a new LocationArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationAreaWithDefaults() *LocationArea {
	this := LocationArea{}
	return &this
}

// GetGeographicAreas returns the GeographicAreas field value if set, zero value otherwise.
func (o *LocationArea) GetGeographicAreas() []GeographicArea {
	if o == nil || IsNil(o.GeographicAreas) {
		var ret []GeographicArea
		return ret
	}
	return o.GeographicAreas
}

// GetGeographicAreasOk returns a tuple with the GeographicAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetGeographicAreasOk() ([]GeographicArea, bool) {
	if o == nil || IsNil(o.GeographicAreas) {
		return nil, false
	}
	return o.GeographicAreas, true
}

// HasGeographicAreas returns a boolean if a field has been set.
func (o *LocationArea) HasGeographicAreas() bool {
	if o != nil && !IsNil(o.GeographicAreas) {
		return true
	}

	return false
}

// SetGeographicAreas gets a reference to the given []GeographicArea and assigns it to the GeographicAreas field.
func (o *LocationArea) SetGeographicAreas(v []GeographicArea) {
	o.GeographicAreas = v
}

// GetCivicAddresses returns the CivicAddresses field value if set, zero value otherwise.
func (o *LocationArea) GetCivicAddresses() []CivicAddress {
	if o == nil || IsNil(o.CivicAddresses) {
		var ret []CivicAddress
		return ret
	}
	return o.CivicAddresses
}

// GetCivicAddressesOk returns a tuple with the CivicAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetCivicAddressesOk() ([]CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddresses) {
		return nil, false
	}
	return o.CivicAddresses, true
}

// HasCivicAddresses returns a boolean if a field has been set.
func (o *LocationArea) HasCivicAddresses() bool {
	if o != nil && !IsNil(o.CivicAddresses) {
		return true
	}

	return false
}

// SetCivicAddresses gets a reference to the given []CivicAddress and assigns it to the CivicAddresses field.
func (o *LocationArea) SetCivicAddresses(v []CivicAddress) {
	o.CivicAddresses = v
}

// GetNwAreaInfo returns the NwAreaInfo field value if set, zero value otherwise.
func (o *LocationArea) GetNwAreaInfo() NetworkAreaInfo {
	if o == nil || IsNil(o.NwAreaInfo) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NwAreaInfo
}

// GetNwAreaInfoOk returns a tuple with the NwAreaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetNwAreaInfoOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.NwAreaInfo) {
		return nil, false
	}
	return o.NwAreaInfo, true
}

// HasNwAreaInfo returns a boolean if a field has been set.
func (o *LocationArea) HasNwAreaInfo() bool {
	if o != nil && !IsNil(o.NwAreaInfo) {
		return true
	}

	return false
}

// SetNwAreaInfo gets a reference to the given NetworkAreaInfo and assigns it to the NwAreaInfo field.
func (o *LocationArea) SetNwAreaInfo(v NetworkAreaInfo) {
	o.NwAreaInfo = &v
}

// GetUmtTime returns the UmtTime field value if set, zero value otherwise.
func (o *LocationArea) GetUmtTime() UmtTime {
	if o == nil || IsNil(o.UmtTime) {
		var ret UmtTime
		return ret
	}
	return *o.UmtTime
}

// GetUmtTimeOk returns a tuple with the UmtTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationArea) GetUmtTimeOk() (*UmtTime, bool) {
	if o == nil || IsNil(o.UmtTime) {
		return nil, false
	}
	return o.UmtTime, true
}

// HasUmtTime returns a boolean if a field has been set.
func (o *LocationArea) HasUmtTime() bool {
	if o != nil && !IsNil(o.UmtTime) {
		return true
	}

	return false
}

// SetUmtTime gets a reference to the given UmtTime and assigns it to the UmtTime field.
func (o *LocationArea) SetUmtTime(v UmtTime) {
	o.UmtTime = &v
}

func (o LocationArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeographicAreas) {
		toSerialize["geographicAreas"] = o.GeographicAreas
	}
	if !IsNil(o.CivicAddresses) {
		toSerialize["civicAddresses"] = o.CivicAddresses
	}
	if !IsNil(o.NwAreaInfo) {
		toSerialize["nwAreaInfo"] = o.NwAreaInfo
	}
	if !IsNil(o.UmtTime) {
		toSerialize["umtTime"] = o.UmtTime
	}
	return toSerialize, nil
}

type NullableLocationArea struct {
	value *LocationArea
	isSet bool
}

func (v NullableLocationArea) Get() *LocationArea {
	return v.value
}

func (v *NullableLocationArea) Set(val *LocationArea) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationArea) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationArea(val *LocationArea) *NullableLocationArea {
	return &NullableLocationArea{value: val, isSet: true}
}

func (v NullableLocationArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


