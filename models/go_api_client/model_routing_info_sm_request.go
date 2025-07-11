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

// checks if the RoutingInfoSmRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInfoSmRequest{}

// RoutingInfoSmRequest Request body of the send-routing-info-sm custom operation
type RoutingInfoSmRequest struct {
	IpSmGwInd *bool `json:"ipSmGwInd,omitempty"`
	CorrelationId *string `json:"correlationId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
}

// NewRoutingInfoSmRequest instantiates a new RoutingInfoSmRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInfoSmRequest() *RoutingInfoSmRequest {
	this := RoutingInfoSmRequest{}
	var ipSmGwInd bool = false
	this.IpSmGwInd = &ipSmGwInd
	return &this
}

// NewRoutingInfoSmRequestWithDefaults instantiates a new RoutingInfoSmRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInfoSmRequestWithDefaults() *RoutingInfoSmRequest {
	this := RoutingInfoSmRequest{}
	var ipSmGwInd bool = false
	this.IpSmGwInd = &ipSmGwInd
	return &this
}

// GetIpSmGwInd returns the IpSmGwInd field value if set, zero value otherwise.
func (o *RoutingInfoSmRequest) GetIpSmGwInd() bool {
	if o == nil || IsNil(o.IpSmGwInd) {
		var ret bool
		return ret
	}
	return *o.IpSmGwInd
}

// GetIpSmGwIndOk returns a tuple with the IpSmGwInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoSmRequest) GetIpSmGwIndOk() (*bool, bool) {
	if o == nil || IsNil(o.IpSmGwInd) {
		return nil, false
	}
	return o.IpSmGwInd, true
}

// HasIpSmGwInd returns a boolean if a field has been set.
func (o *RoutingInfoSmRequest) HasIpSmGwInd() bool {
	if o != nil && !IsNil(o.IpSmGwInd) {
		return true
	}

	return false
}

// SetIpSmGwInd gets a reference to the given bool and assigns it to the IpSmGwInd field.
func (o *RoutingInfoSmRequest) SetIpSmGwInd(v bool) {
	o.IpSmGwInd = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *RoutingInfoSmRequest) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoSmRequest) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *RoutingInfoSmRequest) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *RoutingInfoSmRequest) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RoutingInfoSmRequest) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoSmRequest) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RoutingInfoSmRequest) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RoutingInfoSmRequest) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o RoutingInfoSmRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInfoSmRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpSmGwInd) {
		toSerialize["ipSmGwInd"] = o.IpSmGwInd
	}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableRoutingInfoSmRequest struct {
	value *RoutingInfoSmRequest
	isSet bool
}

func (v NullableRoutingInfoSmRequest) Get() *RoutingInfoSmRequest {
	return v.value
}

func (v *NullableRoutingInfoSmRequest) Set(val *RoutingInfoSmRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInfoSmRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInfoSmRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInfoSmRequest(val *RoutingInfoSmRequest) *NullableRoutingInfoSmRequest {
	return &NullableRoutingInfoSmRequest{value: val, isSet: true}
}

func (v NullableRoutingInfoSmRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInfoSmRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


