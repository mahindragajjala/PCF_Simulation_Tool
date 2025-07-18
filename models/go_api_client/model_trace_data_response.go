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

// checks if the TraceDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraceDataResponse{}

// TraceDataResponse struct for TraceDataResponse
type TraceDataResponse struct {
	TraceData NullableTraceData `json:"traceData,omitempty"`
	SharedTraceDataId *string `json:"sharedTraceDataId,omitempty" validate:"regexp=^[0-9]{5,6}-.+$"`
}

// NewTraceDataResponse instantiates a new TraceDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceDataResponse() *TraceDataResponse {
	this := TraceDataResponse{}
	return &this
}

// NewTraceDataResponseWithDefaults instantiates a new TraceDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceDataResponseWithDefaults() *TraceDataResponse {
	this := TraceDataResponse{}
	return &this
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TraceDataResponse) GetTraceData() TraceData {
	if o == nil || IsNil(o.TraceData.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TraceDataResponse) GetTraceDataOk() (*TraceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *TraceDataResponse) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData and assigns it to the TraceData field.
func (o *TraceDataResponse) SetTraceData(v TraceData) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *TraceDataResponse) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *TraceDataResponse) UnsetTraceData() {
	o.TraceData.Unset()
}

// GetSharedTraceDataId returns the SharedTraceDataId field value if set, zero value otherwise.
func (o *TraceDataResponse) GetSharedTraceDataId() string {
	if o == nil || IsNil(o.SharedTraceDataId) {
		var ret string
		return ret
	}
	return *o.SharedTraceDataId
}

// GetSharedTraceDataIdOk returns a tuple with the SharedTraceDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceDataResponse) GetSharedTraceDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.SharedTraceDataId) {
		return nil, false
	}
	return o.SharedTraceDataId, true
}

// HasSharedTraceDataId returns a boolean if a field has been set.
func (o *TraceDataResponse) HasSharedTraceDataId() bool {
	if o != nil && !IsNil(o.SharedTraceDataId) {
		return true
	}

	return false
}

// SetSharedTraceDataId gets a reference to the given string and assigns it to the SharedTraceDataId field.
func (o *TraceDataResponse) SetSharedTraceDataId(v string) {
	o.SharedTraceDataId = &v
}

func (o TraceDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraceDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TraceData.IsSet() {
		toSerialize["traceData"] = o.TraceData.Get()
	}
	if !IsNil(o.SharedTraceDataId) {
		toSerialize["sharedTraceDataId"] = o.SharedTraceDataId
	}
	return toSerialize, nil
}

type NullableTraceDataResponse struct {
	value *TraceDataResponse
	isSet bool
}

func (v NullableTraceDataResponse) Get() *TraceDataResponse {
	return v.value
}

func (v *NullableTraceDataResponse) Set(val *TraceDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDataResponse(val *TraceDataResponse) *NullableTraceDataResponse {
	return &NullableTraceDataResponse{value: val, isSet: true}
}

func (v NullableTraceDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


