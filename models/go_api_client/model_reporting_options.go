/*
Nudm_EE

Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ReportingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportingOptions{}

// ReportingOptions struct for ReportingOptions
type ReportingOptions struct {
	ReportMode *EventReportMode `json:"reportMode,omitempty"`
	MaxNumOfReports *int32 `json:"maxNumOfReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SamplingRatio *int32 `json:"samplingRatio,omitempty"`
	// indicating a time in seconds.
	GuardTime *int32 `json:"guardTime,omitempty"`
	// indicating a time in seconds.
	ReportPeriod *int32 `json:"reportPeriod,omitempty"`
	NotifFlag *NotificationFlag `json:"notifFlag,omitempty"`
	MutingExcInstructions *MutingExceptionInstructions `json:"mutingExcInstructions,omitempty"`
	MutingNotSettings *MutingNotificationsSettings `json:"mutingNotSettings,omitempty"`
	VarRepPeriodInfo []VarRepPeriod `json:"varRepPeriodInfo,omitempty"`
}

// NewReportingOptions instantiates a new ReportingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingOptions() *ReportingOptions {
	this := ReportingOptions{}
	return &this
}

// NewReportingOptionsWithDefaults instantiates a new ReportingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingOptionsWithDefaults() *ReportingOptions {
	this := ReportingOptions{}
	return &this
}

// GetReportMode returns the ReportMode field value if set, zero value otherwise.
func (o *ReportingOptions) GetReportMode() EventReportMode {
	if o == nil || IsNil(o.ReportMode) {
		var ret EventReportMode
		return ret
	}
	return *o.ReportMode
}

// GetReportModeOk returns a tuple with the ReportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetReportModeOk() (*EventReportMode, bool) {
	if o == nil || IsNil(o.ReportMode) {
		return nil, false
	}
	return o.ReportMode, true
}

// HasReportMode returns a boolean if a field has been set.
func (o *ReportingOptions) HasReportMode() bool {
	if o != nil && !IsNil(o.ReportMode) {
		return true
	}

	return false
}

// SetReportMode gets a reference to the given EventReportMode and assigns it to the ReportMode field.
func (o *ReportingOptions) SetReportMode(v EventReportMode) {
	o.ReportMode = &v
}

// GetMaxNumOfReports returns the MaxNumOfReports field value if set, zero value otherwise.
func (o *ReportingOptions) GetMaxNumOfReports() int32 {
	if o == nil || IsNil(o.MaxNumOfReports) {
		var ret int32
		return ret
	}
	return *o.MaxNumOfReports
}

// GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetMaxNumOfReportsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNumOfReports) {
		return nil, false
	}
	return o.MaxNumOfReports, true
}

// HasMaxNumOfReports returns a boolean if a field has been set.
func (o *ReportingOptions) HasMaxNumOfReports() bool {
	if o != nil && !IsNil(o.MaxNumOfReports) {
		return true
	}

	return false
}

// SetMaxNumOfReports gets a reference to the given int32 and assigns it to the MaxNumOfReports field.
func (o *ReportingOptions) SetMaxNumOfReports(v int32) {
	o.MaxNumOfReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ReportingOptions) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ReportingOptions) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *ReportingOptions) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetSamplingRatio returns the SamplingRatio field value if set, zero value otherwise.
func (o *ReportingOptions) GetSamplingRatio() int32 {
	if o == nil || IsNil(o.SamplingRatio) {
		var ret int32
		return ret
	}
	return *o.SamplingRatio
}

// GetSamplingRatioOk returns a tuple with the SamplingRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetSamplingRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.SamplingRatio) {
		return nil, false
	}
	return o.SamplingRatio, true
}

// HasSamplingRatio returns a boolean if a field has been set.
func (o *ReportingOptions) HasSamplingRatio() bool {
	if o != nil && !IsNil(o.SamplingRatio) {
		return true
	}

	return false
}

// SetSamplingRatio gets a reference to the given int32 and assigns it to the SamplingRatio field.
func (o *ReportingOptions) SetSamplingRatio(v int32) {
	o.SamplingRatio = &v
}

// GetGuardTime returns the GuardTime field value if set, zero value otherwise.
func (o *ReportingOptions) GetGuardTime() int32 {
	if o == nil || IsNil(o.GuardTime) {
		var ret int32
		return ret
	}
	return *o.GuardTime
}

// GetGuardTimeOk returns a tuple with the GuardTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetGuardTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.GuardTime) {
		return nil, false
	}
	return o.GuardTime, true
}

// HasGuardTime returns a boolean if a field has been set.
func (o *ReportingOptions) HasGuardTime() bool {
	if o != nil && !IsNil(o.GuardTime) {
		return true
	}

	return false
}

// SetGuardTime gets a reference to the given int32 and assigns it to the GuardTime field.
func (o *ReportingOptions) SetGuardTime(v int32) {
	o.GuardTime = &v
}

// GetReportPeriod returns the ReportPeriod field value if set, zero value otherwise.
func (o *ReportingOptions) GetReportPeriod() int32 {
	if o == nil || IsNil(o.ReportPeriod) {
		var ret int32
		return ret
	}
	return *o.ReportPeriod
}

// GetReportPeriodOk returns a tuple with the ReportPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetReportPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportPeriod) {
		return nil, false
	}
	return o.ReportPeriod, true
}

// HasReportPeriod returns a boolean if a field has been set.
func (o *ReportingOptions) HasReportPeriod() bool {
	if o != nil && !IsNil(o.ReportPeriod) {
		return true
	}

	return false
}

// SetReportPeriod gets a reference to the given int32 and assigns it to the ReportPeriod field.
func (o *ReportingOptions) SetReportPeriod(v int32) {
	o.ReportPeriod = &v
}

// GetNotifFlag returns the NotifFlag field value if set, zero value otherwise.
func (o *ReportingOptions) GetNotifFlag() NotificationFlag {
	if o == nil || IsNil(o.NotifFlag) {
		var ret NotificationFlag
		return ret
	}
	return *o.NotifFlag
}

// GetNotifFlagOk returns a tuple with the NotifFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetNotifFlagOk() (*NotificationFlag, bool) {
	if o == nil || IsNil(o.NotifFlag) {
		return nil, false
	}
	return o.NotifFlag, true
}

// HasNotifFlag returns a boolean if a field has been set.
func (o *ReportingOptions) HasNotifFlag() bool {
	if o != nil && !IsNil(o.NotifFlag) {
		return true
	}

	return false
}

// SetNotifFlag gets a reference to the given NotificationFlag and assigns it to the NotifFlag field.
func (o *ReportingOptions) SetNotifFlag(v NotificationFlag) {
	o.NotifFlag = &v
}

// GetMutingExcInstructions returns the MutingExcInstructions field value if set, zero value otherwise.
func (o *ReportingOptions) GetMutingExcInstructions() MutingExceptionInstructions {
	if o == nil || IsNil(o.MutingExcInstructions) {
		var ret MutingExceptionInstructions
		return ret
	}
	return *o.MutingExcInstructions
}

// GetMutingExcInstructionsOk returns a tuple with the MutingExcInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetMutingExcInstructionsOk() (*MutingExceptionInstructions, bool) {
	if o == nil || IsNil(o.MutingExcInstructions) {
		return nil, false
	}
	return o.MutingExcInstructions, true
}

// HasMutingExcInstructions returns a boolean if a field has been set.
func (o *ReportingOptions) HasMutingExcInstructions() bool {
	if o != nil && !IsNil(o.MutingExcInstructions) {
		return true
	}

	return false
}

// SetMutingExcInstructions gets a reference to the given MutingExceptionInstructions and assigns it to the MutingExcInstructions field.
func (o *ReportingOptions) SetMutingExcInstructions(v MutingExceptionInstructions) {
	o.MutingExcInstructions = &v
}

// GetMutingNotSettings returns the MutingNotSettings field value if set, zero value otherwise.
func (o *ReportingOptions) GetMutingNotSettings() MutingNotificationsSettings {
	if o == nil || IsNil(o.MutingNotSettings) {
		var ret MutingNotificationsSettings
		return ret
	}
	return *o.MutingNotSettings
}

// GetMutingNotSettingsOk returns a tuple with the MutingNotSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetMutingNotSettingsOk() (*MutingNotificationsSettings, bool) {
	if o == nil || IsNil(o.MutingNotSettings) {
		return nil, false
	}
	return o.MutingNotSettings, true
}

// HasMutingNotSettings returns a boolean if a field has been set.
func (o *ReportingOptions) HasMutingNotSettings() bool {
	if o != nil && !IsNil(o.MutingNotSettings) {
		return true
	}

	return false
}

// SetMutingNotSettings gets a reference to the given MutingNotificationsSettings and assigns it to the MutingNotSettings field.
func (o *ReportingOptions) SetMutingNotSettings(v MutingNotificationsSettings) {
	o.MutingNotSettings = &v
}

// GetVarRepPeriodInfo returns the VarRepPeriodInfo field value if set, zero value otherwise.
func (o *ReportingOptions) GetVarRepPeriodInfo() []VarRepPeriod {
	if o == nil || IsNil(o.VarRepPeriodInfo) {
		var ret []VarRepPeriod
		return ret
	}
	return o.VarRepPeriodInfo
}

// GetVarRepPeriodInfoOk returns a tuple with the VarRepPeriodInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingOptions) GetVarRepPeriodInfoOk() ([]VarRepPeriod, bool) {
	if o == nil || IsNil(o.VarRepPeriodInfo) {
		return nil, false
	}
	return o.VarRepPeriodInfo, true
}

// HasVarRepPeriodInfo returns a boolean if a field has been set.
func (o *ReportingOptions) HasVarRepPeriodInfo() bool {
	if o != nil && !IsNil(o.VarRepPeriodInfo) {
		return true
	}

	return false
}

// SetVarRepPeriodInfo gets a reference to the given []VarRepPeriod and assigns it to the VarRepPeriodInfo field.
func (o *ReportingOptions) SetVarRepPeriodInfo(v []VarRepPeriod) {
	o.VarRepPeriodInfo = v
}

func (o ReportingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportMode) {
		toSerialize["reportMode"] = o.ReportMode
	}
	if !IsNil(o.MaxNumOfReports) {
		toSerialize["maxNumOfReports"] = o.MaxNumOfReports
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.SamplingRatio) {
		toSerialize["samplingRatio"] = o.SamplingRatio
	}
	if !IsNil(o.GuardTime) {
		toSerialize["guardTime"] = o.GuardTime
	}
	if !IsNil(o.ReportPeriod) {
		toSerialize["reportPeriod"] = o.ReportPeriod
	}
	if !IsNil(o.NotifFlag) {
		toSerialize["notifFlag"] = o.NotifFlag
	}
	if !IsNil(o.MutingExcInstructions) {
		toSerialize["mutingExcInstructions"] = o.MutingExcInstructions
	}
	if !IsNil(o.MutingNotSettings) {
		toSerialize["mutingNotSettings"] = o.MutingNotSettings
	}
	if !IsNil(o.VarRepPeriodInfo) {
		toSerialize["varRepPeriodInfo"] = o.VarRepPeriodInfo
	}
	return toSerialize, nil
}

type NullableReportingOptions struct {
	value *ReportingOptions
	isSet bool
}

func (v NullableReportingOptions) Get() *ReportingOptions {
	return v.value
}

func (v *NullableReportingOptions) Set(val *ReportingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingOptions(val *ReportingOptions) *NullableReportingOptions {
	return &NullableReportingOptions{value: val, isSet: true}
}

func (v NullableReportingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


