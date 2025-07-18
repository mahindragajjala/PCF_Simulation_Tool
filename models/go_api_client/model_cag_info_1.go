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

// checks if the CagInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CagInfo1{}

// CagInfo1 contains CAG IDs.
type CagInfo1 struct {
	CagList []string `json:"cagList"`
}

type _CagInfo1 CagInfo1

// NewCagInfo1 instantiates a new CagInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCagInfo1(cagList []string) *CagInfo1 {
	this := CagInfo1{}
	this.CagList = cagList
	return &this
}

// NewCagInfo1WithDefaults instantiates a new CagInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCagInfo1WithDefaults() *CagInfo1 {
	this := CagInfo1{}
	return &this
}

// GetCagList returns the CagList field value
func (o *CagInfo1) GetCagList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CagList
}

// GetCagListOk returns a tuple with the CagList field value
// and a boolean to check if the value has been set.
func (o *CagInfo1) GetCagListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CagList, true
}

// SetCagList sets field value
func (o *CagInfo1) SetCagList(v []string) {
	o.CagList = v
}

func (o CagInfo1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CagInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cagList"] = o.CagList
	return toSerialize, nil
}

func (o *CagInfo1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cagList",
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

	varCagInfo1 := _CagInfo1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCagInfo1)

	if err != nil {
		return err
	}

	*o = CagInfo1(varCagInfo1)

	return err
}

type NullableCagInfo1 struct {
	value *CagInfo1
	isSet bool
}

func (v NullableCagInfo1) Get() *CagInfo1 {
	return v.value
}

func (v *NullableCagInfo1) Set(val *CagInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableCagInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableCagInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCagInfo1(val *CagInfo1) *NullableCagInfo1 {
	return &NullableCagInfo1{value: val, isSet: true}
}

func (v NullableCagInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCagInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


