/*
Nudm_PP

Nudm Parameter Provision Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ExpectedUeBehaviour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpectedUeBehaviour{}

// ExpectedUeBehaviour struct for ExpectedUeBehaviour
type ExpectedUeBehaviour struct {
	AfInstanceId string `json:"afInstanceId"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	ReferenceId int32 `json:"referenceId"`
	StationaryIndication NullableStationaryIndication `json:"stationaryIndication,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	CommunicationDurationTime NullableInt32 `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationType NullableScheduledCommunicationType `json:"scheduledCommunicationType,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	PeriodicTime NullableInt32 `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime NullableScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	// Identifies the UE's expected geographical movement. The attribute is only applicable in 5G. 
	ExpectedUmts []LocationArea `json:"expectedUmts,omitempty"`
	TrafficProfile NullableTrafficProfile `json:"trafficProfile,omitempty"`
	BatteryIndication NullableBatteryIndication `json:"batteryIndication,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
}

type _ExpectedUeBehaviour ExpectedUeBehaviour

// NewExpectedUeBehaviour instantiates a new ExpectedUeBehaviour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpectedUeBehaviour(afInstanceId string, referenceId int32) *ExpectedUeBehaviour {
	this := ExpectedUeBehaviour{}
	this.AfInstanceId = afInstanceId
	this.ReferenceId = referenceId
	return &this
}

// NewExpectedUeBehaviourWithDefaults instantiates a new ExpectedUeBehaviour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpectedUeBehaviourWithDefaults() *ExpectedUeBehaviour {
	this := ExpectedUeBehaviour{}
	return &this
}

// GetAfInstanceId returns the AfInstanceId field value
func (o *ExpectedUeBehaviour) GetAfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfInstanceId
}

// GetAfInstanceIdOk returns a tuple with the AfInstanceId field value
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetAfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfInstanceId, true
}

// SetAfInstanceId sets field value
func (o *ExpectedUeBehaviour) SetAfInstanceId(v string) {
	o.AfInstanceId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *ExpectedUeBehaviour) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *ExpectedUeBehaviour) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetStationaryIndication returns the StationaryIndication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetStationaryIndication() StationaryIndication {
	if o == nil || IsNil(o.StationaryIndication.Get()) {
		var ret StationaryIndication
		return ret
	}
	return *o.StationaryIndication.Get()
}

// GetStationaryIndicationOk returns a tuple with the StationaryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetStationaryIndicationOk() (*StationaryIndication, bool) {
	if o == nil {
		return nil, false
	}
	return o.StationaryIndication.Get(), o.StationaryIndication.IsSet()
}

// HasStationaryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasStationaryIndication() bool {
	if o != nil && o.StationaryIndication.IsSet() {
		return true
	}

	return false
}

// SetStationaryIndication gets a reference to the given NullableStationaryIndication and assigns it to the StationaryIndication field.
func (o *ExpectedUeBehaviour) SetStationaryIndication(v StationaryIndication) {
	o.StationaryIndication.Set(&v)
}
// SetStationaryIndicationNil sets the value for StationaryIndication to be an explicit nil
func (o *ExpectedUeBehaviour) SetStationaryIndicationNil() {
	o.StationaryIndication.Set(nil)
}

// UnsetStationaryIndication ensures that no value is present for StationaryIndication, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetStationaryIndication() {
	o.StationaryIndication.Unset()
}

// GetCommunicationDurationTime returns the CommunicationDurationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetCommunicationDurationTime() int32 {
	if o == nil || IsNil(o.CommunicationDurationTime.Get()) {
		var ret int32
		return ret
	}
	return *o.CommunicationDurationTime.Get()
}

// GetCommunicationDurationTimeOk returns a tuple with the CommunicationDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetCommunicationDurationTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommunicationDurationTime.Get(), o.CommunicationDurationTime.IsSet()
}

// HasCommunicationDurationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasCommunicationDurationTime() bool {
	if o != nil && o.CommunicationDurationTime.IsSet() {
		return true
	}

	return false
}

// SetCommunicationDurationTime gets a reference to the given NullableInt32 and assigns it to the CommunicationDurationTime field.
func (o *ExpectedUeBehaviour) SetCommunicationDurationTime(v int32) {
	o.CommunicationDurationTime.Set(&v)
}
// SetCommunicationDurationTimeNil sets the value for CommunicationDurationTime to be an explicit nil
func (o *ExpectedUeBehaviour) SetCommunicationDurationTimeNil() {
	o.CommunicationDurationTime.Set(nil)
}

// UnsetCommunicationDurationTime ensures that no value is present for CommunicationDurationTime, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetCommunicationDurationTime() {
	o.CommunicationDurationTime.Unset()
}

// GetScheduledCommunicationType returns the ScheduledCommunicationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetScheduledCommunicationType() ScheduledCommunicationType {
	if o == nil || IsNil(o.ScheduledCommunicationType.Get()) {
		var ret ScheduledCommunicationType
		return ret
	}
	return *o.ScheduledCommunicationType.Get()
}

// GetScheduledCommunicationTypeOk returns a tuple with the ScheduledCommunicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTypeOk() (*ScheduledCommunicationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledCommunicationType.Get(), o.ScheduledCommunicationType.IsSet()
}

// HasScheduledCommunicationType returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasScheduledCommunicationType() bool {
	if o != nil && o.ScheduledCommunicationType.IsSet() {
		return true
	}

	return false
}

// SetScheduledCommunicationType gets a reference to the given NullableScheduledCommunicationType and assigns it to the ScheduledCommunicationType field.
func (o *ExpectedUeBehaviour) SetScheduledCommunicationType(v ScheduledCommunicationType) {
	o.ScheduledCommunicationType.Set(&v)
}
// SetScheduledCommunicationTypeNil sets the value for ScheduledCommunicationType to be an explicit nil
func (o *ExpectedUeBehaviour) SetScheduledCommunicationTypeNil() {
	o.ScheduledCommunicationType.Set(nil)
}

// UnsetScheduledCommunicationType ensures that no value is present for ScheduledCommunicationType, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetScheduledCommunicationType() {
	o.ScheduledCommunicationType.Unset()
}

// GetPeriodicTime returns the PeriodicTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetPeriodicTime() int32 {
	if o == nil || IsNil(o.PeriodicTime.Get()) {
		var ret int32
		return ret
	}
	return *o.PeriodicTime.Get()
}

// GetPeriodicTimeOk returns a tuple with the PeriodicTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetPeriodicTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodicTime.Get(), o.PeriodicTime.IsSet()
}

// HasPeriodicTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasPeriodicTime() bool {
	if o != nil && o.PeriodicTime.IsSet() {
		return true
	}

	return false
}

// SetPeriodicTime gets a reference to the given NullableInt32 and assigns it to the PeriodicTime field.
func (o *ExpectedUeBehaviour) SetPeriodicTime(v int32) {
	o.PeriodicTime.Set(&v)
}
// SetPeriodicTimeNil sets the value for PeriodicTime to be an explicit nil
func (o *ExpectedUeBehaviour) SetPeriodicTimeNil() {
	o.PeriodicTime.Set(nil)
}

// UnsetPeriodicTime ensures that no value is present for PeriodicTime, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetPeriodicTime() {
	o.PeriodicTime.Unset()
}

// GetScheduledCommunicationTime returns the ScheduledCommunicationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTime() ScheduledCommunicationTime {
	if o == nil || IsNil(o.ScheduledCommunicationTime.Get()) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.ScheduledCommunicationTime.Get()
}

// GetScheduledCommunicationTimeOk returns a tuple with the ScheduledCommunicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetScheduledCommunicationTimeOk() (*ScheduledCommunicationTime, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledCommunicationTime.Get(), o.ScheduledCommunicationTime.IsSet()
}

// HasScheduledCommunicationTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasScheduledCommunicationTime() bool {
	if o != nil && o.ScheduledCommunicationTime.IsSet() {
		return true
	}

	return false
}

// SetScheduledCommunicationTime gets a reference to the given NullableScheduledCommunicationTime and assigns it to the ScheduledCommunicationTime field.
func (o *ExpectedUeBehaviour) SetScheduledCommunicationTime(v ScheduledCommunicationTime) {
	o.ScheduledCommunicationTime.Set(&v)
}
// SetScheduledCommunicationTimeNil sets the value for ScheduledCommunicationTime to be an explicit nil
func (o *ExpectedUeBehaviour) SetScheduledCommunicationTimeNil() {
	o.ScheduledCommunicationTime.Set(nil)
}

// UnsetScheduledCommunicationTime ensures that no value is present for ScheduledCommunicationTime, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetScheduledCommunicationTime() {
	o.ScheduledCommunicationTime.Unset()
}

// GetExpectedUmts returns the ExpectedUmts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetExpectedUmts() []LocationArea {
	if o == nil {
		var ret []LocationArea
		return ret
	}
	return o.ExpectedUmts
}

// GetExpectedUmtsOk returns a tuple with the ExpectedUmts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetExpectedUmtsOk() ([]LocationArea, bool) {
	if o == nil || IsNil(o.ExpectedUmts) {
		return nil, false
	}
	return o.ExpectedUmts, true
}

// HasExpectedUmts returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasExpectedUmts() bool {
	if o != nil && !IsNil(o.ExpectedUmts) {
		return true
	}

	return false
}

// SetExpectedUmts gets a reference to the given []LocationArea and assigns it to the ExpectedUmts field.
func (o *ExpectedUeBehaviour) SetExpectedUmts(v []LocationArea) {
	o.ExpectedUmts = v
}

// GetTrafficProfile returns the TrafficProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetTrafficProfile() TrafficProfile {
	if o == nil || IsNil(o.TrafficProfile.Get()) {
		var ret TrafficProfile
		return ret
	}
	return *o.TrafficProfile.Get()
}

// GetTrafficProfileOk returns a tuple with the TrafficProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetTrafficProfileOk() (*TrafficProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficProfile.Get(), o.TrafficProfile.IsSet()
}

// HasTrafficProfile returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasTrafficProfile() bool {
	if o != nil && o.TrafficProfile.IsSet() {
		return true
	}

	return false
}

// SetTrafficProfile gets a reference to the given NullableTrafficProfile and assigns it to the TrafficProfile field.
func (o *ExpectedUeBehaviour) SetTrafficProfile(v TrafficProfile) {
	o.TrafficProfile.Set(&v)
}
// SetTrafficProfileNil sets the value for TrafficProfile to be an explicit nil
func (o *ExpectedUeBehaviour) SetTrafficProfileNil() {
	o.TrafficProfile.Set(nil)
}

// UnsetTrafficProfile ensures that no value is present for TrafficProfile, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetTrafficProfile() {
	o.TrafficProfile.Unset()
}

// GetBatteryIndication returns the BatteryIndication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpectedUeBehaviour) GetBatteryIndication() BatteryIndication {
	if o == nil || IsNil(o.BatteryIndication.Get()) {
		var ret BatteryIndication
		return ret
	}
	return *o.BatteryIndication.Get()
}

// GetBatteryIndicationOk returns a tuple with the BatteryIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpectedUeBehaviour) GetBatteryIndicationOk() (*BatteryIndication, bool) {
	if o == nil {
		return nil, false
	}
	return o.BatteryIndication.Get(), o.BatteryIndication.IsSet()
}

// HasBatteryIndication returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasBatteryIndication() bool {
	if o != nil && o.BatteryIndication.IsSet() {
		return true
	}

	return false
}

// SetBatteryIndication gets a reference to the given NullableBatteryIndication and assigns it to the BatteryIndication field.
func (o *ExpectedUeBehaviour) SetBatteryIndication(v BatteryIndication) {
	o.BatteryIndication.Set(&v)
}
// SetBatteryIndicationNil sets the value for BatteryIndication to be an explicit nil
func (o *ExpectedUeBehaviour) SetBatteryIndicationNil() {
	o.BatteryIndication.Set(nil)
}

// UnsetBatteryIndication ensures that no value is present for BatteryIndication, not even an explicit nil
func (o *ExpectedUeBehaviour) UnsetBatteryIndication() {
	o.BatteryIndication.Unset()
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *ExpectedUeBehaviour) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *ExpectedUeBehaviour) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpectedUeBehaviour) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *ExpectedUeBehaviour) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *ExpectedUeBehaviour) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

func (o ExpectedUeBehaviour) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpectedUeBehaviour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afInstanceId"] = o.AfInstanceId
	toSerialize["referenceId"] = o.ReferenceId
	if o.StationaryIndication.IsSet() {
		toSerialize["stationaryIndication"] = o.StationaryIndication.Get()
	}
	if o.CommunicationDurationTime.IsSet() {
		toSerialize["communicationDurationTime"] = o.CommunicationDurationTime.Get()
	}
	if o.ScheduledCommunicationType.IsSet() {
		toSerialize["scheduledCommunicationType"] = o.ScheduledCommunicationType.Get()
	}
	if o.PeriodicTime.IsSet() {
		toSerialize["periodicTime"] = o.PeriodicTime.Get()
	}
	if o.ScheduledCommunicationTime.IsSet() {
		toSerialize["scheduledCommunicationTime"] = o.ScheduledCommunicationTime.Get()
	}
	if o.ExpectedUmts != nil {
		toSerialize["expectedUmts"] = o.ExpectedUmts
	}
	if o.TrafficProfile.IsSet() {
		toSerialize["trafficProfile"] = o.TrafficProfile.Get()
	}
	if o.BatteryIndication.IsSet() {
		toSerialize["batteryIndication"] = o.BatteryIndication.Get()
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	return toSerialize, nil
}

func (o *ExpectedUeBehaviour) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"afInstanceId",
		"referenceId",
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

	varExpectedUeBehaviour := _ExpectedUeBehaviour{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExpectedUeBehaviour)

	if err != nil {
		return err
	}

	*o = ExpectedUeBehaviour(varExpectedUeBehaviour)

	return err
}

type NullableExpectedUeBehaviour struct {
	value *ExpectedUeBehaviour
	isSet bool
}

func (v NullableExpectedUeBehaviour) Get() *ExpectedUeBehaviour {
	return v.value
}

func (v *NullableExpectedUeBehaviour) Set(val *ExpectedUeBehaviour) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectedUeBehaviour) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectedUeBehaviour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectedUeBehaviour(val *ExpectedUeBehaviour) *NullableExpectedUeBehaviour {
	return &NullableExpectedUeBehaviour{value: val, isSet: true}
}

func (v NullableExpectedUeBehaviour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectedUeBehaviour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


