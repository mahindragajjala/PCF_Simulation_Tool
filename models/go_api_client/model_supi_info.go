/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SupiInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupiInfo{}

// SupiInfo List of Supis.
type SupiInfo struct {
	SupiList []string `json:"supiList"`
}

type _SupiInfo SupiInfo

// NewSupiInfo instantiates a new SupiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupiInfo(supiList []string) *SupiInfo {
	this := SupiInfo{}
	this.SupiList = supiList
	return &this
}

// NewSupiInfoWithDefaults instantiates a new SupiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupiInfoWithDefaults() *SupiInfo {
	this := SupiInfo{}
	return &this
}

// GetSupiList returns the SupiList field value
func (o *SupiInfo) GetSupiList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupiList
}

// GetSupiListOk returns a tuple with the SupiList field value
// and a boolean to check if the value has been set.
func (o *SupiInfo) GetSupiListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupiList, true
}

// SetSupiList sets field value
func (o *SupiInfo) SetSupiList(v []string) {
	o.SupiList = v
}

func (o SupiInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supiList"] = o.SupiList
	return toSerialize, nil
}

func (o *SupiInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"supiList",
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

	varSupiInfo := _SupiInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupiInfo)

	if err != nil {
		return err
	}

	*o = SupiInfo(varSupiInfo)

	return err
}

type NullableSupiInfo struct {
	value *SupiInfo
	isSet bool
}

func (v NullableSupiInfo) Get() *SupiInfo {
	return v.value
}

func (v *NullableSupiInfo) Set(val *SupiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSupiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSupiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupiInfo(val *SupiInfo) *NullableSupiInfo {
	return &NullableSupiInfo{value: val, isSet: true}
}

func (v NullableSupiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


