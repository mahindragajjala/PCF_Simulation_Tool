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

// checks if the FrameRouteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameRouteInfo{}

// FrameRouteInfo struct for FrameRouteInfo
type FrameRouteInfo struct {
	// \"String identifying a IPv4 address mask formatted in the 'dotted decimal' notation as defined in RFC 1166.\" 
	Ipv4Mask *string `json:"ipv4Mask,omitempty" validate:"regexp=^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))$"`
	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// NewFrameRouteInfo instantiates a new FrameRouteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameRouteInfo() *FrameRouteInfo {
	this := FrameRouteInfo{}
	return &this
}

// NewFrameRouteInfoWithDefaults instantiates a new FrameRouteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameRouteInfoWithDefaults() *FrameRouteInfo {
	this := FrameRouteInfo{}
	return &this
}

// GetIpv4Mask returns the Ipv4Mask field value if set, zero value otherwise.
func (o *FrameRouteInfo) GetIpv4Mask() string {
	if o == nil || IsNil(o.Ipv4Mask) {
		var ret string
		return ret
	}
	return *o.Ipv4Mask
}

// GetIpv4MaskOk returns a tuple with the Ipv4Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameRouteInfo) GetIpv4MaskOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Mask) {
		return nil, false
	}
	return o.Ipv4Mask, true
}

// HasIpv4Mask returns a boolean if a field has been set.
func (o *FrameRouteInfo) HasIpv4Mask() bool {
	if o != nil && !IsNil(o.Ipv4Mask) {
		return true
	}

	return false
}

// SetIpv4Mask gets a reference to the given string and assigns it to the Ipv4Mask field.
func (o *FrameRouteInfo) SetIpv4Mask(v string) {
	o.Ipv4Mask = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *FrameRouteInfo) GetIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameRouteInfo) GetIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6Prefix) {
		return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *FrameRouteInfo) HasIpv6Prefix() bool {
	if o != nil && !IsNil(o.Ipv6Prefix) {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the Ipv6Prefix field.
func (o *FrameRouteInfo) SetIpv6Prefix(v Ipv6Prefix) {
	o.Ipv6Prefix = &v
}

func (o FrameRouteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameRouteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Mask) {
		toSerialize["ipv4Mask"] = o.Ipv4Mask
	}
	if !IsNil(o.Ipv6Prefix) {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix
	}
	return toSerialize, nil
}

type NullableFrameRouteInfo struct {
	value *FrameRouteInfo
	isSet bool
}

func (v NullableFrameRouteInfo) Get() *FrameRouteInfo {
	return v.value
}

func (v *NullableFrameRouteInfo) Set(val *FrameRouteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameRouteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameRouteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameRouteInfo(val *FrameRouteInfo) *NullableFrameRouteInfo {
	return &NullableFrameRouteInfo{value: val, isSet: true}
}

func (v NullableFrameRouteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameRouteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


