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

// checks if the PtwParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PtwParameters{}

// PtwParameters struct for PtwParameters
type PtwParameters struct {
	OperationMode OperationMode `json:"operationMode"`
	PtwValue string `json:"ptwValue" validate:"regexp=^([0-1]{4})$"`
	ExtendedPtwValue *string `json:"extendedPtwValue,omitempty" validate:"regexp=^([0-1]{8})$"`
}

type _PtwParameters PtwParameters

// NewPtwParameters instantiates a new PtwParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPtwParameters(operationMode OperationMode, ptwValue string) *PtwParameters {
	this := PtwParameters{}
	this.OperationMode = operationMode
	this.PtwValue = ptwValue
	return &this
}

// NewPtwParametersWithDefaults instantiates a new PtwParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPtwParametersWithDefaults() *PtwParameters {
	this := PtwParameters{}
	return &this
}

// GetOperationMode returns the OperationMode field value
func (o *PtwParameters) GetOperationMode() OperationMode {
	if o == nil {
		var ret OperationMode
		return ret
	}

	return o.OperationMode
}

// GetOperationModeOk returns a tuple with the OperationMode field value
// and a boolean to check if the value has been set.
func (o *PtwParameters) GetOperationModeOk() (*OperationMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationMode, true
}

// SetOperationMode sets field value
func (o *PtwParameters) SetOperationMode(v OperationMode) {
	o.OperationMode = v
}

// GetPtwValue returns the PtwValue field value
func (o *PtwParameters) GetPtwValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PtwValue
}

// GetPtwValueOk returns a tuple with the PtwValue field value
// and a boolean to check if the value has been set.
func (o *PtwParameters) GetPtwValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PtwValue, true
}

// SetPtwValue sets field value
func (o *PtwParameters) SetPtwValue(v string) {
	o.PtwValue = v
}

// GetExtendedPtwValue returns the ExtendedPtwValue field value if set, zero value otherwise.
func (o *PtwParameters) GetExtendedPtwValue() string {
	if o == nil || IsNil(o.ExtendedPtwValue) {
		var ret string
		return ret
	}
	return *o.ExtendedPtwValue
}

// GetExtendedPtwValueOk returns a tuple with the ExtendedPtwValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PtwParameters) GetExtendedPtwValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedPtwValue) {
		return nil, false
	}
	return o.ExtendedPtwValue, true
}

// HasExtendedPtwValue returns a boolean if a field has been set.
func (o *PtwParameters) HasExtendedPtwValue() bool {
	if o != nil && !IsNil(o.ExtendedPtwValue) {
		return true
	}

	return false
}

// SetExtendedPtwValue gets a reference to the given string and assigns it to the ExtendedPtwValue field.
func (o *PtwParameters) SetExtendedPtwValue(v string) {
	o.ExtendedPtwValue = &v
}

func (o PtwParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PtwParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationMode"] = o.OperationMode
	toSerialize["ptwValue"] = o.PtwValue
	if !IsNil(o.ExtendedPtwValue) {
		toSerialize["extendedPtwValue"] = o.ExtendedPtwValue
	}
	return toSerialize, nil
}

func (o *PtwParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operationMode",
		"ptwValue",
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

	varPtwParameters := _PtwParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPtwParameters)

	if err != nil {
		return err
	}

	*o = PtwParameters(varPtwParameters)

	return err
}

type NullablePtwParameters struct {
	value *PtwParameters
	isSet bool
}

func (v NullablePtwParameters) Get() *PtwParameters {
	return v.value
}

func (v *NullablePtwParameters) Set(val *PtwParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePtwParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePtwParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtwParameters(val *PtwParameters) *NullablePtwParameters {
	return &NullablePtwParameters{value: val, isSet: true}
}

func (v NullablePtwParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtwParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


