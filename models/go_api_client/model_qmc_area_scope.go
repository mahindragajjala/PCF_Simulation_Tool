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

// checks if the QmcAreaScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QmcAreaScope{}

// QmcAreaScope This IE contains the area in Cells or Tracking Areas where the QMC data collection shall take place. 
type QmcAreaScope struct {
	NrCellIdList []string `json:"nrCellIdList,omitempty"`
	TacList []string `json:"tacList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	PlmnList []PlmnId `json:"plmnList,omitempty"`
}

// NewQmcAreaScope instantiates a new QmcAreaScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQmcAreaScope() *QmcAreaScope {
	this := QmcAreaScope{}
	return &this
}

// NewQmcAreaScopeWithDefaults instantiates a new QmcAreaScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQmcAreaScopeWithDefaults() *QmcAreaScope {
	this := QmcAreaScope{}
	return &this
}

// GetNrCellIdList returns the NrCellIdList field value if set, zero value otherwise.
func (o *QmcAreaScope) GetNrCellIdList() []string {
	if o == nil || IsNil(o.NrCellIdList) {
		var ret []string
		return ret
	}
	return o.NrCellIdList
}

// GetNrCellIdListOk returns a tuple with the NrCellIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QmcAreaScope) GetNrCellIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.NrCellIdList) {
		return nil, false
	}
	return o.NrCellIdList, true
}

// HasNrCellIdList returns a boolean if a field has been set.
func (o *QmcAreaScope) HasNrCellIdList() bool {
	if o != nil && !IsNil(o.NrCellIdList) {
		return true
	}

	return false
}

// SetNrCellIdList gets a reference to the given []string and assigns it to the NrCellIdList field.
func (o *QmcAreaScope) SetNrCellIdList(v []string) {
	o.NrCellIdList = v
}

// GetTacList returns the TacList field value if set, zero value otherwise.
func (o *QmcAreaScope) GetTacList() []string {
	if o == nil || IsNil(o.TacList) {
		var ret []string
		return ret
	}
	return o.TacList
}

// GetTacListOk returns a tuple with the TacList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QmcAreaScope) GetTacListOk() ([]string, bool) {
	if o == nil || IsNil(o.TacList) {
		return nil, false
	}
	return o.TacList, true
}

// HasTacList returns a boolean if a field has been set.
func (o *QmcAreaScope) HasTacList() bool {
	if o != nil && !IsNil(o.TacList) {
		return true
	}

	return false
}

// SetTacList gets a reference to the given []string and assigns it to the TacList field.
func (o *QmcAreaScope) SetTacList(v []string) {
	o.TacList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *QmcAreaScope) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QmcAreaScope) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *QmcAreaScope) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *QmcAreaScope) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetPlmnList returns the PlmnList field value if set, zero value otherwise.
func (o *QmcAreaScope) GetPlmnList() []PlmnId {
	if o == nil || IsNil(o.PlmnList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnList
}

// GetPlmnListOk returns a tuple with the PlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QmcAreaScope) GetPlmnListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnList) {
		return nil, false
	}
	return o.PlmnList, true
}

// HasPlmnList returns a boolean if a field has been set.
func (o *QmcAreaScope) HasPlmnList() bool {
	if o != nil && !IsNil(o.PlmnList) {
		return true
	}

	return false
}

// SetPlmnList gets a reference to the given []PlmnId and assigns it to the PlmnList field.
func (o *QmcAreaScope) SetPlmnList(v []PlmnId) {
	o.PlmnList = v
}

func (o QmcAreaScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QmcAreaScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NrCellIdList) {
		toSerialize["nrCellIdList"] = o.NrCellIdList
	}
	if !IsNil(o.TacList) {
		toSerialize["tacList"] = o.TacList
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.PlmnList) {
		toSerialize["plmnList"] = o.PlmnList
	}
	return toSerialize, nil
}

type NullableQmcAreaScope struct {
	value *QmcAreaScope
	isSet bool
}

func (v NullableQmcAreaScope) Get() *QmcAreaScope {
	return v.value
}

func (v *NullableQmcAreaScope) Set(val *QmcAreaScope) {
	v.value = val
	v.isSet = true
}

func (v NullableQmcAreaScope) IsSet() bool {
	return v.isSet
}

func (v *NullableQmcAreaScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQmcAreaScope(val *QmcAreaScope) *NullableQmcAreaScope {
	return &NullableQmcAreaScope{value: val, isSet: true}
}

func (v NullableQmcAreaScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQmcAreaScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


