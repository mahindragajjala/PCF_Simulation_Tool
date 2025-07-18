/*
Nudm_UECM

Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PcscfAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcscfAddress{}

// PcscfAddress Contains the addressing information (IP addresses and/or FQDN) of the P-CSCF
type PcscfAddress struct {
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`
	// Fully Qualified Domain Name
	Fqdn *string `json:"fqdn,omitempty"`
}

// NewPcscfAddress instantiates a new PcscfAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcscfAddress() *PcscfAddress {
	this := PcscfAddress{}
	return &this
}

// NewPcscfAddressWithDefaults instantiates a new PcscfAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcscfAddressWithDefaults() *PcscfAddress {
	this := PcscfAddress{}
	return &this
}

// GetIpv4Addrs returns the Ipv4Addrs field value if set, zero value otherwise.
func (o *PcscfAddress) GetIpv4Addrs() []string {
	if o == nil || IsNil(o.Ipv4Addrs) {
		var ret []string
		return ret
	}
	return o.Ipv4Addrs
}

// GetIpv4AddrsOk returns a tuple with the Ipv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfAddress) GetIpv4AddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4Addrs) {
		return nil, false
	}
	return o.Ipv4Addrs, true
}

// HasIpv4Addrs returns a boolean if a field has been set.
func (o *PcscfAddress) HasIpv4Addrs() bool {
	if o != nil && !IsNil(o.Ipv4Addrs) {
		return true
	}

	return false
}

// SetIpv4Addrs gets a reference to the given []string and assigns it to the Ipv4Addrs field.
func (o *PcscfAddress) SetIpv4Addrs(v []string) {
	o.Ipv4Addrs = v
}

// GetIpv6Addrs returns the Ipv6Addrs field value if set, zero value otherwise.
func (o *PcscfAddress) GetIpv6Addrs() []Ipv6Addr {
	if o == nil || IsNil(o.Ipv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6Addrs
}

// GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfAddress) GetIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6Addrs) {
		return nil, false
	}
	return o.Ipv6Addrs, true
}

// HasIpv6Addrs returns a boolean if a field has been set.
func (o *PcscfAddress) HasIpv6Addrs() bool {
	if o != nil && !IsNil(o.Ipv6Addrs) {
		return true
	}

	return false
}

// SetIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the Ipv6Addrs field.
func (o *PcscfAddress) SetIpv6Addrs(v []Ipv6Addr) {
	o.Ipv6Addrs = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *PcscfAddress) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcscfAddress) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *PcscfAddress) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *PcscfAddress) SetFqdn(v string) {
	o.Fqdn = &v
}

func (o PcscfAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcscfAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Addrs) {
		toSerialize["ipv4Addrs"] = o.Ipv4Addrs
	}
	if !IsNil(o.Ipv6Addrs) {
		toSerialize["ipv6Addrs"] = o.Ipv6Addrs
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	return toSerialize, nil
}

type NullablePcscfAddress struct {
	value *PcscfAddress
	isSet bool
}

func (v NullablePcscfAddress) Get() *PcscfAddress {
	return v.value
}

func (v *NullablePcscfAddress) Set(val *PcscfAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePcscfAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePcscfAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcscfAddress(val *PcscfAddress) *NullablePcscfAddress {
	return &NullablePcscfAddress{value: val, isSet: true}
}

func (v NullablePcscfAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcscfAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


