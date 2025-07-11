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

// checks if the GeographicalCoordinates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeographicalCoordinates{}

// GeographicalCoordinates Geographical coordinates.
type GeographicalCoordinates struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type _GeographicalCoordinates GeographicalCoordinates

// NewGeographicalCoordinates instantiates a new GeographicalCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeographicalCoordinates(lon float64, lat float64) *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	this.Lon = lon
	this.Lat = lat
	return &this
}

// NewGeographicalCoordinatesWithDefaults instantiates a new GeographicalCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeographicalCoordinatesWithDefaults() *GeographicalCoordinates {
	this := GeographicalCoordinates{}
	return &this
}

// GetLon returns the Lon field value
func (o *GeographicalCoordinates) GetLon() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Lon
}

// GetLonOk returns a tuple with the Lon field value
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLonOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lon, true
}

// SetLon sets field value
func (o *GeographicalCoordinates) SetLon(v float64) {
	o.Lon = v
}

// GetLat returns the Lat field value
func (o *GeographicalCoordinates) GetLat() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value
// and a boolean to check if the value has been set.
func (o *GeographicalCoordinates) GetLatOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lat, true
}

// SetLat sets field value
func (o *GeographicalCoordinates) SetLat(v float64) {
	o.Lat = v
}

func (o GeographicalCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeographicalCoordinates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lon"] = o.Lon
	toSerialize["lat"] = o.Lat
	return toSerialize, nil
}

func (o *GeographicalCoordinates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lon",
		"lat",
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

	varGeographicalCoordinates := _GeographicalCoordinates{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGeographicalCoordinates)

	if err != nil {
		return err
	}

	*o = GeographicalCoordinates(varGeographicalCoordinates)

	return err
}

type NullableGeographicalCoordinates struct {
	value *GeographicalCoordinates
	isSet bool
}

func (v NullableGeographicalCoordinates) Get() *GeographicalCoordinates {
	return v.value
}

func (v *NullableGeographicalCoordinates) Set(val *GeographicalCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicalCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicalCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicalCoordinates(val *GeographicalCoordinates) *NullableGeographicalCoordinates {
	return &NullableGeographicalCoordinates{value: val, isSet: true}
}

func (v NullableGeographicalCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicalCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


