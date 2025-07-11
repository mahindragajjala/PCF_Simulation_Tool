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

// checks if the SpatialValidityCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpatialValidityCond{}

// SpatialValidityCond Contains the Spatial Validity Condition.
type SpatialValidityCond struct {
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`
	Countries []string `json:"countries,omitempty"`
	GeographicalServiceArea *GeoServiceArea `json:"geographicalServiceArea,omitempty"`
}

// NewSpatialValidityCond instantiates a new SpatialValidityCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpatialValidityCond() *SpatialValidityCond {
	this := SpatialValidityCond{}
	return &this
}

// NewSpatialValidityCondWithDefaults instantiates a new SpatialValidityCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpatialValidityCondWithDefaults() *SpatialValidityCond {
	this := SpatialValidityCond{}
	return &this
}

// GetTrackingAreaList returns the TrackingAreaList field value if set, zero value otherwise.
func (o *SpatialValidityCond) GetTrackingAreaList() []Tai {
	if o == nil || IsNil(o.TrackingAreaList) {
		var ret []Tai
		return ret
	}
	return o.TrackingAreaList
}

// GetTrackingAreaListOk returns a tuple with the TrackingAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond) GetTrackingAreaListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TrackingAreaList) {
		return nil, false
	}
	return o.TrackingAreaList, true
}

// HasTrackingAreaList returns a boolean if a field has been set.
func (o *SpatialValidityCond) HasTrackingAreaList() bool {
	if o != nil && !IsNil(o.TrackingAreaList) {
		return true
	}

	return false
}

// SetTrackingAreaList gets a reference to the given []Tai and assigns it to the TrackingAreaList field.
func (o *SpatialValidityCond) SetTrackingAreaList(v []Tai) {
	o.TrackingAreaList = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *SpatialValidityCond) GetCountries() []string {
	if o == nil || IsNil(o.Countries) {
		var ret []string
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond) GetCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *SpatialValidityCond) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []string and assigns it to the Countries field.
func (o *SpatialValidityCond) SetCountries(v []string) {
	o.Countries = v
}

// GetGeographicalServiceArea returns the GeographicalServiceArea field value if set, zero value otherwise.
func (o *SpatialValidityCond) GetGeographicalServiceArea() GeoServiceArea {
	if o == nil || IsNil(o.GeographicalServiceArea) {
		var ret GeoServiceArea
		return ret
	}
	return *o.GeographicalServiceArea
}

// GetGeographicalServiceAreaOk returns a tuple with the GeographicalServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpatialValidityCond) GetGeographicalServiceAreaOk() (*GeoServiceArea, bool) {
	if o == nil || IsNil(o.GeographicalServiceArea) {
		return nil, false
	}
	return o.GeographicalServiceArea, true
}

// HasGeographicalServiceArea returns a boolean if a field has been set.
func (o *SpatialValidityCond) HasGeographicalServiceArea() bool {
	if o != nil && !IsNil(o.GeographicalServiceArea) {
		return true
	}

	return false
}

// SetGeographicalServiceArea gets a reference to the given GeoServiceArea and assigns it to the GeographicalServiceArea field.
func (o *SpatialValidityCond) SetGeographicalServiceArea(v GeoServiceArea) {
	o.GeographicalServiceArea = &v
}

func (o SpatialValidityCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpatialValidityCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingAreaList) {
		toSerialize["trackingAreaList"] = o.TrackingAreaList
	}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !IsNil(o.GeographicalServiceArea) {
		toSerialize["geographicalServiceArea"] = o.GeographicalServiceArea
	}
	return toSerialize, nil
}

type NullableSpatialValidityCond struct {
	value *SpatialValidityCond
	isSet bool
}

func (v NullableSpatialValidityCond) Get() *SpatialValidityCond {
	return v.value
}

func (v *NullableSpatialValidityCond) Set(val *SpatialValidityCond) {
	v.value = val
	v.isSet = true
}

func (v NullableSpatialValidityCond) IsSet() bool {
	return v.isSet
}

func (v *NullableSpatialValidityCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpatialValidityCond(val *SpatialValidityCond) *NullableSpatialValidityCond {
	return &NullableSpatialValidityCond{value: val, isSet: true}
}

func (v NullableSpatialValidityCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpatialValidityCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


