/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SdmSubsModification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SdmSubsModification{}

// SdmSubsModification struct for SdmSubsModification
type SdmSubsModification struct {
	// string with format 'date-time' as defined in OpenAPI.
	Expires *time.Time `json:"expires,omitempty"`
	MonitoredResourceUris []string `json:"monitoredResourceUris,omitempty"`
	// A map(list of key-value pairs) where a valid JSON pointer serves as key of expectedUeBehaviourThresholds
	ExpectedUeBehaviourThresholds *map[string]ExpectedUeBehaviourThreshold `json:"expectedUeBehaviourThresholds,omitempty"`
}

// NewSdmSubsModification instantiates a new SdmSubsModification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdmSubsModification() *SdmSubsModification {
	this := SdmSubsModification{}
	return &this
}

// NewSdmSubsModificationWithDefaults instantiates a new SdmSubsModification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdmSubsModificationWithDefaults() *SdmSubsModification {
	this := SdmSubsModification{}
	return &this
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *SdmSubsModification) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires) {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubsModification) GetExpiresOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *SdmSubsModification) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *SdmSubsModification) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetMonitoredResourceUris returns the MonitoredResourceUris field value if set, zero value otherwise.
func (o *SdmSubsModification) GetMonitoredResourceUris() []string {
	if o == nil || IsNil(o.MonitoredResourceUris) {
		var ret []string
		return ret
	}
	return o.MonitoredResourceUris
}

// GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubsModification) GetMonitoredResourceUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.MonitoredResourceUris) {
		return nil, false
	}
	return o.MonitoredResourceUris, true
}

// HasMonitoredResourceUris returns a boolean if a field has been set.
func (o *SdmSubsModification) HasMonitoredResourceUris() bool {
	if o != nil && !IsNil(o.MonitoredResourceUris) {
		return true
	}

	return false
}

// SetMonitoredResourceUris gets a reference to the given []string and assigns it to the MonitoredResourceUris field.
func (o *SdmSubsModification) SetMonitoredResourceUris(v []string) {
	o.MonitoredResourceUris = v
}

// GetExpectedUeBehaviourThresholds returns the ExpectedUeBehaviourThresholds field value if set, zero value otherwise.
func (o *SdmSubsModification) GetExpectedUeBehaviourThresholds() map[string]ExpectedUeBehaviourThreshold {
	if o == nil || IsNil(o.ExpectedUeBehaviourThresholds) {
		var ret map[string]ExpectedUeBehaviourThreshold
		return ret
	}
	return *o.ExpectedUeBehaviourThresholds
}

// GetExpectedUeBehaviourThresholdsOk returns a tuple with the ExpectedUeBehaviourThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdmSubsModification) GetExpectedUeBehaviourThresholdsOk() (*map[string]ExpectedUeBehaviourThreshold, bool) {
	if o == nil || IsNil(o.ExpectedUeBehaviourThresholds) {
		return nil, false
	}
	return o.ExpectedUeBehaviourThresholds, true
}

// HasExpectedUeBehaviourThresholds returns a boolean if a field has been set.
func (o *SdmSubsModification) HasExpectedUeBehaviourThresholds() bool {
	if o != nil && !IsNil(o.ExpectedUeBehaviourThresholds) {
		return true
	}

	return false
}

// SetExpectedUeBehaviourThresholds gets a reference to the given map[string]ExpectedUeBehaviourThreshold and assigns it to the ExpectedUeBehaviourThresholds field.
func (o *SdmSubsModification) SetExpectedUeBehaviourThresholds(v map[string]ExpectedUeBehaviourThreshold) {
	o.ExpectedUeBehaviourThresholds = &v
}

func (o SdmSubsModification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SdmSubsModification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !IsNil(o.MonitoredResourceUris) {
		toSerialize["monitoredResourceUris"] = o.MonitoredResourceUris
	}
	if !IsNil(o.ExpectedUeBehaviourThresholds) {
		toSerialize["expectedUeBehaviourThresholds"] = o.ExpectedUeBehaviourThresholds
	}
	return toSerialize, nil
}

type NullableSdmSubsModification struct {
	value *SdmSubsModification
	isSet bool
}

func (v NullableSdmSubsModification) Get() *SdmSubsModification {
	return v.value
}

func (v *NullableSdmSubsModification) Set(val *SdmSubsModification) {
	v.value = val
	v.isSet = true
}

func (v NullableSdmSubsModification) IsSet() bool {
	return v.isSet
}

func (v *NullableSdmSubsModification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdmSubsModification(val *SdmSubsModification) *NullableSdmSubsModification {
	return &NullableSdmSubsModification{value: val, isSet: true}
}

func (v NullableSdmSubsModification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdmSubsModification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


