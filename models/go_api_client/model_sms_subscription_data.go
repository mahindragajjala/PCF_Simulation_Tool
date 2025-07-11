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

// checks if the SmsSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsSubscriptionData{}

// SmsSubscriptionData struct for SmsSubscriptionData
type SmsSubscriptionData struct {
	SmsSubscribed *bool `json:"smsSubscribed,omitempty"`
	SharedSmsSubsDataId *string `json:"sharedSmsSubsDataId,omitempty" validate:"regexp=^[0-9]{5,6}-.+$"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
}

// NewSmsSubscriptionData instantiates a new SmsSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsSubscriptionData() *SmsSubscriptionData {
	this := SmsSubscriptionData{}
	return &this
}

// NewSmsSubscriptionDataWithDefaults instantiates a new SmsSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsSubscriptionDataWithDefaults() *SmsSubscriptionData {
	this := SmsSubscriptionData{}
	return &this
}

// GetSmsSubscribed returns the SmsSubscribed field value if set, zero value otherwise.
func (o *SmsSubscriptionData) GetSmsSubscribed() bool {
	if o == nil || IsNil(o.SmsSubscribed) {
		var ret bool
		return ret
	}
	return *o.SmsSubscribed
}

// GetSmsSubscribedOk returns a tuple with the SmsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsSubscriptionData) GetSmsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.SmsSubscribed) {
		return nil, false
	}
	return o.SmsSubscribed, true
}

// HasSmsSubscribed returns a boolean if a field has been set.
func (o *SmsSubscriptionData) HasSmsSubscribed() bool {
	if o != nil && !IsNil(o.SmsSubscribed) {
		return true
	}

	return false
}

// SetSmsSubscribed gets a reference to the given bool and assigns it to the SmsSubscribed field.
func (o *SmsSubscriptionData) SetSmsSubscribed(v bool) {
	o.SmsSubscribed = &v
}

// GetSharedSmsSubsDataId returns the SharedSmsSubsDataId field value if set, zero value otherwise.
func (o *SmsSubscriptionData) GetSharedSmsSubsDataId() string {
	if o == nil || IsNil(o.SharedSmsSubsDataId) {
		var ret string
		return ret
	}
	return *o.SharedSmsSubsDataId
}

// GetSharedSmsSubsDataIdOk returns a tuple with the SharedSmsSubsDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsSubscriptionData) GetSharedSmsSubsDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSmsSubsDataId) {
		return nil, false
	}
	return o.SharedSmsSubsDataId, true
}

// HasSharedSmsSubsDataId returns a boolean if a field has been set.
func (o *SmsSubscriptionData) HasSharedSmsSubsDataId() bool {
	if o != nil && !IsNil(o.SharedSmsSubsDataId) {
		return true
	}

	return false
}

// SetSharedSmsSubsDataId gets a reference to the given string and assigns it to the SharedSmsSubsDataId field.
func (o *SmsSubscriptionData) SetSharedSmsSubsDataId(v string) {
	o.SharedSmsSubsDataId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *SmsSubscriptionData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsSubscriptionData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *SmsSubscriptionData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *SmsSubscriptionData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o SmsSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SmsSubscribed) {
		toSerialize["smsSubscribed"] = o.SmsSubscribed
	}
	if !IsNil(o.SharedSmsSubsDataId) {
		toSerialize["sharedSmsSubsDataId"] = o.SharedSmsSubsDataId
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableSmsSubscriptionData struct {
	value *SmsSubscriptionData
	isSet bool
}

func (v NullableSmsSubscriptionData) Get() *SmsSubscriptionData {
	return v.value
}

func (v *NullableSmsSubscriptionData) Set(val *SmsSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsSubscriptionData(val *SmsSubscriptionData) *NullableSmsSubscriptionData {
	return &NullableSmsSubscriptionData{value: val, isSet: true}
}

func (v NullableSmsSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


