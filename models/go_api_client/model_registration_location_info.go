/*
Nudm_UECM

Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RegistrationLocationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationLocationInfo{}

// RegistrationLocationInfo This data type contains the information about serving AMF, optional VGMLC and access type related information used by (H)GMLC. 
type RegistrationLocationInfo struct {
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId string `json:"amfInstanceId"`
	Guami *Guami `json:"guami,omitempty"`
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	VgmlcAddress *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []AccessType `json:"accessTypeList"`
}

type _RegistrationLocationInfo RegistrationLocationInfo

// NewRegistrationLocationInfo instantiates a new RegistrationLocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationLocationInfo(amfInstanceId string, accessTypeList []AccessType) *RegistrationLocationInfo {
	this := RegistrationLocationInfo{}
	this.AmfInstanceId = amfInstanceId
	this.AccessTypeList = accessTypeList
	return &this
}

// NewRegistrationLocationInfoWithDefaults instantiates a new RegistrationLocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationLocationInfoWithDefaults() *RegistrationLocationInfo {
	this := RegistrationLocationInfo{}
	return &this
}

// GetAmfInstanceId returns the AmfInstanceId field value
func (o *RegistrationLocationInfo) GetAmfInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfInstanceId
}

// GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field value
// and a boolean to check if the value has been set.
func (o *RegistrationLocationInfo) GetAmfInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfInstanceId, true
}

// SetAmfInstanceId sets field value
func (o *RegistrationLocationInfo) SetAmfInstanceId(v string) {
	o.AmfInstanceId = v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *RegistrationLocationInfo) GetGuami() Guami {
	if o == nil || IsNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationLocationInfo) GetGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.Guami) {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *RegistrationLocationInfo) HasGuami() bool {
	if o != nil && !IsNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *RegistrationLocationInfo) SetGuami(v Guami) {
	o.Guami = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *RegistrationLocationInfo) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationLocationInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *RegistrationLocationInfo) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *RegistrationLocationInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetVgmlcAddress returns the VgmlcAddress field value if set, zero value otherwise.
func (o *RegistrationLocationInfo) GetVgmlcAddress() VgmlcAddress {
	if o == nil || IsNil(o.VgmlcAddress) {
		var ret VgmlcAddress
		return ret
	}
	return *o.VgmlcAddress
}

// GetVgmlcAddressOk returns a tuple with the VgmlcAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationLocationInfo) GetVgmlcAddressOk() (*VgmlcAddress, bool) {
	if o == nil || IsNil(o.VgmlcAddress) {
		return nil, false
	}
	return o.VgmlcAddress, true
}

// HasVgmlcAddress returns a boolean if a field has been set.
func (o *RegistrationLocationInfo) HasVgmlcAddress() bool {
	if o != nil && !IsNil(o.VgmlcAddress) {
		return true
	}

	return false
}

// SetVgmlcAddress gets a reference to the given VgmlcAddress and assigns it to the VgmlcAddress field.
func (o *RegistrationLocationInfo) SetVgmlcAddress(v VgmlcAddress) {
	o.VgmlcAddress = &v
}

// GetAccessTypeList returns the AccessTypeList field value
func (o *RegistrationLocationInfo) GetAccessTypeList() []AccessType {
	if o == nil {
		var ret []AccessType
		return ret
	}

	return o.AccessTypeList
}

// GetAccessTypeListOk returns a tuple with the AccessTypeList field value
// and a boolean to check if the value has been set.
func (o *RegistrationLocationInfo) GetAccessTypeListOk() ([]AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTypeList, true
}

// SetAccessTypeList sets field value
func (o *RegistrationLocationInfo) SetAccessTypeList(v []AccessType) {
	o.AccessTypeList = v
}

func (o RegistrationLocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationLocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amfInstanceId"] = o.AmfInstanceId
	if !IsNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.VgmlcAddress) {
		toSerialize["vgmlcAddress"] = o.VgmlcAddress
	}
	toSerialize["accessTypeList"] = o.AccessTypeList
	return toSerialize, nil
}

func (o *RegistrationLocationInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amfInstanceId",
		"accessTypeList",
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

	varRegistrationLocationInfo := _RegistrationLocationInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegistrationLocationInfo)

	if err != nil {
		return err
	}

	*o = RegistrationLocationInfo(varRegistrationLocationInfo)

	return err
}

type NullableRegistrationLocationInfo struct {
	value *RegistrationLocationInfo
	isSet bool
}

func (v NullableRegistrationLocationInfo) Get() *RegistrationLocationInfo {
	return v.value
}

func (v *NullableRegistrationLocationInfo) Set(val *RegistrationLocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationLocationInfo(val *RegistrationLocationInfo) *NullableRegistrationLocationInfo {
	return &NullableRegistrationLocationInfo{value: val, isSet: true}
}

func (v NullableRegistrationLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


